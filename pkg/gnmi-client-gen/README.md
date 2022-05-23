# gNMI Client Gen

`gnmi-client-gen` is a utility to generate a `gNMI` client starting from the `yang` models.

This simplifies the interactions between a client and `onos-config`, creating a tool that abstract the generation
of `gNMI/yang` specific path providing a model-first library that you can import in your code.

_NOTE that this is a preliminary implementation, so future bugfixes and new features are to be expected_

## Roadmap

- [X] Release the first version of the compiler
- [ ] Augment `pkg/compiler` to generate the required code
- [ ] Augment the `make models-images` target to generate the client for each model
- [ ] Combine the `gnmi-client-gen` and `oapi-gen` target inside the model compiler so that a separate step is not required

### Missing features
- [ ] Support for `Yidentityref` and `Yenum` types

## Testing (before it's plumbed together)

If you want to test the `gnmi-client-gen` tool from this current patch you need take a few steps:

- Uncomment code in `pkg/compiler/compiler.go` from line 136 to 141
- Uncomment code in `templates/go.mod.tpl` at line 18: `// replace github.com/onosproject/config-models => ../../`
- Rebuild the compiler image with `make images`
- Regenerate the plugin code with `make models`
- Genereate the gNMI client with `make models-gnmi-client`