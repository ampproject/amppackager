package transformer

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
	"github.com/google/go-cmp/cmp"
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
		{rpb.Request_DEFAULT, 12},
		{rpb.Request_NONE, 0},
		{rpb.Request_VALIDATION, 1},
		{rpb.Request_CUSTOM, 0},
	}

	for _, tc := range tests {
		r := rpb.Request{Html: "<html ⚡><lemur>", Config: tc.config}
		html, metadata, err := Process(&r)
		if err != nil {
			t.Fatalf("Process(%v) unexpectedly failed %v", tc.config, err)
		}

		expectedHTML := "<html ⚡><head></head><body><lemur></lemur></body></html>"
		if html != expectedHTML {
			t.Errorf("Process(%v) = %q, want = %q", tc.config, html, expectedHTML)
		}

		if metadata == nil {
			t.Errorf("Process(%v) metadata unexpectedly nil", tc.config)
		}

		if len(fns) != tc.expectedLen {
			t.Errorf("Process(%v) number of transformers, got=%d, want=%d", tc.config, len(fns), tc.expectedLen)
		}
	}
}

func TestPreloads(t *testing.T) {
	// Programmatically prepare the `> maxPreloads` test case.
	var manyScriptsHTML strings.Builder
	manyScriptsPreloads := []*rpb.Metadata_Preload{}
	manyScriptsHTML.WriteString("<html ⚡>")
	for i := 0; i <= maxPreloads; i++ {
		fmt.Fprintf(&manyScriptsHTML, `<script src="foo%d"></script>`, i)
		if i < maxPreloads {
			manyScriptsPreloads = append(manyScriptsPreloads, &rpb.Metadata_Preload{Url: fmt.Sprintf("foo%d", i), As: "script"})
		}
	}

	tests := []struct {
		html             string
		expectedPreloads []*rpb.Metadata_Preload
	}{
		{
			"<html ⚡><script>",
			[]*rpb.Metadata_Preload{},
		},
		{
			"<html ⚡><script src=foo>",
			[]*rpb.Metadata_Preload{{Url: "foo", As: "script"}},
		},
		{
			"<html ⚡><link rel=foaf href=foo>",
			[]*rpb.Metadata_Preload{},
		},
		{
			"<html ⚡><link rel=stylesheet href=foo>",
			[]*rpb.Metadata_Preload{{Url: "foo", As: "style"}},
		},
		{ // case-insensitive
			"<html ⚡><link rel=STYLEsheet href=foo>",
			[]*rpb.Metadata_Preload{{Url: "foo", As: "style"}},
		},
		{
			"<html ⚡><link rel=stylesheet href=foo><script src=bar>",
			[]*rpb.Metadata_Preload{{Url: "foo", As: "style"}, {Url: "bar", As: "script"}},
		},
		{
			manyScriptsHTML.String(),
			manyScriptsPreloads,
		},
	}

	for _, test := range tests {
		_, metadata, err := Process(&rpb.Request{Html: test.html, Config: rpb.Request_NONE})
		if err != nil {
			t.Fatalf("Process(%q) unexpectedly failed: %v", test.html, err)
		}

		if diff := cmp.Diff(test.expectedPreloads, metadata.Preloads); diff != "" {
			t.Errorf("Process(%q) preloads differ (-want +got):\n%s", test.html, diff)
		}
	}
}

