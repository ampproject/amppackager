package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/WICG/webpackage/go/signedexchange/certurl"
	pb "github.com/ampproject/amppackager/cmd/gateway_server/gateway"
	"github.com/ampproject/amppackager/packager/rtv"
	"github.com/ampproject/amppackager/packager/signer"
	"github.com/ampproject/amppackager/packager/util"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

var (
	port                = flag.Int("port", 9000, "Gateway server port")
	publisherServerPort = flag.Int("publisher_server_port", 10000,
		"Publisher server port.")
	privateKey = flag.String("private_key", "",
		"Path to private key. Must be same private key used to generate certs.")
)

type gatewayServer struct{}

func shouldPackage() bool {
	return true
}

func errorToSXGResponse(err error) *pb.SXGResponse {
	response := &pb.SXGResponse{
		Error:            true,
		ErrorDescription: err.Error(),
	}
	return response
}

func (s *gatewayServer) GenerateSXG(ctx context.Context, request *pb.SXGRequest) (*pb.SXGResponse, error) {
	rtvCache, err := rtv.New()
	if err != nil {
		return errorToSXGResponse(err), nil
	}

	certs, err := signedexchange.ParseCertificates(request.PublicCert)
	if err != nil {
		return errorToSXGResponse(err), nil
	}

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

	packager, err := signer.New(certs[0], privateKey, urlSets, rtvCache, shouldPackage, signUrl, false)

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
	packager.ServeHTTP(httpresp, httpreq, httprouter.Params{})

	// TODO(amaltas): Capture error when signer returns unsigned document.
	if httpresp.Code != 200 {
		// TODO(amaltas): Add counter.
		return &pb.SXGResponse{
			Error:            true,
			ErrorDescription: "Packager error.",
		}, nil
	}

	// Creates cbor data.
	ocspDer := []byte("ocsp")
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
	keyPem, err := ioutil.ReadFile(*privateKey)
	if err != nil {
		log.Fatalf("Error reading private key file.")
	}
	key, err := util.ParsePrivateKey(keyPem)
	if err != nil {
		log.Fatalf("Error reading parsed private key string.")
	}
	if key == nil {
		log.Fatalf("Key is nil.")
	}

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

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGatewayServiceServer(grpcServer, &gatewayServer{})
	fmt.Println("Starting server on port: ", *port)
	grpcServer.Serve(listener)
}
