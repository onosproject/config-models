// Copyright 2021-present Open Networking Foundation.
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

package testdevice_2_0_0

import (
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_WalkAndValidateMust(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice2-config.json")
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
