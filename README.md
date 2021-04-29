# config-models
Legacy support for Model Plugins for `onos-config`.

This has now been replaced by the `onos-operator` loader for Config Model Plugins.

* The Model Plugins YANG files are now held in the helm Charts at
https://github.com/onosproject/onos-helm-charts/tree/master/config-models
* The Config Model compiler is at
https://github.com/onosproject/onos-config-model

This repo is for holding compiled versions of some model plugins

# aether-1.0.0, aether-2.0.0, aether-2.1.0
These folders hold the compiled version of `generated.go` which allows an **OpenAPI 3** model to
be generated from the compiled YANG. In each case this can be called like e.g.:
```bash
cd modelplugin/aether-2.1.0
./generator.sh
go run cmd/openapi-gen.go
```
> The output is sent to the console - it should be saved to https://github.com/onosproject/aether-roc-api/tree/master/api

# devicesim-1.0.0, testdevice-1.0.0, testdevice-2.0.0
These folders hold the compiled version of `generated.go` which are used in unit tests for `onos-config`
For this reason each one has it's own `go.mod`.

When the YGOT version is changed in these go modules, the version used by `onos-config` will have to be updated too,
and in turn the `onos-config-model` will have to be updated and re-released.

# Top level
The main purpose of the top level `go.mod` is for the common library - **pkg/openapi-gen**
