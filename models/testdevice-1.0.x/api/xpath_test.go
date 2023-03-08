// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"fmt"
	"github.com/SeanCondon/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_XPathSelectNamespace(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config.json")
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name:  "test key1",
			XPath: "/t1:cont1a/t1e:list5/@t1e:key1",
			Expected: []string{
				"Iter Value: key1: eight",
				"Iter Value: key1: five",
				"Iter Value: key1: five",
				"Iter Value: key1: two",
			},
		},
		{
			Name:  "test key2",
			XPath: "/t1:cont1a/t1e:list5/@t1e:key2",
			Expected: []string{
				"Iter Value: key2: 8",
				"Iter Value: key2: 6",
				"Iter Value: key2: 7",
				"Iter Value: key2: 1",
			},
		},
		{
			Name:  "test key2 eight",
			XPath: "/t1:cont1a/t1e:list5[@t1e:key1='eight'][@t1e:key2=8]/t1e:leaf5a",
			Expected: []string{
				"Iter Value: leaf5a: 5a eight-8",
			},
		},
		{
			Name:  "test list4 select REFERENCE by list 2a tx-power",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id=/t1:cont1a/t1:list2a[t1:tx-power=6]/@t1:name]/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a2",
			},
		},
		{
			Name:  "test list4 select REFERENCE by list 2a tx-power relative path",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id=../t1:list2a[t1:tx-power=6]/@t1:name]/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a2",
			},
		},
		{
			Name:  "test list4 select REFERENCE by list 2a tx-power 5 relative path",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id=../t1:list2a[t1:tx-power=5]/@t1:name]/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a1",
			},
		},
		{
			Name:  "test list4 1 list4a",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a",
			Expected: []string{
				"Iter Value: list4a: value of list4a",
				"Iter Value: list4a: value of list4a",
				"Iter Value: list4a: value of list4a",
			},
		},
		{
			Name:  "test list4 1 list4a displayname",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
				"Iter Value: displayname: Value l2a1-six-6",
			},
		},
		{
			Name:  "test list4 1 list4a fives displayname",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a[@t1e:fkey1='five']/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name:  "test list4 1 list4a 1 displayname",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a[@t1e:fkey1='five'][@t1e:fkey2=7]/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name:  "test list4 1 list4a 1 displayname by reference to list 5 entry 5a five-7",
			XPath: "/t1:cont1a/t1e:list4[@t1e:id=../t1:list2a[t1:tx-power=5]/@t1:name]/t1e:list4a[@t1e:fkey1=../../t1e:list5[t1e:leaf5a='5a five-7']/@t1e:key1][@t1e:fkey2=../../t1e:list5[t1e:leaf5a='5a five-7']/@t1e:key2]/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.XPath)
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

func Test_XPathSelect(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config.json")
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name:  "test key1",
			XPath: "/cont1a/list5/@key1",
			Expected: []string{
				"Iter Value: key1: eight",
				"Iter Value: key1: five",
				"Iter Value: key1: five",
				"Iter Value: key1: two",
			},
		},
		{
			Name:  "test key2",
			XPath: "/cont1a/list5/@key2",
			Expected: []string{
				"Iter Value: key2: 8",
				"Iter Value: key2: 6",
				"Iter Value: key2: 7",
				"Iter Value: key2: 1",
			},
		},
		{
			Name:  "test key2 eight",
			XPath: "/cont1a/list5[@key1='eight'][@key2=8]/leaf5a",
			Expected: []string{
				"Iter Value: leaf5a: 5a eight-8",
			},
		},
		{
			Name:  "test list4 select REFERENCE by list 2a tx-power",
			XPath: "/cont1a/list4[@id=/cont1a/list2a[tx-power=6]/@name]/leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a2",
			},
		},
		{
			Name:  "test list4 select REFERENCE by list 2a tx-power relative path",
			XPath: "/cont1a/list4[@id=../list2a[tx-power=6]/@name]/leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a2",
			},
		},
		{
			Name:  "test list4 select REFERENCE by list 2a tx-power 5 relative path",
			XPath: "/cont1a/list4[@id=../list2a[tx-power=5]/@name]/leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a1",
			},
		},
		{
			Name:  "test list4 1 list4a",
			XPath: "/cont1a/list4[@id='l2a1']/list4a",
			Expected: []string{
				"Iter Value: list4a: value of list4a",
				"Iter Value: list4a: value of list4a",
				"Iter Value: list4a: value of list4a",
			},
		},
		{
			Name:  "test list4 1 list4a displayname",
			XPath: "/cont1a/list4[@id='l2a1']/list4a/displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
				"Iter Value: displayname: Value l2a1-six-6",
			},
		},
		{
			Name:  "test list4 1 list4a fives displayname",
			XPath: "/cont1a/list4[@id='l2a1']/list4a[@fkey1='five']/displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name:  "test list4 1 list4a 1 displayname",
			XPath: "/cont1a/list4[@id='l2a1']/list4a[@fkey1='five'][@fkey2=7]/displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name:  "test list4 1 list4a 1 displayname by reference to list 5 entry 5a five-7",
			XPath: "/cont1a/list4[@id=../list2a[tx-power=5]/@name]/list4a[@fkey1=../../list5[leaf5a='5a five-7']/@key1][@fkey2=../../list5[leaf5a='5a five-7']/@key2]/displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.XPath)
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
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config.json")
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name:  "test key1",
			XPath: "list2a[1]/tx-power",
			Expected: []string{
				"Iter Value: tx-power: 5",
			},
		},
		{
			Name:     "test preceding-sibling",
			XPath:    "list2a[1]/preceding-sibling::node()/tx-power", // There is no preceding sibling
			Expected: []string{},
		},
		{
			Name:  "test following-sibling",
			XPath: "list2a[1]/following-sibling::node()/tx-power", // There are 4 following siblings, but only 3 contain tx-power
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
			XPath:    "list2a[set-contains(following-sibling::list2a/tx-power, tx-power)]/@name",
			Expected: []string{},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.XPath)
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
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config.json")
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get key1",
			XPath:    "count(/cont1a/list5/@key1)",
			Expected: float64(4),
		},
		{
			Name:     "test get key1 for 'five'", // There are 2 fives
			XPath:    "count(/cont1a/list5[@key1='five']/@key1)",
			Expected: float64(2),
		},
		{
			Name:     "test extract key1 for five",
			XPath:    "string(/cont1a/list5[@key1='five'][@key2=7]/@key1)",
			Expected: "five",
		},
		{
			Name:     "test extract key2 for five",
			XPath:    "number(/cont1a/list5[@key1='five'][@key2=7]/@key2)",
			Expected: float64(7),
		},
		{
			Name: "test concat string",
			XPath: "concat(concat('5e ', string(/cont1a/list5[@key1='five'][@key2=7]/@key1)), " +
				"concat('-', string(/cont1a/list5[@key1='five'][@key2=7]/@key2)))",
			Expected: "5e five-7",
		},
		{
			Name:     "test list2a entry name when tx-power=5",
			XPath:    "string(/cont1a/list2a[tx-power=5]/@name)",
			Expected: "l2a1",
		},
		{
			Name:     "test set-contains detects string in a set",
			XPath:    "set-contains(/cont1a/list2a[tx-power>10]/@name, 'l2a6')",
			Expected: true,
		},
		{
			Name:     "test set-contains detects string is missing from set",
			XPath:    "set-contains(/cont1a/list2a[tx-power>10]/@name, 'l2a3')",
			Expected: false,
		},
		{
			Name:     "test set-contains detects that some items from first set are in second set",
			XPath:    "set-contains(/cont1a/list2a[tx-power>10]/@name, /cont1a/list2a[tx-power>11]/@name)",
			Expected: true,
		},
		{
			Name:     "test set-contains detects that no items from first set are in second set",
			XPath:    "set-contains(/cont1a/list2a[tx-power>8]/@name, /cont1a/list2a[tx-power<=8]/@name)",
			Expected: false,
		},
		{
			Name:     "test set-equals detects that all items from first set are in second set",
			XPath:    "set-equals(/cont1a/list2a[tx-power>8]/@name, /cont1a/list2a[tx-power>8]/@name)",
			Expected: true,
		},
		{
			Name:     "test set-equals detects that all items from first set are not found in second set",
			XPath:    "set-equals(/cont1a/list2a[tx-power>8]/@name, /cont1a/list2a[tx-power>=8]/@name)",
			Expected: false,
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.XPath)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}

// Test_XPathEvaluateRelativePath - start each test from cont1a - the thing that contains all the list2a entries
func Test_XPathEvaluateRelativePath(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config.json")
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get key1",
			XPath:    "number(list2a[1]/tx-power)",
			Expected: float64(5), // query gives a node-set, which is converted to string, which extracts the value of first node "5" and then converts to number
		},
		{
			Name:     "get the tx-power value of next entry in list",
			XPath:    "number(list2a[1]/following-sibling::node()/tx-power)",
			Expected: float64(6), // Takes the first entry of the 5 node set
		},
		{
			Name:     "get the count of all the tx-power in following nodes",
			XPath:    "count(list2a[1]/following-sibling::node()/tx-power)",
			Expected: float64(4), // The result is a node set - we count it here
		},
		{
			Name:     "get the count of all the tx-power same as current",
			XPath:    "count(list2a[tx-power = following-sibling::list2a/tx-power])",
			Expected: float64(0), // node set is cast to string, which means only first entry is compared
		},
		{
			Name:     "check if tx-power is unique",
			XPath:    "boolean(list2a[set-contains(following-sibling::list2a/tx-power, tx-power)])",
			Expected: false, // means the node-set is not-empty
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.XPath)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild()) // cont1a

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}
