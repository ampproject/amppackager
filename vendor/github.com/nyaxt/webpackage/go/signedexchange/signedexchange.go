package signedexchange

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/nyaxt/webpackage/go/signedexchange/cbor"
	"github.com/nyaxt/webpackage/go/signedexchange/mice"
)

type Exchange struct {
	// Request
	RequestUri     *url.URL
	RequestHeaders http.Header

	// Response
	ResponseStatus  int
	ResponseHeaders http.Header

	// Payload
	Payload []byte
}

var (
	keyMethod = []byte(":method")
	keyURL    = []byte(":url")
	keyStatus = []byte(":status")

	valueGet = []byte("GET")
)

func NewExchange(uri *url.URL, requestHeaders http.Header, status int, responseHeaders http.Header, payload []byte, miRecordSize int) (*Exchange, error) {
	e := &Exchange{
		RequestUri:      uri,
		ResponseStatus:  status,
		RequestHeaders:  requestHeaders,
		ResponseHeaders: responseHeaders,
	}
	if err := e.miEncode(payload, miRecordSize); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Exchange) miEncode(payload []byte, recordSize int) error {
	var buf bytes.Buffer
	mi, err := mice.Encode(&buf, payload, recordSize)
	if err != nil {
		return err
	}
	e.Payload = buf.Bytes()
	e.ResponseHeaders.Add("Content-Encoding", "mi-sha256")
	e.ResponseHeaders.Add("MI", mi)
	return nil
}

func (e *Exchange) AddSignatureHeader(s *Signer) error {
	h, err := s.signatureHeaderValue(e)
	if err != nil {
		return err
	}
	e.ResponseHeaders.Add("Signature", h)
	return nil
}

func (e *Exchange) encodeRequestCommon(enc *cbor.Encoder) []*cbor.MapEntryEncoder {
	return []*cbor.MapEntryEncoder{
		cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
			keyE.EncodeByteString(keyMethod)
			valueE.EncodeByteString(valueGet)
		}),
		cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
			keyE.EncodeByteString(keyURL)
			valueE.EncodeByteString([]byte(e.RequestUri.String()))
		}),
	}
}

func (e *Exchange) encodeRequest(enc *cbor.Encoder) error {
	mes := e.encodeRequestCommon(enc)
	return enc.EncodeMap(mes)
}

func (e *Exchange) decodeRequest(dec *cbor.Decoder) error {
	nelem, err := dec.DecodeMapHeader()
	if err != nil {
		return err
	}

	for i := uint64(0); i < nelem; i++ {
		key, err := dec.DecodeByteString()
		if err != nil {
			return fmt.Errorf("signedexchange: Failed to decode key bytestring: %s", err)
		}
		value, err := dec.DecodeByteString()
		if err != nil {
			return fmt.Errorf("signedexchange: Failed to decode value bytestring: %s", err)
		}
		// TODO: add key/value str validation?

		if bytes.Equal(key, keyMethod) {
			if !bytes.Equal(value, valueGet) {
				// TODO: Consider alternative to log.Printf to communicate ill-formed signed-exchange
				log.Printf("Request map key %q: Expected %q, got %q", keyMethod, valueGet, value)
			}
		} else if bytes.Equal(key, keyURL) {
			e.RequestUri, err = url.Parse(string(value))
			if err != nil {
				// TODO: Consider alternative to log.Printf to communicate ill-formed signed-exchange
				log.Printf("Failed to parse URI: %q", value)
			}
		} else {
			// TODO: dup chk
			e.RequestHeaders.Add(string(key), string(value))
		}
	}
	return nil
}

func normalizeHeaderValues(values []string) string {
	// RFC 2616 - Hypertext Transfer Protocol -- HTTP/1.1
	// 4.2 Message Headers
	// https://tools.ietf.org/html/rfc2616#section-4.2
	//
	// Multiple message-header fields with the same field-name MAY be
	// present in a message if and only if the entire field-value for that
	// header field is defined as a comma-separated list [i.e., #(values)].
	// It MUST be possible to combine the multiple header fields into one
	// "field-name: field-value" pair, without changing the semantics of the
	// message, by appending each subsequent field-value to the first, each
	// separated by a comma. The order in which header fields with the same
	// field-name are received is therefore significant to the
	// interpretation of the combined field value, and thus a proxy MUST NOT
	// change the order of these field values when a message is forwarded.
	return strings.Join(values, ",")
}

func (e *Exchange) encodeRequestWithHeaders(enc *cbor.Encoder) error {
	mes := e.encodeRequestCommon(enc)
	for name, value := range e.RequestHeaders {
		mes = append(mes,
			cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
				keyE.EncodeByteString([]byte(strings.ToLower(name)))
				valueE.EncodeByteString([]byte(normalizeHeaderValues(value)))
			}))
	}
	return enc.EncodeMap(mes)
}

func (e *Exchange) encodeResponseHeaders(enc *cbor.Encoder) error {
	mes := []*cbor.MapEntryEncoder{
		cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
			keyE.EncodeByteString(keyStatus)
			valueE.EncodeByteString([]byte(strconv.Itoa(e.ResponseStatus)))
		}),
	}
	for name, value := range e.ResponseHeaders {
		mes = append(mes,
			cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
				keyE.EncodeByteString([]byte(strings.ToLower(name)))
				valueE.EncodeByteString([]byte(normalizeHeaderValues(value)))
			}))
	}
	return enc.EncodeMap(mes)
}

