package amppackager

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rtv = "1234"
const paddedRtv = "000000000001234"
const css = "css contents"

func TestRTVPoll(t *testing.T) {
	// Reset the cache
	RTVCache = new(RTVCacheStruct)
	ts := httptest.NewServer(fakeRTVServer())
	defer ts.Close()
	rtvHost = ts.URL

	assert.Equal(t, "", RTVCache.RTV)
	assert.Equal(t, "", RTVCache.CSS)
	//	rtvPoll()
	//	assert.Equal(t, paddedRtv, RTVCache.RTV)
	//	assert.Equal(t, css, RTVCache.CSS)
}

func TestRTVPollDieOnInit(t *testing.T) {
	// Reset the cache
	RTVCache = new(RTVCacheStruct)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))

	defer ts.Close()
	rtvHost = ts.URL
	// Remember the original die function and reinstate after this test.
	origDie := die
	defer func() { die = origDie }()
	var errors string
	die = func(format string, args ...interface{}) {
		errors = fmt.Sprintf(format, args)
	}
	assert.Equal(t, "", RTVCache.RTV)
	assert.Equal(t, "", RTVCache.CSS)
	rtvPoll()
	if errors == "" {
		t.Errorf("Expected die to be called, but wasn't!")
	}
}

func TestRTVPollWarn(t *testing.T) {
	// Initialize the cache to some values
	RTVCache.RTV = paddedRtv
	RTVCache.CSS = css
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))

	defer ts.Close()
	rtvHost = ts.URL

	assert.Equal(t, paddedRtv, RTVCache.RTV)
	assert.Equal(t, css, RTVCache.CSS)
	rtvPoll()
	// Values should not change, despite HTTP error.
	assert.Equal(t, paddedRtv, RTVCache.RTV)
	assert.Equal(t, css, RTVCache.CSS)
}

func fakeRTVServer() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v0/version.txt" {
			fmt.Fprint(w, rtv)
			return
		}
		if r.URL.Path == "/rtv/"+paddedRtv+"/v0.css" {
			fmt.Fprint(w, css)
			return
		}
	})
}
