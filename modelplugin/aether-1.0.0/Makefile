models:
	go generate
	pyang -f tree yang/*.yang > aether-1.0.0.tree || echo "ignore pyang erors about paths not found"
