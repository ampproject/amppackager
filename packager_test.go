package amppackager

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pkg/errors"
)

func newPackager(t *testing.T, urlSets []URLSet) *Packager {
	handler, err := NewPackager(cert, key, "https://example.com/", urlSets)
	if err != nil {
		t.Fatal(errors.WithStack(err))
	}
	return handler
}

func stringPtr(s string) *string { return &s }

func TestSimple(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("yum yum yum"))
	}))
	url, _ := url.Parse(server.URL)
	defer server.Close()
	urlSets := []URLSet{URLSet{
		SamePath: true,
		Fetch:    URLPattern{[]string{"http"}, url.Host, stringPtr("/amp/.*"), []string{}, stringPtr(""), false},
		Sign:     URLPattern{[]string{"https"}, "example.com", stringPtr("/amp/.*"), []string{}, stringPtr(""), false}}}
	resp := get(t, newPackager(t, urlSets), `/priv/doc?fetch=http%3A%2F%2F`+url.Host+`%2Famp%2Fsecret-life-of-pine-trees.html&sign=https%3A%2F%2Fexample.com%2Famp%2Fsecret-life-of-pine-trees.html`)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status: %#v", resp)
	}
	_, _ = ioutil.ReadAll(resp.Body)
	// TODO(twifkak): Test the body somehow.
}

// TODO(twifkak): Write lots more tests.
// TODO(twifkak): Fuzz-test.
