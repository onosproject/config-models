/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package api

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_Validate(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../examples/full-config-example-1.json")
	assert.NoError(t, err)
	assert.NotNil(t, sampleConfig)

	device := new(Device)
	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	assert.NoError(t, err)

	err = device.Validate()
	assert.NoError(t, err)
}
