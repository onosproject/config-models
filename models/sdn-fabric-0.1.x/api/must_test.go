/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package api

import (
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_WalkAndValidateMustSucceed(t *testing.T) {
	sampleConfig, err := os.ReadFile("../examples/full-config-example-1.json")
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
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.NoError(t, validateErr)
}

func Test_WalkAndValidateMustFailPortChannel(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-broken-must-port-channel.json")
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
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.EqualError(t, validateErr, `port channel-number exceeds max-channel of corresponding switch-model/port. Must statement 'number(.) <= number(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/max-channel)' to true. Container(s): [context: channel-number=4 cage-number=4]`)
}

func Test_WalkAndValidateMustFailPortCage(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-broken-must-port-cage.json")
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
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.EqualError(t, validateErr, `port cage-number must be present in corresponding switch-model/port. Must statement 'set-contains(/switch-model[@switch-model-id=$this/../../model-id]/port/@cage-number, .)' to true. Container(s): [context: cage-number=3 switch-model-id=super-switch-2100]`)
}

func Test_WalkAndValidateMustFailPortSpeed(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-broken-must-port-speed.json")
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
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.EqualError(t, validateErr, `port speed must be present in corresponding switch-model/port. Must statement 'contains(/switch-model[@switch-model-id=$this/../../model-id]/port[@cage-number=$this/../@cage-number]/speeds, string($this))' to true. Container(s): [context: speed=speed-100g cage-number=4]`)
}
