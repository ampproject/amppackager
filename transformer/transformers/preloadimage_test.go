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
        //tt "github.com/ampproject/amppackager/transformer/internal/testing"
        "github.com/ampproject/amppackager/transformer/transformers"
        "golang.org/x/net/html"
)

func transformAndPrint(input string) (string, error) {
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
		DOM: inputDOM,
		BaseURL: baseURL,
		DocumentURL: documentURL,
	})
	var output strings.Builder
	if err := html.Render(&output, inputDoc); err != nil {
		return "", err
	}
	return output.String(), nil
}

func testNoPrefetchImagesAdded(input string, t *testing.T) {
        output, err := transformAndPrint(input)
	if err != nil {
		t.Errorf("Unexpected error %q", err)
	}

	if strings.Contains(output, "<link rel=\"prefetch\"") {
		t.Errorf("inline data:image added as prefetch image link. Transformed HTML: %s", output)
	}
}

func testPrefetchImagesAdded(input string, prefetchTag string, t *testing.T) {
        output, err := transformAndPrint(input)
	if err != nil {
		t.Errorf("Unexpected error %q", err)
	}

	if !strings.Contains(output, prefetchTag) {
		t.Errorf("Expected Tag: %s. Not found. Transformed HTML: %s", prefetchTag, output)
	}
}

func testNoSuchPrefetchImagesAdded(input string, unexpectedPrefetchTag string, t *testing.T) {
        output, err := transformAndPrint(input)
	if err != nil {
		t.Errorf("Unexpected error %q", err)
	}

	if strings.Contains(output, unexpectedPrefetchTag) {
		t.Errorf("UnExpected Tag: %s. Transformed HTML: %s", unexpectedPrefetchTag, output)
	}
}

func TestDataImageIgnored(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <amp-img src="data:image/png;somebase64encodeddata;" width="300" height="300">
  </amp-img>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestTinyImagesIgnored(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
<amp-img src="http://cdn.mycdn.com/foo.jpg" width="20" height="20">
  </amp-img>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestImagesNoSrcIgnored(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
<amp-img src="" width="200" height="200"></amp-img>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestImagesnodisplayLayoutIgnored(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
<amp-img src="/foo.jpg" layout="nodisplay" width="200" height="200"></amp-img>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestPlaceholderForNoIframe(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
<!-- Placeholder image for nothing. -->
  <amp-img placeholder layout="fill" width="200" height="200" src="/trump.jpg"></amp-img>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestIframePlaceholderImage(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
<amp-iframe src="foo.html">
  <amp-img placeholder layout="fill" width="200" height="200" src="/obama.jpg"></amp-img>
</amp-iframe>
</body>
</html>
`
        testPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/obama.jpg\"/>", t)
}

func TestVideoPlaceholderImage(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <amp-video layout="responsive" width="700" height="400" src="/obama.mp4" poster="/obama-video-img.jpg"></amp-video>
</body>
</html>
`
        testPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/obama-video-img.jpg\"/>", t)
}

func TestVideoIframePlaceholderImage(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <amp-video-iframe layout="responsive" width="700" height="400" src="/trump.mp4" poster="/trump-video-img.jpg"></amp-video-iframe>
</body>
</html>
`
        testPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/trump-video-img.jpg\"/>", t)
}

func TestVideoIframeButNoPlaceholderImage(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <amp-video-iframe layout="responsive" width="700" height="400" src="/trump.mp4"></amp-video-iframe>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestVideoButNoPlaceholderImage(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <amp-video layout="responsive" width="700" height="400" src="/obama.mp4"></amp-video>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestNoImageDimensionsParentContainerOK(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <a href="/foo.html" layout="responsive" width="200" height="200">
    <amp-img src="/foo-from-parent-container.jpg" layout="fill"></amp-img>
  </a>
</body>
</html>
`
        testPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/foo-from-parent-container.jpg\"/>", t)
}

func TestNoImageDimensionsParentContainerSmall(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <a href="/foo.html" layout="responsive" width="20" height="20">
    <amp-img src="/foo-from-parent-container.jpg" layout="fill"></amp-img>
  </a>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestImageTooSmall(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <div>
    <amp-img src="/foo-from-parent-container.jpg" layout="fill" width="149" height="149"></amp-img>
  </div>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestImageNoSizeParentTooSmall(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="149" height="149">
    <amp-img src="/foo-from-parent-container.jpg"></amp-img>
  </div>
</body>
</html>
`
        testNoPrefetchImagesAdded(input, t)
}

func TestImageNoSizeParentSizeOK(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="150" height="150">
    <amp-img src="/foo-from-parent-container.jpg"></amp-img>
  </div>
</body>
</html>
`
        testPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/foo-from-parent-container.jpg\"/>", t)
}

func TestFirstCandidateImageSelectedForPreloading(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="150" height="150">
    <amp-img src="/first-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="/second-candidate-image.jpg"></amp-img>
</body>
</html>
`
        testPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/first-candidate-image.jpg\"/>", t)
        testNoSuchPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/second-candidate-image.jpg\"/>", t)
}

func TestMultipleImagesOnPage(t *testing.T) {
	input := `
<html amp>
<head>
</head>
<body>
  <div layout="responsive" width="50" height="50">
    <amp-img src="/first-unqualified-candidate-image.jpg"></amp-img>
  </div>
  <div>foo</div>
  <amp-img src="/second-unqualified-candidate-image.jpg"></amp-img>
  <amp-img src="/third-qualified-candidate-image.jpg" width="150" height="150"></amp-img>
</body>
</html>
`
        testPrefetchImagesAdded(input, "<link rel=\"prefetch\" href=\"/third-qualified-candidate-image.jpg\"/>", t)
}

