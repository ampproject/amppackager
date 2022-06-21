module github.com/ampproject/amppackager

go 1.13

require (
	github.com/WICG/webpackage 83e9041e8f1f
	github.com/ampproject/amphtml 15781b541aae
	github.com/go-acme/lego/v4 v4.7.0
	github.com/gofrs/flock v0.8.1
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8
	github.com/kylelemons/godebug v1.1.0
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.1.0
	github.com/prometheus/client_golang v1.12.2
	github.com/prometheus/common v0.35.0
	github.com/stretchr/testify v1.7.4
	github.com/twifkak/crypto v0.0.0-20210326012946-1fce8924335d
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	golang.org/x/net 355a448f1bc9
	google.golang.org/grpc v1.47.0
	gopkg.in/square/go-jose.v2 v2.6.0
)

replace github.com/WICG/webpackage v0.0.0-20220613002059-40128674f59f => github.com/WICG/webpackage v0.0.0-20220530033255-ba99f6be9166
