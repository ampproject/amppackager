package certcache

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/gofrs/flock"
	"github.com/pkg/errors"
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

// Check whether a file or directory exists.
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func (this *LocalFile) Read(ctx context.Context, isExpired func([]byte) bool, update func([]byte) []byte) ([]byte, error) {
	// Use independent .lock file; necessary on Windows to avoid "The process cannot
	// access the file because another process has locked a portion of the file."
	lockPath := this.path + ".lock"
	lock := flock.New(lockPath)
	locked, err := lock.TryRLock()
	if err != nil {
		return nil, errors.Wrapf(err, "obtaining shared lock for %s", lockPath)
	}
	if !locked {
		return nil, errors.Errorf("unable to obtain shared lock for %s", lockPath)
	}
	defer func() {
		if err = lock.Unlock(); err != nil {
			log.Printf("Error unlocking %s; %+v", lockPath, err)
		}
	}()

	// Check whether OCSP cache file exists.
	pathExists, err := exists(this.path)
	if err != nil {
		return nil, errors.Wrapf(err, "checking file exists %s", this.path)
	}

	// Initialize empty contents.
	var contents []byte

	// If cache file exists, read it and check freshness. Note that zero-length
	// contents are considered "expired" by isExpired(). If an attempt is made
	// to read the file before it exists on Windows, error "The system cannot
	// find the file specified." is thrown.
	if pathExists {
		contents, err = ioutil.ReadFile(this.path)
		if err != nil {
			return nil, errors.Wrapf(err, "reading %s", this.path)
		}
	}

	// At first glance, this looks like "broken" double-checked locking, as in
	// http://www.cs.umd.edu/~pugh/java/memoryModel/DoubleCheckedLocking.html.
	// However, the difference is that a read lock is established first, so
	// that this shouldn't be looking at a partially-written file.
	select {
	case <-ctx.Done():
		return nil, errors.Wrapf(ctx.Err(), "while reading %s", this.path)
	default:
		if !isExpired(contents) {
			return contents, nil
		}
		// Upgrade to a write-lock. It seems this may or may not be atomic, depending on the system.
		// Windows does not handle a lock "upgrade", hence unlock before lock.
		if runtime.GOOS == "windows" {
			if err = lock.Unlock(); err != nil {
				return nil, errors.Wrapf(err, "Error unlocking %s", lockPath)
			}
		}
		locked, err = lock.TryLock()
		if err != nil {
			return nil, errors.Wrapf(err, "obtaining exclusive lock for %s", lockPath)
		}
		if !locked {
			return nil, errors.Errorf("unable to obtain exclusive lock for %s", lockPath)
		}

		// Reread the file while in write-lock, to make the
		// read-modify-write atomic, and thus reduce the chance of
		// multiple calls to update() in parallel.
		if pathExists {
			contents, err := ioutil.ReadFile(this.path)
			if err != nil {
				return nil, errors.Wrapf(err, "rereading %s", this.path)
			}
			if !isExpired(contents) {
				return contents, nil
			}
		}

		contents = update(contents)
		// TODO(twifkak): Should I write to a tempfile in the same dir and move into place, instead?
		if err = ioutil.WriteFile(this.path, contents, 0600); err != nil {
			return nil, errors.Wrapf(err, "writing %s", this.path)
		}
		return contents, nil
	}
}

// Represents an in-memory copy of a file.
type InMemory struct {
	mu       sync.RWMutex
	contents []byte
}

func (this *InMemory) read() []byte {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.contents
}

func (this *InMemory) Read(ctx context.Context, isExpired func([]byte) bool, update func([]byte) []byte) ([]byte, error) {
	contents := this.read()
	// The note above about double-checked locking applies here.
	if !isExpired(contents) {
		return contents, nil
	}
	this.mu.Lock()
	defer this.mu.Unlock()
	if !isExpired(this.contents) {
		return this.contents, nil
	}
	this.contents = update(this.contents)
	return this.contents, nil
}

// Represents a file backed by two updateables. If the first is expired, then
// the second is consulted, and only if both are expired is update() run (and
// the contents of both updateables updated).
type Chained struct {
	first, second Updateable
}

func (this *Chained) Read(ctx context.Context, isExpired func([]byte) bool, update func([]byte) []byte) ([]byte, error) {
	return this.first.Read(ctx, isExpired, func([]byte) []byte {
		contents, err := this.second.Read(ctx, isExpired, update)
		if err != nil {
			log.Printf("%+v", err)
			return nil
		}
		return contents
	})
}
