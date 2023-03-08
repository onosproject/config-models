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

func Test_XPathSelectSwitch(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-example-1.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	assert.NoError(t, err)

	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name:  "test key",
			XPath: "/switch-model/@switch-model-id",
			Expected: []string{
				"Iter Value: switch-model-id: super-switch-1610",
				"Iter Value: switch-model-id: super-switch-2100",
			},
		},
		{
			Name:  "test ordinary attribute",
			XPath: "/switch-model/description",
			Expected: []string{
				"Iter Value: description: Super Switch Model 1610 16-port 10Gb on mars fabric",
				"Iter Value: description: Super Switch Model NIC 100 2-port 100Gb",
			},
		},
		{
			Name:  "test switch-models description",
			XPath: "/switch-model/description",
			Expected: []string{
				"Iter Value: description: Super Switch Model 1610 16-port 10Gb on mars fabric",
				"Iter Value: description: Super Switch Model NIC 100 2-port 100Gb",
			},
		},
		{
			Name:  "test switch model[0] desc",
			XPath: "/switch-model[@switch-model-id='super-switch-1610']/description",
			Expected: []string{
				"Iter Value: description: Super Switch Model 1610 16-port 10Gb on mars fabric",
			},
		},
		{
			Name:  "test switch model[0] desc",
			XPath: "/switch-model[@switch-model-id='super-switch-2100']/description",
			Expected: []string{
				"Iter Value: description: Super Switch Model NIC 100 2-port 100Gb",
			},
		},
		{
			Name:  "test switch model[0] ports",
			XPath: "/switch-model[@switch-model-id='super-switch-1610']/port/@cage-number",
			Expected: []string{
				"Iter Value: cage-number: 1", // duplicate
				"Iter Value: cage-number: 2", // duplicate
				"Iter Value: cage-number: 3",
				"Iter Value: cage-number: 4",
			},
		},
		{
			Name:  "test switch model[0] port 3",
			XPath: "/switch-model[@switch-model-id='super-switch-1610']/port[@cage-number=3]/@cage-number",
			Expected: []string{
				"Iter Value: cage-number: 3",
			},
		},
		{
			Name:  "test switch[0] port 3 name",
			XPath: "/switch[@switch-id='san-jose-edge-tor-1S']/port[@cage-number=3]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 3/0",
			},
		},
		{
			Name:  "test switch by model port 3 name", // 2 switches with this model
			XPath: "/switch[model-id='super-switch-1610']/port[@cage-number=3]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 3/0",
				"Iter Value: display-name: Port 3/0",
			},
		},
		{
			Name:  "test switch-model by switch name",
			XPath: "/switch-model[@switch-model-id=/switch[@switch-id='san-jose-edge-tor-1S']/model-id]/display-name",
			Expected: []string{
				"Iter Value: display-name: Super Switch 1610",
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

// Test_XPathSelectRelativeStart - start each test from switch[1] - the thing that contains all the port entries
func Test_XPathSelectSwitchRelativeStart(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-example-1.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	assert.NoError(t, err)

	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)
	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)

	tests := []navigator.XpathSelect{
		{
			Name:  "test name",
			XPath: "./display-name",
			Expected: []string{
				"Iter Value: display-name: SJ TOR S2",
			},
		},
		{
			Name:  "port[2] name by id",
			XPath: "port[@cage-number=2][@channel-number=0]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 2/0 on switch 2",
			},
		},
		{
			Name:  "port[2] name by number",
			XPath: "port[2]/display-name",
			Expected: []string{
				"Iter Value: display-name: Port 2/0 on switch 2",
			},
		},
		{
			Name:     "port[2] preceding-sibling",
			XPath:    "port[@cage-number=2][@channel-number=0]/preceding-sibling::node()/display-name",
			Expected: []string{}, // There is no preceding sibling
		},
		{
			Name:  "port[2] following-sibling",
			XPath: "port[@cage-number=2][@channel-number=0]/following-sibling::node()/display-name", // There are 4 following siblings, but only 1 is a port
			Expected: []string{
				"Iter Value: display-name: Port 3/0", // lists following nodes regardless of type
			},
		},
		{
			Name: "test following-sibling who has same tx-power as current",
			// following-sibling below returns a node-set which is inadvertently cast to a string
			// which will extract only the first entry and then cast to string = "6"
			// and it will detect a match only when processing node l2a5 (as l2a6) has a similar value.
			// This means that to detect duplicate nodes the nodes will have to be sorted in order of tx-power
			// This is not currently done, as nodes are sorted by their @name attribute
			XPath:    "port[set-contains(following-sibling::port/display-name, display-name)]/display-name",
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
		expr, err1 := xpath.Compile(test.XPath)
		assert.NoError(t, err1, test.Name)
		if err1 != nil {
			t.FailNow()
		}
		assert.NotNil(t, expr, test.Name)

		navigateErr := ynn.NavigateTo("/switch[switch-id=san-jose-edge-tor-2S]")
		assert.NoError(t, navigateErr)

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.Expected)-1, test.Name, ". More results than expected")
			// Note if you get errors here after adding things to ../examples/switch-config-example-1.json
			// you should run this in DEBUG and monitor the value of "ynn.Name" for each of the steps after ynn.MoveToRoot() above
			// The value in the comment is where you should be when debugging that line
			assert.Equal(t, test.Expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.Name)
			resultCount++
		}
		assert.Equal(t, len(test.Expected), resultCount, "%s. Did not receive all the expected results", test.Name)
	}
}

