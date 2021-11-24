redo: clean copyin models openapi-gen plproxy-3.0.0-openapi3.yaml copyout

copyin:
#	rm -rf yang
#	cp -a ../../../sdran-helm-charts/config-models/plproxy-3.x/files/yang .

models:
	go generate
	pyang -f tree ../../../roc-helm-charts/config-models/plproxy-1.x/files/yang/*.yang > plproxy-1.0.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

plproxy-1.0.0-openapi3.yaml: openapi-gen
	./openapi-gen > plproxy-1.0.0-openapi3.yaml

copyout:
	cp plproxy-1.0.0-openapi3.yaml ../../../aether-roc-api/api/
	cp plproxy-1.0.0.tree ../../../roc-helm-charts/config-models/plproxy-1.x/files/

clean:
	rm -f openapi-gen plproxy-1.0.0-openapi3.yaml
