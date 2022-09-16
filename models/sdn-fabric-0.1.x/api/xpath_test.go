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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "test key",
			Path: "/switch-model/@switch-model-id",
			Expected: []string{
				"Iter Value: switch-model-id: super-switch-1610",
				"Iter Value: switch-model-id: super-switch-2100",
			},
		},
		{
			Name: "test ordinary attribute",
			Path: "/switch-model/pipeline",
			Expected: []string{
				"Iter Value: pipeline: dual",
				"Iter Value: pipeline: dual",
			},
		},
		{
			Name: "test switch-models description",
			Path: "/switch-model/description",
			Expected: []string{
				"Iter Value: description: Super Switch Model 1610 16-port 10Gb on mars fabric",
				"Iter Value: description: Super Switch Model NIC 100 2-port 100Gb",
			},
		},
		{
			Name: "test switch model[0] desc",
			Path: "/switch-model[@switch-model-id='super-switch-1610']/description",
			Expected: []string{
				"Iter Value: description: Super Switch Model 1610 16-port 10Gb on mars fabric",
			},
		},
		{
			Name: "test switch model[0] desc",
			Path: "/switch-model[@switch-model-id='super-switch-2100']/description",
			Expected: []string{
				"Iter Value: description: Super Switch Model NIC 100 2-port 100Gb",
			},
		},
		{
			Name: "test switch model[0] ports",
			Path: "/switch-model[@switch-model-id='super-switch-1610']/port/@cage-number",
			Expected: []string{
				"Iter Value: cage-number: 1", // duplicate
				"Iter Value: cage-number: 2", // duplicate
				"Iter Value: cage-number: 3",
				"Iter Value: cage-number: 4",
			},
		},
		{
			Name: "test switch model[0] port 3",
			Path: "/switch-model[@switch-model-id='super-switch-1610']/port[@cage-number=3]/@cage-number",
			Expected: []string{
				"Iter Value: cage-number: 3",
			},
		},
		{
			Name: "test switch[0] port 3 name",
			Path: "/switch[@switch-id='san-jose-edge-tor-1S']/port[@cage-number=3]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 3/0",
			},
		},
		{
			Name: "test switch by model port 3 name", // 2 switches with this model
			Path: "/switch[model-id='super-switch-1610']/port[@cage-number=3]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 3/0",
				"Iter Value: display-name: Port 3/0",
			},
		},
		{
			Name: "test switch-model by switch name",
			Path: "/switch-model[@switch-model-id=/switch[@switch-id='san-jose-edge-tor-1S']/model-id]/display-name",
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "test name",
			Path: "./display-name",
			Expected: []string{
				"Iter Value: display-name: SJ TOR S2",
			},
		},
		{
			Name: "port[2] name by id",
			Path: "port[@cage-number=2][@channel-number=0]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 2/0 on switch 2",
			},
		},
		{
			Name: "port[2] name by number",
			Path: "port[2]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 2/0 on switch 2",
			},
		},
		{
			Name:     "port[2] preceding-sibling",
			Path:     "port[@cage-number=2][@channel-number=0]/preceding-sibling::node()/display-name",
			Expected: []string{}, // There is no preceding sibling
		},
		{
			Name: "port[2] following-sibling",
			Path: "port[@cage-number=2][@channel-number=0]/following-sibling::node()/display-name", // There are 4 following siblings, but only 1 is a port
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
			Path:     "port[set-contains(following-sibling::port/display-name, display-name)]/display-name",
			Expected: []string{},
		},
		//{
		//	Name: "this variable test basic",
		//	// using the $this variable - the node which is current at the start of the query
		//	Path: "$this/display-name",
		//	Expected: []string{
		//		"Iter Value: display-name: SJ TOR S2",
		//	},
		//},
		//{
		//	Name: "this variable test advanced",
		//	Path: "/switch-model[@switch-model-id=$this/model-id]/display-name",
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "ports on switch-model[1]",
			Path:     "count(/switch-model[1]/port/@cage-number)",
			Expected: float64(4),
		},
		{
			Name:     "ports on switch[2]",
			Path:     "count(/switch[2]/port/@cage-number)",
			Expected: float64(5),
		},
		{
			Name:     "port[5] on switch[2] - name",
			Path:     "string(/switch[2]/port[5]/display-name)",
			Expected: "Port 4/1",
		},
		{
			Name:     "port[5] on switch[2] - channel",
			Path:     "number(/switch[2]/port[5]/@channel-number)",
			Expected: float64(1),
		},
		{
			Name:     "test concat string",
			Path:     "concat('Port ', string(/switch[2]/port[5]/@cage-number), '/', string(/switch[2]/port[5]/@channel-number))",
			Expected: "Port 4/1",
		},
		{
			Name:     "test set-contains detects string in a set",
			Path:     "set-contains(/switch/model-id, 'super-switch-1610')",
			Expected: true,
		},
		{
			Name:     "test set-contains detects string is missing from set",
			Path:     "set-contains(/switch/model-id, 'super-switch-1611')",
			Expected: false,
		},
		{
			Name:     "test set-contains to see all model types are used in switches",
			Path:     "set-contains(/switch-model/@switch-model-id, /switch/model-id)",
			Expected: true,
		},
		{
			Name:     "test set-contains to see all model types are used in switches only once",
			Path:     "set-equals(/switch-model/@switch-model-id, /switch/model-id)",
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "cage",
			Path:     "number(@cage-number)",
			Expected: float64(2),
		},
		{
			Name:     "channel",
			Path:     "number(@channel-number)",
			Expected: float64(2),
		},
		{
			Name:     "switch port display-name",
			Path:     "string(display-name)",
			Expected: "Port 2/0 on switch 1",
		},
		{
			Name:     "switch model",
			Path:     "string(../model-id)",
			Expected: "super-switch-1610",
		},
		{
			Name:     "model name - absolute ref",
			Path:     "string(/switch-model[@switch-model-id='super-switch-1610']/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "name of corresponding model for this port - relative ref", // $this is the port
			Path:     "string(/switch-model[@switch-model-id=$this/../model-id]/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "number of ports on corresponding model - relative ref", // $this is the port
			Path:     "count(/switch-model[@switch-model-id=$this/../model-id]/port)",
			Expected: float64(4),
		},
		{
			Name:     "max channels on corresponding model port - relative ref", // $this is the port
			Path:     "number(/switch-model[@switch-model-id=$this/../model-id]/port[@cage-number=$this/@cage-number]/max-channel)",
			Expected: float64(2),
		},
		{
			Name:     "check channel-num <= max channels on corresponding model port - relative ref", // $this is the port
			Path:     "number(./@channel-number) <= number(/switch-model[@switch-model-id=$this/../model-id]/port[@cage-number=$this/@cage-number]/max-channel)",
			Expected: true,
		},
		{
			Name:     "check cage numbers - relative ref", // $this is the port
			Path:     "set-contains(/switch-model[@switch-model-id=$this/../model-id]/port/@cage-number, ./@cage-number)",
			Expected: true,
		},
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

