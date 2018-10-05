package util

import (
	"bytes"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorsSuite struct {
	suite.Suite
	logOut bytes.Buffer
}

func (this *ErrorsSuite) SetupTest() {
	this.logOut.Reset()
	log.SetOutput(&this.logOut)
}

func (this *ErrorsSuite) TearDownTest() {
	log.SetOutput(os.Stderr)
}

func (this *ErrorsSuite) TestError() {
	this.Assert().Equal("Coffee grinder is broken", NewHTTPError(418, "Coffee grinder is broken").Error())
}

func (this *ErrorsSuite) TestLogAndRespond() {
	resp := httptest.NewRecorder()
	NewHTTPError(418, "Coffee grinder is broken").LogAndRespond(resp)
	this.Assert().Equal(418, resp.Code)
	this.Assert().Equal("no-store", resp.Header().Get("Cache-Control"))
	this.Assert().Equal("I'm a teapot\n", resp.Body.String())
}

func TestErrorsSuite(t *testing.T) {
	suite.Run(t, new(ErrorsSuite))
}
