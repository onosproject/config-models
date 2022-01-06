# config-models
This repository contains configuration model compiler and a few sample configuration models.

The config model compiler operates as a docker image that can be run on a volume where the model
meta data and the desired model YANG files are contained as inputs. The compiler will produce a set
of outputs from which a docker image can be built and can be started as a sidecar to `onos-config`.
The outputs include generated Go code used for validating configurations, generated main and NB API 
that is used by the `onos-config`.

# Building model compiler image
To build the configuration model compiler, run:
```
> make model-compiler-docker 
```
The above will compile and assemble the compiler docker image.

# Sample models

The repository also include several sample configuration models:
* 
* devicesim-1.0.0
* testdevice-1.0.0
* testdevice-2.0.0

These are located in the `models` directory and demonstrate the structure of the configuration models.

To generate the artifacts for the `devicesim-1.0.0 configuration model, run the following:
```
> docker run -v $(pwd)/models/devicesim-1.0.0:/config-model onosproject/model-compiler:latest
```

Afterwards, to compile and assemble the configuration model docker image, simply run:
```
> cd models/devicesim-1.0.0 && make
```
