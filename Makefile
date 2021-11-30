export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

KIND_CLUSTER_NAME   ?= kind
MODEL_COMPILER_VERSION ?= latest

build: # @HELP build all libraries
build:
	go build -o build/_output/model-compiler ./cmd/model-compiler

linters: golang-ci # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.42.0

license_check: build-tools # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/)"

test: # @HELP run go test on projects
test: build linters license_check gofmt
	go test ./pkg/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build-tools deps license_check linters

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./cmd/...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

all: # @HELP build all libraries
all: build

model-compiler-docker: # @HELP build model-compiler Docker image
	docker build . -t onosproject/model-compiler:${MODEL_COMPILER_VERSION} -f build/model-compiler/Dockerfile

publish: # @HELP publish version on github, and PyPI
	./../build-tools/publish-version ${VERSION}

jenkins-publish: build-tools jenkins-tools # @HELP Jenkins calls this to publish artifacts
	../build-tools/release-merge-commit

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