func Test_XPathSwitchEvaluate(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-example-1.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	assert.NoError(t, err)

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
			XPath:    "count(/switch-model[1]/port/@cage-number)",
			Expected: float64(4),
		},
		{
			Name:     "ports on switch[2]",
			XPath:    "count(/switch[2]/port/@cage-number)",
			Expected: float64(5),
		},
		{
			Name:     "port[5] on switch[2] - name",
			XPath:    "string(/switch[2]/port[5]/display-name)",
			Expected: "Port 4/1",
		},
		{
			Name:     "port[5] on switch[2] - channel",
			XPath:    "number(/switch[2]/port[5]/@channel-number)",
			Expected: float64(1),
		},
		{
			Name:     "test concat string",
			XPath:    "concat('Port ', string(/switch[2]/port[5]/@cage-number), '/', string(/switch[2]/port[5]/@channel-number))",
			Expected: "Port 4/1",
		},
		{
			Name:     "test set-contains detects string in a set",
			XPath:    "set-contains(/switch/model-id, 'super-switch-1610')",
			Expected: true,
		},
		{
			Name:     "test set-contains detects string is missing from set",
			XPath:    "set-contains(/switch/model-id, 'super-switch-1611')",
			Expected: false,
		},
		{
			Name:     "test set-contains to see all model types are used in switches",
			XPath:    "set-contains(/switch-model/@switch-model-id, /switch/model-id)",
			Expected: true,
		},
		{
			Name:     "test set-contains to see all model types are used in switches only once",
			XPath:    "set-equals(/switch-model/@switch-model-id, /switch/model-id)",
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

// Test_XPathEvaluateRelativePath - start each test from switch[1]/port[2]
func Test_XPathEvaluateSwitchRelativePath(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-example-1.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	assert.NoError(t, err)

	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)
	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "cage",
			XPath:    "number(@cage-number)",
			Expected: float64(2),
		},
		{
			Name:     "channel",
			XPath:    "number(@channel-number)",
			Expected: float64(2),
		},
		{
			Name:     "switch port display-name",
			XPath:    "string(display-name)",
			Expected: "Port 2/0 on switch 1",
		},
		{
			Name:     "switch model",
			XPath:    "string(../model-id)",
			Expected: "super-switch-1610",
		},
		{
			Name:     "model name - absolute ref",
			XPath:    "string(/switch-model[@switch-model-id='super-switch-1610']/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "name of corresponding model for this port - relative ref", // $this is the port
			XPath:    "string(/switch-model[@switch-model-id=$this/../model-id]/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "number of ports on corresponding model - relative ref", // $this is the port
			XPath:    "count(/switch-model[@switch-model-id=$this/../model-id]/port)",
			Expected: float64(4),
		},
		{
			Name:     "max channels on corresponding model port - relative ref", // $this is the port
			XPath:    "number(/switch-model[@switch-model-id=$this/../model-id]/port[@cage-number=$this/@cage-number]/max-channel)",
			Expected: float64(2),
		},
		{
			Name:     "check channel-num <= max channels on corresponding model port - relative ref", // $this is the port
			XPath:    "number(./@channel-number) <= number(/switch-model[@switch-model-id=$this/../model-id]/port[@cage-number=$this/@cage-number]/max-channel)",
			Expected: true,
		},
		{
			Name:     "check cage numbers - relative ref", // $this is the port
			XPath:    "set-contains(/switch-model[@switch-model-id=$this/../model-id]/port/@cage-number, ./@cage-number)",
			Expected: true,
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.XPath)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		navigateErr := ynn.NavigateTo("/switch[switch-id=san-jose-edge-tor-1S]/port[cage-number=2][channel-number=2]")
		assert.NoError(t, navigateErr)

		result := expr.Evaluate(ynn)
		// Note if you get errors here after adding things to ../examples/switch-config-example-1.json
		// you should run this in DEBUG and monitor the value of "ynn.Name" for each of the steps after ynn.MoveToRoot() above
		// The value in the comment is where you should be when debugging that line
		assert.Equal(t, test.Expected, result, test.Name)
	}
}

