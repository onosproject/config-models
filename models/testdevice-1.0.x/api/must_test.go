// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func Test_WalkAndValidateMustSucceed(t *testing.T) {
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
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, true)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.NoError(t, validateErr)
}

func Test_WalkAndValidateMustMinMaxFailure(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config-min-max.json")
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
	assert.True(t, strings.HasPrefix(validateErr.Error(), "range-min must be less than or equal to range-max. Must statement 'number(./range-min) <= number(./range-max)' to true. Container(s): [context: list2a="), validateErr)
	assert.True(t, strings.HasSuffix(validateErr.Error(), "name=l2a1]"), validateErr)
}

func Test_WalkAndValidateMustFailureList2a(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config-must-list2a-false.json")
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
	assert.True(t, strings.HasPrefix(validateErr.Error(), "tx-power must not be repeated in list2a. Must statement 'not(list2a[set-contains(following-sibling::list2a/tx-power, tx-power)])' to true. Container(s): [context:"), validateErr)
	assert.True(t, strings.HasSuffix(validateErr.Error(), "name=l2a2]"), validateErr)

}

func Test_WalkAndValidateMustFailureLeaf4a(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config-must-list4akeys-false.json")
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
	assert.True(t, strings.HasPrefix(validateErr.Error(), "displayname must be formatted string like 'Value <../id>-<fkey1>-<fkey2>'. Must statement 'concat('Value ', string(../@id), '-', string(./@fkey1), '-', string(./@fkey2)) = string(./displayname)' to true. Container(s): [context: list4a="), validateErr)
	assert.True(t, strings.HasSuffix(validateErr.Error(), "fkey1=five fkey2=6]"), validateErr)
}

func Test_WalkAndValidateMustFailureLeaf5a(t *testing.T) {
	sampleConfig, err := os.ReadFile("../testdata/sample-testdevice-1-config-must-leaf5a-false.json")
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
	assert.True(t, strings.HasPrefix(validateErr.Error(), "leaf5a must be formatted string like '5a <key1>-<key2>'. Must statement 'concat('5a ', string(./@key1), '-', string(./@key2)) = string(./leaf5a)' to true. Container(s): [context: list5="), validateErr)
	assert.True(t, strings.HasSuffix(validateErr.Error(), "key1=eight key2=8]"), validateErr)
}
