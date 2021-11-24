FROM onosproject/golang-build:v1.0 as build

ENV GO111MODULE=on
COPY . /go/src/github.com/config-models/models/{{ .Model.Name }}_{{ .Model.Version | replace "." "_" }}
RUN cd /go/src/github.com/config-models/models/{{ .Model.Name }}_{{ .Model.Version | replace "." "_" }} && GOFLAGS=-mod=vendor make build

FROM alpine:3.11
RUN apk add libc6-compat

USER nobody

COPY --from=build /go/src/github.com/config-models/models/{{ .Model.Name }}_{{ .Model.Version | replace "." "_" }}/build/_output/{{ .Model.Name }}-{{ .Model.Version }} /usr/local/bin/{{ .Model.Name }}-{{ .Model.Version }}

ENTRYPOINT ["{{ .Model.Name }}-{{ .Model.Version }}"]
