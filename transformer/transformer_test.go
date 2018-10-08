package transformer

import (
	"errors"
	"testing"

	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
)

func TestProcess(t *testing.T) {
	var fns []func(*transformers.Context) error
	// Remember the original function and reinstate after test
	orig := runTransformers
	defer func() { runTransformers = orig }()
	runTransformers = func(e *transformers.Context, fs []func(*transformers.Context) error) error {
		fns = fs
		return nil
	}

	// TODO(alin04): Test for func identity equality.
	tests := []struct {
		config      rpb.Request_TransformersConfig
		expectedLen int
	}{
		{rpb.Request_DEFAULT, 5},
		{rpb.Request_NONE, 0},
		{rpb.Request_VALIDATION, 1},
		{rpb.Request_CUSTOM, 0},
	}

	for _, tc := range tests {
		r := rpb.Request{Html: "<html ⚡><lemur>", Config: tc.config}
		got, err := Process(&r)
		if err != nil {
			t.Fatalf("Process(%v) unexpectedly failed %v", tc.config, err)
		}

		want := "<html ⚡><head></head><body><lemur></lemur></body></html>"
		if got != want {
			t.Errorf("Process(%v) = %q, want = %q", tc.config, got, want)
		}

		if len(fns) != tc.expectedLen {
			t.Errorf("Process(%v) number of transformers, get=%d, want=%d", tc.config, len(fns), tc.expectedLen)
		}
	}
}

func TestCustom(t *testing.T) {
	var fns []func(*transformers.Context) error
	// Remember the original function and reinstate after test
	orig := runTransformers
	defer func() { runTransformers = orig }()
	runTransformers = func(e *transformers.Context, fs []func(*transformers.Context) error) error {
		fns = fs
		return nil
	}

	// Case insensitive
	tests := []string{
		"aMpBoIlerplate",
		"AMPRuntimeCSS",
		"linktag",
		"metaTag",
		"NODECLEANUP",
		"reorderHead",
		"serverSideRendering",
		"transformedIdentifier",
		"uRl",
	}
	for _, tc := range tests {
		r := rpb.Request{Html: "<html ⚡><lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{tc}}
		if _, err := Process(&r); err != nil {
			t.Fatalf("Process(%v) unexpectedly failed %v", tc, err)
		}

		if len(fns) != 1 {
			t.Errorf("Process(%v) expected successful transformer lookup", tc)
		}
	}
}

func TestCustomFail(t *testing.T) {
	r := rpb.Request{Html: "<html ⚡><lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{"does_not_exist"}}
	if got, err := Process(&r); err == nil {
		t.Fatalf("Process(%v) = %s, nil; want error", r, got)
	}
}

func TestError(t *testing.T) {
	s := "something happened!"
	// Remember the original function and reinstate after test
	orig := runTransformers
	defer func() { runTransformers = orig }()
	runTransformers = func(e *transformers.Context, fs []func(*transformers.Context) error) error {
		return errors.New(s)
	}

	r := rpb.Request{Html: "<html ⚡><lemur>", Config: rpb.Request_DEFAULT}
	_, err := Process(&r)
	if err == nil {
		t.Fatalf("Process() unexpectedly succeeded")
	}
	if err.Error() != s {
		t.Fatalf("mismatched error. got=%s, want=%s", err.Error(), s)
	}
}

func TestRequireAMPAttribute(t *testing.T) {
	tests := []struct {
		desc     string
		html     string
		expectedError bool
	}{
		{
			"⚡",
			"<html ⚡><head></head><body></body></html>",
			false,
		},
		{
			"amp",
			"<html amp><head></head><body></body></html>",
			false,
		},
		{
			"AMP",
			"<HTML AMP><HEAD></HEAD><BODY></BODY></HTML>",
			false,
		},
		{
			"⚡4ads",
			"<html ⚡4ads><head></head><body></body></html>",
			false,
		},
		{
			"amp4ads",
			"<html amp4ads><head></head><body></body></html>",
			false,
		},
		{
			"AMP4ADS",
			"<HTML AMP4ADS><HEAD></HEAD><BODY></BODY></HTML>",
			false,
		},
		{
			"⚡4email",
			"<html ⚡4email><head></head><body></body></html>",
			false,
		},
		{
			"amp4email",
			"<html amp4email><head></head><body></body></html>",
			false,
		},
		{
			"AMP4EMAIL",
			"<HTML AMP4EMAIL><HEAD></HEAD><BODY></BODY></HTML>",
			false,
		},
		{
			"not AMP",
			"<html><head></head><body></body></html>",
			true,
		},
	}
	for _, test := range tests {
		r := rpb.Request{Html: test.html, Config: rpb.Request_NONE}
		_, err := Process(&r)
		if (err != nil) != test.expectedError {
			t.Errorf("%s: RequireAMPAttribute() has error=%#v want=%t", test.desc, err, test.expectedError)
		}
	}
 }
