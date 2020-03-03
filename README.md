# config-models
Model Plugins for `onos-config` YANG Files compiled with YGOT loadable as Shared Object Libraries

## Building the modules
The plugins can be built individually (use `make help` to see the target names),
or altogether with `make images`.

## Running
The plugins are loaded in Kubernetes by declaring them 

1. as additional (sidecar) containers in the `onos-config` Pod AND
2. as a -modelPlugin option on the startup of `onos-config`

These are usually specified in the `onos-config` [Helm Chart](https://github.com/onosproject/onos-helm-charts/tree/master/onos-config)
or in [onit](https://github.com/onosproject/onos-test).

As the plugin containers startup in Kubernetes the default command is `copylibandstay`.
This is a tiny program that copies the library in to a shared folder where
`onos-config` can refer to it as a file e.g. `/usr/local/lib/shared/devicesim.so.1.0.0`

## Extending
See the guide [Extending onos-config with Model Plugins](https://docs.onosproject.org/onos-config/docs/modelplugin/) for more details.

## Troubleshooting
There are some tips at the end of the Extending guide [Troubleshooting](https://docs.onosproject.org/onos-config/docs/modelplugin/#troubleshooting)

## Additional notes
The first 3 plugins

* testdevice-1.0.0
* testdevice-2.0.0
* devicesim-1.0.0

are used as Go module packages in their own right from inside the Unit tests of `onos-config`.
For this reason each one has it's own `go.mod`.

The plugin

* stratum-1.0.0

has its own `go.mod` to that it is ignored by the top level `go.mod`.
> Also in future it is intended that the `stratum_1_0_0` package will be generated at compile time from YANG, (using YGOT
>Generate) and removed from Git versioning

The main purpose of the top level `go.mod` is to provide a set of go modules (as **vendor**) to the Docker build container.

The versions of the packages in this top level `go.mod` must exactly match those of `onos-config`'s `go.mod`, or the plugin
will not be loadable by `onos-config` at run time.