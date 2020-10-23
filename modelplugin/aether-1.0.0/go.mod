module github.com/onosproject/config-models/modelplugin/aether-1.0.0

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.3.3
	github.com/onosproject/config-models v0.6.9
	github.com/openconfig/gnmi v0.0.0-20190823184014-89b2bf29312c
	github.com/openconfig/goyang v0.0.0-20200115183954-d0a48929f0ea
	github.com/openconfig/ygot v0.6.1-0.20200103195725-e3c44fa43926
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)

replace github.com/onosproject/config-models => ../../
