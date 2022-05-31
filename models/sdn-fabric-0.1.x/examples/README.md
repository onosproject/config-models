<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation

SPDX-License-Identifier: Apache-2.0
-->

# Using examples with gNMI

Deploy the `fabric-umbrella` Helm Chart (after the prerequisites 
from [here](https://docs.aetherproject.org/master/developer/roc.html#installing-prerequisites))

```bash
helm -n micro-onos install fabric-umbrella aetherproject/fabric-umbrella
```

This creates 2 fabrics in the deployment `mars` and `venus`.

To create a switch model use the `gnmi_cli set` command to push the contents of **switch-model-example.gnmi** to onos-config.

> For reference the structure of this file follows the [message SetRequest](https://github.com/openconfig/gnmi/blob/6eb133c65a13a4521601d533a0fe2be5daf3033f/proto/gnmi/gnmi.proto#L344) in the gNMI standard.

# gNMI samples

There are a suite of gNMI files available in this examples folder:

* In switch-model:
  * set-switch-model-example.gnmi - creates the switch model `super-switch-1610`
  * set-switch-model-port-* - creates ports in this switch model
  * get-switch-model-ports-example.gnmi - queries all attributes of all ports
  * get-switch-model-port-1-0-example.gnmi - query only the display-name and speed of port 1/0
  * get-switch-model-port-star-example.gnmi - query the display name of all ports (wildcard)
* In switch (switch-model must be created first):
  * set-switch-example.gnmi - creates the switch
  * set-switch-port-* - creates ports on the switch
  * set-switch-vlan-* - creates vlans on the switch
  * get-switch-example.gnmi - get the display-name of the switch
* In route:
  * set-route-1-example.gnmi - creates a global route


# applying a **set-** gNMI file 
For example, with the file locally copied over to the `onos-cli` pod:
```bash
gnmi_cli -set -address onos-config:5150 -proto "$(cat switch-model-example.gnmi)" -timeout 5s -en PROTO -alsologtostderr -insecure -client_crt /etc/ssl/certs/client1.crt -client_key /etc/ssl/certs/client1.key -ca_crt /etc/ssl/certs/onfca.crt
```

# applying a **get-** gNMI file
For example, with the file locally copied over to the `onos-cli` pod:
```bash
gnmi_cli -get -address onos-config:5150 -proto "$(cat get-switch-model-port-1-example.gnmi)" -timeout 5s -en PROTO -alsologtostderr -insecure -client_crt /etc/ssl/certs/client1.crt -client_key /etc/ssl/certs/client1.key -ca_crt /etc/ssl/certs/onfca.crt
```
