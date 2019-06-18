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
)

func pdata(imgURL string, media string) *transformers.PreloadData {
	imgURLObj, err := url.Parse(imgURL)
	if err != nil {
		return &transformers.PreloadData{}
	}

	return &transformers.PreloadData{URL: imgURLObj, Media: media, As: "image"}

}

func transformAndOutput(input string) (*transformers.Context, error) {
	inputDoc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return nil, err
	}
	inputDOM, err := amphtml.NewDOM(inputDoc)
	if err != nil {
		return nil, err
	}
	baseURL, _ := url.Parse("https://www.example.com")
	documentURL, _ := url.Parse("https://www.example.com/foo")

	context := &transformers.Context{
		DOM:         inputDOM,
		BaseURL:     baseURL,
		DocumentURL: documentURL,
		Preloads:    []transformers.PreloadData{},
	}
	transformers.PreloadImage(context)
	var output strings.Builder
	if err := html.Render(&output, inputDoc); err != nil {
		return nil, err
	}
	return context, nil
}

var testcaseInput = []struct {
	testcaseName    string
	html            string
	noPrefetchImage bool
	preloads        []*transformers.PreloadData
}{
	{"DataImageIgnored",
		`
<html amp>
<head>
</head>
<body>
  <amp-img src="data:image/png;somebase64encodeddata;" width="300" height="300">
  </amp-img>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"TinyImagesIgnored",
		`
<html amp>
<head>
</head>
<body>
<amp-img src="http://cdn.mycdn.com/foo.jpg" width="20" height="20">
  </amp-img>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"ImagesNoSrcIgnored",
		`
<html amp>
<head>
</head>
<body>
<amp-img src="" width="200" height="200"></amp-img>
</body>
</html>
        `,
		true, []*transformers.PreloadData{},
	},
	{"ImagesnodisplayLayoutIgnored",
		`
<html amp>
<head>
</head>
<body>
<amp-img src="https://cdn.com/foo.jpg" layout="nodisplay" width="200" height="200"></amp-img>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"PlaceholderForNoIframe",
		`
<html amp>
<head>
</head>
<body>
<!-- Placeholder image for nothing. -->
<amp-img placeholder layout="fill" width="200" height="200" src="https://cdn.com/trump.jpg"></amp-img>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"IframePlaceholderImage",
		`
<html amp>
<head>
</head>
<body>
<amp-iframe src="https://cdn.com/foo.html">
<amp-img placeholder layout="fill" width="200" height="200" src="https://cdn.com/obama.jpg" srcset="https://cdn.com/obama300.jpg 300w,https//cdn.com/obama600.jpg 600w,https://cdn.com/obama1000.jpg 1000w"></amp-img>
</amp-iframe>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://cdn.com/obama300.jpg", "(max-width: 300)"),
			pdata("https//cdn.com/obama600.jpg", "(min-width: 301) and (max-width: 600)"),
			pdata("https://cdn.com/obama1000.jpg", "(min-width: 601)")},
	},
	{"VideoPlaceholderImage",
		`
<html amp>
<head>
</head>
<body>
<amp-video layout="responsive" width="700" height="400" src="https://cdn.com/obama.mp4" poster="https://cdn.com/obama-video-img.jpg"></amp-video>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://cdn.com/obama-video-img.jpg", "")},
	},
	{"VideoIframePlaceholderImage",
		`
<html amp>
<head>
</head>
<body>
<amp-video-iframe layout="responsive" width="700" height="400" src="https://cdn.com/trump.mp4" poster="https://cdn.com/trump-video-img.jpg"></amp-video-iframe>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://cdn.com/trump-video-img.jpg", "")},
	},
	{"VideoIframeButNoPlaceholderImage",
		`
<html amp>
<head>
</head>
<body>
<amp-video-iframe layout="responsive" width="700" height="400" src="https://cdn.com/trump.mp4"></amp-video-iframe>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"VideoButNoPlaceholderImage",
		`
<html amp>
<head>
</head>
<body>
<amp-video layout="responsive" width="700" height="400" src="https://cdn.com/obama.mp4"></amp-video>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"NoImageDimensionsParentContainerOK",
		`
<html amp>
<head>
</head>
<body>
  <a href="/foo.html" layout="responsive" width="200" height="200">
  <amp-img src="https://cdn.com/foo-from-parent-container.jpg" layout="fill" srcset="https://cdn.com/foo-from-parent-container300.jpg 300w,https://cdn.com/foo-from-parent-container600.jpg 600w"></amp-img>
  </a>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://cdn.com/foo-from-parent-container300.jpg", "(max-width: 300)"),
			pdata("https://cdn.com/foo-from-parent-container600.jpg", "(min-width: 301)")},
	},
	{"NoImageDimensionsParentContainerSmall",
		`
<html amp>
<head>
</head>
<body>
  <a href="/foo.html" layout="responsive" width="20" height="20">
  <amp-img src="https://cdn.com/foo-from-parent-container.jpg" layout="fill"></amp-img>
  </a>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"ImageSmall",
		`
<html amp>
<head>
</head>
<body>
  <div>
  <amp-img src="https://cdn.com/foo-from-parent-container.jpg" layout="fill" width="149" height="149"></amp-img>
  </div>
</body
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"ImageNoSizeParentSmall",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="149" height="149">
  <amp-img src="https://cdn.com/foo-from-parent-container.jpg"></amp-img>
  </div>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"ImageNoSizeParentSizeOK",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="150" height="150">
  <amp-img src="https://cdn.com/foo-from-parent-container.jpg" srcset="https://cdn.com/foo-from-parent-container300.jpg 300w,https://cdn.com/foo-from-parent-container600.jpg 600w"></amp-img>
  </div>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://cdn.com/foo-from-parent-container300.jpg", "(max-width: 300)"),
			pdata("https://cdn.com/foo-from-parent-container600.jpg", "(min-width: 301)")},
	},
	{"FirstCandidateImageSelectedForPreloading",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="150" height="150">
  <amp-img src="https://cdn.com/first-candidate-image.jpg" srcset="https://cdn.com/second-candidate-image300.jpg 300w,https://cdn.com/second-candidate-image600.jpg 600w,https://cdn.com/second-candidate-image900.jpg 900w"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-candidate-image.jpg"></amp-img>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://cdn.com/second-candidate-image300.jpg", "(max-width: 300)"),
			pdata("https://cdn.com/second-candidate-image600.jpg", "(min-width: 301) and (max-width: 600)"),
			pdata("https://cdn.com/second-candidate-image900.jpg", "(min-width: 601)")},
	},
	{"MultipleImagesOnPage",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset="https://www.google.com/foo.png 300w,https://www.google.com/foo600.png 600w,https://www.google.com/foo1000.png 1000w"></amp-img>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://www.google.com/foo.png", "(max-width: 300)"),
			pdata("https://www.google.com/foo600.png", "(min-width: 301) and (max-width: 600)"),
			pdata("https://www.google.com/foo1000.png", "(min-width: 601)")},
	},
	{"SrcsetCommaSeparatorInPlaceOfWhitespace",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset="https://www.google.com/foo.png,300w,https://www.google.com/foo600.png 600w,https://www.google.com/foo1000.png 1000w"></amp-img>
  <!-- -------------------------^ -->
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"SrcsetWhitespaceSeparatorInPlaceOfComma",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset="https://www.google.com/foo.png 300w https://www.google.com/foo600.png 600w,https://www.google.com/foo1000.png 1000w"></amp-img>
  <!-- ------------------------------^ -->
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"EmptySrcset",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset=""></amp-img>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"SrcsetWithWhitespaceInUrls",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset="https://www.google.com/foo.png 300w,https://www.google.com/foo 600.png 600w,https://www.google.com/foo1000.png 1000w"></amp-img>
  <!-- ---------------------------------------------------------^
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"SrcsetWithWhitespaceInUrls",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset="https://www.google.com/foo.png 300w,https://www.google.com/foo 600.png 600w,https://www.google.com/foo1000.png 1000w"></amp-img>
</body>
</html>
	`,
		true, []*transformers.PreloadData{},
	},
	{"SrcsetWithWhitespaceInEncodedUrlsOK",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset="https://www.google.com/foo.png 300w,https://www.google.com/foo%20600.png 600w,https://www.google.com/foo1000.png 1000w"></amp-img>
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://www.google.com/foo.png", "(max-width: 300)"),
			pdata("https://www.google.com/foo%20600.png", "(min-width: 301) and (max-width: 600)"),
			pdata("https://www.google.com/foo1000.png", "(min-width: 601)")},
	},
	{"SrcsetWithImgsizeNotSorted",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
  <amp-img src="https://cdn.com/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150" srcset="https://www.google.com/foo.png 600w,https://www.google.com/foo%20300.png 300w,https://www.google.com/foo1000.png 1000w"></amp-img>
  <!-- -------^
</body>
</html>
	`,
		false, []*transformers.PreloadData{
			pdata("https://www.google.com/foo%20300.png", "(max-width: 300)"),
			pdata("https://www.google.com/foo.png", "(min-width: 301) and (max-width: 600)"),
			pdata("https://www.google.com/foo1000.png", "(min-width: 601)")},
	},
}

func TestAllCases(t *testing.T) {
	for _, tt := range testcaseInput {
		t.Run(tt.testcaseName, func(t *testing.T) {
			context, err := transformAndOutput(tt.html)
			if err != nil {
				t.Fatalf("Unexpected error %q", err)
			}
			if tt.noPrefetchImage {
				if len(context.Preloads) > 0 {
					t.Errorf("Prefetch link added in transformed HTML")
				}
			} else {
				if len(context.Preloads) == 0 {
					t.Errorf("Prefech link missing")
				}
				if len(context.Preloads) != len(tt.preloads) {
					t.Errorf("Number of preload images mismatch. %d vs. %d", len(context.Preloads), len(tt.preloads))
				}
				for i, p := range context.Preloads {
					if p.URL.String() != tt.preloads[i].URL.String() {
						t.Errorf("URL order wrong. %s vs. %s", p.URL.String(), tt.preloads[i].URL.String())
					}
					if p.Media != tt.preloads[i].Media {
						t.Errorf("Preload media attribute mismatch. %s vs. %s", p.Media, tt.preloads[i].Media)
					}
				}
			}
		})
	}
}
