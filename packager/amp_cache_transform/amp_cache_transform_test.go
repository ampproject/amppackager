package amp_cache_transform

import (
	"testing"

	"github.com/ampproject/amppackager/transformer"
	rpb "github.com/ampproject/amppackager/transformer/request"
	"github.com/stretchr/testify/assert"
)

func header(header string, version int64) string {
	return header
}

func version(header string, version int64) int64 {
	return version
}

func TestShouldSendSXG(t *testing.T) {
	orig := transformer.SupportedVersions
	defer func() { transformer.SupportedVersions = orig }()
	transformer.SupportedVersions = []*rpb.VersionRange{{Max: 1, Min: 1}}

	// Success.
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG("any")))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG("foobar, any")))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG("any, google")))
	// Technically this is a false positive -- trailing whitespace is not
	// allowed by the spec. (This is occurring because the comma parsing
	// code doesn't fail if it finds WS and no comma.) Meh, close enough:
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG("any ")))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG("google")))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG("google, any")))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG("google, foobar")))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG("google,foobar")))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG("google,\tfoobar")))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG("google ,foobar")))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG("google, a*b-c_d/e")))

	// Failure.
	assert.Equal(t, "", header(ShouldSendSXG("")))
	assert.Equal(t, "", header(ShouldSendSXG(" any")))
	assert.Equal(t, "", header(ShouldSendSXG("foobar")))
	assert.Equal(t, "", header(ShouldSendSXG("foobar, baz")))
	assert.Equal(t, "", header(ShouldSendSXG("googleany")))
	assert.Equal(t, "", header(ShouldSendSXG("google;any")))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="1,any`)))
	assert.Equal(t, "", header(ShouldSendSXG("google,123")))
	assert.Equal(t, "", header(ShouldSendSXG("google, eh!")))
	assert.Equal(t, "", header(ShouldSendSXG("ABC,google")))

	// Version spec success.
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG(`google;v="1"`)))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG(`google;v="1..2"`)))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG(`google;v="1,2..3,5"`)))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG(`google ; v="1"`)))

	// Version spec parse failure.
	assert.Equal(t, "", header(ShouldSendSXG(`google;`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v=`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v=`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="\`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v=" 1"`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="1 "`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="1 2"`)))
	assert.Equal(t, "", header(ShouldSendSXG("google;v=1,any")))
	assert.Equal(t, "", header(ShouldSendSXG(`google,v="1",any`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v ="1",any`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v= "1",any`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="\a",any`)))
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="\1",any`)))
	assert.Equal(t, "", header(ShouldSendSXG("google;v=\"\t\",any")))

	// Version spec semantic failure or mismatch.
	assert.Equal(t, "", header(ShouldSendSXG(`google;v="2"`)))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG(`google;v="2",any`)))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG(`google;v="2..frog",any`)))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG(`google;v="-1..1",any`)))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG(`google;v="1,-1",any`)))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG(`google;v="1..2,2..3",any`)))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG(`google;v="1..\"2",any`)))
	assert.Equal(t, `any;v="1"`, header(ShouldSendSXG(`google;x="1",any`)))
	assert.Equal(t, `google;v="1"`, header(ShouldSendSXG(`google;v="2",google`)))
}

func TestShouldSendSXG_Version(t *testing.T) {
	orig := transformer.SupportedVersions
	defer func() { transformer.SupportedVersions = orig }()
	transformer.SupportedVersions = []*rpb.VersionRange{{Max: 2, Min: 1}}

	assert.Equal(t, int64(1), version(ShouldSendSXG(`google;v="1"`)))
	assert.Equal(t, int64(2), version(ShouldSendSXG(`google;v="2"`)))
}
