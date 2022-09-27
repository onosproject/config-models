module github.com/onosproject/config-models/models/sdn-fabric-0.1.x

go 1.16

require (
	github.com/SeanCondon/xpath v0.0.0-20220821123841-6149b14eb04f
	github.com/getkin/kin-openapi v0.20.0
	github.com/ghodss/yaml v1.0.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/onosproject/config-models v0.10.38
	github.com/onosproject/onos-api/go v0.9.46
	github.com/onosproject/onos-lib-go v0.9.3
	github.com/openconfig/gnmi v0.0.0-20220617175856-41246b1b3507
	github.com/openconfig/goyang v1.1.0
	github.com/openconfig/ygot v0.24.4
	github.com/stretchr/testify v1.7.1
	google.golang.org/grpc v1.46.0
)

replace github.com/onosproject/config-models => ../..
