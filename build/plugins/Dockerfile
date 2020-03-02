ARG PLUGIN_BUILD_VERSION=stable

FROM onosproject/golang-build:$PLUGIN_BUILD_VERSION as pluginbuild
ENV GO111MODULE=on
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY Makefile go.mod go.sum /go/src/github.com/onosproject/onos-config/
COPY vendor/ /go/src/github.com/onosproject/onos-config/vendor/

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/config-models/
COPY cmd/ /go/src/github.com/onosproject/config-models/cmd/
COPY vendor/ /go/src/github.com/onosproject/config-models/vendor/

# "go mod vendor" will not copy over anything in a "main" package, so we have
# to get them and copy them over in to vendor folder
RUN cd /go/src/github.com/onosproject/onos-config && \
go mod download github.com/onosproject/config-models/modelplugin/${PLUGIN_MAKE_TARGET}-${PLUGIN_MAKE_VERSION} && \
cp /go/pkg/mod/github.com/onosproject/config-models/modelplugin/${PLUGIN_MAKE_TARGET}-${PLUGIN_MAKE_VERSION}*/modelmain.go vendor/github.com/onosproject/config-models/modelplugin/${PLUGIN_MAKE_TARGET}-${PLUGIN_MAKE_VERSION}

RUN cd /go/src/github.com/onosproject/onos-config && GOFLAGS=-mod=vendor make build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION}

RUN cd /go/src/github.com/onosproject/config-models && GOFLAGS=-mod=vendor make build/_output/copylibandstay

FROM gcr.io/distroless/base-debian10
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY --from=pluginbuild /go/src/github.com/onosproject/onos-config/build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} /
COPY --from=pluginbuild /go/src/github.com/onosproject/config-models/build/_output/copylibandstay /

CMD ["/copylibandstay"]
