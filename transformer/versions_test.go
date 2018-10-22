package transformer

import (
	"testing"

	rpb "github.com/ampproject/amppackager/transformer/request"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		rs []*rpb.VersionRange
		expectedError        bool
	}{
		{
			rs:            nil,
			expectedError: false,
		},
		{
			rs:            []*rpb.VersionRange{},
			expectedError: false,
		},
		{
			rs:            []*rpb.VersionRange{{Max: 1, Min: 2}},
			expectedError: true, // Malformed.
		},
		{
			rs:            []*rpb.VersionRange{{Max: 1, Min: -1}},
			expectedError: true, // Negative.
		},
		{
			rs:            []*rpb.VersionRange{{Max: 1, Min: 1}},
			expectedError: false,
		},
		{
			rs:            []*rpb.VersionRange{{Max: 2, Min: 1}, {Max: 1, Min: 1}},
			expectedError: true, // Overlapping.
		},
		{
			rs:            []*rpb.VersionRange{{Max: 3, Min: 2}, {Max: 1, Min: 1}},
			expectedError: false,
		},
		{
			rs:            []*rpb.VersionRange{{Max: 1, Min: 1}, {Max: 3, Min: 2}},
			expectedError: true, // Out of order.
		},
	}
	for _, test := range tests {
		err := validateRequest(test.rs)
		if test.expectedError != (err != nil) {
			t.Errorf("validateRequest(%+v) unexpected err = %v", test.rs, err)
		}
	}
}

func TestSelectVersion(t *testing.T) {
	tests := []struct {
		requested, supported []*rpb.VersionRange
		expectedVersion      int64
		expectedError        bool
	}{
		// No supported.
		{
			requested:     []*rpb.VersionRange{{Max: 4, Min: 1}},
			supported:     []*rpb.VersionRange{},
			expectedError: true,
		},
		// No request, highest version supported.
		{
			requested:       []*rpb.VersionRange{},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			expectedVersion: 4,
		},
		// No request (nil), highest version supported.
		{
			requested:       nil,
			supported:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			expectedVersion: 4,
		},
		// One version, match.
		{
			requested:       []*rpb.VersionRange{{Max: 1, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 1, Min: 1}},
			expectedVersion: 1,
		},
		// One version, mismatch.
		{
			requested:     []*rpb.VersionRange{{Max: 1, Min: 1}},
			supported:     []*rpb.VersionRange{{Max: 2, Min: 2}},
			expectedError: true,
		},
		// Supported is a subset of requested.
		{
			requested:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 2, Min: 2}},
			expectedVersion: 2,
		},
		// Requested is a subset of supported.
		{
			requested:       []*rpb.VersionRange{{Max: 2, Min: 2}},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			expectedVersion: 2,
		},
		// Requested and supported abut.
		{
			requested:       []*rpb.VersionRange{{Max: 4, Min: 2}},
			supported:       []*rpb.VersionRange{{Max: 2, Min: 1}},
			expectedVersion: 2,
		},
		// Requested includes a hole in supported.
		{
			requested:       []*rpb.VersionRange{{Max: 2, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 3}, {Max: 1, Min: 1}},
			expectedVersion: 1,
		},
		// Requested requires a hole in supported.
		{
			requested:     []*rpb.VersionRange{{Max: 2, Min: 2}},
			supported:     []*rpb.VersionRange{{Max: 4, Min: 3}, {Max: 1, Min: 1}},
			expectedError: true,
		},
		// Overlap is between the second elements of both requested and supported.
		{
			requested:       []*rpb.VersionRange{{Max: 6, Min: 5}, {Max: 2, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 3}, {Max: 1, Min: 1}},
			expectedVersion: 1,
		},
	}
	for _, test := range tests {
		v, err := selectVersion(test.requested, test.supported)
		if test.expectedError {
			if err == nil {
				t.Errorf("selectVersion(%+v, %+v), want err, got v = %d", test.requested, test.supported, v)
			}
		} else if err != nil {
			t.Errorf("selectVersion(%+v, %+v), want v = %d, got err = %v", test.requested, test.supported, test.expectedVersion, err)
		} else if v != test.expectedVersion {
			t.Errorf("selectVersion(%+v, %+v), want v = %d, got v = %v", test.requested, test.supported, test.expectedVersion, v)
		}
	}
}
