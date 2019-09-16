module github.com/ampproject/amppackager

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/WICG/webpackage v0.0.0-20190215052515-70386c3750f2
	github.com/ampproject/amphtml v0.0.0-20180912232012-d3df64d07ae9
	github.com/gofrs/flock v0.7.1
	github.com/golang/protobuf v1.3.0
	github.com/google/go-cmp v0.2.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pelletier/go-toml v1.1.0
	github.com/pkg/errors v0.8.0
	github.com/pquerna/cachecontrol v0.0.0-20180306154005-525d0eb5f91d
	github.com/stretchr/testify v1.2.2
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace github.com/davecgh/go-spew => github.com/davecgh/go-spew v1.1.0

replace github.com/stretchr/testify => github.com/stretchr/testify v1.2.1

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac

replace golang.org/x/net => golang.org/x/net v0.0.0-20180808004115-f9ce57c11b24

replace golang.org/x/text => golang.org/x/text v0.3.0
