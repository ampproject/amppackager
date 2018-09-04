package amppackager

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
	"github.com/theckman/go-flock"
)

// This is an abstraction over a single file on a remote storage mechanism. It
// is meant for use-cases where there will be mostly reads. The update callback
// is assumed to be expensive, and thus it should be coordinated among all
// replicas and only done once.
type Updateable interface {
	// Reads the contents of the file. Calls isExpired(contents); if true,
	// then it calls update() and writes the returned contents back to the
	// file.
	Read(ctx context.Context, isExpired func([]byte) bool, update func([]byte) []byte) ([]byte, error)
}

// Uses the OS's file locking mechanisms to obtain shared/exclusive locks to
// ensure update() is only called once. This is probably good enough for a few
// processes running on one server.
//
// For more processes than that, or for a distributed deployment over NFS, it
// would require more reading / testing to see if this is OK. I'm not an expert
// on distributed systems and http://0pointer.de/blog/projects/locking.html and
// https://gavv.github.io/blog/file-locks/ have lots of warnings, and I haven't
// found any documentation on how NFS decides on an exclusive lock owner if
// there's contention. https://tools.ietf.org/html/rfc3530#section-8.1.5
// suggests NFSv4 supports some lock sequencing mechanism that I assume won't
// result in starvation, but I don't know how well that's supported by various
// clients & servers.
//
// Users interested in scaling this widely may want to implement their own
// Updateable using some reasonable remote storage / leader election libraries.
type LocalFile struct {
	path string
}

func (this *LocalFile) Read(ctx context.Context, isExpired func([]byte) bool, update func([]byte) []byte) ([]byte, error) {
	lock := flock.NewFlock(this.path)
	locked, err := lock.TryRLock()
	if err != nil {
		return nil, errors.Wrapf(err, "obtaining shared lock for %s", this.path)
	}
	if !locked {
		return nil, errors.Errorf("unable to obtain shared lock for %s", this.path)
	}
	defer func() {
		if err = lock.Unlock(); err != nil {
			log.Printf("Error unlocking %s; %+v", this.path, err)
		}
	}()
	contents, err := ioutil.ReadFile(this.path)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", this.path)
	}
	select {
	case <-ctx.Done():
		return nil, errors.Wrapf(ctx.Err(), "while reading %s", this.path)
	default:
		if !isExpired(contents) {
			return contents, nil
		}
		// Upgrade to a write-lock. It seems this may or may not be atomic, depending on the system.
		locked, err = lock.TryLock()
		if err != nil {
			return nil, errors.Wrapf(err, "obtaining exclusive lock for %s", this.path)
		}
		if !locked {
			return nil, errors.Errorf("unable to obtain exclusive lock for %s", this.path)
		}
		contents = update(contents)
		// TODO(twifkak): Should I write to a tempfile in the same dir and move into place, instead?
		if err = ioutil.WriteFile(this.path, contents, 0700); err != nil {
			return nil, errors.Wrapf(err, "writing %s", this.path)
		}
		return contents, nil
	}
}
