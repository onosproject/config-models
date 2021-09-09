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
package testdevice_1_0_0

import (
	"gotest.tools/assert"
	"testing"
)

func Test_generated(t *testing.T) {
	assert.Equal(t, 9, len(SchemaTree))

	for k, v := range SchemaTree {
		switch k {
		case "Device":
			assert.Equal(t, "", v.Description)
		case "OnfTest1_Cont1A":
			assert.Equal(t, "The top level container", v.Description)
		case "OnfTest1_Cont1A_Cont2A":
			assert.Equal(t, "The 2nd level container", v.Description)
		case "OnfTest1_Cont1A_List2A":
			assert.Equal(t, "A simple list of configuration items", v.Description)
		case "OnfTest1_Cont1BState":
			assert.Equal(t, "A second top level container - this one for state attributes. Edit symbol should not be visible", v.Description)
		case "OnfTest1_Cont1BState_List2B":
			assert.Equal(t, "A simple list of state items", v.Description)
		case "OnfTest1_Cont1A_List4":
			assert.Equal(t, "A list with a leafref index", v.Description)
		case "OnfTest1_Cont1A_List5":
			assert.Equal(t, "A list with 2 keys", v.Description)
		case "OnfTest1_Cont1A_List4_List4A":
			assert.Equal(t, "A list within a list with 2 keys as leaf refs", v.Description)
		default:
			t.Errorf("unexpected schema entry %s", k)
		}
	}
}
