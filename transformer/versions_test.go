package transformer

import (
	"testing"

	rpb "github.com/ampproject/amppackager/transformer/request"
)

func TestValidateRequest(t *testing.T) {
	tcs := []struct {
		desc          string
		rs            []*rpb.VersionRange
		expectedError bool
	}{
		{
			desc:          "nil",
			rs:            nil,
			expectedError: false,
		},
		{
			desc:          "empty",
			rs:            []*rpb.VersionRange{},
			expectedError: false,
		},
		{
			desc:          "malformed",
			rs:            []*rpb.VersionRange{{Max: 1, Min: 2}},
			expectedError: true,
		},
		{
			desc:          "negative",
			rs:            []*rpb.VersionRange{{Max: 1, Min: -1}},
			expectedError: true,
		},
		{
			desc:          "same",
			rs:            []*rpb.VersionRange{{Max: 1, Min: 1}},
			expectedError: false,
		},
		{
			desc:          "overlapping",
			rs:            []*rpb.VersionRange{{Max: 2, Min: 1}, {Max: 1, Min: 1}},
			expectedError: true,
		},
		{
			desc:          "valid",
			rs:            []*rpb.VersionRange{{Max: 3, Min: 2}, {Max: 1, Min: 1}},
			expectedError: false,
		},
		{
			desc:          "out of order",
			rs:            []*rpb.VersionRange{{Max: 1, Min: 1}, {Max: 3, Min: 2}},
			expectedError: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			err := validateRequest(tc.rs)
			if tc.expectedError != (err != nil) {
				t.Errorf("validateRequest(%+v) unexpected err = %v", tc.rs, err)
			}
		})
	}
}

func TestSelectVersion(t *testing.T) {
	tcs := []struct {
		desc                 string
		requested, supported []*rpb.VersionRange
		expectedVersion      int64
		expectedError        bool
	}{
		{
			desc:          "not supported",
			requested:     []*rpb.VersionRange{{Max: 4, Min: 1}},
			supported:     []*rpb.VersionRange{},
			expectedError: true,
		},
		{
			desc:            "No request, highest version supported.",
			requested:       []*rpb.VersionRange{},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			expectedVersion: 4,
		},
		{
			desc:            "No request (nil), highest version supported.",
			requested:       nil,
			supported:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			expectedVersion: 4,
		},
		{
			desc:            "One version, match.",
			requested:       []*rpb.VersionRange{{Max: 1, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 1, Min: 1}},
			expectedVersion: 1,
		},

		{
			desc:          "One version, mismatch.",
			requested:     []*rpb.VersionRange{{Max: 1, Min: 1}},
			supported:     []*rpb.VersionRange{{Max: 2, Min: 2}},
			expectedError: true,
		},

		{
			desc:            "Supported is a subset of requested.",
			requested:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 2, Min: 2}},
			expectedVersion: 2,
		},

		{
			desc:            "Requested is a subset of supported.",
			requested:       []*rpb.VersionRange{{Max: 2, Min: 2}},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 1}},
			expectedVersion: 2,
		},

		{
			desc:            "Requested and supported abut.",
			requested:       []*rpb.VersionRange{{Max: 4, Min: 2}},
			supported:       []*rpb.VersionRange{{Max: 2, Min: 1}},
			expectedVersion: 2,
		},

		{
			desc:            "Requested includes a hole in supported.",
			requested:       []*rpb.VersionRange{{Max: 2, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 3}, {Max: 1, Min: 1}},
			expectedVersion: 1,
		},

		{
			desc:          "Requested requires a hole in supported.",
			requested:     []*rpb.VersionRange{{Max: 2, Min: 2}},
			supported:     []*rpb.VersionRange{{Max: 4, Min: 3}, {Max: 1, Min: 1}},
			expectedError: true,
		},

		{
			desc:            "Overlap is between the second elements of both requested and supported.",
			requested:       []*rpb.VersionRange{{Max: 6, Min: 5}, {Max: 2, Min: 1}},
			supported:       []*rpb.VersionRange{{Max: 4, Min: 3}, {Max: 1, Min: 1}},
			expectedVersion: 1,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, err := selectVersion(tc.requested, tc.supported)
			t.Logf("selectVersion(%+v, %+v) = %d", tc.requested, tc.supported, v)
			if tc.expectedError {
				if err == nil {
					t.Error("wanted, but didn't get err")
				}
			} else if err != nil {
				t.Errorf("got unexpected err = %v", err)
			} else if v != tc.expectedVersion {
				t.Errorf("retval is unexpectedly %v", v)
			}
		})
	}
}
