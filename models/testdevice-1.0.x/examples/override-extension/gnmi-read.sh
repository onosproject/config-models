#!/bin/sh

#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#
ENCODING=${ENCODING:=PROTO}
CERTS=../../../../../onos-cli/pkg/certs
GNMI_ARGS="--address localhost:5150 -timeout 5s -en ${ENCODING} -alsologtostderr -insecure -client_crt ${CERTS}/client1.crt -client_key ${CERTS}/client1.key -ca_crt ${CERTS}/onfca.crt"

gnmi_get_files=(
get-testdevice-1.0.0-cont1a.gnmi
get-testdevice-1.0.0-cont1a_list2a.gnmi
)

for gnmi_get in ${gnmi_get_files[@]}; do
  echo "\n\nReading with $gnmi_get. Writing to $gnmi_get.out"
  gnmi_cli ${GNMI_ARGS} -get -proto "$(cat $gnmi_get)" > $gnmi_get.out
done
