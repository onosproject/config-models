module {{ .GoPackage }}

go 1.16

require (
	github.com/ghodss/yaml v1.0.0
	github.com/onosproject/config-models v0.10.15
	github.com/onosproject/onos-api/go v0.9.11
	github.com/onosproject/onos-config v0.10.34
	github.com/onosproject/onos-lib-go v0.8.13
	github.com/openconfig/gnmi v0.0.0-20210914185457-51254b657b7d
	github.com/openconfig/goyang v1.0.0
	github.com/openconfig/ygot v0.22.1
	google.golang.org/grpc v1.41.0
)
