package amppackager

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	rtv       = "1234"
	paddedRtv = "000000000001234"
	css       = "css contents"
)

type fakeServer struct {
	rtvCalls, cssCalls     int
	rtvHandler, cssHandler func(*fakeServer, http.ResponseWriter, *http.Request)
}

type RTVTestSuite struct {
	suite.Suite
	f  *fakeServer
	ts *httptest.Server
}

func defaultRTVHandler(f *fakeServer, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rtv)
}

func defaultCSSHandler(f *fakeServer, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, css)
}

func (f *fakeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/v0/version.txt" {
		f.rtvCalls++
		f.rtvHandler(f, w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/rtv/") {
		f.cssCalls++
		f.cssHandler(f, w, r)
	}
}

// Before the entire suite, start the test server
func (t *RTVTestSuite) SetupSuite() {
	t.f = &fakeServer{}
	t.ts = httptest.NewServer(t.f)
	rtvHost = t.ts.URL
}

// Before every test, reset counters and reset default handlers.
func (t *RTVTestSuite) SetupTest() {
	t.f.cssCalls = 0
	t.f.cssHandler = defaultCSSHandler
	t.f.rtvCalls = 0
	t.f.rtvHandler = defaultRTVHandler
}

// After the suite, tear down test server.
func (t *RTVTestSuite) TearDownSuite() {
	t.ts.Close()
}

func TestRTVTestSuite(t *testing.T) {
	suite.Run(t, new(RTVTestSuite))
}

func (t *RTVTestSuite) TestNewRTV() {
	r, err := NewRTV()
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), paddedRtv, r.GetRTV())
	assert.Equal(t.T(), css, r.GetCSS())
}

func (t *RTVTestSuite) TestRTVPollSameValue() {
	r, err := NewRTV()
	assert.NoError(t.T(), err)

	err = r.poll()
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), paddedRtv, r.GetRTV())
	assert.Equal(t.T(), css, r.GetCSS())
	assert.Equal(t.T(), 2, t.f.rtvCalls)
	assert.Equal(t.T(), 1, t.f.cssCalls) // css should only be requested once since rtv value didn't change.
}

func (t *RTVTestSuite) TestRTVPollErrorOnInit() {
	t.f.rtvHandler = func(f *fakeServer, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}

	_, err := NewRTV()
	assert.Error(t.T(), err)
}

func (t *RTVTestSuite) TestRTVPollSkipsCSSOnError() {
	r, err := NewRTV()
	assert.NoError(t.T(), err)

	// Set up the next call to error out.
	t.f.rtvHandler = func(f *fakeServer, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}

	err = r.poll()
	assert.Error(t.T(), err)
	// Values should not change, despite HTTP error.
	assert.Equal(t.T(), paddedRtv, r.GetRTV())
	assert.Equal(t.T(), css, r.GetCSS())
	// Verify css was not called
	assert.Equal(t.T(), 2, t.f.rtvCalls)
	assert.Equal(t.T(), 1, t.f.cssCalls, "css was fetched when it shouldn't have been!")
}

func (t *RTVTestSuite) TestRTVPollRollback() {
	r, err := NewRTV()
	assert.NoError(t.T(), err)

	t.f.rtvHandler = func(f *fakeServer, w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "9999") // a new rtv value
	}
	t.f.cssHandler = func(f *fakeServer, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}

	err = r.poll()
	assert.Error(t.T(), err)
	// Values should not change, despite HTTP error.
	assert.Equal(t.T(), paddedRtv, r.GetRTV())
	assert.Equal(t.T(), css, r.GetCSS())
	assert.Equal(t.T(), 2, t.f.rtvCalls)
	assert.Equal(t.T(), 2, t.f.cssCalls)
}
