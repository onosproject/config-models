#!/usr/bin/env bash

mkdir -p plproxy_1_0_0

go run github.com/openconfig/ygot/generator -path=../../../roc-helm-charts/config-models/plproxy-1.x/files/yang \
       -output_file=plproxy_1_0_0/generated.go -package_name=plproxy_1_0_0 -generate_fakeroot --include_descriptions\
       prom-label-proxy.yang


sedi=(-i)
case "$(uname)" in
  Darwin*) sedi=(-i "")
esac

lf=$'\n'; sed "${sedi[@]}" -e "1s/^/\/\/ Code generated by YGOT. DO NOT EDIT.\\$lf/" plproxy_1_0_0/generated.go