func (e *Exchange) decodeResponseHeaders(dec *cbor.Decoder) error {
	nelem, err := dec.DecodeMapHeader()
	if err != nil {
		return err
	}

	for i := uint64(0); i < nelem; i++ {
		key, err := dec.DecodeByteString()
		if err != nil {
			return fmt.Errorf("signedexchange: Failed to decode key bytestring: %s", err)
		}
		value, err := dec.DecodeByteString()
		if err != nil {
			return fmt.Errorf("signedexchange: Failed to decode value bytestring: %s", err)
		}
		// TODO: add key/value str validation?

		if bytes.Equal(key, keyStatus) {
			// TODO: add value str validation that it only contains [0-9]
			e.ResponseStatus, err = strconv.Atoi(string(value))
			if err != nil {
				log.Printf("Failed to parse responseStatus: %q", value)
			}
		} else {
			// TODO: dup chk
			e.ResponseHeaders.Add(string(key), string(value))
		}
	}
	return nil
}

// draft-yasskin-http-origin-signed-responses.html#rfc.section.3.4
func (e *Exchange) encodeExchangeHeaders(enc *cbor.Encoder) error {
	if err := enc.EncodeArrayHeader(2); err != nil {
		return fmt.Errorf("signedexchange: failed to encode top-level array header: %v", err)
	}
	if err := e.encodeRequest(enc); err != nil {
		return err
	}
	if err := e.encodeResponseHeaders(enc); err != nil {
		return err
	}
	return nil
}

// draft-yasskin-http-origin-signed-responses.html#application-http-exchange
func WriteExchangeFile(w io.Writer, e *Exchange) error {
	buf := &bytes.Buffer{}
	enc := cbor.NewEncoder(buf)
	if err := enc.EncodeArrayHeader(2); err != nil {
		return err
	}
	if err := e.encodeRequestWithHeaders(enc); err != nil {
		return err
	}
	if err := e.encodeResponseHeaders(enc); err != nil {
		return err
	}

	// 1. The first 3 bytes of the content represents the length of the CBOR
	// encoded section, encoded in network byte (big-endian) order.
	cborBytes := buf.Bytes()
	if len(cborBytes) >= 524288 {
		return fmt.Errorf("signedexchange: request headers too big: %d bytes", len(cborBytes))
	}
	if _, err := w.Write([]byte{
		byte(len(cborBytes) >> 16),
		byte(len(cborBytes) >> 8),
		byte(len(cborBytes)),
	}); err != nil {
		return err
	}

	// 2. Then, immediately follows a CBOR-encoded array containing 2 elements:
	// - a map of request header field names to values, encoded as byte strings,
	//   with ":method", and ":url" pseudo header fields
	// - a map from response header field names to values, encoded as byte strings,
	//   with a ":status" pseudo-header field containing the status code (encoded
	//   as 3 ASCII letter byte string)
	if _, err := w.Write(cborBytes); err != nil {
		return err
	}

	// 3. Then, immediately follows the response body, encoded in MI.
	// (note that this doesn't have the length 3 bytes like the CBOR section does)
	if _, err := w.Write(e.Payload); err != nil {
		return err
	}

	// FIXME: Support "trailer"

	return nil
}

func ReadExchangeFile(r io.Reader) (*Exchange, error) {
	var encodedCborLength [3]byte
	if _, err := io.ReadFull(r, encodedCborLength[:]); err != nil {
		return nil, fmt.Errorf("signedexchange: Failed to read length header")
	}
	cborLength := int(encodedCborLength[0])<<16 |
		int(encodedCborLength[1])<<8 |
		int(encodedCborLength[2])

	cborBytes := make([]byte, cborLength)
	if _, err := io.ReadFull(r, cborBytes); err != nil {
		return nil, fmt.Errorf("signedexchange: Failed to read CBOR header binary")
	}

	buf := bytes.NewBuffer(cborBytes)
	dec := cbor.NewDecoder(buf)
	nelem, err := dec.DecodeArrayHeader()
	if err != nil {
		return nil, fmt.Errorf("signedexchange: Failed to read CBOR header array")
	}
	if nelem != 2 {
		// TODO: Consider alternative to log.Printf to communicate ill-formed signed-exchange
		log.Printf("Expected 2 elements in top-level array, but got %d elements", nelem)
	}

	e := &Exchange{
		RequestHeaders:  http.Header{},
		ResponseHeaders: http.Header{},
	}
	if err := e.decodeRequest(dec); err != nil {
		return nil, fmt.Errorf("signedexchange: Failed to decode request map: %v", err)
	}
	if err := e.decodeResponseHeaders(dec); err != nil {
		return nil, fmt.Errorf("signedexchange: Failed to decode response headers map: %v", err)
	}

	miHeaderValue := e.ResponseHeaders.Get("mi")
	var payloadBuf bytes.Buffer
	if err := mice.Decode(&payloadBuf, r, miHeaderValue); err != nil {
		return nil, fmt.Errorf("signedexchange: Failed to mice decode payload: %v", err)
	}
	e.Payload = payloadBuf.Bytes()

	return e, nil
}

func (e *Exchange) PrettyPrint(w io.Writer) {
	fmt.Fprintln(w, "request:")
	fmt.Fprintf(w, "  uri: %s\n", e.RequestUri.String())
	fmt.Fprintln(w, "  headers:")
	for k, _ := range e.RequestHeaders {
		fmt.Fprintf(w, "    %s: %s\n", k, e.ResponseHeaders.Get(k))
	}
	fmt.Fprintln(w, "response:")
	fmt.Fprintf(w, "  status: %d\n", e.ResponseStatus)
	fmt.Fprintln(w, "  headers:")
	for k, _ := range e.ResponseHeaders {
		fmt.Fprintf(w, "    %s: %s\n", k, e.ResponseHeaders.Get(k))
	}
	fmt.Fprintf(w, "payload [%d bytes]:\n", len(e.Payload))
	w.Write(e.Payload)
}
