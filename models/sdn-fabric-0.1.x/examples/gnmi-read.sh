#!/bin/sh

#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

ENCODING=${ENCODING:=PROTO}
CERTS=../../../../onos-cli/pkg/certs
GNMI_ARGS="--address localhost:5150 -timeout 5s -en ${ENCODING} -alsologtostderr -insecure -client_crt ${CERTS}/client1.crt -client_key ${CERTS}/client1.key -ca_crt ${CERTS}/onfca.crt"

gnmi_get_files=(
switch-model/get-switch-model-example.gnmi
switch-model/get-switch-model-port-1-0-example.gnmi
switch-model/get-switch-model-port-star-example.gnmi
switch-model/get-switch-model-ports-example.gnmi
switch/get-switch-example.gnmi
switch/get-switch-vlans-display-name-example.gnmi
switch/get-switch-vlans-example.gnmi
switch/get-switch-ports-example.gnmi
dhcp-server/get-dhcp-example.gnmi
)

for gnmi_get in ${gnmi_get_files[@]}; do
  echo "\n\nReading with $gnmi_get. Writing to $gnmi_get.out"
  gnmi_cli ${GNMI_ARGS} -get -proto "$(cat $gnmi_get)" > $gnmi_get.out
done
