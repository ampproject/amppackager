package amp_cache_transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSendSXG(t *testing.T) {
	assert.Equal(t, "any", ShouldSendSXG("any"))
	assert.Equal(t, "any", ShouldSendSXG("foobar, any"))
	assert.Equal(t, "any", ShouldSendSXG("any, google"))
	assert.Equal(t, "google", ShouldSendSXG("google"))
	assert.Equal(t, "google", ShouldSendSXG("google, any"))
	assert.Equal(t, "google", ShouldSendSXG("google, foobar"))
	assert.Equal(t, "google", ShouldSendSXG("google,foobar"))
	assert.Equal(t, "google", ShouldSendSXG("google,\tfoobar"))
	assert.Equal(t, "google", ShouldSendSXG("google ,foobar"))
	assert.Equal(t, "google", ShouldSendSXG("google, a*b-c_d/e"))

	assert.Equal(t, "", ShouldSendSXG(""))
	assert.Equal(t, "", ShouldSendSXG(" any"))
	assert.Equal(t, "", ShouldSendSXG("any "))
	assert.Equal(t, "", ShouldSendSXG("foobar"))
	assert.Equal(t, "", ShouldSendSXG("foobar, baz"))
	assert.Equal(t, "", ShouldSendSXG("googleany"))
	assert.Equal(t, "", ShouldSendSXG("google;any"))
	assert.Equal(t, "", ShouldSendSXG("google;v=1"))
	assert.Equal(t, "", ShouldSendSXG("google,123"))
	assert.Equal(t, "", ShouldSendSXG("google, eh!"))
	assert.Equal(t, "", ShouldSendSXG("ABC,google"))
}
