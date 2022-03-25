<!--
SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# config-models
This repository contains configuration model compiler and a few sample configuration models.

The config model compiler operates as a docker image that can be run on a volume where the model
meta data and the desired model YANG files are contained as inputs. The compiler will produce a set
of outputs from which a docker image can be built and can be started as a sidecar to `onos-config`.
The outputs include generated Go code used for validating configurations, generated main and NB API 
that is used by the `onos-config`.

This structure allows configuration models to be hosted at arbitrary locations, while providing the neccessary
toolchain conveniently contained wihin the compiler docker image.

## Building model compiler image
To build the configuration model compiler, run:
```shell
make model-compiler-docker 
```
The above will compile and assemble the compiler docker image.

# Sample models
The repository also include several sample configuration models located in the `models` directory
and serving to demonstrate the structure of the configuration models.

* `devicesim-1.0.x`
* `testdevice-1.0.x`
* `testdevice-2.0.x`

## Building sample models
Building configuration models is easy. For example, to generate the artifacts for the `devicesim-1.0.0` configuration 
model, run the following:
```shell
docker run -v $(pwd)/models/devicesim-1.0.x:/config-model onosproject/model-compiler:latest
```

Afterwards, to compile and assemble the configuration model docker image, simply run:
```shell
cd models/devicesim-1.0.x && make
```
