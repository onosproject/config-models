# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

SHELL = bash -e -o pipefail
export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

KIND_CLUSTER_NAME   ?= kind
MODEL_COMPILER_VERSION ?= latest
PLATFORM ?= linux/x86_64
LOCAL_BIN ?= .local/bin

mod-update: # @HELP Download the dependencies to the vendor folder
	go mod tidy
	go mod vendor

build: # @HELP build all libraries
build:
	go build -mod=vendor -o build/_output/model-compiler ./cmd/model-compiler

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

test: # @HELP run go test on projects
test: mod-update build linters license gofmt images models models-openapi models-version-check
	go test ./pkg/...
	@bash test/generated.sh
	@cd models && for model in *; do pushd $$model; make test; popd; done

.PHONY: models
models: # @HELP make demo and test device models
models:
	@for model in models/*; do \
		echo "Generating $$model:"; \
		docker run -v $$(pwd)/$$model:/config-model onosproject/model-compiler:${MODEL_COMPILER_VERSION}; \
	done

models-openapi: # @HELP generates the openapi specs for the models
	@for model in models/*; do \
		echo -e "Building OpenApi Specs for $$model:\n"; \
		make -C $$model openapi; \
		docker run -v $$(pwd)/$$model:/config-model --entrypoint /usr/bin/openapi-spec-validator onosproject/model-compiler:${MODEL_COMPILER_VERSION} /config-model/openapi.yaml; \
	done

# the gNMI client generator is on hold at the moment, disabling it for the moment
#models-gnmi-client: # @HELP generates the gnmi-client for the models
#	@for model in models/*; do echo -e "Building gNMI Client for $$model:\n"; pushd $$model; rm -f api/gnmi_client.go; make gnmi-gen; popd; echo -e "\n\n"; done

models-images: models models-openapi # @HELP Build Docker containers for all the models
	@for model in models/*; do echo -e "Building container for $$model:\n"; pushd $$model; make image; popd; echo -e "\n\n"; done

models-version-check:
	# TODO this fails as the output of the ygot generation has some variablity (see https://jira.opennetworking.org/browse/SDRAN-1473)
	@for model in models/*; do echo -e "Validating VERSION for $$model:\n"; pushd $$model; bash ../../test/model-version.sh $$model; popd; echo -e "\n\n"; done

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
	@for model in models/*; do make -C $$model kind; done

check-models-tag: # @HELP check that the
	@for model in models/*; do make -C $$model check-tag; done

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: deps mod-update build linters check-models-tag images models models-openapi license
	go test ./pkg/...
	# TODO add test/generated.sh once the ygot issue is resolved (https://jira.opennetworking.org/browse/SDRAN-1473)
	@for model in models/*; do make -C $$model test; done

all: # @HELP build all libraries
all: build

hadolint: #Lint the Dockerfile. Install hadolint if not present - on macos it is recommended install tool in advance with "brew install hadolint"
	@hadolint --version || (mkdir -p $$HOME/${LOCAL_BIN} && \
	    curl -L -o $$HOME/${LOCAL_BIN}/hadolint https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-$(shell uname -s)-$(shell uname -m) && \
	    chmod +x $$HOME/${LOCAL_BIN}/hadolint && \
	    echo "hadolint downloaded to $$HOME/${LOCAL_BIN} - ensure this dir is in PATH")
	hadolint build/model-compiler/Dockerfile

model-compiler-docker: hadolint mod-update # @HELP build model-compiler Docker image
	DOCKER_BUILDKIT=1 docker image build --platform ${PLATFORM} . -t onosproject/model-compiler:${MODEL_COMPILER_VERSION} -f build/model-compiler/Dockerfile

images: model-compiler-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/model-compiler:${MODEL_COMPILER_VERSION}

publish: # @HELP publish version on github (called by release-merge-commit)
	./build/build-tools/publish-version ${VERSION} onosproject/model-compiler

jenkins-publish: docker-login # @HELP Jenkins calls this to publish artifacts
	for model in models/*; do make -C $$model publish; done
	./build/build-tools/release-merge-commit

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor
	go clean -testcache github.com/onosproject/config-models/...
	for model in models/*; do make -C $$model clean; done

