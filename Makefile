export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

KIND_CLUSTER_NAME   ?= kind
DOCKER_REPOSITORY   ?= onosproject/
ONOS_CONFIG_VERSION ?= latest
ONOS_BUILD_VERSION  := v0.6.3

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

build/_output/e2node.so.1.0.0: # @HELP build e2node.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/e2node.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/e2node-1.0.0

build/_output/rbac.so.1.0.0: # @HELP build rbac.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/rbac.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/rbac-1.0.0

build/_output/aether.so.1.0.0: # @HELP build aether.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/aether.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/aether-1.0.0

build/_output/aether.so.2.0.0: # @HELP build aether.so.2.0.0
	CGO_ENABLED=1 go build -o build/_output/aether.so.2.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/aether-2.0.0	

build/_output/ric.so.1.0.0: # @HELP build ric.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/ric.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/ric-1.0.0

build/_output/enterprise.so.1.0.0: # @HELP build enterprise.so.1.0.0
	CGO_ENABLED=1 go build -o build/_output/enterprise.so.1.0.0 -buildmode=plugin github.com/onosproject/config-models/modelplugin/enterprise-1.0.0



linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/)"

PHONY:build
build: # @HELP build all libraries
build: linters license_check gofmt \
    build/_output/copylibandstay \
    build/_output/testdevice.so.1.0.0 \
    build/_output/testdevice.so.2.0.0 \
    build/_output/devicesim.so.1.0.0 \
    build/_output/stratum.so.1.0.0 \
    build/_output/e2node.so.1.0.0 \
    build/_output/rbac.so.1.0.0 \
    build/_output/aether.so.1.0.0 \
    build/_output/ric.so.1.0.0 \
    build/_output/enterprise.so.1.0.0

PHONY: config-plugin-docker-testdevice-1.0.0
config-plugin-docker-testdevice-1.0.0: # @HELP build testdevice 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=testdevice \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-testdevice-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-testdevice-2.0.0
config-plugin-docker-testdevice-2.0.0: # @HELP build testdevice 2.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=testdevice \
		--build-arg PLUGIN_MAKE_VERSION=2.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-testdevice-2.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-devicesim-1.0.0
config-plugin-docker-devicesim-1.0.0: # @HELP build devicesim 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=devicesim \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-devicesim-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-stratum-1.0.0
config-plugin-docker-stratum-1.0.0: # @HELP build stratum 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=stratum \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-stratum-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-e2node-1.0.0
config-plugin-docker-e2node-1.0.0: # @HELP build e2node 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2node \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-e2node-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-rbac-1.0.0
config-plugin-docker-rbac-1.0.0: # @HELP build rbac 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=rbac \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-rbac-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-aether-1.0.0
config-plugin-docker-aether-1.0.0: # @HELP build aether 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=aether \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-aether-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-aether-2.0.0
config-plugin-docker-aether-2.0.0: # @HELP build aether 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=aether \
		--build-arg PLUGIN_MAKE_VERSION=2.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-aether-2.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor


PHONY: config-plugin-docker-ric-1.0.0
config-plugin-docker-ric-1.0.0: # @HELP build ric 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=ric \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-ric-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: config-plugin-docker-enterprise-1.0.0
config-plugin-docker-enterprise-1.0.0: # @HELP build enterprise 1.0.0 plugin Docker image
	@go mod vendor
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=enterprise \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t ${DOCKER_REPOSITORY}config-model-enterprise-1.0.0:${ONOS_CONFIG_VERSION}
	@rm -rf vendor

PHONY: images
images: config-plugin-docker-testdevice-1.0.0 \
        config-plugin-docker-testdevice-2.0.0 \
        config-plugin-docker-devicesim-1.0.0 \
        config-plugin-docker-stratum-1.0.0 \
        config-plugin-docker-e2node-1.0.0 \
        config-plugin-docker-rbac-1.0.0 \
        config-plugin-docker-aether-1.0.0 \
		config-plugin-docker-aether-2.0.0 \
		config-plugin-docker-ric-1.0.0 \
		config-plugin-docker-enterprise-1.0.0


kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images kind-only

kind-only:
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image ${DOCKER_REPOSITORY}config-model-testdevice-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-testdevice-2.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-devicesim-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-stratum-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-e2node-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-rbac-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-aether-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-aether-2.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-ric-1.0.0:${ONOS_CONFIG_VERSION}
	kind load docker-image ${DOCKER_REPOSITORY}config-model-enterprise-1.0.0:${ONOS_CONFIG_VERSION}


all: # @HELP build all libraries and all docker images
all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} \
		${DOCKER_REPOSITORY}config-model-testdevice-1.0.0 \
		${DOCKER_REPOSITORY}config-model-testdevice-2.0.0 \
		${DOCKER_REPOSITORY}config-model-devicesim-1.0.0 \
		${DOCKER_REPOSITORY}config-model-stratum-1.0.0 \
		${DOCKER_REPOSITORY}config-model-e2node-1.0.0 \
		${DOCKER_REPOSITORY}config-model-rbac-1.0.0 \
		${DOCKER_REPOSITORY}config-model-aether-1.0.0 \
		${DOCKER_REPOSITORY}config-model-aether-2.0.0 \
		${DOCKER_REPOSITORY}config-model-ric-1.0.0 \
		${DOCKER_REPOSITORY}config-model-enterprise-1.0.0


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