// Test_XPathEvaluateRelativePath - start each test from switch[1]/port[2]
func Test_XPathEvaluateRelativePathChannelNumber(t *testing.T) {
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "cage",
			Path:     "number(../@cage-number)",
			Expected: float64(2),
		},
		{
			Name:     "channel",
			Path:     "number(.)",
			Expected: float64(2),
		},
		{
			Name:     "switch port display-name",
			Path:     "string(../display-name)",
			Expected: "Port 2/0 on switch 1",
		},
		{
			Name:     "switch model",
			Path:     "string(../../model-id)",
			Expected: "super-switch-1610",
		},
		{
			Name:     "model name - absolute ref",
			Path:     "string(/switch-model[@switch-model-id='super-switch-1610']/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "name of corresponding model for this port - relative ref", // $this is the port
			Path:     "string(/switch-model[@switch-model-id=$this/../../model-id]/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "number of ports on corresponding model - relative ref", // $this is the port
			Path:     "count(/switch-model[@switch-model-id=$this/../../model-id]/port)",
			Expected: float64(4),
		},
		{
			Name:     "max channels on corresponding model port - relative ref", // $this is the port
			Path:     "number(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this]/max-channel)",
			Expected: float64(2),
		},
		{
			Name:     "check channel-num <= max channels on corresponding model port - relative ref", // $this is the port
			Path:     "number(.) <= number(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/max-channel)",
			Expected: true,
		},
		{
			Name:     "check cage numbers - relative ref", // $this is the port
			Path:     "set-contains(/switch-model[@switch-model-id=$this/../../model-id]/port/@cage-number, .)",
			Expected: true,
		},
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
		assert.True(t, ynn.MoveToChild()) // switch[2]/port[2]/cage-number
		assert.True(t, ynn.MoveToNext())  // switch[2]/port[2]/channel-number

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}
