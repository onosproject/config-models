/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package api

import (
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_WalkAndValidateMustFailPortChannel(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-broken-must-port-channel.json")
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
	validateErr := ynn.WalkAndValidateMust()
	assert.EqualError(t, validateErr, `port channel-number exceeds max-channel of corresponding switch-model/port. Must statement 'number(.) <= number(/sm:switch-model[@sm:switch-model-id=$this/../../model-id]/sm:port[@cage-number=$this/../@cage-number]/sm:max-channel)' to true. Container(s): [context: channel-number=4 cage-number=4]`)
}

func Test_WalkAndValidateMustFailPortCage(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-broken-must-port-cage.json")
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
	validateErr := ynn.WalkAndValidateMust()
	assert.EqualError(t, validateErr, `port cage-number must be present in corresponding switch-model/port. Must statement 'set-contains(/sm:switch-model[@sm:switch-model-id=$this/../../model-id]/sm:port/@sm:cage-number, .)' to true. Container(s): [context: cage-number=3 switch-model-id=super-switch-2100]`)
}

func Test_WalkAndValidateMustFailPortSpeed(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/switch-config-broken-must-port-speed.json")
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
	validateErr := ynn.WalkAndValidateMust()
	assert.EqualError(t, validateErr, `port speed must be present in corresponding switch-model/port. Must statement 'contains(/sm:switch-model[@sm:switch-model-id=$this/../../model-id]/sm:port[@cage-number=$this/../@cage-number]/sm:speeds, string($this))' to true. Container(s): [context: speed=speed-100g cage-number=4]`)
}
