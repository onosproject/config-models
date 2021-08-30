module github.com/onosproject/config-models/modelplugin/aether-2.1.0

go 1.16

require (
	github.com/ghodss/yaml v1.0.0
	github.com/onosproject/config-models v0.6.9
	github.com/openconfig/gnmi v0.0.0-20210707145734-c69a5df04b53
	github.com/openconfig/goyang v0.2.9
	github.com/openconfig/ygot v0.12.0
)

replace github.com/onosproject/config-models => ../../
