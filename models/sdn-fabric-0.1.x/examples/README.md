<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation

SPDX-License-Identifier: Apache-2.0
-->

# Using examples with gNMI

To use this model with onos-config you will need to compile it as a docker image
```bash
cd models/sdn-fabric-0.1.x
make image
```

then if you are running `onos-umbrella` edit the values file to contain this plugin:
```yaml
  modelPlugins:
    - name: testdevice-2
      image: onosproject/testdevice-2.0.x:0.5.9-testdevice-2.0.x
      endpoint: localhost
      port: 5153
    - name: sdn-fabric
      image: onosproject/sdn-fabric-0.1.x:0.1.0-dev-sdn-fabric-latest
      endpoint: localhost
      port: 5154
```

Run `make deps` in the `onos-umbrella` folder and then run `onos-umbrella` in the normal way e.g.
```bash
helm -n micro-onos install onos-umbrella ./onos-umbrella
```

Run `onos-cli` with
```bash
kubectl -n micro-onos exec -it deployment/onos-cli -- /bin/bash
```

Verify that the plugin is loaded correctly and create a Kind and an Entity in `onos-topo`:
```bash
onos config get plugins -v

onos topo create kind fabric fabric

onos topo create entity new-fabric \
    -a onos.topo.Configurable='{"address":"sdn-fabric-adapter-v0-1:5150","version":"0.1.x","type":"sdn-fabric"}' \
    -a onos.topo.TLSOptions='{"insecure":true}' \
    -a onos.topo.Asset='{"name":"New Fabric"}' \
    -a onos.topo.MastershipState='{}' \
    -k fabric
```

To create a switch model use the `gnmi_cli set` command to push the contents of **switch-model-example.gnmi** to onos-config.

> For reference the structure of this file follows the [message SetRequest](https://github.com/openconfig/gnmi/blob/6eb133c65a13a4521601d533a0fe2be5daf3033f/proto/gnmi/gnmi.proto#L344) in the gNMI standard.

For example, with the file locally copied over to the `onos-cli` pod:
```bash
gnmi_cli -address onos-config:5150 -set -proto "$(cat switch-model-example.gnmi)" -timeout 5s -en PROTO -alsologtostderr -insecure -client_crt /etc/ssl/certs/client1.crt -client_key /etc/ssl/certs/client1.key -ca_crt /etc/ssl/certs/onfca.crt
```

should give a response showing that the switch-model config passed validation and was created properly with pipeline and form-factor attributes
```
response: <
  path: <
    elem: <
      name: "switch-model"
      key: <
        key: "switch-model-id"
        value: "super-switch-1610"
      >
    >
    elem: <
      name: "switch-model-id"
    >
    target: "new-fabric"
  >
  op: UPDATE
>
response: <
  path: <
    elem: <
      name: "switch-model"
      key: <
        key: "switch-model-id"
        value: "super-switch-1610"
      >
    >
    elem: <
      name: "pipeline"
    >
    target: "new-fabric"
  >
  op: UPDATE
>
...
...
timestamp: 1653384582
extension: <
  registered_ext: <
    id: 110
    msg: "\n)uuid:013790ee-b61b-48e2-b805-70b5c25a8c87\020\001"
  >
>
```
