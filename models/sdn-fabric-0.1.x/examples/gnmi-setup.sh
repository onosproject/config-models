#!/bin/sh

#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

kubectl -n micro-onos exec -it deployment/onos-cli -- \
onos topo create kind fabric fabric

kubectl -n micro-onos exec -it deployment/onos-cli -- \
onos topo create entity mars \
    -a onos.topo.Configurable='{"address":"fabric-adapter-v0-1:5150","version":"0.1.x","type":"sdn-fabric"}' \
    -a onos.topo.TLSOptions='{"insecure":true}' \
    -a onos.topo.Asset='{"name":"Mars Fabric"}' \
    -a onos.topo.MastershipState='{}' \
    -k fabric

kubectl -n micro-onos exec -it deployment/onos-cli -- \
onos topo create entity venus \
    -a onos.topo.Configurable='{"address":"fabric-adapter-v0-1:5150","version":"0.1.x","type":"sdn-fabric"}' \
    -a onos.topo.TLSOptions='{"insecure":true}' \
    -a onos.topo.Asset='{"name":"Venus Fabric"}' \
    -a onos.topo.MastershipState='{}' \
    -k fabric

echo "mars and venus targets created - now"
echo "use 'go install github.com/onosproject/onos-cli/cmd/gnmi_cli' to install gnmi_cli"

echo "This expects onos-config to be port-forwared to localhost on port 5150"
