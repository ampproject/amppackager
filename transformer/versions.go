package transformer

import (
	"math"

	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/pkg/errors"
)

// SupportedVersions is a set of all transform versions supported by this
// snapshot of the library. This should not include transforms not yet
// snapshotted to a finalized version.
// The ranges should be non-overlapping and in descending order.
// Visible for test.
var SupportedVersions = []*rpb.VersionRange{{Min: 1, Max: 1}}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// intervalsOverlap returns true if there exists at least one value in both `a`
// and `b`. This assumes `a` and `b` are not malformed.
func intervalsOverlap(a, b *rpb.VersionRange) bool {
	return a.Max >= b.Min && a.Min <= b.Max
}

// isGreater returns true if all of the values in `a` are greater than all of
// the values in `b`. This assumes `a` and `b` are not malformed.
func isGreater(a, b *rpb.VersionRange) bool {
	return a.Min > b.Max
}

// validateRequest verifies the request is properly formed.
func validateRequest(rs []*rpb.VersionRange) error {
	var lastMin int64 = math.MaxInt64
	for i, r := range rs {
		if r.Min > r.Max {
			return errors.Errorf("malformed range %v at index %d", r, i)
		}
		if r.Min < 1 || r.Max < 1 {
			return errors.Errorf("non-positive value in range %v at index %d", r, i)
		}
		if r.Max >= lastMin {
			return errors.Errorf("overlapping or out-of-order range %v at index %d", r, i)
		}
		lastMin = r.Min
	}
	return nil
}

func selectVersion(requested, supported []*rpb.VersionRange) (int64, error) {
	// If request is invalid, then error.
	if err := validateRequest(requested); err != nil {
		return 0, err
	}
	// If no versions are supported, then huh.
	if len(supported) == 0 {
		return 0, errors.New("no version at all is supported by this transformer")
	}
	// If no requested versions are specified, then return the highest
	// support version.
	if len(requested) == 0 {
		return supported[0].Max, nil
	}
	// Otherwise, search for the highest number in both sets.
	for r, s := 0, 0; r < len(requested) && s < len(supported); {
		if intervalsOverlap(requested[r], supported[s]) {
			return min(requested[r].Max, supported[s].Max), nil
		} else if isGreater(requested[r], supported[s]) {
			r++
		} else {
			s++
		}
	}
	return 0, errors.New("no requested version is supported by this transformer")
}

// SelectVersion returns the highest requested version number that the
// transformer supports, or an error if no such version exists. The requested
// list must consist of non-overlapping ranges in descending order.
//
// The transfomer will select the highest version from requested that it
// supports. If requested is nil or empty, the transfomer will select the
// highest version it supports. In both cases, it shouldn't include
// in-development versions.
func SelectVersion(requested []*rpb.VersionRange) (int64, error) {
	return selectVersion(requested, SupportedVersions)
}
