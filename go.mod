module github.com/ampproject/amppackager

go 1.13

require (
	github.com/WICG/webpackage v0.0.0-20190215052515-70386c3750f2
	github.com/ampproject/amphtml v0.0.0-20180912232012-d3df64d07ae9
	github.com/go-acme/lego/v3 v3.2.0
	github.com/gofrs/flock v0.7.1
	github.com/golang/protobuf v1.3.2
	github.com/google/go-cmp v0.3.0
	github.com/pelletier/go-toml v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/pquerna/cachecontrol v0.0.0-20180306154005-525d0eb5f91d
	github.com/prometheus/client_golang v1.1.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net v0.0.0-20190930134127-c5a3c61f89f3
	google.golang.org/grpc v1.20.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/square/go-jose.v2 v2.3.1
)

replace github.com/davecgh/go-spew => github.com/davecgh/go-spew v1.1.0

replace github.com/stretchr/testify => github.com/stretchr/testify v1.2.1

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac

replace golang.org/x/net => golang.org/x/net v0.0.0-20180808004115-f9ce57c11b24

replace golang.org/x/text => golang.org/x/text v0.3.0
