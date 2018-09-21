package packager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSignHttpExchanges(t *testing.T) {
	// Leaf node has the extension.
	assert.True(t, CanSignHttpExchanges(certs[0]))
	// CA node does not.
	assert.False(t, CanSignHttpExchanges(certs[1]))
}
