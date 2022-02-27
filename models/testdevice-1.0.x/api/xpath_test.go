// Copyright 2022-present Open Networking Foundation.
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

package api

import (
	"fmt"
	"github.com/SeanCondon/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_XPathSelect(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config.json")
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

	tests := []navigator.XpathSelect{
		{
			Name: "test key1",
			Path: "/t1:cont1a/t1e:list5/@t1e:key1",
			Expected: []string{
				"Iter Value: key1: eight",
				"Iter Value: key1: five",
				"Iter Value: key1: five",
				"Iter Value: key1: two",
			},
		},
		{
			Name: "test key2",
			Path: "/t1:cont1a/t1e:list5/@t1e:key2",
			Expected: []string{
				"Iter Value: key2: 8",
				"Iter Value: key2: 6",
				"Iter Value: key2: 7",
				"Iter Value: key2: 1",
			},
		},
		{
			Name: "test key2 eight",
			Path: "/t1:cont1a/t1e:list5[@t1e:key1='eight'][@t1e:key2=8]/t1e:leaf5a",
			Expected: []string{
				"Iter Value: leaf5a: 5a eight-8",
			},
		},
		{
			Name: "test list4 select REFERENCE by list 2a tx-power",
			Path: "/t1:cont1a/t1e:list4[@t1e:id=/t1:cont1a/t1:list2a[t1:tx-power=6]/@t1:name]/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a2",
			},
		},
		{
			Name: "test list4 select REFERENCE by list 2a tx-power relative path",
			Path: "/t1:cont1a/t1e:list4[@t1e:id=../t1:list2a[t1:tx-power=6]/@t1:name]/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a2",
			},
		},
		{
			Name: "test list4 select REFERENCE by list 2a tx-power 5 relative path",
			Path: "/t1:cont1a/t1e:list4[@t1e:id=../t1:list2a[t1:tx-power=5]/@t1:name]/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a1",
			},
		},
		{
			Name: "test list4 1 list4a",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a",
			Expected: []string{
				"Iter Value: list4a: value of list4a",
				"Iter Value: list4a: value of list4a",
				"Iter Value: list4a: value of list4a",
			},
		},
		{
			Name: "test list4 1 list4a displayname",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
				"Iter Value: displayname: Value l2a1-six-6",
			},
		},
		{
			Name: "test list4 1 list4a fives displayname",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a[@t1e:fkey1='five']/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name: "test list4 1 list4a 1 displayname",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a[@t1e:fkey1='five'][@t1e:fkey2=7]/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name: "test list4 1 list4a 1 displayname by reference to list 5 entry 5a five-7",
			Path: "/t1:cont1a/t1e:list4[@t1e:id=../t1:list2a[t1:tx-power=5]/@t1:name]/t1e:list4a[@t1e:fkey1=../../t1e:list5[t1e:leaf5a='5a five-7']/@t1e:key1][@t1e:fkey2=../../t1e:list5[t1e:leaf5a='5a five-7']/@t1e:key2]/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.Path)
		assert.NoError(t, err, test.Name)
		assert.NotNil(t, expr, test.Name)

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.Expected)-1, test.Name, ". More results than expected")
			assert.Equal(t, test.Expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.Name)
			resultCount++
		}
		assert.Equal(t, len(test.Expected), resultCount, "%s. Did not receive all the expected results", test.Name)
	}
}

// Test_XPathSelectRelativeStart - start each test from cont1a - the thing that contains all the list2a entries
func Test_XPathSelectRelativeStart(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config.json")
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

	tests := []navigator.XpathSelect{
		{
			Name: "test key1",
			Path: "t1:list2a[1]/t1:tx-power",
			Expected: []string{
				"Iter Value: tx-power: 5",
			},
		},
		{
			Name:     "test preceding-sibling",
			Path:     "t1:list2a[1]/preceding-sibling::node()/t1:tx-power", // There is no preceding sibling
			Expected: []string{},
		},
		{
			Name: "test following-sibling",
			Path: "t1:list2a[1]/following-sibling::node()/t1:tx-power", // There are 4 following siblings, but only 3 contain tx-power
			Expected: []string{
				"Iter Value: tx-power: 6", // l2a2
				"Iter Value: tx-power: 8", // l2a3
				// there's no tx-power on l2a4
				"Iter Value: tx-power: 11", // l2a5
				"Iter Value: tx-power: 12", // l2a6
			},
		},
		{
			Name: "test following-sibling who has same tx-power as current",
			// following-sibling below returns a node-set which is inadvertently cast to a string
			// which will extract only the first entry and then cast to string = "6"
			// and it will detect a match only when processing node l2a5 (as l2a6) has a similar value.
			// This means that to detect duplicate nodes the nodes will have to be sorted in order of tx-power
			// This is not currently done, as nodes are sorted by their @name attribute
			Path:     "t1:list2a[set-contains(following-sibling::t1:list2a/t1:tx-power, t1:tx-power)]/@t1:name",
			Expected: []string{},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.Path)
		assert.NoError(t, err, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild()) // cont1a

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.Expected)-1, test.Name, ". More results than expected")
			assert.Equal(t, test.Expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.Name)
			resultCount++
		}
		assert.Equal(t, len(test.Expected), resultCount, "%s. Did not receive all the expected results", test.Name)
	}
}

