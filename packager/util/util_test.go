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

func TestCanSignHttpExchangesExtension(t *testing.T) {
	// Before grace period, to allow the >90-day lifetime.
	now := time.Date(2019, time.July, 31, 0, 0, 0, 0, time.UTC)

	// Leaf node has the extension.
	assert.Nil(t, util.CanSignHttpExchanges(pkgt.Certs[0], now))
	// CA node does not.
	assert.EqualError(t, util.CanSignHttpExchanges(pkgt.Certs[1], now), "Certificate is missing CanSignHttpExchanges extension")
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

func TestParse90DaysCertificateAfterGracePeriod(t *testing.T) {
	now := time.Date(2019, time.August, 1, 0, 0, 0, 1, time.UTC)
	assert.Nil(t, util.CanSignHttpExchanges(pkgt.B3Certs[0], now))
}

func TestParse91DaysCertificate(t *testing.T) {
	assert.Contains(t, errorFrom(util.CanSignHttpExchanges(pkgt.B3Certs91Days[0],
		time.Now())), "Certificate MUST have a Validity Period no greater than 90 days")
}

func TestParseCertificateIssuedBeforeMay1InGarcePeriod(t *testing.T) {
	now := time.Date(2019, time.July, 31, 0, 0, 0, 0, time.UTC)
	assert.Nil(t, util.CanSignHttpExchanges(pkgt.Certs[0], now))
}

func TestParseCertificateIssuedBeforeMay1AfterGracePeriod(t *testing.T) {
	now := time.Date(2019, time.August, 1, 0, 0, 0, 1, time.UTC)
	assert.Contains(t, errorFrom(util.CanSignHttpExchanges(pkgt.Certs[0],
		now)), "Certificate MUST have a Validity Period no greater than 90 days")
}
