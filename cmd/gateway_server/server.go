// gRPC server to be used only for automated integration testing.

package main

import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/WICG/webpackage/go/signedexchange/certurl"
	pb "github.com/ampproject/amppackager/cmd/gateway_server/gateway"
	"github.com/ampproject/amppackager/packager/certcache"
	"github.com/ampproject/amppackager/packager/rtv"
	"github.com/ampproject/amppackager/packager/signer"
	"github.com/ampproject/amppackager/packager/util"
	"golang.org/x/crypto/ocsp"
	"google.golang.org/grpc"
)

var (
	port                = flag.Int("port", 9000, "Gateway server port")
	publisherServerPort = flag.Int("publisher_server_port", 10000,
		"Publisher server port.")
)

type gatewayServer struct {
	rtvCache *rtv.RTVCache
}

func shouldPackage() error {
	return nil
}

func errorToSXGResponse(err error) *pb.SXGResponse {
	response := &pb.SXGResponse{
		Error:            true,
		ErrorDescription: err.Error(),
	}
	return response
}

func createOCSPResponse(cert *x509.Certificate, key crypto.Signer) ([]byte, error) {
	thisUpdate := time.Now()

	// Construct args to ocsp.CreateResponse.
	template := ocsp.Response{
		SerialNumber: cert.SerialNumber,
		Status:       ocsp.Good,
		ThisUpdate:   thisUpdate,
		NextUpdate:   thisUpdate.Add(time.Hour * 24 * 7),
		IssuerHash:   crypto.SHA256,
	}
	return ocsp.CreateResponse(cert /*issuer*/, cert /*responderCert*/, template, key)
}

func (s *gatewayServer) GenerateSXG(ctx context.Context, request *pb.SXGRequest) (*pb.SXGResponse, error) {
	log.Println("Handling request with fetchUrl =", request.FetchUrl, "; signUrl =", request.SignUrl)

	certs, err := signedexchange.ParseCertificates(request.PublicCert)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	// Note: do not initialize certCache, we just want it to hold the certs for now.
	certCache := certcache.New(certs, nil, []string{""}, "", "", "", nil, time.Now)

	privateKey, err := util.ParsePrivateKey(request.PrivateKey)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	signUrl, err := url.Parse(request.SignUrl)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	var dotStarPattern = ".*"
	signUrlPattern := util.URLPattern{
		Domain:  signUrl.Host,
		QueryRE: &dotStarPattern,
	}
	err = util.ValidateSignURLPattern(&signUrlPattern)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	fetchUrlPattern := util.URLPattern{
		Scheme:                 []string{"http"},
		Domain:                 fmt.Sprintf("localhost:%d", *publisherServerPort),
		ErrorOnStatefulHeaders: false,
		QueryRE:                &dotStarPattern,
		SamePath:               new(bool),
	}
	*fetchUrlPattern.SamePath = false
	err = util.ValidateFetchURLPattern(&fetchUrlPattern)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	urlSets := []util.URLSet{
		{
			Sign:  &signUrlPattern,
			Fetch: &fetchUrlPattern,
		},
	}

	packager, err := signer.New(certCache, privateKey, urlSets, s.rtvCache, shouldPackage, signUrl, false, []string{}, time.Now)

	if err != nil {
		return errorToSXGResponse(err), nil
	}

	if packager == nil {
		return errorToSXGResponse(err), nil
	}

	baseUrl, err := url.Parse(fmt.Sprintf("http://localhost:%d", *port))
	if err != nil {
		return errorToSXGResponse(err), nil
	}
	q := baseUrl.Query()
	q.Set("fetch", request.FetchUrl)
	q.Set("sign", request.SignUrl)
	baseUrl.RawQuery = q.Encode()

	httpreq, err := http.NewRequest("GET", baseUrl.String(), nil)
	httpresp := httptest.NewRecorder()
	packager.ServeHTTP(httpresp, httpreq)

	// TODO(amaltas): Capture error when signer returns unsigned document.
	if httpresp.Code != 200 {
		// TODO(amaltas): Add counter.
		return &pb.SXGResponse{
			Error:            true,
			ErrorDescription: "Packager error.",
		}, nil
	}

	// Create cert-chain+cbor.
	var ocspDer []byte
	if len(certs) > 1 {
		// Attach an OCSP response, signed with the second cert in the
		// chain (assumed to be the issuer and using the same private
		// key as the leaf cert).
		var err error
		ocspDer, err = createOCSPResponse(certs[1], privateKey.(*ecdsa.PrivateKey))
		if err != nil {
			return errorToSXGResponse(err), nil
		}
	} else {
		// Make up an invalid OCSP response.
		ocspDer = []byte("ocsp")
	}
	var sctList []byte
	certChain, err := certurl.NewCertChain(certs, ocspDer, sctList)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	buf := &bytes.Buffer{}
	err = certChain.Write(buf)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	// Record http headers from the packager.
	http_headers := map[string]string{}
	for header_key, header_value := range httpresp.Header() {
		// Ignores multiple header values.
		http_headers[strings.ToLower(header_key)] =
			string(header_value[0])
	}

	return &pb.SXGResponse{
		Sxg:         httpresp.Body.Bytes(),
		Cbor:        buf.Bytes(),
		HttpHeaders: http_headers}, nil
}

func main() {
	flag.Parse()

	if *port == -1 {
		log.Fatalf("Set flag -port")
	}
	if *publisherServerPort == -1 {
		log.Fatalf("Set flag -publisher-server-port")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	rtvCache, err := rtv.New()
	if err != nil {
		log.Fatalf("Error initializing RTVCache: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGatewayServiceServer(grpcServer, &gatewayServer{rtvCache: rtvCache})
	log.Println("Starting server on port: ", *port)
	grpcServer.Serve(listener)
}
