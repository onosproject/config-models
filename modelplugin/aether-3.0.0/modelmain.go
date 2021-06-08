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
package main

import (
	"github.com/onosproject/config-models/modelplugin/aether-3.0.0/modelplugin"
)

//go:generate ./generator.sh

// ModelPlugin is the exported symbol that gives an entry point to this shared module
var ModelPlugin modelplugin.Modelplugin
