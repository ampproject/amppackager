module github.com/ampproject/amppackager

go 1.13

require (
	github.com/WICG/webpackage d26bc590e1b5
	github.com/ampproject/amphtml 9e920c549ea5
	github.com/go-acme/lego/v4 v4.7.0
	github.com/gofrs/flock v0.8.1
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8
	github.com/kylelemons/godebug v1.1.0
	github.com/miekg/dns v1.1.47 // indirect
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.1.0
	github.com/prometheus/client_golang v1.12.2
	github.com/prometheus/common v0.35.0
	github.com/stretchr/testify v1.7.5
	github.com/twifkak/crypto v0.0.0-20210326012946-1fce8924335d
	golang.org/x/crypto 05595931fe9d
	golang.org/x/net 1bab6f366d9e
	google.golang.org/grpc v1.47.0
	gopkg.in/square/go-jose.v2 v2.6.0
)

replace github.com/WICG/webpackage v0.0.0-20220621081514-83e9041e8f1f => github.com/WICG/webpackage v0.0.0-20220530033255-ba99f6be9166
