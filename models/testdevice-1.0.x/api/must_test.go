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
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_WalkAndValidateMustSucceed(t *testing.T) {
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
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.NoError(t, validateErr)
}

func Test_WalkAndValidateMustMinMaxFailure(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config-min-max.json")
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
	assert.Error(t, validateErr)
}

func Test_WalkAndValidateMustFailureList2a(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config-must-list2a-false.json")
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
	assert.EqualError(t, validateErr, "tx-power must not be repeated in list2a. Must statement 'not(t1:list2a[set-contains(following-sibling::t1:list2a/t1:tx-power, t1:tx-power)])' to true. Container(s): [name=l2a2]")
}

func Test_WalkAndValidateMustFailureLeaf4a(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config-must-list4akeys-false.json")
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
	assert.EqualError(t, validateErr, "displayname must be formatted string like 'Value <../id>-<fkey1>-<fkey2>'. Must statement 'concat('Value ', string(../@t1e:id), '-', string(./@t1e:fkey1), '-', string(./@t1e:fkey2)) = string(./t1e:displayname)' to true. Container(s): [fkey1=five fkey2=6]")
}

func Test_WalkAndValidateMustFailureLeaf5a(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config-must-leaf5a-false.json")
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
	assert.EqualError(t, validateErr, "leaf5a must be formatted string like '5a <key1>-<key2>'. Must statement 'concat('5a ', string(./@t1e:key1), '-', string(./@t1e:key2)) = string(./t1e:leaf5a)' to true. Container(s): [key1=eight key2=8]")
}
