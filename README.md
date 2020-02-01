# config-models
Model Plugins for onos-config YANG Files compiled with YGOT loadable as Shared Object Libraries

For each plugin, calling `bazel build <plugin target>` causes this to:

* uses **YGOT** `generate` to compile the YANG files in to Go code `generated.go`
* compiles this `generated.go` along with `modelmain.go` in to a lib `.a` file
* links the `.a` file in to a `.so` shared library
* creates a Docker image containing only the shared library in to a `tar` file,
which may be loaded in to Docker with `docker import`

> Note: there's no Dockerfile included here

To build each model
```bash
bazel build //modelplugin/devicesim-1.0.0 --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```
or
```bash
bazel build //modelplugin/stratum-1.0.0 --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```

It is suggested to use the name image name `onosproject/config-model-<plugin>-<ver>:latest` when loading in to docker.
```bash
docker import bazel-bin/modelplugin/devicesim-1.0.0/devicesim-1.0.0-layer.tar onosproject/config-model-devicesim-1-0-0:latest
```
or
```bash
docker import bazel-bin/modelplugin/stratum-1.0.0/stratum-1.0.0-layer.tar onosproject/config-model-stratum-1-0-0:latest
```

> It is possible to build and load in one action (using `bazel run` instead of
>`bazel build` above) but the label of the image in Docker will be in the style
> `bazel/modelplugin/stratum-1.0.0:stratum-1.0.0` which is not a good
>naming scheme

## Troubleshooting
### gnmi_ext
There is currently a problem with `gnmi_ext` in running `bazel build`.
The workaround is to copy the `gnmi_ext` folder in `external/com_github_openconfig_gnmi`
to where Bazel expects to find it.

First find the name of your bazel cache:
```bash
bazel info output_base
```
Then change directory to the `gnmi` folder under this, and copy the folder
```bash
cd <bazel_output_base>/external/com_github_openconfig_gnmi
mkdir -p github.com/openconfig/gnmi/proto
cp -R proto/gnmi_ext/ github.com/openconfig/gnmi/proto/
```

### go_image
An note in the documentation says that `go_image` cannot be run on Mac. This has not been tested.
An alternative is to use `go_binary` and then `container_image` targets if necessary.

[https://github.com/bazelbuild/rules_docker#language-rules](https://github.com/bazelbuild/rules_docker#language-rules)
