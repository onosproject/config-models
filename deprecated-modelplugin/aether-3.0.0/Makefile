OPENCORD ?= ${HOME}/git/opencord

redo: clean copyin models openapi-gen aether-3.0.0-openapi3.yaml copyout

copyin:
#	rm -rf yang
#	cp -a ${OPENCORD}/roc-helm-charts/config-models/aether-3.x/files/yang .

models:
	go generate
	pyang -f tree ${OPENCORD}/roc-helm-charts/config-models/aether-3.x/files/yang/*.yang > aether-3.0.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

aether-3.0.0-openapi3.yaml: openapi-gen
	./openapi-gen > aether-3.0.0-openapi3.yaml

copyout:
	cp aether-3.0.0-openapi3.yaml ../../../aether-roc-api/api/
	cp aether-3.0.0.tree ${OPENCORD}/roc-helm-charts/config-models/aether-3.x/files/

clean:
	rm -f openapi-gen aether-3.0.0-openapi3.yaml
