FROM golang:1.16-alpine3.13

RUN apk upgrade --update --no-cache && apk add --update make gcc musl-dev git

RUN mkdir -p /build/deps && \
    cd /build/deps && \
    echo "module deps" > go.mod && \
    go get -u -d github.com/onosproject/onos-config && \
    go get -u -d github.com/openconfig/ygot/generator && \
    go mod download -x && \
    rm -r /build/deps

WORKDIR /build

ENTRYPOINT ["make"]