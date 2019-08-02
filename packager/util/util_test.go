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

func errorFrom(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func TestCertName(t *testing.T) {
	assert.Equal(t, "Qk83Jo8qB8cEtxfb_7eit0SWVt0pdj5e7oDCqEgf77o", util.CertName(pkgt.B3Certs[0]))
}

// ParsePrivateKey() is tested indirectly via the definition of pkgt.B3Key.
func TestParsePrivateKey(t *testing.T) {
	require.IsType(t, &ecdsa.PrivateKey{}, pkgt.B3Key)
	assert.Equal(t, elliptic.P256(), pkgt.B3Key.(*ecdsa.PrivateKey).PublicKey.Curve)
}

func TestCanSignHttpExchangesExtension(t *testing.T) {
	// Leaf node has the extension.
	assert.Nil(t, util.CanSignHttpExchanges(pkgt.B3Certs[0]))
	// CA node does not.
	assert.EqualError(t, util.CanSignHttpExchanges(pkgt.B3Certs[1]), "Certificate is missing CanSignHttpExchanges extension")
}

func TestParseCertificate(t *testing.T) {
	assert.Nil(t, util.CertificateMatches(pkgt.B3Certs[0], pkgt.B3Key, "amppackageexample.com"))
}

func TestParseCertificateSubjectAltName(t *testing.T) {
	assert.Nil(t, util.CertificateMatches(pkgt.B3Certs[0], pkgt.B3Key, "www.amppackageexample.com"))
}

func TestParseCertificateNotMatchX(t *testing.T) {
	assert.Contains(t, errorFrom(util.CertificateMatches(pkgt.B3Certs[0],
		pkgt.B3Key2, "amppackageexample.com")), "PublicKey.X not match")
}

func TestParseCertificateNotMatchCurve(t *testing.T) {
	assert.Contains(t, errorFrom(util.CertificateMatches(pkgt.B3Certs[0],
		pkgt.B3KeyP521, "amppackageexample.com")), "PublicKey.Curve not match")
}

func TestParseCertificateNotMatchDomain(t *testing.T) {
	assert.Contains(t, errorFrom(util.CertificateMatches(pkgt.B3Certs2[0],
		pkgt.B3Key2, "amppackageexample.com")), "x509: certificate is valid for amppackageexample2.com, www.amppackageexample2.com, not amppackageexample.com")
}

func TestParse91DaysCertificate(t *testing.T) {
	assert.Contains(t, errorFrom(util.CanSignHttpExchanges(pkgt.B3Certs91Days[0])),
	"Certificate MUST have a Validity Period no greater than 90 days")
}
