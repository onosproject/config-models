module github.com/onosproject/config-models

go 1.16

require (
	github.com/antchfx/xpath v1.2.0
	github.com/getkin/kin-openapi v0.20.0
	github.com/openconfig/goyang v0.3.1
	github.com/openconfig/ygot v0.12.4
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.5.1
	gotest.tools v2.2.0+incompatible
)

replace github.com/openconfig/goyang => ../../openconfig/goyang

replace github.com/openconfig/ygot => ../../openconfig/ygot
