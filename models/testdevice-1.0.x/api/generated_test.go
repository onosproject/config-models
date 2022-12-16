// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package api

import (
	"gotest.tools/assert"
	"testing"
)

func Test_generated(t *testing.T) {
	assert.Equal(t, 10, len(SchemaTree))

	for k, v := range SchemaTree {
		switch k {
		case "Device":
			assert.Equal(t, "", v.Description)
		case "OnfTest1_List1A":
			assert.Equal(t, "A list at the top level", v.Description)
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
