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

// A plugin for the YGOT model of aether-3.0.0.
package modelplugin

import (
	"fmt"

	_ "github.com/golang/protobuf/proto"
	"github.com/onosproject/config-models/modelplugin/aether-3.0.0/aether_3_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	_ "github.com/openconfig/ygot/genutil"
	_ "github.com/openconfig/ygot/ygen"
	"github.com/openconfig/ygot/ygot"
	_ "github.com/openconfig/ygot/ytypes"
)

type Modelplugin string

const modeltype = "Aether"
const modelversion = "3.0.0"
const modulename = "aether.so.3.0.0"

var ModelData = []*gnmi.ModelData{
	{Name: "connectivity-service", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "enterprise", Organization: "Open Networking Foundation", Version: "2021-06-02"},

	{Name: "aether-types", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "ap-list", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "application", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "device-group", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "device-model-list", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "ip-domain", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "network", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "site", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "upf", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "vcs", Organization: "Open Networking Foundation", Version: "2021-06-02"},
	{Name: "template", Organization: "Open Networking Foundation", Version: "2021-06-02"},
}

func (m Modelplugin) ModelData() (string, string, []*gnmi.ModelData, string) {
	return modeltype, modelversion, ModelData, modulename
}

// UnmarshallConfigValues allows Device to implement the Unmarshaller interface
func (m Modelplugin) UnmarshalConfigValues(jsonTree []byte) (*ygot.ValidatedGoStruct, error) {
	device := &aether_3_0_0.Device{}
	vgs := ygot.ValidatedGoStruct(device)

	if err := aether_3_0_0.Unmarshal([]byte(jsonTree), device); err != nil {
		return nil, err
	}

	return &vgs, nil
}

func (m Modelplugin) Validate(ygotModel *ygot.ValidatedGoStruct, opts ...ygot.ValidationOption) error {
	deviceDeref := *ygotModel
	device, ok := deviceDeref.(*aether_3_0_0.Device)
	if !ok {
		return fmt.Errorf("unable to convert model in to aether_3_0_0")
	}
	return device.Validate()
}

func (m Modelplugin) Schema() (map[string]*yang.Entry, error) {
	return aether_3_0_0.UnzipSchema()
}

// GetStateMode returns an int - we do not use the enum because we do not want a
// direct dependency on onos-config code (for build optimization)
func (m Modelplugin) GetStateMode() int {
	return 0 // modelregistry.GetStateNone
}
