# Code generated by model-compiler. DO NOT EDIT.

DOCKER_REPOSITORY ?= onosproject/
VERSION           ?= $(shell cat ./VERSION)-{{ .Name }}-{{ .Version }}
LATEST_VERSION           ?= $(shell cat ./VERSION)-{{ .Name }}-latest
KIND_CLUSTER_NAME ?= kind

export CGO_ENABLED=0
export GO111MODULE=on

all: help

help: # @HELP Print the command options
	@echo
	@echo "\033[0;31m    Model Plugin: {{ .Name }} \033[0m"
	@echo
	@grep -E '^.*: .* *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": .* *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '

image: openapi # @HELP Build the docker image (available parameters: DOCKER_REPOSITORY, VERSION)
	docker build . -t ${DOCKER_REPOSITORY}{{ .ArtifactName }}:${VERSION}
	docker tag ${DOCKER_REPOSITORY}{{ .ArtifactName }}:${VERSION} ${DOCKER_REPOSITORY}{{ .ArtifactName }}:${LATEST_VERSION}

build: # @HELP Build the executable (available parameters: VERSION)
	go mod tidy
	go build -o _bin/{{ .Name }} ./plugin

.PHONY: openapi
openapi: # @HELP Generate OpenApi specs
	go mod download
	go mod tidy
	go run openapi/openapi-gen.go -o openapi.yaml

test: build # @HELP Run the unit tests
	go test ./...

publish: image  # @HELP Builds and publish the docker image (available parameters: DOCKER_REPOSITORY, VERSION)
	docker push ${DOCKER_REPOSITORY}{{ .ArtifactName }}:${VERSION}
	docker push ${DOCKER_REPOSITORY}{{ .ArtifactName }}:${LATEST_VERSION}

kind-only: # @HELP Loads the docker image into the kind cluster  (available parameters: KIND_CLUSTER_NAME, DOCKER_REPOSITORY, VERSION)
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image --name ${KIND_CLUSTER_NAME} ${DOCKER_REPOSITORY}{{ .ArtifactName }}:${VERSION}

kind: # @HELP build the docker image and loads it into the currently configured kind cluster (available parameters: KIND_CLUSTER_NAME, DOCKER_REPOSITORY, VERSION)
kind: image kind-only