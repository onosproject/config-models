module {{ .GoPackage }}

go 1.19

require (
	github.com/SeanCondon/xpath v0.0.0-20221217195644-773fbeaef469
	github.com/getkin/kin-openapi v0.114.0
	github.com/ghodss/yaml v1.0.0
	github.com/onosproject/config-models v0.11.11
	github.com/onosproject/onos-api/go v0.10.31
	github.com/onosproject/onos-lib-go v0.10.17
	github.com/openconfig/gnmi v0.9.1
	github.com/openconfig/goyang v1.4.0
	github.com/openconfig/ygot v0.28.3
	github.com/stretchr/testify v1.8.2
	google.golang.org/grpc v1.54.0
	gotest.tools v2.2.0+incompatible
)
