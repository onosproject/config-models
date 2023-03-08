// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package api

import (
	"gotest.tools/assert"
	"strings"
	"testing"
)

func Test_generated(t *testing.T) {
	assert.Equal(t, 21, len(SchemaTree))

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
		case "OnfTest1Choice_Vehicle_ElectricMotor":
			assert.Equal(t, "Motor configuration - demonstrates a list inside a choice", v.Description)
		case "OnfSwitch_Switch":
			assert.Assert(t, strings.HasPrefix(v.Description, "A managed device in the fabric"))
		case "OnfSwitch_Switch_State":
			assert.Equal(t, "Op state attributes", v.Description)
		case "OnfSwitchModel_SwitchModel":
			assert.Assert(t, strings.HasPrefix(v.Description, "A model of switch"), v.Description)
		case "OnfTest1Choice_Vehicle_UnderCarriage":
			assert.Equal(t, "Traction details", v.Description)
		case "OnfSwitchModel_SwitchModel_Attribute":
			assert.Equal(t, "a map of extra attributes: string-string", v.Description)
		case "OnfSwitchModel_SwitchModel_Port":
			assert.Equal(t, "A port in a switch - this demonstrates a list within a list. Each port has a description\nand a display-name", v.Description)
		case "OnfSwitch_Switch_Port":
			assert.Assert(t, strings.HasPrefix(v.Description, "A port in a switch. This demonstrates a lot advanced functionality in the ROC-GUI."))
		case "OnfSwitch_Switch_Attribute":
			assert.Equal(t, "a map of extra attributes: string-string", v.Description)
		case "OnfTest1Choice_Vehicle":
			assert.Equal(t, "A list of vehicles", v.Description)
		case "OnfTest1Choice_Vehicle_Battery":
			assert.Equal(t, "Battery configuration", v.Description)
		default:
			t.Errorf("unexpected schema entry %s", k)
		}
	}
}
