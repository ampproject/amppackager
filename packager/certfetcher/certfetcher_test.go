package certfetcher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-acme/lego/v3/acme"
	"github.com/go-acme/lego/v3/platform/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	jose "gopkg.in/square/go-jose.v2"
)

// CertResponseMock is just any valid SXG cert generated via:
// https://docs.digicert.com/manage-certificates/certificate-profile-options/get-your-signed-http-exchange-certificate/
const CertResponseMock = `-----BEGIN CERTIFICATE-----
MIIFZDCCBOqgAwIBAgIQBxgJcqaHzEUEVOBlp+mFJTAKBggqhkjOPQQDAjBMMQsw
CQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMSYwJAYDVQQDEx1EaWdp
Q2VydCBFQ0MgU2VjdXJlIFNlcnZlciBDQTAeFw0xODA4MzEwMDAwMDBaFw0yMDA5
MDQxMjAwMDBaMHAxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMREw
DwYDVQQHEwhTYW4gSm9zZTEZMBcGA1UEChMQR3JlZ29yeSBHcm90aGF1czEeMBwG
A1UEAxMVYXplaS1wYWNrYWdlLXRlc3QuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0D
AQcDQgAEc1QETMcI6mWbyAa4y026CLY/OnVGutWCrvTjO8WFZIZ16dxzO7UIsnPc
LdPVxnJQkY7uZnzfFYqLTBgHcgwE4KOCA4gwggOEMB8GA1UdIwQYMBaAFKOd5h/5
2jlPwG7okcuVpdox4gqfMB0GA1UdDgQWBBQdgF7QKodsaqDTJ71z8z8hcXT0LTA7
BgNVHREENDAyghVhemVpLXBhY2thZ2UtdGVzdC5jb22CGXd3dy5hemVpLXBhY2th
Z2UtdGVzdC5jb20wDgYDVR0PAQH/BAQDAgeAMB0GA1UdJQQWMBQGCCsGAQUFBwMB
BggrBgEFBQcDAjBpBgNVHR8EYjBgMC6gLKAqhihodHRwOi8vY3JsMy5kaWdpY2Vy
dC5jb20vc3NjYS1lY2MtZzEuY3JsMC6gLKAqhihodHRwOi8vY3JsNC5kaWdpY2Vy
dC5jb20vc3NjYS1lY2MtZzEuY3JsMEwGA1UdIARFMEMwNwYJYIZIAYb9bAEBMCow
KAYIKwYBBQUHAgEWHGh0dHBzOi8vd3d3LmRpZ2ljZXJ0LmNvbS9DUFMwCAYGZ4EM
AQIDMHsGCCsGAQUFBwEBBG8wbTAkBggrBgEFBQcwAYYYaHR0cDovL29jc3AuZGln
aWNlcnQuY29tMEUGCCsGAQUFBzAChjlodHRwOi8vY2FjZXJ0cy5kaWdpY2VydC5j
b20vRGlnaUNlcnRFQ0NTZWN1cmVTZXJ2ZXJDQS5jcnQwDAYDVR0TAQH/BAIwADAQ
BgorBgEEAdZ5AgEWBAIFADCCAX4GCisGAQQB1nkCBAIEggFuBIIBagFoAHYApLkJ
kLQYWBSHuxOizGdwCjw1mAT5G9+443fNDsgN3BAAAAFlklqt1QAABAMARzBFAiAs
fo+czC5jBghS0acCZ8mLoMNFnnbvBnNCwXhIzohVQAIhALNlBFxwlJ7gift1XNtA
PLq3mI2vjtIpgtQ2azuVX/gBAHcAh3W/51l8+IxDmV+9827/Vo1HVjb/SrVgwbTq
/16ggw8AAAFlklquoQAABAMASDBGAiEAsG3VGlFzdghriY5qT3Mg3pEnLUHASVpu
bJXGfHBluUUCIQDIz36lErTWCOM5rJ7n5xWW15I1rumYCJzrUDZWjRShVgB1ALvZ
37wfinG1k5Qjl6qSe0c4V5UKq1LoGpCWZDaOHtGFAAABZZJardkAAAQDAEYwRAIg
S6oxgvn++wCfZ6wxt/lC2GoQX2LIJl5mrmHrMStgqxgCIF375hwD9aCMlv9SbfkL
GS2Mka/kMMtZrQVIQsyi3lhbMAoGCCqGSM49BAMCA2gAMGUCMBo9NIEu38bvGcKy
P9oN2ELBL3dgIXDq3oU85vX/8rEuwNLvsC4lMtk/QJap3dxuSAIxAMro/ZXw3lVP
YO0x/svBxXf6vDC01lO7LTJvjqA2Hfa/7GI5gbUr3sRTU09aO9ixOA==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDrDCCApSgAwIBAgIQCssoukZe5TkIdnRw883GEjANBgkqhkiG9w0BAQwFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0xMzAzMDgxMjAwMDBaFw0yMzAzMDgxMjAwMDBaMEwxCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxJjAkBgNVBAMTHURpZ2lDZXJ0IEVDQyBT
ZWN1cmUgU2VydmVyIENBMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE4ghC6nfYJN6g
LGSkE85AnCNyqQIKDjc/ITa4jVMU9tWRlUvzlgKNcR7E2Munn17voOZ/WpIRllNv
68DLP679Wz9HJOeaBy6Wvqgvu1cYr3GkvXg6HuhbPGtkESvMNCuMo4IBITCCAR0w
EgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwNAYIKwYBBQUHAQEE
KDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wQgYDVR0f
BDswOTA3oDWgM4YxaHR0cDovL2NybDMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0R2xv
YmFsUm9vdENBLmNybDA9BgNVHSAENjA0MDIGBFUdIAAwKjAoBggrBgEFBQcCARYc
aHR0cHM6Ly93d3cuZGlnaWNlcnQuY29tL0NQUzAdBgNVHQ4EFgQUo53mH/naOU/A
buiRy5Wl2jHiCp8wHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUwDQYJ
KoZIhvcNAQEMBQADggEBAMeKoENL7HTJxavVHzA1Nm6YVntIrAVjrnuaVyRXzG/6
3qttnMe2uuzO58pzZNvfBDcKAEmzP58mrZGMIOgfiA4q+2Y3yDDo0sIkp0VILeoB
UEoxlBPfjV/aKrtJPGHzecicZpIalir0ezZYoyxBEHQa0+1IttK7igZFcTMQMHp6
mCHdJLnsnLWSB62DxsRq+HfmNb4TDydkskO/g+l3VtsIh5RHFPVfKK+jaEyDj2D3
loB5hWp2Jp2VDCADjT7ueihlZGak2YPqmXTNbk19HOuNssWvFhtOyPNV6og4ETQd
Ea8/B6hPatJ0ES8q/HO3X8IVQwVs1n3aAr0im0/T+Xc=
-----END CERTIFICATE-----
`

