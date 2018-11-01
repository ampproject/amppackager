package util_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"testing"

	pkgt "github.com/ampproject/amppackager/packager/testing"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCertName(t *testing.T) {
	assert.Equal(t, "hwEGDUdv5LgqOr9jM-5nhe6vgg65n5upqt1CpAlLP8U", util.CertName(pkgt.Certs[0]))
}

// ParsePrivateKey() is tested indirectly via the definition of pkgt.Key.
func TestParsePrivateKey(t *testing.T) {
	require.IsType(t, &ecdsa.PrivateKey{}, pkgt.Key)
	assert.Equal(t, elliptic.P256(), pkgt.Key.(*ecdsa.PrivateKey).PublicKey.Curve)
}

func TestCanSignHttpExchanges(t *testing.T) {
	// Leaf node has the extension.
	assert.True(t, util.CanSignHttpExchanges(pkgt.Certs[0]))
	// CA node does not.
	assert.False(t, util.CanSignHttpExchanges(pkgt.Certs[1]))
}
