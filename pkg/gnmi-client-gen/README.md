<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation

SPDX-License-Identifier: Apache-2.0
-->

# gNMI Client Gen

`gnmi-client-gen` is a utility to generate a `gNMI` client starting from the `yang` models.

This simplifies the interactions between a client and `onos-config`, creating a tool that abstract the generation
of `gNMI/yang` specific path providing a model-first library that you can import in your code.

_NOTE that this is a preliminary implementation, so future bugfixes and new features are to be expected_

## Roadmap

- [X] Release the first version of the compiler
- [X] Augment `pkg/compiler` to generate the required code
- [X] Augment the `make models-images` target to generate the client for each model
- [ ] Combine the `gnmi-client-gen` and `oapi-gen` target inside the model compiler so that a separate step is not required

### Missing features
- [ ] Support for `Yidentityref` and `Yenum` types