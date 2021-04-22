// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transformers_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
	"github.com/kylelemons/godebug/diff"
)

type TestCase = struct {
	testcaseName string
	input        string
	expected     string
}

func transformAndOutput(input string, version int64) (string, error) {
	inputDoc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return "", err
	}
	inputDOM, err := amphtml.NewDOM(inputDoc)
	if err != nil {
		return "", err
	}
	baseURL, _ := url.Parse("https://www.example.com")
	documentURL, _ := url.Parse("https://www.example.com/foo")

	context := &transformers.Context{
		DOM:         inputDOM,
		BaseURL:     baseURL,
		DocumentURL: documentURL,
		Version:     version,
	}
	transformers.PreloadImage(context)
	var output strings.Builder
	if err := html.Render(&output, inputDoc); err != nil {
		return "", err
	}
	return output.String(), nil
}

func testCases(t *testing.T, cases []TestCase, version int64) {
	for _, tc := range cases {
		t.Run(tc.testcaseName, func(t *testing.T) {
			output, err := transformAndOutput(strings.TrimSpace(tc.input), version)
			if err != nil {
				t.Fatalf("Unexpected error %q", err)
			}
			if diff := diff.Diff(strings.TrimSpace(tc.expected), output); diff != "" {
				t.Errorf("PreloadImage transformer produced unexpected output:\n%s", diff)
			}
		})
	}
}

