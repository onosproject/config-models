models:
	go generate
	pyang -f tree yang/*.yang > aether-2.0.0.tree

openapi-gen:
	go build cmd/openapi-gen.go

clean:
	rm -f openapi-gen
