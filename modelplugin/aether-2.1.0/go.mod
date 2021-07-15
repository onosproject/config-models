module github.com/onosproject/config-models/modelplugin/aether-2.1.0

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/onosproject/config-models v0.6.9
	github.com/openconfig/gnmi v0.0.0-20200508230933-d19cebf5e7be
	github.com/openconfig/goyang v0.2.5
	github.com/openconfig/ygot v0.11.2
)

replace github.com/onosproject/config-models => ../../