func Test_XPathEvaluate(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config.json")
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

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get key1",
			Path:     "count(/t1:cont1a/t1e:list5/@t1e:key1)",
			Expected: float64(4),
		},
		{
			Name:     "test get key1 for 'five'", // There are 2 fives
			Path:     "count(/t1:cont1a/t1e:list5[@t1e:key1='five']/@t1e:key1)",
			Expected: float64(2),
		},
		{
			Name:     "test extract key1 for five",
			Path:     "string(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key1)",
			Expected: "five",
		},
		{
			Name:     "test extract key2 for five",
			Path:     "number(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key2)",
			Expected: float64(7),
		},
		{
			Name: "test concat string",
			Path: "concat(concat('5e ', string(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key1)), " +
				"concat('-', string(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key2)))",
			Expected: "5e five-7",
		},
		{
			Name:     "test list2a entry name when tx-power=5",
			Path:     "string(/t1:cont1a/t1:list2a[t1:tx-power=5]/@t1:name)",
			Expected: "l2a1",
		},
		{
			Name:     "test set-contains detects string in a set",
			Path:     "set-contains(/t1:cont1a/t1:list2a[t1:tx-power>10]/@t1:name, 'l2a6')",
			Expected: true,
		},
		{
			Name:     "test set-contains detects string is missing from set",
			Path:     "set-contains(/t1:cont1a/t1:list2a[t1:tx-power>10]/@t1:name, 'l2a3')",
			Expected: false,
		},
		{
			Name:     "test set-contains detects that some items from first set are in second set",
			Path:     "set-contains(/t1:cont1a/t1:list2a[t1:tx-power>10]/@t1:name, /t1:cont1a/t1:list2a[t1:tx-power>11]/@t1:name)",
			Expected: true,
		},
		{
			Name:     "test set-contains detects that no items from first set are in second set",
			Path:     "set-contains(/t1:cont1a/t1:list2a[t1:tx-power>8]/@t1:name, /t1:cont1a/t1:list2a[t1:tx-power<=8]/@t1:name)",
			Expected: false,
		},
		{
			Name:     "test set-equals detects that all items from first set are in second set",
			Path:     "set-equals(/t1:cont1a/t1:list2a[t1:tx-power>8]/@t1:name, /t1:cont1a/t1:list2a[t1:tx-power>8]/@t1:name)",
			Expected: true,
		},
		{
			Name:     "test set-equals detects that all items from first set are not found in second set",
			Path:     "set-equals(/t1:cont1a/t1:list2a[t1:tx-power>8]/@t1:name, /t1:cont1a/t1:list2a[t1:tx-power>=8]/@t1:name)",
			Expected: false,
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}

// Test_XPathEvaluateRelativePath - start each test from cont1a - the thing that contains all the list2a entries
func Test_XPathEvaluateRelativePath(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config.json")
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

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get key1",
			Path:     "number(t1:list2a[1]/t1:tx-power)",
			Expected: float64(5), // query gives a node-set, which is converted to string, which extracts the value of first node "5" and then converts to number
		},
		{
			Name:     "get the tx-power value of next entry in list",
			Path:     "number(t1:list2a[1]/following-sibling::node()/t1:tx-power)",
			Expected: float64(6), // Takes the first entry of the 5 node set
		},
		{
			Name:     "get the count of all the tx-power in following nodes",
			Path:     "count(t1:list2a[1]/following-sibling::node()/t1:tx-power)",
			Expected: float64(4), // The result is a node set - we count it here
		},
		{
			Name:     "get the count of all the tx-power same as current",
			Path:     "count(t1:list2a[t1:tx-power = following-sibling::t1:list2a/t1:tx-power])",
			Expected: float64(0), // node set is cast to string, which means only first entry is compared
		},
		{
			Name:     "check if tx-power is unique",
			Path:     "boolean(t1:list2a[set-contains(following-sibling::t1:list2a/t1:tx-power, t1:tx-power)])",
			Expected: false, // means the node-set is not-empty
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild()) // cont1a

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}