func TestVersion(t *testing.T) {
	// context is the context provided by Process() to runTransformers().
	var context *transformers.Context
	// Remember the original function and reinstate after test
	orig := runTransformers
	defer func() { runTransformers = orig }()
	runTransformers = func(e *transformers.Context, fs []func(*transformers.Context) error) error {
		context = e
		return nil
	}

	_, _, err := Process(&rpb.Request{Html: "<html ⚡>", Config: rpb.Request_NONE, Version: SupportedVersions[0].Max})
	if err != nil {
		t.Fatalf("Process() unexpectedly failed: %v", err)
	}
	if context.Version != SupportedVersions[0].Max {
		t.Errorf("Incorrect context.Version = %d", context.Version)
	}

	// Construct an unsatisfied version request.
	badVersion := SupportedVersions[0].Max + 1
	_, _, err = Process(&rpb.Request{Html: "", Config: rpb.Request_NONE, Version: badVersion})
	if err == nil {
		t.Fatal("Process() unexpectedly succeeded")
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
		"aBsolUTEuRl",
		"urlRewriTE",
	}
	for _, tc := range tests {
		r := rpb.Request{Html: "<html ⚡><lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{tc}}
		if _, _, err := Process(&r); err != nil {
			t.Fatalf("Process(%v) unexpectedly failed %v", tc, err)
		}

		if len(fns) != 1 {
			t.Errorf("Process(%v) expected successful transformer lookup", tc)
		}
	}
}

func TestCustomFail(t *testing.T) {
	r := rpb.Request{Html: "<html ⚡><lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{"does_not_exist"}}
	if html, _, err := Process(&r); err == nil {
		t.Fatalf("Process(%v) = %s, nil; want error", r, html)
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
	_, _, err := Process(&r)
	if err == nil {
		t.Fatalf("Process() unexpectedly succeeded")
	}
	if err.Error() != s {
		t.Fatalf("mismatched error. got=%s, want=%s", err.Error(), s)
	}
}

func TestRequireAMPAttribute(t *testing.T) {
	tests := []struct {
		desc                     string
		html                     string
		expectedError            bool
		expectedErrorInAMP       bool
		expectedErrorInAMP4Ads   bool
		expectedErrorInAMP4Email bool
	}{
		{
			"⚡",
			"<html ⚡><head></head><body></body></html>",
			false, false, true, true,
		},
		{
			"amp",
			"<html amp><head></head><body></body></html>",
			false, false, true, true,
		},
		{
			"AMP",
			"<HTML AMP><HEAD></HEAD><BODY></BODY></HTML>",
			false, false, true, true,
		},
		{
			"⚡4ads",
			"<html ⚡4ads><head></head><body></body></html>",
			false, true, false, true,
		},
		{
			"amp4ads",
			"<html amp4ads><head></head><body></body></html>",
			false, true, false, true,
		},
		{
			"AMP4ADS",
			"<HTML AMP4ADS><HEAD></HEAD><BODY></BODY></HTML>",
			false, true, false, true,
		},
		{
			"⚡4email",
			"<html ⚡4email><head></head><body></body></html>",
			false, true, true, false,
		},
		{
			"amp4email",
			"<html amp4email><head></head><body></body></html>",
			false, true, true, false,
		},
		{
			"AMP4EMAIL",
			"<HTML AMP4EMAIL><HEAD></HEAD><BODY></BODY></HTML>",
			false, true, true, false,
		},
		{
			"amp4ads amp4email",
			"<html amp4ads amp4email><head></head><body></body></html>",
			false, true, false, false,
		},
		{
			"amp4",
			"<html amp4><head></head><body></body></html>",
			true, true, true, true,
		},
		{
			"not AMP",
			"<html><head></head><body></body></html>",
			true, true, true, true,
		},
	}
	for _, test := range tests {
		r := rpb.Request{Html: test.html, Config: rpb.Request_NONE}
		_, _, err := Process(&r)
		if (err != nil) != test.expectedError {
			t.Errorf("%s: Process() has error=%#v want=%t", test.desc, err, test.expectedError)
		}

		r = rpb.Request{Html: test.html, Config: rpb.Request_NONE, AllowedFormats: []rpb.Request_HtmlFormat{rpb.Request_AMP}}
		_, _, err = Process(&r)
		if (err != nil) != test.expectedErrorInAMP {
			t.Errorf("%s: Process(AMP) has error=%#v want=%t", test.desc, err, test.expectedErrorInAMP)
		}

		r = rpb.Request{Html: test.html, Config: rpb.Request_NONE, AllowedFormats: []rpb.Request_HtmlFormat{rpb.Request_AMP4ADS}}
		_, _, err = Process(&r)
		if (err != nil) != test.expectedErrorInAMP4Ads {
			t.Errorf("%s: Process(AMP4Ads) has error=%#v want=%t", test.desc, err, test.expectedErrorInAMP4Ads)
		}

		r = rpb.Request{Html: test.html, Config: rpb.Request_NONE, AllowedFormats: []rpb.Request_HtmlFormat{rpb.Request_AMP4EMAIL}}
		_, _, err = Process(&r)
		if (err != nil) != test.expectedErrorInAMP4Email {
			t.Errorf("%s: Process(AMP4Email) has error=%#v want=%t", test.desc, err, test.expectedErrorInAMP4Email)
		}
	}
}

func TestBaseURL(t *testing.T) {
	docURL := "http://example.com/a/page.html"
	tests := []struct {
		desc, base, expected string
	}{
		{
			"no base href",
			"<base target=_top>",
			docURL,
		},
		{
			"absolute",
			"<base href=https://www.foo.com>",
			"https://www.foo.com",
		},
		{
			"relative",
			"<base href=\"./child/to/a\">",
			"http://example.com/a/child/to/a",
		},
		{
			"relative to root",
			"<base href=\"/\">",
			"http://example.com/",
		},
	}
	for _, test := range tests {
		// Remember the original function and reinstate after test
		orig := runTransformers
		defer func() { runTransformers = orig }()
		runTransformers = func(e *transformers.Context, fs []func(*transformers.Context) error) error {
			if e.BaseURL.String() != test.expected {
				t.Errorf("%s : setBaseURL(%s)=%s, want=%s", test.desc, test.base, e.BaseURL, test.expected)
			}
			return nil
		}
		r := rpb.Request{Html: "<html amp><head>" + test.base + "</head></html>", DocumentUrl: docURL, Config: rpb.Request_NONE}
		_, _, err := Process(&r)
		if err != nil {
			t.Fatalf("Process(%v) unexpectedly failed %v", test.desc, err)
		}
	}
}
