# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

# syntax=docker/dockerfile:1.3

FROM onosproject/golang-build:v1.3.0 as build

ENV GO111MODULE=on

RUN --mount=type=cache,target=/root/.cache/go-build \
    go install github.com/openconfig/ygot/generator@v0.24.4
COPY . /go/src/github.com/onosproject/config-models
WORKDIR /go/src/github.com/onosproject/config-models
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -v -mod=vendor -o build/_output/model-compiler ./cmd/model-compiler

FROM alpine:3.17
RUN apk add libc6-compat libc-dev gcc libxml2-dev libxslt-dev python3-dev py3-wheel py3-pip && pip3 install pyang

COPY --from=build /go/src/github.com/onosproject/config-models/build/_output/model-compiler /usr/local/bin/model-compiler
COPY --from=build /go/src/github.com/onosproject/config-models/templates /var/model-compiler/templates
COPY --from=build /go/src/github.com/onosproject/config-models/yang-base /var/model-compiler/yang-base
COPY --from=build /go/bin/generator /usr/local/bin/generator

WORKDIR /var/model-compiler

ENTRYPOINT ["model-compiler"]
