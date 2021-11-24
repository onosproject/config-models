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

package aether_4_0_0

import (
	"fmt"
	"github.com/antchfx/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_XPathSelect(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-aether4-config.json")
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
			Name: "test applications",
			Path: "/app:application/app:application",
			Expected: []string{
				"Iter Value: application: value of application",
				"Iter Value: application: value of application",
				"Iter Value: application: value of application",
			},
		},
		{
			Name: "test application acme",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/@app:id",
			Expected: []string{
				"Iter Value: id: acme-dataacquisition",
			},
		},
		{
			Name: "test application acme id",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/@app:id",
			Expected: []string{
				"Iter Value: id: acme-dataacquisition",
			},
		},
		{
			Name: "test application acme address",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/app:address",
			Expected: []string{
				"Iter Value: address: da.acme.com",
			},
		},
		{
			Name: "test application acme endpoints",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/app:endpoint",
			Expected: []string{
				"Iter Value: endpoint: value of endpoint",
				"Iter Value: endpoint: value of endpoint",
			},
		},
		{
			Name: "test application acme endpoints ids",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/app:endpoint/@app:endpoint-id",
			Expected: []string{
				"Iter Value: endpoint-id: da1",
				"Iter Value: endpoint-id: da2",
			},
		},
		{
			Name: "test application acme endpoint da1 port-start",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/app:endpoint[@app:endpoint-id='da1']/app:port-start",
			Expected: []string{
				"Iter Value: port-start: 1230",
			},
		},
		{
			Name: "test application acme endpoint da2 port-end",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/app:endpoint[@app:endpoint-id='da2']/app:port-end",
			Expected: []string{
				"Iter Value: port-end: 1010",
			},
		},
		{
			Name: "test application acme endpoint da2 mbr downlink",
			Path: "/app:application/app:application[@app:id='acme-dataacquisition']/app:endpoint[@app:endpoint-id='da2']/app:mbr/app:downlink",
			Expected: []string{
				"Iter Value: downlink: 3000000",
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
		if resultCount < len(test.Expected) {
			t.FailNow()
		}
	}
}
