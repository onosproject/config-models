export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

KIND_CLUSTER_NAME   ?= kind
DOCKER_REPOSITORY   ?= onosproject/
ONOS_CONFIG_VERSION ?= latest
ONOS_BUILD_VERSION  := v0.6.3

build/_output/copylibandstay: # @HELP build the copylibandstay utility
	CGO_ENABLED=1 go build -o build/_output/copylibandstay github.com/onosproject/config-models/cmd

linters: golang-ci # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.36.0

license_check: build-tools # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/)"

test: # @HELP run go test on projects
test: build linters license_check gofmt
	cd modelplugin/testdevice-1.0.0/ && (go test ./... || cd ..)

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build-tools deps license_check linters
	GODEBUG=cgocheck=0 TEST_PACKAGES=github.com/onosproject/onos-e2t/... ./../build-tools/build/jenkins/make-unit

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
