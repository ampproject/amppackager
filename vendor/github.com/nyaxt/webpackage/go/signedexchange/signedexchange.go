package signedexchange

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/WICG/webpackage/go/signedexchange/cbor"
	"github.com/WICG/webpackage/go/signedexchange/mice"
)

type Exchange struct {
	// Request
	requestUri     *url.URL
	requestHeaders http.Header

	// Response
	responseStatus  int
	responseHeaders http.Header

	// Payload
	payload []byte
}

func NewExchange(uri *url.URL, requestHeaders http.Header, status int, responseHeaders http.Header, payload []byte, miRecordSize int) (*Exchange, error) {
	e := &Exchange{
		requestUri:      uri,
		responseStatus:  status,
		requestHeaders:  requestHeaders,
		responseHeaders: responseHeaders,
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
	e.payload = buf.Bytes()
	e.responseHeaders.Add("Content-Encoding", "mi-sha256")
	e.responseHeaders.Add("MI", mi)
	return nil
}

func (e *Exchange) AddSignatureHeader(s *Signer) error {
	h, err := s.signatureHeaderValue(e)
	if err != nil {
		return err
	}
	e.responseHeaders.Add("Signature", h)
	return nil
}

func (e *Exchange) encodeRequestCommon(enc *cbor.Encoder) []*cbor.MapEntryEncoder {
	return []*cbor.MapEntryEncoder{
		cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
			keyE.EncodeByteString([]byte(":method"))
			valueE.EncodeByteString([]byte("GET"))
		}),
		cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
			keyE.EncodeByteString([]byte(":url"))
			valueE.EncodeByteString([]byte(e.requestUri.String()))
		}),
	}
}

func (e *Exchange) encodeRequest(enc *cbor.Encoder) error {
	mes := e.encodeRequestCommon(enc)
	return enc.EncodeMap(mes)
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
	for name, value := range e.requestHeaders {
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
			keyE.EncodeByteString([]byte(":status"))
			valueE.EncodeByteString([]byte(strconv.Itoa(e.responseStatus)))
		}),
	}
	for name, value := range e.responseHeaders {
		mes = append(mes,
			cbor.GenerateMapEntry(func(keyE *cbor.Encoder, valueE *cbor.Encoder) {
				keyE.EncodeByteString([]byte(strings.ToLower(name)))
				valueE.EncodeByteString([]byte(normalizeHeaderValues(value)))
			}))
	}
	return enc.EncodeMap(mes)
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
	if len(cborBytes) >= 1 << 24 {
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
	if _, err := w.Write(e.payload); err != nil {
		return err
	}

	// FIXME: Support "trailer"

	return nil
}