var testcaseInferSize = []TestCase{
	{
		"inferred-size: Has hero image.",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Dimensions too small",
		`<html><head></head><body><amp-img height="100" src="https://example.com/foo.png" width="100"></amp-img></body></html>`,
		`<html><head></head><body><amp-img height="100" src="https://example.com/foo.png" width="100"></amp-img></body></html>`,
	},
	{
		"inferred-size: Crossorigin attribute.",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" crossorigin="anonymous"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" crossorigin="anonymous"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" crossorigin="anonymous" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" crossorigin="anonymous" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Referrerpolicy attribute.",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" referrerpolicy="origin"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" referrerpolicy="origin"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" referrerpolicy="origin" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" referrerpolicy="origin" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Srcset attribute.",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Srcset without src.",
		`<html><head></head><body><amp-img width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-img width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Sizes",
		`<html><head></head><body><amp-img width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" sizes="100vw"></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w" imagesizes="100vw"/></head><body><amp-img width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" sizes="100vw" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" sizes="100vw"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Several images",
		`<html><head></head><body><amp-img height="100" src="https://example.com/bar.png" width="100"></amp-img><amp-img width="500" height="400" src="https://example.com/foo.png"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img height="100" src="https://example.com/bar.png" width="100"></amp-img><amp-img width="500" height="400" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Several images none qualifies. (all tiny)",
		`<html><head></head><body><amp-img height="100" src="https://example-com.cdn.ampproject.org/bar.png" width="100"></amp-img><amp-img width="100" height="100" src="https://example.com/foo.png"></body></html>`,
		`<html><head></head><body><amp-img height="100" src="https://example-com.cdn.ampproject.org/bar.png" width="100"></amp-img><amp-img width="100" height="100" src="https://example.com/foo.png"></amp-img></body></html>`,
	},
	{
		"inferred-size: Iframe placeholder",
		`<html><head></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/bar.png"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/bar.png"/></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/bar.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/bar.png"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"inferred-size: Iframe placeholder crossorigin",
		`<html><head></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/bar.png" crossorigin="anonymous"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/bar.png" crossorigin="anonymous"/></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/bar.png" crossorigin="anonymous" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" crossorigin="anonymous" src="https://example.com/bar.png"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"inferred-size: Iframe placeholder referrerpolicy",
		`<html><head></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/bar.png" referrerpolicy="origin"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/bar.png" referrerpolicy="origin"/></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/bar.png" referrerpolicy="origin" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" referrerpolicy="origin" src="https://example.com/bar.png"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"inferred-size: Iframe placeholder srcset",
		`<html><head></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"inferred-size: No placeholder image",
		`<html><head></head><body><amp-iframe src="/foo.html"></amp-iframe></body></html>`,
		`<html><head></head><body><amp-iframe src="/foo.html"></amp-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video placeholder",
		`<html><head></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/foo.png"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video placeholder crossorigin",
		`<html><head></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/foo.png" crossorigin="anonymous"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" crossorigin="anonymous"/></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/foo.png" crossorigin="anonymous" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" crossorigin="anonymous" src="https://example.com/foo.png"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video placeholder referrerpolicy",
		`<html><head></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/foo.png" referrerpolicy="origin"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" referrerpolicy="origin"/></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/foo.png" referrerpolicy="origin" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" referrerpolicy="origin" src="https://example.com/foo.png"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video placeholder srcset",
		`<html><head></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></amp-img></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-video-iframe height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video poster",
		`<html><head></head><body><amp-video-iframe height="500" width="500" src="/foo.html" poster="https://example.com/foo.png"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe height="500" width="500" src="/foo.html" poster="https://example.com/foo.png"></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video poster crossorigin",
		`<html><head></head><body><amp-video-iframe height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" crossorigin="anonymous"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" crossorigin="anonymous"></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video poster referrerpolicy",
		`<html><head></head><body><amp-video-iframe height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" referrerpolicy="origin"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" referrerpolicy="origin"></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: No placeholder image",
		`<html><head></head><body><amp-video-iframe src="/foo.html"></amp-video-iframe></body></html>`,
		`<html><head></head><body><amp-video-iframe src="/foo.html"></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: Video poster",
		`<html><head></head><body><amp-video poster="https://example.com/foo.png" width="400" height="400"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video poster="https://example.com/foo.png" width="400" height="400"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"inferred-size: Video poster crossorigin",
		`<html><head></head><body><amp-video poster="https://example.com/foo.png" width="400" height="400" crossorigin="anonymous"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video poster="https://example.com/foo.png" width="400" height="400" crossorigin="anonymous"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"inferred-size: Video poster referrerpolicy",
		`<html><head></head><body><amp-video poster="https://example.com/foo.png" width="400" height="400" referrerpolicy="origin"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video poster="https://example.com/foo.png" width="400" height="400" referrerpolicy="origin"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"inferred-size: amp-video with missing poster.",
		`<html><head></head><body><amp-video width="400" height="400"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head></head><body><amp-video width="400" height="400"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"inferred-size: No display layout",
		`<html><head></head><body><amp-img height="500" src="https://example.com/foo.png" width="500" layout="nodisplay"></amp-img></body></html>`,
		`<html><head></head><body><amp-img height="500" src="https://example.com/foo.png" width="500" layout="nodisplay"></amp-img></body></html>`,
	},
	{
		"inferred-size: Same as above with nodisplay layout removed.",
		`<html><head></head><body><amp-img height="500" src="https://example.com/foo.png" width="500"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img height="500" src="https://example.com/foo.png" width="500" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Invalid protocol",
		`<html><head></head><body><amp-img width="500" height="400" src="ftp://example.com/ftp.png"></body></html>`,
		`<html><head></head><body><amp-img width="500" height="400" src="ftp://example.com/ftp.png"></amp-img></body></html>`,
	},
	{
		"inferred-size: Srcset validity. Empty srcset.",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset=""></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset="" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset=""/></amp-img></body></html>`,
	},
	{
		"inferred-size: Invalid srcset",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset="foo bar baz"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset="foo bar baz" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset="foo bar baz"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Invalid srcset duplicates.",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset="foo 10w, bar 10w, baz 100w"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" srcset="foo 10w, bar 10w, baz 100w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset="foo 10w, bar 10w, baz 100w"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Hero image dimensions from parent container.",
		`<html><head></head><body><div width="500" height="500"><amp-img layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><div width="500" height="500"><amp-img layout="fill" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></div></body></html>`,
	},
	{
		"inferred-size: Hero image dimesions from parent container, too small.",
		`<html><head></head><body><div width="50" height="50"><amp-img layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head></head><body><div width="50" height="50"><amp-img layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
	},
	{
		"inferred-size: No dimensions in parent containers.",
		`<html><head></head><body><div><amp-img layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head></head><body><div><amp-img layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
	},
	{
		"inferred-size: No dimension from parent because layout is not responsive or fill",
		`<html><head></head><body><div><amp-img src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head></head><body><div><amp-img src="https://example.com/foo.png"></amp-img></div></body></html>`,
	},
}

