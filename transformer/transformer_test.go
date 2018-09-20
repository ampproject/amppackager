package transformer

import (
	"testing"

	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
)

func TestProcess(t *testing.T) {
	var fns []func(*transformers.Context)
	// Remember the original function and reinstate after test
	orig := runTransformers
	defer func() { runTransformers = orig }()
	runTransformers = func(e *transformers.Context, fs []func(*transformers.Context)) {
		fns = fs
	}

	// TODO(alin04): Test for func identity equality.
	tests := []struct {
		config      rpb.Request_TransformersConfig
		expectedLen int
	}{
		{rpb.Request_DEFAULT, 5},
		{rpb.Request_NONE, 1},
		{rpb.Request_VALIDATION, 1},
		{rpb.Request_CUSTOM, 0},
	}

	for _, tc := range tests {
		r := rpb.Request{Html: "<lemur>", Config: tc.config}
		got, err := Process(&r)
		if err != nil {
			t.Fatalf("Process(%v) unexpectedly failed %v", tc.config, err)
		}

		want := "<html><head></head><body><lemur></lemur></body></html>"
		if got != want {
			t.Errorf("Process(%v) = %q, want = %q", tc.config, got, want)
		}

		if len(fns) != tc.expectedLen {
			t.Errorf("Process(%v) number of transformers, get=%d, want=%d", tc.config, len(fns), tc.expectedLen)
		}
	}
}

func TestCustom(t *testing.T) {
	var fns []func(*transformers.Context)
	// Remember the original function and reinstate after test
	orig := runTransformers
	defer func() { runTransformers = orig }()
	runTransformers = func(e *transformers.Context, fs []func(*transformers.Context)) {
		fns = fs
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
		r := rpb.Request{Html: "<lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{tc}}
		if _, err := Process(&r); err != nil {
			t.Fatalf("Process(%v) unexpectedly failed %v", tc, err)
		}

		if len(fns) != 1 {
			t.Errorf("Process(%v) expected successful transformer lookup", tc)
		}
	}
}

func TestCustomFail(t *testing.T) {
	r := rpb.Request{Html: "<lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{"does_not_exist"}}
	if got, err := Process(&r); err == nil {
		t.Fatalf("Process(%v) = %s, nil; want error", r, got)
	}
}
