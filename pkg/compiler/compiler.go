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
	"fmt"
	api "github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	_ "github.com/openconfig/gnmi/proto/gnmi" // gnmi
	_ "github.com/openconfig/goyang/pkg/yang" // yang
	_ "github.com/openconfig/ygot/genutil"    // genutil
	_ "github.com/openconfig/ygot/ygen"       // ygen
	_ "github.com/openconfig/ygot/ygot"       // ygot
	_ "github.com/openconfig/ygot/ytypes"     // ytypes
	_ "google.golang.org/protobuf/proto"      // proto
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var log = logging.GetLogger("config-model", "compiler")

// NewCompiler creates a new config model compiler
func NewCompiler() *ModelCompiler {
	return &ModelCompiler{}
}

// ModelCompiler is a model plugin compiler
type ModelCompiler struct {
	MetaData *api.ModelInfo
}

// Compile compiles the config model
func (c *ModelCompiler) Compile(path string) error {
	log.Infof("Compiling config model at '%s'", path)
	var err error

	// Make sure inputs are present: meta-data file and YANG files directory
	// Read model meta-data
	c.MetaData, err = c.readModelMetaData(path)
	if err != nil {
		log.Errorf("Unable to read model meta-data: %+v", err)
		return err
	}

	// Generate Golang bindings for the YANG files
	err = c.generateGolangBindings(path)
	if err != nil {
		log.Errorf("Unable to generate Golang bindings: %+v", err)
		return err
	}

	// Generate YANG model tree
	err = c.generateModelTree(path)
	if err != nil {
		log.Errorf("Unable to generate YANG model tree: %+v", err)
		return err
	}

	// Generate model plugin artifacts from generic templates
	// - main
	// - gRPC PluginService NB
	// - model plugin Dockerfile from template
	err = c.generatePluginArtifacts(path)
	if err != nil {
		log.Errorf("Unable to generate model plugin artifacts: %+v", err)
		return err
	}

	// TODO: Generate OpenAPI for RBAC
	return nil
}

func (c *ModelCompiler) readModelMetaData(path string) (*api.ModelInfo, error) {
	metaData := &MetaData{}
	if err := LoadMetaData(path, metaData); err != nil {
		return nil, err
	}
	return &api.ModelInfo{Name: metaData.Name, Version: metaData.Version}, nil
}

func (c *ModelCompiler) generateGolangBindings(path string) error {
	pkg := c.getModelPackage()
	file := filepath.Join(path, pkg, "generated.go")
	log.Infof("Generating YANG bindings '%s'", file)
	args := []string{
		"run",
		"github.com/openconfig/ygot/generator",
		fmt.Sprintf("-path=%s/yang", path),
		fmt.Sprintf("-output_file=%s", file),
		fmt.Sprintf("-package_name=%s", pkg),
		"-generate_fakeroot",
	}

	// Append all YANG files to the command-line arguments
	yangDir := filepath.Join(path, "yang")
	files, err := ioutil.ReadDir(yangDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		args = append(args, filepath.Join(yangDir, file.Name()))
	}

	log.Infof("Executing %s", path, strings.Join(args, " "))
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *ModelCompiler) generateModelTree(path string) error {
	return nil
}

func (c *ModelCompiler) generatePluginArtifacts(path string) error {
	return nil
}

func (c *ModelCompiler) getModelPackage() string {
	return sanitized(c.MetaData.Name) + "_" + sanitized(c.MetaData.Version)
}

func sanitized(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, ".", "_"), "-", "_")
}
