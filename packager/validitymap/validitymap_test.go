package validitymap

import (
	"io/ioutil"
	"testing"

	pkgt "github.com/ampproject/amppackager/packager/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidityMap(t *testing.T) {
	handler, err := New()
	require.NoError(t, err)

	resp := pkgt.Get(t, handler, "/")
	defer resp.Body.Close()
	assert.Equal(t, "application/cbor", resp.Header.Get("Content-Type"))
	assert.Equal(t, "public, max-age=604800", resp.Header.Get("Cache-Control"))
	assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, []byte("\xA0"), body)
}
