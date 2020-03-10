package transformer

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/ampproject/amppackager/transformer/transformers"
	"github.com/google/go-cmp/cmp"
	"github.com/golang/protobuf/proto"
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
		{rpb.Request_DEFAULT, 13},
		{rpb.Request_NONE, 0},
		{rpb.Request_VALIDATION, 1},
		{rpb.Request_CUSTOM, 0},
	}

	for _, tc := range tests {
		t.Run(tc.config.String(), func(t *testing.T) {
			r := rpb.Request{Html: "<html ⚡><lemur>", Config: tc.config}
			html, metadata, err := Process(&r)
			if err != nil {
				t.Fatalf("unexpected failure %v", err)
			}

			expectedHTML := "<html ⚡><head></head><body><lemur></lemur></body></html>"
			if html != expectedHTML {
				t.Errorf("got = %q, want = %q", html, expectedHTML)
			}

			if metadata == nil {
				t.Error("metadata unexpectedly nil")
			}

			if len(fns) != tc.expectedLen {
				t.Errorf("number of transformers, got=%d, want=%d", len(fns), tc.expectedLen)
			}
		})
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

	tcs := []struct {
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

	for _, tc := range tcs {
		t.Run(tc.html, func(t *testing.T) {
			_, metadata, err := Process(&rpb.Request{Html: tc.html, Config: rpb.Request_NONE})
			if err != nil {
				t.Fatalf("unexpected failure: %v", err)
			}

			if diff := cmp.Diff(tc.expectedPreloads, metadata.Preloads, cmp.Comparer(proto.Equal)); diff != "" {
				t.Errorf("preloads differ (-want +got):\n%s", diff)
			}
		})
	}
}

func TestMaxAge(t *testing.T) {
	tcs := []struct {
		html               string
		expectedMaxAgeSecs int32
	}{
		{
			// No amp-scripts; no constraints on signing duration.
			"<html ⚡>",
			604800,
		},
		{
			// amp-script but not inline; no constraints.
			"<html ⚡><amp-script>",
			604800,
		},
		{
			// Inline amp-script; default to 1-day duration.
			"<html ⚡><amp-script script=foo>",
			86400,
		},
		{
			// Inline amp-script with explicit 4-day duration.
			"<html ⚡><amp-script script=foo max-age=345600>",
			345600,
		},
		{
			// Inline amp-script with explicit 1-year duration; capped at 7 days.
			"<html ⚡><amp-script script=foo max-age=31536000>",
			604800,
		},
		{
			// Inline amp-script with invalid duration; use default.
			"<html ⚡><amp-script script=foo max-age=aaaaaa>",
			86400,
		},
		{
			// Inline amp-script with negative duration; use 0.
			"<html ⚡><amp-script script=foo max-age=-86400>",
			0,
		},
		{
			// Two inline amp-scripts, use min of both.
			"<html ⚡><amp-script script=foo max-age=600000><amp-script script=foo max-age=500000>",
			500000,
		},
		{
			// Two inline amp-scripts, explicit > implicit.
			"<html ⚡><amp-script script=foo max-age=600000><amp-script script=foo>",
			86400,
		},
		{
			// Two inline amp-scripts, explicit < implicit.
			"<html ⚡><amp-script script=foo max-age=1><amp-script script=foo>",
			1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.html, func(t *testing.T) {
			_, metadata, err := Process(&rpb.Request{Html: tc.html, Config: rpb.Request_NONE})
			if err != nil {
				t.Fatalf("unexpected failure: %v", err)
			}

			if metadata.MaxAgeSecs != tc.expectedMaxAgeSecs {
				t.Errorf("maxAgeSecs differs; got=%d, want=%d", metadata.MaxAgeSecs, tc.expectedMaxAgeSecs)
			}
		})
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
	tcs := []string{
		"aMpBoIlerplate",
		"AMPRuntimeCSS",
		"linktag",
		"NODECLEANUP",
		"reorderHead",
		"serverSideRendering",
		"transformedIdentifier",
		"aBsolUTEuRl",
		"urlRewriTE",
	}
	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			r := rpb.Request{Html: "<html ⚡><lemur>", Config: rpb.Request_CUSTOM, Transformers: []string{tc}}
			if _, _, err := Process(&r); err != nil {
				t.Fatalf("unexpected failure %v", err)
			}

			if len(fns) != 1 {
				t.Error("expected successful transformer lookup")
			}
		})
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

func TestInvalidUTF8(t *testing.T) {
	tcs := []struct{ html, expectedError string }{
		{"<html ⚡><le\003mur>", "character U+0003 at position 13 is not allowed in AMPHTML"},
		{"<html ⚡><le\xc0mur>", "invalid UTF-8 at byte position 13"},
	}
	for _, tc := range tcs {
		t.Run(tc.html, func(t *testing.T) {
			r := rpb.Request{Html: tc.html, Config: rpb.Request_DEFAULT}
			_, _, err := Process(&r)
			if err == nil {
				t.Fatal("unexpected success")
			}
			if err.Error() != tc.expectedError {
				t.Fatalf("mismatched error. got=%s, want=%s", err.Error(), tc.expectedError)
			}
		})
	}
}

func TestRequireAMPAttribute(t *testing.T) {
	tcs := []struct {
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
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			r := rpb.Request{Html: tc.html, Config: rpb.Request_NONE}
			_, _, err := Process(&r)
			if (err != nil) != tc.expectedError {
				t.Errorf("Process() has error=%#v want=%t", err, tc.expectedError)
			}

			r = rpb.Request{Html: tc.html, Config: rpb.Request_NONE, AllowedFormats: []rpb.Request_HtmlFormat{rpb.Request_AMP}}
			_, _, err = Process(&r)
			if (err != nil) != tc.expectedErrorInAMP {
				t.Errorf("Process(AMP) has error=%#v want=%t", err, tc.expectedErrorInAMP)
			}

			r = rpb.Request{Html: tc.html, Config: rpb.Request_NONE, AllowedFormats: []rpb.Request_HtmlFormat{rpb.Request_AMP4ADS}}
			_, _, err = Process(&r)
			if (err != nil) != tc.expectedErrorInAMP4Ads {
				t.Errorf("Process(AMP4Ads) has error=%#v want=%t", err, tc.expectedErrorInAMP4Ads)
			}

			r = rpb.Request{Html: tc.html, Config: rpb.Request_NONE, AllowedFormats: []rpb.Request_HtmlFormat{rpb.Request_AMP4EMAIL}}
			_, _, err = Process(&r)
			if (err != nil) != tc.expectedErrorInAMP4Email {
				t.Errorf("Process(AMP4Email) has error=%#v want=%t", err, tc.expectedErrorInAMP4Email)
			}
		})
	}
}

func TestBaseURL(t *testing.T) {
	docURL := "http://example.com/a/page.html"
	tcs := []struct {
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
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			// Remember the original function and reinstate after test
			orig := runTransformers
			defer func() { runTransformers = orig }()
			runTransformers = func(e *transformers.Context, fs []func(*transformers.Context) error) error {
				if e.BaseURL.String() != tc.expected {
					t.Errorf("setBaseURL(%s)=%s, want=%s", tc.base, e.BaseURL, tc.expected)
				}
				return nil
			}
			r := rpb.Request{Html: "<html amp><head>" + tc.base + "</head></html>", DocumentUrl: docURL, Config: rpb.Request_NONE}
			_, _, err := Process(&r)
			if err != nil {
				t.Fatalf("unexpected failure %v", err)
			}
		})
	}
}
