// Copyright 2019-present Open Networking Foundation.
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

// A plugin for the YGOT model of enterprise-1.0.0.
package modelplugin

import (
	"fmt"
	_ "github.com/golang/protobuf/proto"
	"github.com/onosproject/config-models/modelplugin/enterprise-1.0.0/enterprise_1_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	_ "github.com/openconfig/ygot/genutil"
	_ "github.com/openconfig/ygot/ygen"
	"github.com/openconfig/ygot/ygot"
	_ "github.com/openconfig/ygot/ytypes"
)

type ModelPlugin string

const modeltype = "Enterprise"
const modelversion = "1.0.0"
const modulename = "enterprise.so.1.0.0"

var ModelData = []*gnmi.ModelData{
	{Name: "device-range", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "device-group", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "enterprise", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "site", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "use-case", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "application", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "deployment-environment", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "qos", Organization: "Open Networking Foundation", Version: "2020-11-18"},
	{Name: "visibility", Organization: "Open Networking Foundation", Version: "2020-11-18"},
}

func (m ModelPlugin) ModelData() (string, string, []*gnmi.ModelData, string) {
	return modeltype, modelversion, ModelData, modulename
}

// UnmarshallConfigValues allows Device to implement the Unmarshaller interface
func (m ModelPlugin) UnmarshalConfigValues(jsonTree []byte) (*ygot.ValidatedGoStruct, error) {
	device := &enterprise_1_0_0.Device{}
	vgs := ygot.ValidatedGoStruct(device)

	if err := enterprise_1_0_0.Unmarshal([]byte(jsonTree), device); err != nil {
		return nil, err
	}

	return &vgs, nil
}

func (m ModelPlugin) Validate(ygotModel *ygot.ValidatedGoStruct, opts ...ygot.ValidationOption) error {
	deviceDeref := *ygotModel
	device, ok := deviceDeref.(*enterprise_1_0_0.Device)
	if !ok {
		return fmt.Errorf("unable to convert model in to enterprise_1_0_0")
	}
	return device.Validate()
}

func (m ModelPlugin) Schema() (map[string]*yang.Entry, error) {
	return enterprise_1_0_0.UnzipSchema()
}

// GetStateMode returns an int - we do not use the enum because we do not want a
// direct dependency on onos-config code (for build optimization)
func (m ModelPlugin) GetStateMode() int {
	return 0 // modelregistry.GetStateNone
}