func TestNewFetcher(t *testing.T) {
	_, apiURL, tearDown := tester.SetupFakeAPI()
	defer tearDown()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Could not generate test key")

	csr := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   "test.example.com",
			Organization: []string{"Acme Co"},
		},
		DNSNames: []string{"test.example.com"},
	}

	fetcher, err := New("test@test.com", "eab.kid", "eab.hmac", &csr,
		privateKey, apiURL+"/dir", 5002, "", 0, "", false)
	assert.Nil(t, err)
	assert.NotNil(t, fetcher.legoClient)
	assert.Equal(t, "test@test.com", fetcher.AcmeUser.Email)
	assert.Equal(t, privateKey, fetcher.AcmeUser.key)
}

func TestFetchCertSuccess(t *testing.T) {
	mux, apiURL, tearDown := tester.SetupFakeAPI()
	defer tearDown()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Could not generate test key")

	setupMux(mux, apiURL, privateKey)
	mux.HandleFunc("/certificate", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(CertResponseMock))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	csr := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   "test.example.com",
			Organization: []string{"Acme Co"},
		},
		DNSNames: []string{"test.example.com"},
	}

	fetcher, err := New("test@test.com", "eab.kid", "eab.hmac", &csr,
		privateKey, apiURL+"/dir", 5002, "", 0, "", false)
	assert.Nil(t, err)
	assert.NotNil(t, fetcher)

	cert, err := fetcher.FetchNewCert()
	assert.Nil(t, err)
	assert.NotNil(t, cert)
}

func TestFetchCertFail(t *testing.T) {
	mux, apiURL, tearDown := tester.SetupFakeAPI()
	defer tearDown()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Could not generate test key")

	setupMux(mux, apiURL, privateKey)
	mux.HandleFunc("/certificate", func(w http.ResponseWriter, _ *http.Request) {
		// Intentionally return an error.
		http.Error(w, "", http.StatusInternalServerError)
	})

	csr := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   "test.example.com",
			Organization: []string{"Acme Co"},
		},
		DNSNames: []string{"test.example.com"},
	}

	fetcher, err := New("test@test.com", "eab.kid", "eab.hmac", &csr,
		privateKey, apiURL+"/dir", 5002, "", 0, "", false)
	assert.Nil(t, err)
	assert.NotNil(t, fetcher)

	cert, err := fetcher.FetchNewCert()
	assert.NotNil(t, err)
	assert.Nil(t, cert)
}

func setupMux(mux *http.ServeMux, apiURL string, privateKey *rsa.PrivateKey) {
	mux.HandleFunc("/newOrder", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		body, err := readSignedBody(r, privateKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		order := acme.Order{}
		err = json.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = tester.WriteJSONResponse(w, acme.Order{
			Status:      acme.StatusValid,
			Finalize:    apiURL + "/finalize",
			Identifiers: order.Identifiers,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/finalize", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		body, err := readSignedBody(r, privateKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		order := acme.Order{}
		err = json.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = tester.WriteJSONResponse(w, acme.Order{
			Status:      acme.StatusValid,
			Identifiers: order.Identifiers,
			Certificate: apiURL + "/certificate",
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func readSignedBody(r *http.Request, privateKey *rsa.PrivateKey) ([]byte, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	jws, err := jose.ParseSigned(string(reqBody))
	if err != nil {
		return nil, err
	}

	body, err := jws.Verify(&jose.JSONWebKey{
		Key:       privateKey.Public(),
		Algorithm: "RSA",
	})
	if err != nil {
		return nil, err
	}

	return body, nil
}
