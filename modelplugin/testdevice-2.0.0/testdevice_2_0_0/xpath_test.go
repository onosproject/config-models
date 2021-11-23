// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testdevice_2_0_0

import (
	"github.com/antchfx/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_XPath(t *testing.T) {

	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice2-config.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "device", ynn.LocalName())
	assert.Equal(t, "", ynn.Prefix())
	assert.Equal(t, "", ynn.Value())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "cont1a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "cont2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "leaf2a", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "1", ynn.Value())

	assert.False(t, ynn.MoveToChild())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2b", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "0.4321", ynn.Value())

	// Skips leaf2c and leaf2d as they have no values
	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2e", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType()) // Leaf list
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2f", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2g", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "true", ynn.Value())

	// no next exists so returns false
	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToPrevious())
	assert.Equal(t, "leaf2f", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", ynn.Value())

	assert.True(t, ynn.MoveToFirst())
	assert.Equal(t, "leaf2a", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "1", ynn.Value())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "cont2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf1a", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "leaf1aval", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "name", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "l2a1", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "rx-power", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "25", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "tx-power", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "5", ynn.Value())

	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "name", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "l2a2", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "rx-power", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "26", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "tx-power", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "6", ynn.Value())

	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "cont1a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "cont1b-state", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "list2b", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "index1", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "10", ynn.Value())

	assert.False(t, ynn.MoveToChild())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "index2", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "20", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf3c", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "3c 10-20 test", ynn.Value())

	assert.False(t, ynn.MoveToNext()) // There's no leaf3d present

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "list2b", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToNext()) // the next list entry

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "index1", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "11", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "index2", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "20", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf3c", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "3c 11-20 test", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf3d", ynn.LocalName())
	assert.Equal(t, xpath.TextNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "IDTYPE2", ynn.Value())

	assert.False(t, ynn.MoveToNext()) // No further leaves

}
