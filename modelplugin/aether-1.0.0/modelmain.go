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

// A plugin for the YGOT model of aether-1.0.0.
package main

import (
	"fmt"
	_ "github.com/golang/protobuf/proto"
	"github.com/onosproject/config-models/modelplugin/aether-1.0.0/aether_1_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	_ "github.com/openconfig/ygot/genutil"
	_ "github.com/openconfig/ygot/ygen"
	"github.com/openconfig/ygot/ygot"
	_ "github.com/openconfig/ygot/ytypes"
)

//go:generate go run github.com/openconfig/ygot/generator -path=yang -output_file=aether_1_0_0/generated.go -package_name=aether_1_0_0 -generate_fakeroot aether-subscriber@2020-07-15.yang aether-profile@2020-07-15.yang aether-subscriber-m2m@2020-07-15.yang

type modelplugin string

const modeltype = "Aether"
const modelversion = "1.0.0"
const modulename = "aether.so.1.0.0"

var modelData = []*gnmi.ModelData{
	{Name: "aether-subscriber", Organization: "Open Networking Foundation", Version: "2020-07-15"},
	{Name: "aether-profile", Organization: "Open Networking Foundation", Version: "2020-07-15"},
	{Name: "aether-subscriber-m2m", Organization: "Open Networking Foundation", Version: "2020-07-15"},
}

func (m modelplugin) ModelData() (string, string, []*gnmi.ModelData, string) {
	return modeltype, modelversion, modelData, modulename
}

// UnmarshallConfigValues allows Device to implement the Unmarshaller interface
func (m modelplugin) UnmarshalConfigValues(jsonTree []byte) (*ygot.ValidatedGoStruct, error) {
	device := &aether_1_0_0.Device{}
	vgs := ygot.ValidatedGoStruct(device)

	if err := aether_1_0_0.Unmarshal([]byte(jsonTree), device); err != nil {
		return nil, err
	}

	return &vgs, nil
}

func (m modelplugin) Validate(ygotModel *ygot.ValidatedGoStruct, opts ...ygot.ValidationOption) error {
	deviceDeref := *ygotModel
	device, ok := deviceDeref.(*aether_1_0_0.Device)
	if !ok {
		return fmt.Errorf("unable to convert model in to aether_1_0_0")
	}
	return device.Validate()
}

func (m modelplugin) Schema() (map[string]*yang.Entry, error) {
	return aether_1_0_0.UnzipSchema()
}

// GetStateMode returns an int - we do not use the enum because we do not want a
// direct dependency on onos-config code (for build optimization)
func (m modelplugin) GetStateMode() int {
	return 0 // modelregistry.GetStateNone
}

// ModelPlugin is the exported symbol that gives an entry point to this shared module
var ModelPlugin modelplugin
