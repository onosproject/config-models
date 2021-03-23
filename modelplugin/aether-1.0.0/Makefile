models:
	go generate
	pyang -f tree ../../../sdran-helm-charts/config-models/aether-1.x/files/yang/*.yang > aether-1.0.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

aether-1.0.0-openapi3.yaml: openapi-gen
	./openapi-gen > aether-1.0.0-openapi3.yaml

clean:
	rm -f openapi-gen aether-1.0.0-openapi3.yaml
