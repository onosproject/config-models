#!/bin/sh

#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

kubectl -n micro-onos exec -it deployment/onos-cli -- \
onos topo create kind test-kind test-kind

kubectl -n micro-onos exec -it deployment/onos-cli -- \
onos topo create entity test-target \
    -a onos.topo.Configurable='{"address":"","version":"1.0.x","type":"testdevice"}' \
    -a onos.topo.TLSOptions='{"insecure":true}' \
    -a onos.topo.Asset='{"name":"Test Target"}' \
    -a onos.topo.MastershipState='{}' \
    -k test-kind

echo "test-target targets created - now"
echo "use 'go install github.com/onosproject/onos-cli/cmd/gnmi_cli' to install gnmi_cli"

echo "This expects onos-config to be port-forwared to localhost on port 5150"
