models:
	go generate
	# pyang is complainy about the relative paths
	pyang -f tree yang/*.yang > aether-2.0.0.tree || echo "ignore pyang erors about paths not found"
