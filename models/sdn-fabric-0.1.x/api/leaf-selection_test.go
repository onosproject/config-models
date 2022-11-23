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

func Test_LeafSelection_Cages(t *testing.T) {
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

	ynn, ok := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ok)

	err = ynn.NavigateTo("/switch[switch-id=san-jose-edge-tor-1S]/port[cage-number=2][channel-number=2]/cage-number")
	assert.NoError(t, err)

	selection, err := ynn.LeafSelection()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(selection))
	assert.EqualValues(t, []string{"1", "2", "3", "4"}, selection)
}

func Test_LeafSelection_Speeds(t *testing.T) {
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

	ynn, ok := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ok)

	err = ynn.NavigateTo("/switch[switch-id=san-jose-edge-tor-1S]/port[cage-number=2][channel-number=2]/speed")
	assert.NoError(t, err)

	selection, err := ynn.LeafSelection()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(selection))
	// because it's a leaf list it will have all entries in a string encoded array
	assert.Equal(t, "[speed-1g speed-10g]", selection[0])
}
