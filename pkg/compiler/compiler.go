// Copyright 2020-present Open Networking Foundation.
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

package compiler

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	_ "github.com/openconfig/gnmi/proto/gnmi" // gnmi
	_ "github.com/openconfig/goyang/pkg/yang" // yang
	_ "github.com/openconfig/ygot/genutil"    // genutil
	_ "github.com/openconfig/ygot/ygen"       // ygen
	_ "github.com/openconfig/ygot/ygot"       // ygot
	_ "github.com/openconfig/ygot/ytypes"     // ytypes
	_ "google.golang.org/protobuf/proto"      // proto
)

var log = logging.GetLogger("config-model", "compiler")

// NewCompiler creates a new config model compiler
func NewCompiler() *ModelCompiler {
	return &ModelCompiler{}
}

// ModelCompiler is a model plugin compiler
type ModelCompiler struct {
}

// Compile compiles the config model
func (c *ModelCompiler) Compile(path string) error {
	log.Infof("Compiling config model at '%s'", path)

	// Make sure inputs are present: meta-data file and YANG files directory
	// Read model meta-data

	// Generate Golang bindings for the YANG files
	// Generate YANG model tree

	// Generate model plugin Go code from templates
	// - main
	// - gRPC PluginService NB

	// ?Generate OpenAPI for RBAC

	// Generate model plugin docker file from template

	return nil
}
