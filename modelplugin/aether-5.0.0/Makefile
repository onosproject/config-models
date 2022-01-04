OPENCORD ?= ${HOME}/git/opencord

redo: clean models openapi-gen aether-5.0.0-openapi3.yaml copyout

models:
	go generate
	pyang -f tree ${OPENCORD}/roc-helm-charts/config-models/aether-5.x/files/yang/*.yang > aether-5.0.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

aether-5.0.0-openapi3.yaml: openapi-gen
	./openapi-gen > aether-5.0.0-openapi3.yaml

copyout:
	cp aether-5.0.0-openapi3.yaml ../../../aether-roc-api/api/
	cp aether-5.0.0.tree ${OPENCORD}/roc-helm-charts/config-models/aether-5.x/files/

clean:
	rm -f openapi-gen aether-5.0.0-openapi3.yaml
	rm -rf aether_5_0_0
	rm -f aether-5.0.0.tree
