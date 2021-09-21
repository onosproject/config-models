#!/usr/bin/env bash

mkdir -p aether_3_0_0

OPENCORD="${OPENCORD:-$HOME/git/opencord}"

go run github.com/openconfig/ygot/generator -path=$OPENCORD/roc-helm-charts/config-models/aether-3.x/files/yang \
       -output_file=aether_3_0_0/generated.go -package_name=aether_3_0_0 -generate_fakeroot --include_descriptions \
       enterprise.yang connectivity-service.yang \
       aether-types.yang application.yang device-group.yang ip-domain.yang \
       site.yang upf.yang ap-list.yang template.yang vcs.yang \
       traffic-class.yang


sedi=(-i)
case "$(uname)" in
  Darwin*) sedi=(-i "")
esac

lf=$'\n'; sed "${sedi[@]}" -e "1s/^/\/\/ Code generated by YGOT. DO NOT EDIT.\\$lf/" aether_3_0_0/generated.go
