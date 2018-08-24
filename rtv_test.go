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
	RTVCache = new(rtvCacheStruct)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v0/version.txt" {
			fmt.Fprint(w, rtv)
			return
		}
		if r.URL.Path == "/rtv/"+paddedRtv+"/v0.css" {
			fmt.Fprint(w, css)
			return
		}
	}))
	defer ts.Close()
	rtvHost = ts.URL

	assert.Equal(t, "", RTVCache.RTV)
	assert.Equal(t, "", RTVCache.CSS)
	rtvPoll()
	assert.Equal(t, paddedRtv, RTVCache.RTV)
	assert.Equal(t, css, RTVCache.CSS)
}

func TestRTVPollSameValue(t *testing.T) {
	// Reset the cache
	RTVCache = new(rtvCacheStruct)
	var rtvCalls, cssCalls int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v0/version.txt" {
			rtvCalls++
			fmt.Fprint(w, rtv)
			return
		}
		if r.URL.Path == "/rtv/"+paddedRtv+"/v0.css" {
			cssCalls++
			fmt.Fprint(w, css)
			return
		}
	}))
	defer ts.Close()
	rtvHost = ts.URL

	assert.Equal(t, "", RTVCache.RTV)
	assert.Equal(t, "", RTVCache.CSS)
	rtvPoll()
	assert.Equal(t, paddedRtv, RTVCache.RTV)
	assert.Equal(t, css, RTVCache.CSS)
	rtvPoll()
	assert.Equal(t, paddedRtv, RTVCache.RTV)
	assert.Equal(t, css, RTVCache.CSS)
	assert.Equal(t, 2, rtvCalls)
	assert.Equal(t, 1, cssCalls) // css should only be requested once since rtv value didn't change.
}

func TestRTVPollDieOnInit(t *testing.T) {
	// Reset the cache
	RTVCache = new(rtvCacheStruct)
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

func TestRTVPollSkipsCSSOnError(t *testing.T) {
	// Initialize the cache to some values
	RTVCache.RTV = paddedRtv
	RTVCache.CSS = css
	var rtvCalled, cssCalled bool
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v0/version.txt" {
			rtvCalled = true
			w.WriteHeader(500)
		}
		if r.URL.Path == "/rtv/"+paddedRtv+"/v0.css" {
			cssCalled = true
			fmt.Fprint(w, css)
			return
		}
	}))

	defer ts.Close()
	rtvHost = ts.URL

	assert.Equal(t, paddedRtv, RTVCache.RTV)
	assert.Equal(t, css, RTVCache.CSS)
	rtvPoll()
	// Values should not change, despite HTTP error.
	assert.Equal(t, paddedRtv, RTVCache.RTV)
	assert.Equal(t, css, RTVCache.CSS)
	// Verify css was not called
	assert.True(t, rtvCalled)
	assert.False(t, cssCalled, "css was fetched when it shouldn't have been!")
}

func TestRTVPollRollback(t *testing.T) {
	// Initialize the cache to some values
	RTVCache.RTV = paddedRtv
	RTVCache.CSS = css
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v0/version.txt" {
			fmt.Fprint(w, "9999") // a new rtv value
		}
		if r.URL.Path == "/rtv/000000000009999/v0.css" {
			w.WriteHeader(500)
			return
		}
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

func TestStartCronDieOnInit(t *testing.T) {
	// Reset the cache
	RTVCache = new(rtvCacheStruct)
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
	StartCron()
	if errors == "" {
		t.Errorf("Expected die to be called, but wasn't!")
	}
}
