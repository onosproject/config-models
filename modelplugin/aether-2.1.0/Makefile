models:
	go generate
	pyang -f tree yang/*.yang > aether-2.1.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

aether-2.1.0-openapi3.yaml: openapi-gen
	./openapi-gen > aether-2.1.0-openapi3.yaml

clean:
	rm -f openapi-gen aether-2.1.0-openapi3.yaml
