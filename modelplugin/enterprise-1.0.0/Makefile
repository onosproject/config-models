models:
	go generate
	pyang -f tree yang/*.yang > enterprise.tree

openapi-gen:
	go build cmd/openapi-gen.go

enterprise-1.0.0-openapi3.yaml: openapi-gen
	./openapi-gen > enterprise-1.0.0-openapi3.yaml

clean:
	rm -f openapi-gen enterprise-1.0.0-openapi3.yaml
