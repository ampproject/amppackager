module github.com/exoscale/egoscale

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/deepmap/oapi-codegen v1.6.1
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/jarcoal/httpmock v1.0.6
	github.com/pkg/errors v0.9.1
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

go 1.16

retract v1.19.0 // Published accidentally.
