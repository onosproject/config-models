models:
	go generate
	pyang -f tree yang/*.yang > rbac-1.0.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

rbac-1.0.0-openapi3.yaml: openapi-gen
	./openapi-gen > rbac-1.0.0-openapi3.yaml

clean:
	rm -f openapi-gen rbac-1.0.0-openapi3.yaml