// Test_XPathEvaluateRelativePath - start each test from switch[1]/port[2]
func Test_XPathEvaluateSwitchRelativePathChannelNumber(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-example-1.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	assert.NoError(t, err)

	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)
	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "cage",
			XPath:    "number(../@cage-number)",
			Expected: float64(2),
		},
		{
			Name:     "channel",
			XPath:    "number(.)",
			Expected: float64(2),
		},
		{
			Name:     "switch port display-name",
			XPath:    "string(../display-name)",
			Expected: "Port 2/0 on switch 1",
		},
		{
			Name:     "switch model",
			XPath:    "string(../../model-id)",
			Expected: "super-switch-1610",
		},
		{
			Name:     "model name - absolute ref",
			XPath:    "string(/switch-model[@switch-model-id='super-switch-1610']/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "name of corresponding model for this port - relative ref", // $this is the port
			XPath:    "string(/switch-model[@switch-model-id=$this/../../model-id]/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "number of ports on corresponding model - relative ref", // $this is the port
			XPath:    "count(/switch-model[@switch-model-id=$this/../../model-id]/port)",
			Expected: float64(4),
		},
		{
			Name:     "max channels on corresponding model port - relative ref", // $this is the port
			XPath:    "number(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this]/max-channel)",
			Expected: float64(2),
		},
		{
			Name:     "check channel-num <= max channels on corresponding model port - relative ref", // $this is the port
			XPath:    "number(.) <= number(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/max-channel)",
			Expected: true,
		},
		{
			Name:     "check cage numbers - relative ref", // $this is the port
			XPath:    "set-contains(/switch-model[@switch-model-id=$this/../../model-id]/port/@cage-number, .)",
			Expected: true,
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.XPath)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		navigateErr := ynn.NavigateTo("/switch[switch-id=san-jose-edge-tor-1S]/port[cage-number=2][channel-number=2]/channel-number")
		assert.NoError(t, navigateErr)

		result := expr.Evaluate(ynn)
		// Note if you get errors here after adding things to ../examples/switch-config-example-1.json
		// you should run this in DEBUG and monitor the value of "ynn.Name" for each of the steps after ynn.MoveToRoot() above
		// The value in the comment is where you should be when debugging that line
		assert.Equal(t, test.Expected, result, test.Name)
	}
}

// Test_XPathEvaluateRelativePath - start each test from switch[1]/port[2]
func Test_XPathEvaluateSwitchRelativePathSpeeds(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-example-1.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	assert.NoError(t, err)

	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)
	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)

	tests := []navigator.XpathEvaluate{
		//{
		//	Name:     "cage",
		//	Path:     "number(../@sw:cage-number)",
		//	Expected: float64(2),
		//},
		{
			Name:     "channel",
			XPath:    "string(../speed)",
			Expected: "speed-10g",
		},
		{
			Name:     "current speed",
			XPath:    "string($this)",
			Expected: "speed-10g",
		},
		{
			Name:     "switch port display-name",
			XPath:    "string(../display-name)",
			Expected: "Port 2/0 on switch 1",
		},
		{
			Name:     "switch model",
			XPath:    "string(../../model-id)",
			Expected: "super-switch-1610",
		},
		{
			Name:     "name of corresponding model for this port - relative ref", // $this is the port
			XPath:    "string(/switch-model[@switch-model-id=$this/../../model-id]/display-name)",
			Expected: "Super Switch 1610",
		},
		{
			Name:     "count of speeds for all ports in corresponding switch model - relative ref", // $this is the speed
			XPath:    "count(/switch-model[@switch-model-id=$this/../../model-id]/port/speeds)",
			Expected: float64(4),
		},
		{
			Name:     "count of speeds for current port in corresponding switch model - relative ref", // $this is the speed
			XPath:    "count(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/speeds)",
			Expected: float64(1), // 1 array
		},
		{
			Name:     "list of speeds for current port in corresponding switch model - relative ref", // $this is the speed
			XPath:    "string(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/speeds)",
			Expected: "[speed-1g speed-10g]",
		},
		{
			Name:     "list of speeds for current port in corresponding switch model - use prefixes - relative ref", // $this is the speed
			XPath:    "string(/sm:switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/speeds)",
			Expected: "[speed-1g speed-10g]",
		},
		{
			Name:     "corresponding switch model contains speed-1g - relative ref", // $this is the speed
			XPath:    "contains(/sm:switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/speeds, string('speed-1g'))",
			Expected: bool(true),
		},
		{
			Name:     "corresponding switch model contains speed-100g - relative ref", // $this is the speed
			XPath:    "contains(/sm:switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/speeds, string('speed-100g'))",
			Expected: bool(false),
		},
		{
			Name:     "corresponding switch model contains current speed - relative ref", // $this is the speed
			XPath:    "contains(/sm:switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/speeds, string($this))",
			Expected: bool(true),
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.XPath)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		navigateErr := ynn.NavigateTo("/switch[switch-id=san-jose-edge-tor-1S]/port[cage-number=2][channel-number=2]/speed")
		assert.NoError(t, navigateErr)

		result := expr.Evaluate(ynn)
		// Note if you get errors here after adding things to ../examples/switch-config-example-1.json
		// you should run this in DEBUG and monitor the value of "ynn.Name" for each of the steps after ynn.MoveToRoot() above
		// The value in the comment is where you should be when debugging that line
		assert.Equal(t, test.Expected, result, test.Name)
	}
}
