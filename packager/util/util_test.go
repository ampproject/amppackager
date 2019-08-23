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
	assert.Equal(t, "Qk83Jo8qB8cEtxfb_7eit0SWVt0pdj5e7oDCqEgf77o", util.CertName(pkgt.B3Certs[0]))
}

func TestGetDurationToExpiry(t *testing.T) {
	// Time before the cert validity.
	beforeCert := time.Date(2019, time.May, 8, 0, 0, 0, 0, time.UTC)
	// Time after the cert validity.
	afterCert := time.Date(2019, time.August, 8, 0, 0, 0, 0, time.UTC)
	// Time 2 days before cert validity expiration.
	twoDaysBeforeExpiry := time.Date(2019, time.August, 5, 5, 43, 32, 0, time.UTC)
	// Time 0 days, 1 hour before cert validity expiration.
	oneHourBeforeExpiry := time.Date(2019, time.August, 7, 4, 43, 32, 0, time.UTC)
	// Time 0 days before cert validity expiration.
	zeroDaysBeforeExpiry := time.Date(2019, time.August, 7, 5, 43, 32, 0, time.UTC)

	d, err := util.GetDurationToExpiry(pkgt.B3Certs[0], beforeCert)
	assert.EqualError(t, err, "Certificate is future-dated")
	d, err = util.GetDurationToExpiry(pkgt.B3Certs[0], afterCert)
	assert.EqualError(t, err, "Certificate is expired")

	d, err = util.GetDurationToExpiry(pkgt.B3Certs[0], twoDaysBeforeExpiry)
	assert.Equal(t, time.Duration(2 * time.Hour * 24), d)

	d, err = util.GetDurationToExpiry(pkgt.B3Certs[0], oneHourBeforeExpiry)
	assert.Equal(t, time.Duration(1 * time.Hour), d)

	d, err = util.GetDurationToExpiry(pkgt.B3Certs[0], zeroDaysBeforeExpiry)
	assert.Equal(t, time.Duration(0), d)
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
