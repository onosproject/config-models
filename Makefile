# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

SHELL = bash -e -o pipefail
export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

KIND_CLUSTER_NAME   ?= kind
MODEL_COMPILER_VERSION ?= latest
PLATFORM ?= --platform linux/x86_64

mod-update: # @HELP Download the dependencies to the vendor folder
	go mod tidy
	go mod vendor

build: # @HELP build all libraries
build:
	go build -mod=vendor -o build/_output/model-compiler ./cmd/model-compiler

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

test: # @HELP run go test on projects
test: mod-update build linters license gofmt images models models-version-check
	go test ./pkg/...
	@bash test/generated.sh
	@cd models && for model in *; do pushd $$model; make test; popd; done

.PHONY: models
models: # @HELP make demo and test device models
models:
	@cd models && for model in *; do echo "Generating $$model:"; docker run ${PLATFORM} -v $$(pwd)/$$model:/config-model onosproject/model-compiler:latest; done

models-openapi: # @HELP generates the openapi specs for the models
	@cd models && for model in *; do echo -e "Building OpenApi Specs for $$model:\n"; pushd $$model; make openapi; popd; echo -e "\n\n"; done

models-images: models models-openapi # @HELP Build Docker containers for all the models
	@cd models && for model in *; do echo -e "Building container for $$model:\n"; pushd $$model; make image; popd; echo -e "\n\n"; done

models-version-check:
	# TODO this fails as the output of the ygot generation has some variablity (see https://jira.opennetworking.org/browse/SDRAN-1473)
	@cd models && for model in *; do echo -e "Validating VERSION for $$model:\n"; pushd $$model; bash ../../test/model-version.sh $$model; popd; echo -e "\n\n"; done

docker-login:
ifdef DOCKER_USER
ifdef DOCKER_PASSWORD
	echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USER} --password-stdin
else
	@echo "DOCKER_USER is specified but DOCKER_PASSWORD is missing"
	@exit 1
endif
endif

kind-models:
	@cd models && for model in *; do pushd $$model; make kind; popd; done

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: deps build linters license images models
	go test ./pkg/...
	# TODO add test/generated.sh once the ygot issue is resolved (https://jira.opennetworking.org/browse/SDRAN-1473)
	@cd models && for model in *; do pushd $$model; make test; popd; done

all: # @HELP build all libraries
all: build

model-compiler-docker: mod-update # @HELP build model-compiler Docker image
	docker build ${PLATFORM} . -t onosproject/model-compiler:${MODEL_COMPILER_VERSION} -f build/model-compiler/Dockerfile

images: model-compiler-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/model-compiler:${MODEL_COMPILER_VERSION}

publish: # @HELP publish version on github (called by release-merge-commit)
	./build/build-tools/publish-version ${VERSION} onosproject/model-compiler

jenkins-publish: docker-login # @HELP Jenkins calls this to publish artifacts
	make -C models/ric-1.x publish
	make -C models/e2node-1.x publish
	make -C models/devicesim-1.0.x publish
	make -C models/testdevice-1.0.x publish
	make -C models/testdevice-2.0.x publish
	./build/build-tools/release-merge-commit

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor
	go clean -testcache github.com/onosproject/config-models/...

