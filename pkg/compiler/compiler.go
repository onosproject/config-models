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
	"github.com/openconfig/gnmi/proto/gnmi"
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

const (
	mainTemplate       = "main.go.tpl"
	pathsUtilsTemplate = "paths.go.tpl"
	gomodTemplate      = "go.mod.tpl"
	makefileTemplate   = "Makefile.tpl"
	dockerfileTemplate = "Dockerfile.tpl"
)

// NewCompiler creates a new config model compiler
func NewCompiler() *ModelCompiler {
	return &ModelCompiler{}
}

// ModelCompiler is a model plugin compiler
type ModelCompiler struct {
	metaData  *MetaData
	modelInfo *api.ModelInfo
}

// Compile compiles the config model
func (c *ModelCompiler) Compile(path string) error {
	log.Infof("Compiling config model at '%s'", path)
	var err error

	// Make sure inputs are present: meta-data file and YANG files directory
	// Read model meta-data
	err = c.loadModelMetaData(path)
	if err != nil {
		log.Errorf("Unable to read model meta-data: %+v", err)
		return err
	}

	// Lint YANG files if the model requests lint validation
	if c.metaData.LintModel {
		err = c.lintModel(path)
		if err != nil {
			log.Errorf("YANG files contain issues: %+v", err)
			return err
		}
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
	err = c.generatePluginArtifacts(path)
	if err != nil {
		log.Errorf("Unable to generate model plugin artifacts: %+v", err)
		return err
	}

	// TODO: Generate OpenAPI for RBAC
	return nil
}

func (c *ModelCompiler) loadModelMetaData(path string) error {
	c.metaData = &MetaData{}
	if err := LoadMetaData(path, c.metaData); err != nil {
		return err
	}
	modelData := make([]*gnmi.ModelData, 0, len(c.metaData.Modules))
	for _, module := range c.metaData.Modules {
		modelData = append(modelData, &gnmi.ModelData{
			Name:         module.Name,
			Version:      module.Revision,
			Organization: module.Organization,
		})
	}
	c.modelInfo = &api.ModelInfo{
		Name:         c.metaData.Name,
		Version:      c.metaData.Version,
		ModelData:    modelData,
		GetStateMode: c.metaData.GetStateMode,
	}
	return nil
}

func (c *ModelCompiler) lintModel(path string) error {
	log.Infof("Linting YANG files")

	args := []string{"--lint", "--lint-ensure-hyphenated-names", "-W", "error"}

	// Append the root YANG files to the command-line arguments
	yangDir := filepath.Join(path, "yang")
	for _, module := range c.metaData.Modules {
		args = append(args, filepath.Join(yangDir, module.YangFile))
	}

	log.Infof("Executing %s", path, strings.Join(args, " "))
	cmd := exec.Command("pyang", args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *ModelCompiler) generateGolangBindings(path string) error {
	apiDir := filepath.Join(path, "api")
	c.createDir(apiDir)

	apiFile := filepath.Join(apiDir, "generated.go")
	log.Infof("Generating YANG bindings '%s'", apiFile)

	args := []string{
		fmt.Sprintf("-path=%s/yang", path),
		fmt.Sprintf("-output_file=%s", apiFile),
		"-package_name=api",
		"-generate_fakeroot",
		"--include_descriptions",
	}

	// Append all YANG files to the command-line arguments
	files, err := ioutil.ReadDir(filepath.Join(path, "yang"))
	if err != nil {
		return err
	}
	for _, file := range files {
		args = append(args, file.Name())
	}

	log.Infof("Executing %s", path, strings.Join(args, " "))
	cmd := exec.Command("generator", args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return insertHeaderPrefix(apiFile)
}

func insertHeaderPrefix(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	newContent := []byte("// Code generated by YGOT. DO NOT")
	newContent = append(newContent, []byte("EDIT.\n")...) // HACK: Defeat the license header check
	newContent = append(newContent, content...)
	return ioutil.WriteFile(file, newContent, 0640)
}

func (c *ModelCompiler) generateModelTree(path string) error {
	treeFile := filepath.Join(path, c.modelInfo.Name+"-"+c.modelInfo.Version+".tree")
	log.Infof("Generating YANG tree '%s'", treeFile)

	yangDir := filepath.Join(path, "yang")
	args := []string{"-f", "tree", "-p", yangDir, "-o", treeFile}

	// Append the root YANG files to the command-line arguments
	for _, module := range c.metaData.Modules {
		args = append(args, filepath.Join(yangDir, module.YangFile))
	}

	log.Infof("Executing %s", path, strings.Join(args, " "))
	cmd := exec.Command("pyang", args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *ModelCompiler) generatePluginArtifacts(path string) error {
	// Generate main and paths extraction
	if err := c.generateMain(path); err != nil {
		return err
	}
	if err := c.generatePathsExtraction(path); err != nil {
		return err
	}

	// Generate go.mod from template
	if err := c.generateGoModule(path); err != nil {
		return err
	}

	// Generate Makefile from template
	if err := c.generateMakefile(path); err != nil {
		return err
	}

	// Generate Dockerfile from template
	if err := c.generateDockerfile(path); err != nil {
		return err
	}
	return nil
}

func (c *ModelCompiler) generateMain(path string) error {
	mainDir := filepath.Join(path, "plugin")
	mainFile := filepath.Join(mainDir, "main.go")
	log.Infof("Generating plugin main '%s'", mainFile)
	c.createDir(mainDir)
	return applyTemplate(mainTemplate, c.getTemplatePath(mainTemplate), mainFile, c.modelInfo)
}

func (c *ModelCompiler) generatePathsExtraction(path string) error {
	mainDir := filepath.Join(path, "plugin")
	pathsFile := filepath.Join(mainDir, "paths.go")
	log.Infof("Generating plugin paths extraction utility '%s'", pathsFile)
	c.createDir(mainDir)
	return applyTemplate(pathsUtilsTemplate, c.getTemplatePath(pathsUtilsTemplate), pathsFile, c.modelInfo)
}

func (c *ModelCompiler) generateGoModule(path string) error {
	gomodFile := filepath.Join(path, "go.mod")
	log.Infof("Generating plugin Go module '%s'", gomodFile)
	return applyTemplate(gomodTemplate, c.getTemplatePath(gomodTemplate), gomodFile, c.modelInfo)
}

func (c *ModelCompiler) generateMakefile(path string) error {
	makefileFile := filepath.Join(path, "Makefile")
	log.Infof("Generating plugin Makefile '%s'", makefileFile)
	return applyTemplate(makefileTemplate, c.getTemplatePath(makefileTemplate), makefileFile, c.modelInfo)
}

func (c *ModelCompiler) generateDockerfile(path string) error {
	dockerfileFile := filepath.Join(path, "Dockerfile")
	log.Infof("Generating plugin Dockerfile '%s'", dockerfileFile)
	return applyTemplate(dockerfileTemplate, c.getTemplatePath(dockerfileTemplate), dockerfileFile, c.modelInfo)
}
