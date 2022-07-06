module github.com/ampproject/amppackager

go 1.13

require (
	github.com/WICG/webpackage 46ccff500038
	github.com/ampproject/amphtml 34fbeda57c3c
	github.com/go-acme/lego/v4 v4.8.0
	github.com/gofrs/flock v0.8.1
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8
	github.com/kylelemons/godebug v1.1.0
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.1.0
	github.com/prometheus/client_golang v1.12.2
	github.com/prometheus/common v0.35.0
	github.com/stretchr/testify v1.8.0
	github.com/twifkak/crypto v0.0.0-20210326012946-1fce8924335d
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	golang.org/x/net c90051bbdb60
	google.golang.org/grpc v1.47.0
	gopkg.in/square/go-jose.v2 v2.6.0
)

replace github.com/WICG/webpackage v0.0.0-20220624124119-d26bc590e1b5 => github.com/WICG/webpackage v0.0.0-20220530033255-ba99f6be9166
