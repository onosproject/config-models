// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-example-1.json")
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
			Name: "test key",
			Path: "/sm:switch-model/@sm:switch-model-id",
			Expected: []string{
				"Iter Value: switch-model-id: super-switch-1610",
				"Iter Value: switch-model-id: super-switch-2100",
			},
		},
		{
			Name: "test ordinary attribute",
			Path: "/sm:switch-model/sm:pipeline",
			Expected: []string{
				"Iter Value: pipeline: 2",
				"Iter Value: pipeline: 2",
			},
		},
		{
			Name: "test switch-models description",
			Path: "/sm:switch-model/ft:description",
			Expected: []string{
				"Iter Value: description: Super Switch Model 1610 16-port 10Gb on mars fabric",
				"Iter Value: description: Super Switch Model NIC 100 2-port 100Gb",
			},
		},
		{
			Name: "test switch model[0] desc",
			Path: "/sm:switch-model[@sm:switch-model-id='super-switch-1610']/ft:description",
			Expected: []string{
				"Iter Value: description: Super Switch Model 1610 16-port 10Gb on mars fabric",
			},
		},
		{
			Name: "test switch model[0] desc",
			Path: "/sm:switch-model[@sm:switch-model-id='super-switch-2100']/ft:description",
			Expected: []string{
				"Iter Value: description: Super Switch Model NIC 100 2-port 100Gb",
			},
		},
		{
			Name: "test switch model[0] ports",
			Path: "/sm:switch-model[@sm:switch-model-id='super-switch-1610']/sm:port/@sm:cage-number",
			Expected: []string{
				"Iter Value: cage-number: 1", // duplicate
				"Iter Value: cage-number: 2", // duplicate
				"Iter Value: cage-number: 3",
				"Iter Value: cage-number: 4",
			},
		},
		{
			Name: "test switch model[0] port 3",
			Path: "/sm:switch-model[@sm:switch-model-id='super-switch-1610']/sm:port[@sm:cage-number=3]/@sm:cage-number",
			Expected: []string{
				"Iter Value: cage-number: 3",
			},
		},
		{
			Name: "test switch[0] port 3 name",
			Path: "/sw:switch[@sw:switch-id='san-jose-edge-tor-1S']/sw:port[@sw:cage-number=3]/ft:display-name",
			Expected: []string{
				"Iter Value: display-name: Port 3/0",
			},
		},
		{
			Name: "test switch by model port 3 name", // 2 switches with this model
			Path: "/sw:switch[sw:model-id='super-switch-1610']/sw:port[@sw:cage-number=3]/ft:display-name",
			Expected: []string{
				"Iter Value: display-name: Port 3/0",
				"Iter Value: display-name: Port 3/0",
			},
		},
		{
			Name: "test switch-model by switch name",
			Path: "/sm:switch-model[@sm:switch-model-id=/sw:switch[@sw:switch-id='san-jose-edge-tor-1S']/sw:model-id]/ft:display-name",
			Expected: []string{
				"Iter Value: display-name: Super Switch 1610",
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

// Test_XPathSelectRelativeStart - start each test from switch[1] - the thing that contains all the port entries
func Test_XPathSelectRelativeStart(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-example-1.json")
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
			Name: "test name",
			Path: "./ft:display-name",
			Expected: []string{
				"Iter Value: display-name: SJ TOR S2",
			},
		},
		{
			Name: "port[2] name by id",
			Path: "sw:port[@sw:cage-number=2][@sw:channel-number=0]/ft:display-name",
			Expected: []string{
				"Iter Value: display-name: Port 2/0 on switch 2",
			},
		},
		{
			Name: "port[2] name by number",
			Path: "sw:port[2]/ft:display-name",
			Expected: []string{
				"Iter Value: display-name: Port 2/0 on switch 2",
			},
		},
		{
			Name:     "port[2] preceding-sibling",
			Path:     "sw:port[@sw:cage-number=2][@sw:channel-number=0]/preceding-sibling::node()/ft:display-name",
			Expected: []string{}, // There is no preceding sibling
		},
		{
			Name: "port[2] following-sibling",
			Path: "sw:port[@sw:cage-number=2][@sw:channel-number=0]/following-sibling::node()/ft:display-name", // There are 4 following siblings, but only 1 is a port
			Expected: []string{
				"Iter Value: display-name: Port 3/0",
				"Iter Value: display-name: VLAN 100", // lists following nodes regardless of type
				"Iter Value: display-name: VLAN 101",
				"Iter Value: display-name: VLAN 102",
			},
		},
		{
			Name: "test following-sibling who has same tx-power as current",
			// following-sibling below returns a node-set which is inadvertently cast to a string
			// which will extract only the first entry and then cast to string = "6"
			// and it will detect a match only when processing node l2a5 (as l2a6) has a similar value.
			// This means that to detect duplicate nodes the nodes will have to be sorted in order of tx-power
			// This is not currently done, as nodes are sorted by their @name attribute
			Path:     "sw:port[set-contains(following-sibling::sw:port/ft:display-name, ft:display-name)]/ft:display-name",
			Expected: []string{},
		},
		//{
		//	Name: "this variable test basic",
		//	// using the $this variable - the node which is current at the start of the query
		//	Path: "$this/ft:display-name",
		//	Expected: []string{
		//		"Iter Value: display-name: SJ TOR S2",
		//	},
		//},
		//{
		//	Name: "this variable test advanced",
		//	Path: "/sm:switch-model[@sm:switch-model-id=$this/sw:model-id]/ft:display-name",
		//	Expected: []string{
		//		"Iter Value: display-name: Super Switch 1610",
		//	},
		//},
	}

	for _, test := range tests {
		expr, err1 := xpath.Compile(test.Path)
		assert.NoError(t, err1, test.Name)
		if err1 != nil {
			t.FailNow()
		}
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToNext()) //route
		assert.True(t, ynn.MoveToNext()) //switch-model[0]
		assert.True(t, ynn.MoveToNext()) //switch-model[1]
		assert.True(t, ynn.MoveToNext()) //switch[0]
		assert.True(t, ynn.MoveToNext()) //switch[1]
		assert.True(t, ynn.MoveToNext()) //switch[2]

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
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-example-1.json")
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
			Name:     "ports on switch-model[1]",
			Path:     "count(/sm:switch-model[1]/sm:port/@sm:cage-number)",
			Expected: float64(4),
		},
		{
			Name:     "ports on switch[2]",
			Path:     "count(/sw:switch[2]/sw:port/@sw:cage-number)",
			Expected: float64(5),
		},
		{
			Name:     "port[5] on switch[2] - name",
			Path:     "string(/sw:switch[2]/sw:port[5]/ft:display-name)",
			Expected: "Port 4/1",
		},
		{
			Name:     "port[5] on switch[2] - channel",
			Path:     "number(/sw:switch[2]/sw:port[5]/@sw:channel-number)",
			Expected: float64(1),
		},
		{
			Name:     "test concat string",
			Path:     "concat('Port ', string(/sw:switch[2]/sw:port[5]/@sw:cage-number), '/', string(/sw:switch[2]/sw:port[5]/@sw:channel-number))",
			Expected: "Port 4/1",
		},
		{
			Name:     "test set-contains detects string in a set",
			Path:     "set-contains(/sw:switch/sw:model-id, 'super-switch-1610')",
			Expected: true,
		},
		{
			Name:     "test set-contains detects string is missing from set",
			Path:     "set-contains(/sw:switch/sw:model-id, 'super-switch-1611')",
			Expected: false,
		},
		{
			Name:     "test set-contains to see all model types are used in switches",
			Path:     "set-contains(/sm:switch-model/@sm:switch-model-id, /sw:switch/sw:model-id)",
			Expected: true,
		},
		{
			Name:     "test set-contains to see all model types are used in switches only once",
			Path:     "set-equals(/sm:switch-model/@sm:switch-model-id, /sw:switch/sw:model-id)",
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

// Test_XPathEvaluateRelativePath - start each test from switch[1]/port[2]
func Test_XPathEvaluateRelativePath(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-example-1.json")
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
			Name:     "cage",
			Path:     "number(@sw:cage-number)",
			Expected: float64(2),
		},
		{
			Name:     "channel",
			Path:     "number(@sw:channel-number)",
			Expected: float64(2),
		},
		{
			Name:     "switch port speed",
			Path:     "string(sw:speed)",
			Expected: "speed-40g",
		},
		{
			Name:     "switch port display-name",
			Path:     "string(ft:display-name)",
			Expected: "Port 2/0 on switch 1",
		},
		{
			Name:     "switch model",
			Path:     "string(../sw:model-id)",
			Expected: "super-switch-1610",
		},
		{
			Name:     "model name - absolute ref",
			Path:     "string(/sm:switch-model[@sm:switch-model-id='super-switch-1610']/ft:display-name)",
			Expected: "Super Switch 1610",
		},
		//{
		//	Name:     "name of corresponding model for this port - relative ref", // $this is the port
		//	Path:     "string(/sm:switch-model[@sm:switch-model-id=$this/../sw:model-id]/ft:display-name)",
		//	Expected: "Super Switch 1610",
		//},
		//{
		//	Name:     "number of ports on corresponding model - relative ref", // $this is the port
		//	Path:     "count(/sm:switch-model[@sm:switch-model-id=$this/../sw:model-id]/sm:port)",
		//	Expected: float64(4),
		//},
		//{
		//	Name:     "max channels on corresponding model port - relative ref", // $this is the port
		//	Path:     "number(/sm:switch-model[@sm:switch-model-id=$this/../sw:model-id]/sm:port[@sm:cage-number=$this/@sw:cage-number]/sm:max-channel)",
		//	Expected: float64(2),
		//},
		//{
		//	Name:     "check channel-num <= max channels on corresponding model port - relative ref", // $this is the port
		//	Path:     "./@sw:channel-number <= /sm:switch-model[@sm:switch-model-id=$this/../sw:model-id]/sm:port[@sm:cage-number=$this/@sw:cage-number]/sm:max-channel",
		//	Expected: true,
		//},
		//{
		//	Name:     "check cage numbers - relative ref", // $this is the port
		//	Path:     "set-contains(/sm:switch-model[@sm:switch-model-id=$this/../sw:model-id]/sm:port/@sm:cage-number, ./@sw:cage-number)",
		//	Expected: true,
		//},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild()) // route[1]
		assert.True(t, ynn.MoveToNext())  // route[2]
		assert.True(t, ynn.MoveToNext())  // switch-model[1]
		assert.True(t, ynn.MoveToNext())  // switch-model[2]
		assert.True(t, ynn.MoveToNext())  // switch[1]
		assert.True(t, ynn.MoveToNext())  // switch[2]
		assert.True(t, ynn.MoveToChild()) // switch[2]/attribute[1]
		assert.True(t, ynn.MoveToNext())  // switch[2]/attribute[2]
		assert.True(t, ynn.MoveToNext())  // switch[2]/description
		assert.True(t, ynn.MoveToNext())  // switch[2]/display-name
		assert.True(t, ynn.MoveToNext())  // switch[2]/management
		assert.True(t, ynn.MoveToNext())  // switch[2]/model-id
		assert.True(t, ynn.MoveToNext())  // switch[2]/port[1]
		assert.True(t, ynn.MoveToNext())  // switch[2]/port[2]

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}
