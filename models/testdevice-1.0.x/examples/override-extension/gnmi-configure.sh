#!/bin/sh

#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

CERTS=../../../../../onos-cli/pkg/certs
GNMI_ARGS="--address localhost:5150 -timeout 5s -en PROTO -alsologtostderr -insecure -client_crt ${CERTS}/client1.crt -client_key ${CERTS}/client1.key -ca_crt ${CERTS}/onfca.crt"

gnmi_set_files=(
set-testdevice-1.0.0-cont1a_cont2a.gnmi
set-testdevice-1.0.0-cont1a_list2a-2a-1.gnmi
)

for gnmi_set in ${gnmi_set_files[@]}; do
  echo Loading $gnmi_set
  gnmi_cli ${GNMI_ARGS} -set -proto "$(cat $gnmi_set)"
done

echo "Now you should be able to create a 2.0.x version of 'test-target'"
echo "Please run ../../../testdevice-2.0.x/examples/override-extension/gnmi-configure.sh"

