package util_test

import (
	"testing"

	pkgt "github.com/ampproject/amppackager/packager/testing"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/stretchr/testify/assert"
)

func TestCanSignHttpExchanges(t *testing.T) {
	// Leaf node has the extension.
	assert.True(t, util.CanSignHttpExchanges(pkgt.Certs[0]))
	// CA node does not.
	assert.False(t, util.CanSignHttpExchanges(pkgt.Certs[1]))
}
