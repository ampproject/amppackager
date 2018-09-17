package amp_cache_transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSendSXG(t *testing.T) {
	assert.True(t, ShouldSendSXG("any"))
	assert.True(t, ShouldSendSXG("google"))
	assert.True(t, ShouldSendSXG("google, any"))
	assert.True(t, ShouldSendSXG("google, foobar"))
	assert.True(t, ShouldSendSXG("google,foobar"))
	assert.True(t, ShouldSendSXG("google,\tfoobar"))
	assert.True(t, ShouldSendSXG("google ,foobar"))
	assert.True(t, ShouldSendSXG("google, a*b-c_d/e"))
	assert.True(t, ShouldSendSXG("foobar, any"))

	assert.False(t, ShouldSendSXG(""))
	assert.False(t, ShouldSendSXG(" any"))
	assert.False(t, ShouldSendSXG("any "))
	assert.False(t, ShouldSendSXG("foobar"))
	assert.False(t, ShouldSendSXG("foobar, baz"))
	assert.False(t, ShouldSendSXG("googleany"))
	assert.False(t, ShouldSendSXG("google;any"))
	assert.False(t, ShouldSendSXG("google;v=1"))
	assert.False(t, ShouldSendSXG("google,123"))
	assert.False(t, ShouldSendSXG("google, eh!"))
	assert.False(t, ShouldSendSXG("ABC,google"))
}
