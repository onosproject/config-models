module github.com/onosproject/config-models/models/testdevice/v2

go 1.16

require (
	github.com/SeanCondon/xpath v0.0.0-20220217125907-c2b75876708f
	github.com/ghodss/yaml v1.0.0
	github.com/onosproject/config-models v0.9.11
	github.com/onosproject/onos-api/go v0.8.36
	github.com/onosproject/onos-config v0.10.12
	github.com/onosproject/onos-lib-go v0.8.1
	github.com/openconfig/gnmi v0.0.0-20210914185457-51254b657b7d
	github.com/openconfig/goyang v0.3.1
	github.com/openconfig/ygot v0.12.5
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.41.0
)

replace github.com/onosproject/config-models => ../..
