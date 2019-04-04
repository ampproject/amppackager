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
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/ampproject/amppackager/transformer/internal/amphtml"
	"github.com/ampproject/amppackager/transformer/transformers"
	"golang.org/x/net/html"
)

func transformAndOutput(input string) (string, error) {
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

	transformers.PreloadImage(&transformers.Context{
		DOM:         inputDOM,
		BaseURL:     baseURL,
		DocumentURL: documentURL,
	})
	var output strings.Builder
	if err := html.Render(&output, inputDoc); err != nil {
		return "", err
	}
	return output.String(), nil
}

var testcaseInput = []struct {
	testcaseName    string
	html            string
	noPrefetchImage bool
	prefetchLink    string
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
		true, "",
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
		true, "",
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
		true, "",
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
		true, "",
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
		true, "",
	},
	{"IframePlaceholderImage",
		`
<html amp>
<head>
</head>
<body>
<amp-iframe src="https://cdn.com/foo.html">
<amp-img placeholder layout="fill" width="200" height="200" src="https://cdn.com/obama.jpg"></amp-img>
</amp-iframe>
</body>
</html>
	`,
		false, "https://cdn.com/obama.jpg",
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
		false, "https://cdn.com/obama-video-img.jpg",
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
		false, "https://cdn.com/trump-video-img.jpg",
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
		true, "",
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
		true, "",
	},
	{"NoImageDimensionsParentContainerOK",
		`
<html amp>
<head>
</head>
<body>
  <a href="/foo.html" layout="responsive" width="200" height="200">
  <amp-img src="https://cdn.com/foo-from-parent-container.jpg" layout="fill"></amp-img>
  </a>
</body>
</html>
	`,
		false, "https://cdn.com/foo-from-parent-container.jpg",
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
		true, "",
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
		true, "",
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
		true, "",
	},
	{"ImageNoSizeParentSizeOK",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="150" height="150">
  <amp-img src="https://cdn.com/foo-from-parent-container.jpg"></amp-img>
  </div>
</body>
</html>
	`,
		false, "https://cdn.com/foo-from-parent-container.jpg",
	},
	{"FirstCandidateImageSelectedForPreloading",
		`
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="150" height="150">
  <amp-img src="https://cdn.com/first-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="https://cdn.com/second-candidate-image.jpg"></amp-img>
</body>
</html>
	`,
		false, "https://cdn.com/first-candidate-image.jpg",
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
  <amp-img src="https://cdn.com/third-qualified-candidate-image.jpg" width="150" height="150"></amp-img>
</body>
</html>
	`,
		false, "https://cdn.com/third-qualified-candidate-image.jpg",
	},
}

func TestAllCases(t *testing.T) {
	for _, tt := range testcaseInput {
		t.Run(tt.testcaseName, func(t *testing.T) {
			output, err := transformAndOutput(tt.html)
			if err != nil {
				t.Fatalf("Unexpected error %q", err)
			}
			if tt.noPrefetchImage {
				if strings.Contains(output, "<link rel=\"prefetch\"") {
					t.Errorf("Prefetch link added in transformed HTML: %s", output)
				}
			} else if !strings.Contains(output, fmt.Sprintf("<link rel=\"prefetch\" href=\"%s\"/>", tt.prefetchLink)) {
				t.Errorf("Expected Tag: %s. Not found. Transformed HTML: %s", tt.prefetchLink, output)
			}
		})
	}
}
