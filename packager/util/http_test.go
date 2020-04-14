package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuotedString(t *testing.T) {
	valueFrom := func(s string, _ error) string { return s }
	errorFrom := func(_ string, err error) error { return err }

	assert.EqualError(t, errorFrom(QuotedString("abc\ndef")), "contains non-printable char")
	assert.Equal(t, `"abc"`, valueFrom(QuotedString("abc")))
	assert.Equal(t, `"abc\"\\"`, valueFrom(QuotedString(`abc"\`)))
}
