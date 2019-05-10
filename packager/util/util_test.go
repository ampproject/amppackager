package util_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"testing"
	"time"

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
	assert.Equal(t, "PJ1IwfP1igOlJd2oTUVs2mj4dWIZcOWHMk5jfJYS2Qc", util.CertName(pkgt.Certs[0]))
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

func TestParseCertificate(t *testing.T) {
	assert.Nil(t, util.CheckCertificate(pkgt.B3Certs[0], pkgt.B3Key, "amppackageexample.com"))
}

func TestParseCertificateSubjectAltName(t *testing.T) {
	assert.Nil(t, util.CheckCertificate(pkgt.B3Certs[0], pkgt.B3Key, "www.amppackageexample.com"))
}

func TestParseCertificateNotMatchX(t *testing.T) {
	assert.Contains(t, errorFrom(util.CheckCertificate(pkgt.B3Certs[0],
		pkgt.B3Key2, "amppackageexample.com")), "PublicKey.X not match")
}

func TestParseCertificateNotMatchCurve(t *testing.T) {
	assert.Contains(t, errorFrom(util.CheckCertificate(pkgt.B3Certs[0],
		pkgt.B3KeyP521, "amppackageexample.com")), "PublicKey.Curve not match")
}

func TestParseCertificateNotMatchDomain(t *testing.T) {
	assert.Contains(t, errorFrom(util.CheckCertificate(pkgt.B3Certs2[0],
		pkgt.B3Key2, "amppackageexample.com")), "x509: certificate is valid for amppackageexample2.com, www.amppackageexample2.com, not amppackageexample.com")
}

func TestParse91DaysCertificate(t *testing.T) {
	assert.Contains(t, errorFrom(util.CheckCertificate(pkgt.B3Certs91Days[0],
		pkgt.B3Key, "amppackageexample.com")), "Certificate MUST have a Validity Period no greater than 90 days")
}

func TestParseCertificateIssuedBeforeMay1(t *testing.T) {
	aug1st := time.Date(2009, time.August, 1, 0, 0, 0, 0, time.UTC)
	if time.Now().After(aug1st) {
		assert.Contains(t, errorFrom(util.CheckCertificate(pkgt.Certs[0],
		pkgt.Key, "amppackageexample.com")), "Certificate MUST have a Validity Period no greater than 90 days")
	} else {
		assert.Nil(t, util.CheckCertificate(pkgt.Certs[0], pkgt.Key, "amppackageexample.com"))
	}
}