var testcaseDataHero = []TestCase{
	{
		"data-hero",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"data-hero: Allows multiple heros",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png"></amp-img><amp-img data-hero width="500" height="400" src="https://example.com/bar.png"></amp-img><amp-img data-hero width="500" height="400" src="https://example.com/baz.png"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/><link rel="preload" as="image" href="https://example.com/bar.png"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img><amp-img data-hero="" width="500" height="400" src="https://example.com/bar.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/bar.png"/></amp-img><amp-img data-hero="" width="500" height="400" src="https://example.com/baz.png"></amp-img></body></html>`,
	},
	{
		"data-hero: Prioritizes data-hero",
		`<html><head></head><body><amp-img width="100" height="100" src="https://example.com/foo.png"></amp-img><amp-img data-hero width="500" height="400" src="https://example.com/bar.png"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/bar.png"/></head><body><amp-img width="100" height="100" src="https://example.com/foo.png"></amp-img><amp-img data-hero="" width="500" height="400" src="https://example.com/bar.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/bar.png"/></amp-img></body></html>`,
	},
	{
		"data-hero: Prevents size-inferred hero",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png"></amp-img><amp-img width="500" height="400" src="https://example.com/bar.png"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img><amp-img width="500" height="400" src="https://example.com/bar.png"></amp-img></body></html>`,
	},
	{
		"data-hero: Dimensions too small",
		`<html><head></head><body><amp-img data-hero height="100" src="https://example.com/foo.png" width="100"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" height="100" src="https://example.com/foo.png" width="100" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"data-hero: Crossorigin attribute.",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png" crossorigin="anonymous"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" crossorigin="anonymous"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" crossorigin="anonymous" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" crossorigin="anonymous" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"data-hero: Referrerpolicy referrerpolicy.",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png" referrerpolicy="origin"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" referrerpolicy="origin"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" referrerpolicy="origin" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" referrerpolicy="origin" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"data-hero: Srcset attribute.",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Srcset without src.",
		`<html><head></head><body><amp-img data-hero width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-img data-hero="" width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></body></html>`,
	},
	{
		"inferred-size: Sizes.",
		`<html><head></head><body><amp-img data-hero width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" sizes="100vw"></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w" imagesizes="100vw"/></head><body><amp-img data-hero="" width="500" height="400" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" sizes="100vw" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" sizes="100vw"/></amp-img></body></html>`,
	},
	{
		"data-hero: Iframe placeholder",
		`<html><head></head><body><amp-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/bar.png"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/bar.png"/></head><body><amp-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/bar.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/bar.png"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"data-hero: Iframe placeholder crossorigin",
		`<html><head></head><body><amp-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/bar.png" crossorigin="anonymous"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/bar.png" crossorigin="anonymous"/></head><body><amp-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/bar.png" crossorigin="anonymous" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" crossorigin="anonymous" src="https://example.com/bar.png"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"data-hero: Iframe placeholder referrerpolicy",
		`<html><head></head><body><amp-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/bar.png" referrerpolicy="origin"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/bar.png" referrerpolicy="origin"/></head><body><amp-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/bar.png" referrerpolicy="origin" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" referrerpolicy="origin" src="https://example.com/bar.png"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"inferred-size: Iframe placeholder srcset",
		`<html><head></head><body><amp-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></amp-img></amp-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></amp-iframe></body></html>`,
	},
	{
		"data-hero: No placeholder image",
		`<html><head></head><body><amp-iframe data-hero src="/foo.html"></amp-iframe></body></html>`,
		`<html><head></head><body><amp-iframe data-hero="" src="/foo.html"></amp-iframe></body></html>`,
	},
	{
		"data-hero: iframe video placeholder",
		`<html><head></head><body><amp-video-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/foo.png"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"data-hero: iframe video placeholder crossorigin",
		`<html><head></head><body><amp-video-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/foo.png" crossorigin="anonymous"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" crossorigin="anonymous"/></head><body><amp-video-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/foo.png" crossorigin="anonymous" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" crossorigin="anonymous" src="https://example.com/foo.png"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"data-hero: iframe video placeholder referrerpolicy",
		`<html><head></head><body><amp-video-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" src="https://example.com/foo.png" referrerpolicy="origin"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png" referrerpolicy="origin"/></head><body><amp-video-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" src="https://example.com/foo.png" referrerpolicy="origin" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" referrerpolicy="origin" src="https://example.com/foo.png"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"inferred-size: iframe video placeholder srcset",
		`<html><head></head><body><amp-video-iframe data-hero height="500" width="500" src="/foo.html"><amp-img placeholder layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"></amp-img></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" imagesrcset="https://example.com/foolarge.png 1200w, https://example.com/foomedium.png 800w"/></head><body><amp-video-iframe data-hero="" height="500" width="500" src="/foo.html"><amp-img placeholder="" layout="fill" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" srcset="https://example.com/foomedium.png 800w, https://example.com/foolarge.png 1200w"/></amp-img></amp-video-iframe></body></html>`,
	},
	{
		"data-hero: iframe video poster",
		`<html><head></head><body><amp-video-iframe data-hero height="500" width="500" src="/foo.html" poster="https://example.com/foo.png"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe data-hero="" height="500" width="500" src="/foo.html" poster="https://example.com/foo.png"></amp-video-iframe></body></html>`,
	},
	{
		"data-hero: iframe video poster crossorigin",
		`<html><head></head><body><amp-video-iframe data-hero height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" crossorigin="anonymous"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe data-hero="" height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" crossorigin="anonymous"></amp-video-iframe></body></html>`,
	},
	{
		"data-hero: iframe video poster referrerpolicy",
		`<html><head></head><body><amp-video-iframe data-hero height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" referrerpolicy="origin"></amp-video-iframe></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video-iframe data-hero="" height="500" width="500" src="/foo.html" poster="https://example.com/foo.png" referrerpolicy="origin"></amp-video-iframe></body></html>`,
	},
	{
		"data-hero: No placeholder image",
		`<html><head></head><body><amp-video-iframe data-hero src="/foo.html"></amp-video-iframe></body></html>`,
		`<html><head></head><body><amp-video-iframe data-hero="" src="/foo.html"></amp-video-iframe></body></html>`,
	},
	{
		"data-hero: Video poster",
		`<html><head></head><body><amp-video data-hero poster="https://example.com/foo.png" width="400" height="400"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video data-hero="" poster="https://example.com/foo.png" width="400" height="400"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"data-hero: Video poster crossorigin",
		`<html><head></head><body><amp-video data-hero poster="https://example.com/foo.png" width="400" height="400" crossorigin="anonymous"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video data-hero="" poster="https://example.com/foo.png" width="400" height="400" crossorigin="anonymous"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"data-hero: Video poster referrerpolicy",
		`<html><head></head><body><amp-video data-hero poster="https://example.com/foo.png" width="400" height="400" referrerpolicy="origin"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-video data-hero="" poster="https://example.com/foo.png" width="400" height="400" referrerpolicy="origin"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"data-hero: amp-video with missing poster.",
		`<html><head></head><body><amp-video data-hero width="400" height="400"><source src="foo.mp4" /></amp-video></body></html>`,
		`<html><head></head><body><amp-video data-hero="" width="400" height="400"><source src="foo.mp4"/></amp-video></body></html>`,
	},
	{
		"data-hero: No display layout",
		`<html><head></head><body><amp-img data-hero height="500" src="https://example.com/foo.png" width="500" layout="nodisplay"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" height="500" src="https://example.com/foo.png" width="500" layout="nodisplay" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"data-hero: Same as above with nodisplay layout removed.",
		`<html><head></head><body><amp-img data-hero height="500" src="https://example.com/foo.png" width="500"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" height="500" src="https://example.com/foo.png" width="500" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></body></html>`,
	},
	{
		"data-hero: Invalid protocol",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="ftp://example.com/ftp.png"></body></html>`,
		`<html><head></head><body><amp-img data-hero="" width="500" height="400" src="ftp://example.com/ftp.png"></amp-img></body></html>`,
	},
	{
		"data-hero: Srcset validity. Empty srcset.",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png" srcset=""></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" srcset="" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset=""/></amp-img></body></html>`,
	},
	{
		"data-hero: Invalid srcset",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png" srcset="foo bar baz"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" srcset="foo bar baz" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset="foo bar baz"/></amp-img></body></html>`,
	},
	{
		"data-hero: Invalid srcset duplicates.",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png" srcset="foo 10w, bar 10w, baz 100w"></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" srcset="foo 10w, bar 10w, baz 100w" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" srcset="foo 10w, bar 10w, baz 100w"/></amp-img></body></html>`,
	},
	{
		"data-hero: Hero image dimensions from parent container.",
		`<html><head></head><body><div width="500" height="500"><amp-img data-hero layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><div width="500" height="500"><amp-img data-hero="" layout="fill" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></div></body></html>`,
	},
	{
		"data-hero: Hero image dimesions from parent container, too small.",
		`<html><head></head><body><div width="50" height="50"><amp-img data-hero layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><div width="50" height="50"><amp-img data-hero="" layout="fill" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></div></body></html>`,
	},
	{
		"data-hero: No dimensions in parent containers.",
		`<html><head></head><body><div><amp-img data-hero layout="fill" src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><div><amp-img data-hero="" layout="fill" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></div></body></html>`,
	},
	{
		"data-hero: No dimension from parent because layout is not responsive or fill",
		`<html><head></head><body><div><amp-img data-hero src="https://example.com/foo.png"></amp-img></div></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><div><amp-img data-hero="" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img></div></body></html>`,
	},
}

