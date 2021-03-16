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

package openapi_gen

import (
	"github.com/openconfig/goyang/pkg/yang"
	"gotest.tools/assert"
	"testing"
)

// Test the range min..10 | 20..100
func Test_yangRange(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Kind:  yang.MinNumber,
			Value: 0,
		},
		Max: yang.Number{
			Kind:  yang.Positive,
			Value: 10,
		},
	})
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Kind:  yang.Positive,
			Value: 20,
		},
		Max: yang.Number{
			Kind:  yang.Positive,
			Value: 100,
		},
	})

	min, max, err := yangRange(testRange1)
	assert.NilError(t, err)
	assert.Assert(t, min == nil)
	assert.Assert(t, max != nil)
	if max != nil {
		assert.Equal(t, uint64(100), *max)
	}
}
