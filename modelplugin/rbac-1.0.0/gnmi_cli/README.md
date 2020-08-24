The GNMI CLI may be used to demo the `onos-config` RBAC models. Example files in this directory.

See https://github.com/onosproject/onos-config/blob/master/gnmi_cli/README.md for details on how to use these.

```bash
gnmi_cli -set -address localhost:5150 -timeout 20s -en PROTO -alsologtostderr -insecure -client_crt ../../../onos-helm-charts/onos-config/files/certs/tls.crt -client_key ../../../onos-helm-charts/onos-config/files/certs/tls.key -ca_crt ../../../onos-helm-charts/onos-config/files/certs/tls.cacrt -proto "$(cat gnmi_cli/set.roles-aether-admin.gnmi)"

gnmi_cli -set -address localhost:5150 -timeout 20s -en PROTO -alsologtostderr -insecure -client_crt ../../../onos-helm-charts/onos-config/files/certs/tls.crt -client_key ../../../onos-helm-charts/onos-config/files/certs/tls.key -ca_crt ../../../onos-helm-charts/onos-config/files/certs/tls.cacrt -proto "$(cat gnmi_cli/set.roles-aether-ops.gnmi)"

gnmi_cli -set -address localhost:5150 -timeout 20s -en PROTO -alsologtostderr -insecure -client_crt ../../../onos-helm-charts/onos-config/files/certs/tls.crt -client_key ../../../onos-helm-charts/onos-config/files/certs/tls.key -ca_crt ../../../onos-helm-charts/onos-config/files/certs/tls.cacrt -proto "$(cat gnmi_cli/set.group-menlo-admins.gnmi)"
```