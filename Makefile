export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

KIND_CLUSTER_NAME   ?= kind
DOCKER_REPOSITORY   ?= onosproject/
ONOS_CONFIG_VERSION ?= latest
ONOS_BUILD_VERSION  := v0.6.3

build/_output/copylibandstay: # @HELP build the copylibandstay utility
	CGO_ENABLED=1 go build -o build/_output/copylibandstay github.com/onosproject/config-models/cmd

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/)"

test: # @HELP run go test on projects
test: build linters license_check gofmt

PHONY:build
build: # @HELP build all libraries
build: \
	build/_output/copylibandstay
	go build ./...
	cd modelplugin/devicesim-1.0.0/ && (go build ./... || cd ..)
	cd modelplugin/testdevice-1.0.0/ && (go build ./... || cd ..)
	cd modelplugin/testdevice-2.0.0/ && (go build ./... || cd ..)

all: # @HELP build all libraries and all docker images
all: build

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
