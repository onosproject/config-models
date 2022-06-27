#!/bin/sh

#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

CERTS=../../../../onos-cli/pkg/certs
GNMI_ARGS="--address localhost:5150 -timeout 5s -en PROTO -alsologtostderr -insecure -client_crt ${CERTS}/client1.crt -client_key ${CERTS}/client1.key -ca_crt ${CERTS}/onfca.crt"

gnmi_set_files=(
#switch-model/set-switch-model-example.gnmi
#switch-model/set-switch-model-port-1-0-example.gnmi
#switch-model/set-switch-model-port-2-0-example.gnmi
#switch-model/set-switch-model-port-3-0-example.gnmi
#switch-model/set-switch-model-port-3-1-example.gnmi
#route/set-route-200-example.gnmi
#route/set-route-201-example.gnmi
#dhcp-server/set-dhcp-server-42-example.gnmi
#dhcp-server/set-dhcp-server-43-example.gnmi
#dhcp-server/set-dhcp-server-44-example.gnmi
#switch/set-switch-example.gnmi
#switch/set-switch-vlan-100-example.gnmi
#switch/set-switch-vlan-101-example.gnmi
#switch/set-switch-vlan-102-example.gnmi
#switch/set-switch-port-1-0-example.gnmi
#switch/set-switch-port-2-0-example.gnmi
#switch/set-switch-port-3-0-example.gnmi
#switch/set-switch-port-3-1-example.gnmi
route/set-route-200-example.gnmi
route/set-route-201-example.gnmi
)

for gnmi_set in ${gnmi_set_files[@]}; do
  echo Loading $gnmi_set
  gnmi_cli ${GNMI_ARGS} -set -proto "$(cat $gnmi_set)"
done
