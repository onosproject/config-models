# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

# syntax=docker/dockerfile:1.3

FROM onosproject/golang-build:v1.3.0 as build

ENV GO111MODULE=on

RUN --mount=type=cache,target=/root/.cache/go-build \
    go install github.com/openconfig/ygot/generator@v0.26.2
COPY . /go/src/github.com/onosproject/config-models
WORKDIR /go/src/github.com/onosproject/config-models
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -v -mod=vendor -o build/_output/model-compiler ./cmd/model-compiler

FROM alpine:3.17.2
RUN apk add --no-cache libc6-compat=1.2.3-r5 libc-dev=0.7.2-r3 gcc=12.2.1_git20220924-r4 libxml2-dev=2.10.4-r0 \
    libxslt-dev=1.1.37-r1 python3-dev=3.10.11-r0 py3-wheel=0.38.4-r0 py3-pip=22.3.1-r1 && \
    pip3 install --no-cache-dir pyang==2.5.3 openapi-spec-validator==0.5.5

COPY --from=build /go/src/github.com/onosproject/config-models/build/_output/model-compiler /usr/local/bin/model-compiler
COPY --from=build /go/src/github.com/onosproject/config-models/templates /var/model-compiler/templates
COPY --from=build /go/src/github.com/onosproject/config-models/yang-base /var/model-compiler/yang-base
COPY --from=build /go/bin/generator /usr/local/bin/generator
COPY pyang/plugins/*.py /usr/lib/python3.10/site-packages/pyang/plugins/

WORKDIR /var/model-compiler

ENTRYPOINT ["model-compiler"]
