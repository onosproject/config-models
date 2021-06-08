redo: clean copyin models openapi-gen aether-3.0.0-openapi3.yaml copyout

copyin:
#	rm -rf yang
#	cp -a ../../../sdran-helm-charts/config-models/aether-3.x/files/yang .

models:
	go generate
	pyang -f tree ../../../sdran-helm-charts/config-models/aether-3.x/files/yang/*.yang > aether-3.0.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

aether-3.0.0-openapi3.yaml: openapi-gen
	./openapi-gen > aether-3.0.0-openapi3.yaml

copyout:
	cp aether-3.0.0-openapi3.yaml ../../../aether-roc-api/api/

clean:
	rm -f openapi-gen aether-3.0.0-openapi3.yaml
