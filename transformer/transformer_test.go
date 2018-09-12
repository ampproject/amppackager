package transformer

import (
	"testing"

	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
)

func TestProcess(t *testing.T) {
	var engine *transformers.Engine
	// Remember the original function and reinstate after test
	orig := runTransform
	defer func() { runTransform = orig }()
	runTransform = func(e *transformers.Engine) {
		engine = e
	}

	// TODO(angielin): Test for func identity equality.
	tests := []struct {
		config      rpb.Request_TransformersConfig
		expectedLen int
	}{
		{rpb.Request_DEFAULT, 8},
		{rpb.Request_NONE, 0},
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

		if len(engine.Transformers) != tc.expectedLen {
			t.Errorf("Process(%v) number of transformers, get=%d, want=%d", tc.config, len(engine.Transformers), tc.expectedLen)
		}
	}
}

func TestCustom(t *testing.T) {
	var engine *transformers.Engine
	// Remember the original function and reinstate after test
	orig := runTransform
	defer func() { runTransform = orig }()
	runTransform = func(e *transformers.Engine) {
		engine = e
	}

	// Case insensitive
	tests := []string{
		"aMpBoIlerplate",
		"AMPRuntimeCSS",
		"linkTag",
		"metaTag",
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

		if len(engine.Transformers) != 1 {
			t.Errorf("Process(%v) expected successful transformer lookup", tc)
		}
	}
}

func TestCustomFail(t *testing.T) {
	// Remember the original function and reinstate after test
	orig := runTransform
	defer func() { runTransform = orig }()
	runTransform = func(e *transformers.Engine) {
		// do nothing
	}

	r := rpb.Request{Html: "<lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{"does_not_exist"}}
	if got, err := Process(&r); err == nil {
		t.Fatalf("Process(%v) = %s, nil; want error", r, got)
	}
}
