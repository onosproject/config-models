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

// A plugin for the YGOT model of plproxy-1.0.0.
package modelplugin

import (
	"fmt"

	_ "github.com/golang/protobuf/proto"
	"github.com/onosproject/config-models/modelplugin/plproxy-1.0.0/plproxy_1_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	_ "github.com/openconfig/ygot/genutil"
	_ "github.com/openconfig/ygot/ygen"
	"github.com/openconfig/ygot/ygot"
	_ "github.com/openconfig/ygot/ytypes"
)

type Modelplugin string

const modeltype = "PLProxy"
const modelversion = "1.0.0"
const modulename = "plproxy.so.1.0.0"

var ModelData = []*gnmi.ModelData{
	{Name: "plproxy-service", Organization: "Open Networking Foundation", Version: "2021-09-25"},
}

func (m Modelplugin) ModelData() (string, string, []*gnmi.ModelData, string) {
	return modeltype, modelversion, ModelData, modulename
}

// UnmarshallConfigValues allows Device to implement the Unmarshaller interface
func (m Modelplugin) UnmarshalConfigValues(jsonTree []byte) (*ygot.ValidatedGoStruct, error) {
	device := &plproxy_1_0_0.Device{}
	vgs := ygot.ValidatedGoStruct(device)

	if err := plproxy_1_0_0.Unmarshal([]byte(jsonTree), device); err != nil {
		return nil, err
	}

	return &vgs, nil
}

func (m Modelplugin) Validate(ygotModel *ygot.ValidatedGoStruct, opts ...ygot.ValidationOption) error {
	deviceDeref := *ygotModel
	device, ok := deviceDeref.(*plproxy_1_0_0.Device)
	if !ok {
		return fmt.Errorf("unable to convert model in to plproxy_1_0_0")
	}
	return device.Validate()
}

func (m Modelplugin) Schema() (map[string]*yang.Entry, error) {
	return plproxy_1_0_0.UnzipSchema()
}

// GetStateMode returns an int - we do not use the enum because we do not want a
// direct dependency on onos-config code (for build optimization)
func (m Modelplugin) GetStateMode() int {
	return 0 // modelregistry.GetStateNone
}
