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
	"fmt"
	"github.com/antchfx/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type xpath_test struct {
	name     string
	path     string
	expected []string
}

type xpath_evaluate struct {
	name     string
	path     string
	expected interface{}
}

func Test_XPathSelect(t *testing.T) {
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

	tests := []xpath_test{
		{
			name: "test leaf2a",
			path: "/t1:cont1a/t1:cont2a/t1:leaf2a",
			expected: []string{
				"Iter Value: leaf2a: 1",
			},
		},
		{
			name: "test leaf2b",
			path: "/t1:cont1a/t1:cont2a/t1:leaf2b",
			expected: []string{
				"Iter Value: leaf2b: 0.4321",
			},
		},
		{
			name:     "test leaf2c",
			path:     "/t1:cont1a/t1:cont2a/t1:leaf2c",
			expected: []string{}, // No value, so no response
		},
		{
			name: "test leaf2g",
			path: "/t1:cont1a/t1:cont2a/t1:leaf2g",
			expected: []string{
				"Iter Value: leaf2g: true",
			},
		},
		{
			name: "test leaf1a",
			path: "/t1:cont1a/t1:leaf1a",
			expected: []string{
				"Iter Value: leaf1a: leaf1aval",
			},
		},
		{
			name: "test list2a all instances names",
			path: "/t1:cont1a/t1:list2a/@t1:name", // List indices are always attributes - referred to with @
			expected: []string{
				"Iter Value: name: l2a1",
				"Iter Value: name: l2a2",
			},
		},
		{
			name: "test list2a select 2nd instance rx-power",
			path: "/t1:cont1a/t1:list2a[@t1:name='l2a2']/t1:rx-power", // select with []
			expected: []string{
				"Iter Value: rx-power: 26",
			},
		},
		{
			name: "test list2a select 1st instance tx-power",
			path: "/t1:cont1a/t1:list2a[@t1:name='l2a1']/t1:tx-power", // select with []
			expected: []string{
				"Iter Value: tx-power: 5",
			},
		},
		{
			name: "test index 1 filter",
			path: "/t1:cont1b-state/t1:list2b[@t1:index1=11]/t1:leaf3c",
			expected: []string{
				"Iter Value: leaf3c: 3c 11-20 test",
			},
		},
		{
			name: "test index 2 filter",
			path: "/t1:cont1b-state/t1:list2b[@t1:index2=20]/@t1:index1",
			expected: []string{
				"Iter Value: index1: 10",
				"Iter Value: index1: 11",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.path)
		assert.NoError(t, err, test.name)
		assert.NotNil(t, expr, test.name)

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.expected)-1, test.name, ". More results than expected")
			assert.Equal(t, test.expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.name)
			resultCount++
		}
		assert.Equal(t, len(test.expected), resultCount, "%s. Did not receive all the expected results", test.name)
	}
}

func Test_XPathEvaluate(t *testing.T) {
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

	tests := []xpath_evaluate{
		{
			name: "test leaf2a",
			path: "/t1:cont1a/t1:cont2a/t1:leaf2a = true",
			expected: true,
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.path)
		assert.NoError(t, testErr, test.name)
		assert.NotNil(t, expr, test.name)

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.expected, result)
	}
}

func Test_WalkAndValidate(t *testing.T) {
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
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.NoError(t, validateErr)
}

func Test_XPathNodeNavigation(t *testing.T) {

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
	assert.Equal(t, "value of device", ynn.Value())

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
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "1", ynn.Value())

	assert.False(t, ynn.MoveToChild())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2b", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "0.4321", ynn.Value())

	// Skips leaf2c and leaf2d as they have no values
	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2e", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType()) // Leaf list
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "[5 4 3 2 1]", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2f", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2g", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "true", ynn.Value())

	// no next exists so returns false
	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToPrevious())
	assert.Equal(t, "leaf2f", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", ynn.Value())

	assert.True(t, ynn.MoveToFirst())
	assert.Equal(t, "leaf2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "1", ynn.Value())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "cont2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "value of cont2a", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf1a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
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
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "25", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "tx-power", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
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
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "26", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "tx-power", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
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
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
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
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "3c 11-20 test", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf3d", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "IDTYPE2", ynn.Value())

	assert.False(t, ynn.MoveToNext()) // No further leaves

}