var testLazyLoadImg = []TestCase{
	{
		"data-hero leftover",
		`<html><head></head><body><amp-img data-hero width="500" height="400" src="https://example.com/foo.png"></amp-img><amp-img width="500" height="400" src="https://example.com/bar.png"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img data-hero="" width="500" height="400" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img><amp-img width="500" height="400" src="https://example.com/bar.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/bar.png" loading="lazy"/></amp-img></body></html>`,
	},
	{
		"inferred-size leftover",
		`<html><head></head><body><amp-img width="500" height="400" src="https://example.com/foo.png"></amp-img><amp-img width="100" height="100" src="https://example.com/bar.png"></amp-img></body></html>`,
		`<html><head><link rel="preload" as="image" href="https://example.com/foo.png"/></head><body><amp-img width="500" height="400" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png"/></amp-img><amp-img width="100" height="100" src="https://example.com/bar.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/bar.png" loading="lazy"/></amp-img></body></html>`,
	},
	{
		"no transformed images",
		`<html><head></head><body><amp-img width="100" height="100" src="https://example.com/foo.png"></amp-img></body></html>`,
		`<html><head></head><body><amp-img width="100" height="100" src="https://example.com/foo.png" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" src="https://example.com/foo.png" loading="lazy"/></amp-img></body></html>`,
	},
	{
		"crossorigin propagates",
		`<html><head></head><body><amp-img width="100" height="100" src="https://example.com/foo.png" crossorigin="anonymous"></amp-img></body></html>`,
		`<html><head></head><body><amp-img width="100" height="100" src="https://example.com/foo.png" crossorigin="anonymous" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" crossorigin="anonymous" src="https://example.com/foo.png" loading="lazy"/></amp-img></body></html>`,
	},
	{
		"referrerpolicy propagates",
		`<html><head></head><body><amp-img width="100" height="100" src="https://example.com/foo.png" referrerpolicy="origin"></amp-img></body></html>`,
		`<html><head></head><body><amp-img width="100" height="100" src="https://example.com/foo.png" referrerpolicy="origin" i-amphtml-ssr=""><img class="i-amphtml-fill-content i-amphtml-replaced-content" decoding="async" referrerpolicy="origin" src="https://example.com/foo.png" loading="lazy"/></amp-img></body></html>`,
	},
}

func TestInferSizeCases(t *testing.T) {
	testCases(t, testcaseInferSize, 0)
}

func TestDataHeroCases(t *testing.T) {
	testCases(t, testcaseDataHero, 0)
}

func TestLazyLoadCases(t *testing.T) {
	testCases(t, testLazyLoadImg, 5)
}
