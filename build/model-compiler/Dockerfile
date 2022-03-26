# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v2 as build

ENV GO111MODULE=on
COPY . /go/src/github.com/onosproject/config-models
RUN cd /go/src/github.com/onosproject/config-models && make build && \
    go get github.com/openconfig/ygot/generator@v0.14.0

FROM alpine:3.14
RUN apk add libc6-compat libc-dev gcc libxml2-dev libxslt-dev python3-dev py3-pip && pip3 install pyang

COPY --from=build /go/src/github.com/onosproject/config-models/build/_output/model-compiler /usr/local/bin/model-compiler
COPY --from=build /go/src/github.com/onosproject/config-models/templates /var/model-compiler/templates
COPY --from=build /go/bin/generator /usr/local/bin/generator

WORKDIR /var/model-compiler

ENTRYPOINT ["model-compiler"]
