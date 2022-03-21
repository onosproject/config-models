module {{ .GoPackage }}

go 1.17

require (
	github.com/ghodss/yaml v1.0.0
	github.com/onosproject/config-models v0.9.24
	github.com/onosproject/onos-api/go v0.9.8
	github.com/onosproject/onos-config v0.10.23
	github.com/onosproject/onos-lib-go v0.8.13
	github.com/openconfig/gnmi v0.0.0-20210914185457-51254b657b7d
	github.com/openconfig/goyang v0.4.0
	github.com/openconfig/ygot v0.14.0
	google.golang.org/grpc v1.41.0
)
