#!/bin/sh

#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

CERTS=../../../../../onos-cli/pkg/certs
GNMI_ARGS="--address localhost:5150 -timeout 5s -en PROTO -alsologtostderr -insecure -client_crt ${CERTS}/client1.crt -client_key ${CERTS}/client1.key -ca_crt ${CERTS}/onfca.crt"

gnmi_set_files=(
set-testdevice-2.0.0-cont1a_cont2a.gnmi
set-testdevice-2.0.0-cont1a_list2a.gnmi
)

for gnmi_set in ${gnmi_set_files[@]}; do
  echo Loading $gnmi_set
  gnmi_cli ${GNMI_ARGS} -set -proto "$(cat $gnmi_set)"
done
