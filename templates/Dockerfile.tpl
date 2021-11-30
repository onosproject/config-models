# Code generated by model-compiler. DO NOT EDIT

FROM onosproject/golang-build:v1.0 as build

ENV GO111MODULE=on
COPY . /models/{{ .Name }}-{{ .Version }}
RUN cd /models/{{ .Name }}-{{ .Version }} && make build

FROM alpine:3.11
RUN apk add libc6-compat

USER nobody

COPY --from=build /models/{{ .Name }}-{{ .Version }}/_bin/{{ .Name }}-{{ .Version }} /usr/local/bin/{{ .Name }}-{{ .Version }}

ENTRYPOINT ["{{ .Name }}-{{ .Version }}"]
