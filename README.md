# config-models
Model Plugins for `onos-config` YANG Files compiled with YGOT loadable as Shared Object Libraries

## Building the modules
The plugins are built from the Makefile of `onos-config`.

The build cannot be done from this folder, because Go plugins cannot be loaded in
to a Main module if they are built under a different GOPATH than the main module

> the error `plugin was built with a different version of package internal/cpu`
> (or similar) can be expected, if this is not the case.

> See https://github.com/golang/go/issues/26759

## Running
The plugins are loaded in Kubernetes by declaring them as additional containers
in the `onos-config` Pod - this is specified in the `Helm Chart` for `onos-config`
or in `onit`.

As the plugin containers startup in Kubernetes the default command is `copylibandstay`.
This is a tiny program that copies the library in to a shared folder where
`onos-config` can refer to it as a file e.g. `/usr/local/lib/shared/devicesim.so.1.0.0`
