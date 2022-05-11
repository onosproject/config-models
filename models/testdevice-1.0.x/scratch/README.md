Install `onos-umbrella`
```shell
helm upgrade --install onos-config onosproject/onos-umbrella \
  --set onos-config.logging.loggers.root.level=debug --set onos-config.debug=true
```

Forward the gNMI service from `onos-config` to `localhost`
```shell
kubectl port-forward svc/onos-config 5150
```

Setup the required entities in `onos-topo`
```shell
onos topo create kind testdevice1 testdevice1
onos topo create entity target-foo \
  -a onos.topo.Configurable='{"address":"","version":"1.0.0","type":"testdevice","timeout":"30s"}' \
  -a onos.topo.TLSOptions='{"insecure":true,"plain":true}' -a onos.topo.Asset='{"name":"TargetFoo"}' \
  -a onos.topo.MastershipState='{}' -k testdevice1
```

Run the example:
```shell
go run scratch/gnmi-client-exploration.go
```