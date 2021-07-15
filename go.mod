module github.com/ampproject/amppackager

go 1.13

require (
	github.com/WICG/webpackage v0.0.0-20190215052515-70386c3750f2
	github.com/ampproject/amphtml v0.0.0-20180912232012-d3df64d07ae9
	github.com/go-acme/lego/v4 v4.4.0
	github.com/gofrs/flock v0.7.1
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.5
	github.com/kylelemons/godebug v1.1.0
	github.com/pelletier/go-toml v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.0.0-20180306154005-525d0eb5f91d
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.29.0
	github.com/stretchr/testify v1.7.0
	github.com/twifkak/crypto v0.0.0-20210326012946-1fce8924335d
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	google.golang.org/grpc v1.31.0
	gopkg.in/square/go-jose.v2 v2.5.1
)

replace github.com/davecgh/go-spew => github.com/davecgh/go-spew v1.1.0

replace github.com/stretchr/testify => github.com/stretchr/testify v1.2.1

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac

replace golang.org/x/net => golang.org/x/net v0.0.0-20180808004115-f9ce57c11b24

replace golang.org/x/text => golang.org/x/text v0.3.0
