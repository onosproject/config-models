redo: clean copyin models openapi-gen aether-2.2.0-openapi3.yaml copyout

copyin:
	rm -rf yang
	cp -a ../../../sdran-helm-charts/config-models/aether-2.2.x/files/yang .

models:
	go generate
	pyang -f tree yang/*.yang > aether-2.2.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

aether-2.2.0-openapi3.yaml: openapi-gen
	./openapi-gen > aether-2.2.0-openapi3.yaml

copyout:
	cp aether-2.2.0-openapi3.yaml ../../../aether-roc-api/api/

clean:
	rm -f openapi-gen aether-2.2.0-openapi3.yaml
