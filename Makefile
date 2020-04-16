export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

ONOS_CONFIG_VERSION := latest
ONOS_BUILD_VERSION := v0.6.0

build/_output/copylibandstay: # @HELP build the copylibandstay utility
	CGO_ENABLED=1 go build -o build/_output/copylibandstay github.com/onosproject/config-models/cmd

build/_output/testdevice.so.1.0.0: # @HELP build testdevice.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/testdevice.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/testdevice-1.0.0

build/_output/testdevice.so.2.0.0: # @HELP build testdevice.so.2.0.0
	CGO_ENABLED=1 go build -o build/_output/testdevice.so.2.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/testdevice-2.0.0

build/_output/devicesim.so.1.0.0: # @HELP build devicesim.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/devicesim.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/devicesim-1.0.0

build/_output/stratum.so.1.0.0: # @HELP build stratum.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/stratum.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/stratum-1.0.0

PHONY:build
build: # @HELP build all libraries
build: build/_output/copylibandstay build/_output/testdevice.so.1.0.0 build/_output/testdevice.so.2.0.0 build/_output/devicesim.so.1.0.0 build/_output/stratum.so.1.0.0

PHONY: config-plugin-docker-testdevice-1.0.0
config-plugin-docker-testdevice-1.0.0: # @HELP build testdevice 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=testdevice \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/config-model-testdevice-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-testdevice-2.0.0
config-plugin-docker-testdevice-2.0.0: # @HELP build testdevice 2.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=testdevice \
		--build-arg PLUGIN_MAKE_VERSION=2.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/config-model-testdevice-2.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-devicesim-1.0.0
config-plugin-docker-devicesim-1.0.0: # @HELP build devicesim 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=devicesim \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/config-model-devicesim-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-stratum-1.0.0
config-plugin-docker-stratum-1.0.0: # @HELP build stratum 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=stratum \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/config-model-stratum-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: images
images: config-plugin-docker-testdevice-1.0.0 config-plugin-docker-testdevice-2.0.0 config-plugin-docker-devicesim-1.0.0 config-plugin-docker-stratum-1.0.0

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/config-model-testdevice-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image onosproject/config-model-testdevice-2.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image onosproject/config-model-devicesim-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image onosproject/config-model-stratum-1.0.0:${ONOS_CONFIG_VERSION}

all: # @HELP build all libraries and all docker images
all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} \
		onosproject/config-model-testdevice-1.0.0 \
		onosproject/config-model-testdevice-2.0.0 \
		onosproject/config-model-devicesim-1.0.0 \
		onosproject/config-model-stratum-1.0.0

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor
	go clean -testcache github.com/onosproject/config-models/...

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
