// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var log = logging.GetLogger("config-model", "compiler")

const (
	versionFile        = "VERSION"
	mainTemplate       = "main.go.tpl"
	modelTemplate      = "model.go.tpl"
	gomodTemplate      = "go.mod.tpl"
	makefileTemplate   = "Makefile.tpl"
	dockerfileTemplate = "Dockerfile.tpl"
	openapiGenTemplate = "openapi-gen.go.tpl"
	//gnmiGenTemplate    = "gnmi-gen.go.tpl"
)

// NewCompiler creates a new config model compiler
func NewCompiler() *ModelCompiler {
	return &ModelCompiler{}
}

type Dictionary struct {
	Name               string
	Version            string
	PluginVersion      string
	ArtifactName       string
	GoPackage          string
	ModelData          []*gnmi.ModelData
	Module             string
	GetStateMode       uint32
	ReadOnlyPath       []*api.ReadOnlyPath
	ReadWritePath      []*api.ReadWritePath
	OpenAPITargetAlias string
	ContactName        string
	ContactUrl         string
	ContactEmail       string
	LicenseName        string
	LicenseUrl         string
}

// ModelCompiler is a model plugin compiler
type ModelCompiler struct {
	pluginVersion string
	metaData      *MetaData
	modelInfo     *api.ModelInfo
	dictionary    Dictionary
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

	err = c.loadPluginVersion(path)
	if err != nil {
		log.Errorf("Unable to load model plugin version; defaulting to %s: %+v", c.pluginVersion, err)
	}

	// Lint YANG files if the model requests lint validation
	if c.metaData.LintModel {
		err = c.lintModel(path)
		if err != nil {
			log.Errorf("YANG files contain issues: %+v", err)
			return err
		}
	}

	// Create dictionary from metadata and model info
	c.dictionary = Dictionary{
		Name:               c.modelInfo.Name,
		Version:            c.modelInfo.Version,
		PluginVersion:      c.pluginVersion,
		ArtifactName:       c.metaData.ArtifactName,
		GoPackage:          c.metaData.GoPackage,
		ModelData:          c.modelInfo.ModelData,
		Module:             c.modelInfo.Module,
		GetStateMode:       c.modelInfo.GetStateMode,
		ReadOnlyPath:       c.modelInfo.ReadOnlyPath,
		ReadWritePath:      c.modelInfo.ReadWritePath,
		OpenAPITargetAlias: c.metaData.OpenAPITargetAlias,
		ContactName:        c.metaData.ContactName,
		ContactUrl:         c.metaData.ContactUrl,
		ContactEmail:       c.metaData.ContactEmail,
		LicenseName:        c.metaData.LicenseName,
		LicenseUrl:         c.metaData.LicenseUrl,
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

	// Generate OpenAPI
	err = c.generateOpenApi(path)
	if err != nil {
		log.Errorf("Unable to generate OpenApi specs: %+v", err)
		return err
	}

	// the gNMI client generator is on hold at the moment,
	// disabling it for the moment
	//Generate gNMI Client Generator
	//err = c.generateGnmiClientGenerator(path)
	//if err != nil {
	//	log.Errorf("Unable to generate gNMI Client Generator: %+v", err)
	//	return err
	//}

	// Now generate the gNMI client itself
	//err = c.generateGnmiClient(path)
	//if err != nil {
	//	log.Errorf("Unable to generate gNMI Client: %+v", err)
	//	return err
	//}

	return nil
}

func (c *ModelCompiler) loadModelMetaData(path string) error {
	c.metaData = &MetaData{}
	if err := LoadMetaData(path, "metadata", c.metaData); err != nil {
		return err
	}
	if err := ValidateMetaData(c.metaData); err != nil {
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

func (c *ModelCompiler) loadPluginVersion(path string) error {
	data, err := os.ReadFile(filepath.Join(path, versionFile))
	if err != nil {
		c.pluginVersion = "1.0.0"
	}
	v := string(data)
	c.pluginVersion = strings.Split(strings.ReplaceAll(v, "\r\n", "\n"), "\n")[0]
	return err
}

func (c *ModelCompiler) lintModel(path string) error {
	log.Infof("Linting YANG files")

	args := []string{"--lint", "--lint-ensure-hyphenated-names", "-W", "error", "--ignore-error=XPATH_FUNCTION"}

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
	pathDirs := make([]string, 0)
	err := filepath.Walk(filepath.Join(path, "yang"), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			pathDirs = append(pathDirs, p)
		}
		return nil
	})
	if err != nil {
		return err
	}
	args = append(args, fmt.Sprintf("-path=\"%s\"", strings.Join(pathDirs, ",")))
	for _, module := range c.metaData.Modules {
		args = append(args, filepath.Join(path, "yang", module.YangFile))
	}

	log.Infof("Executing: generator %s", path, strings.Join(args, " "))
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
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	newContent := []byte("// Code generated by YGOT. DO NOT")
	newContent = append(newContent, []byte("EDIT.\n")...) // HACK: Defeat the license header check
	newContent = append(newContent, content...)
	return os.WriteFile(file, newContent, 0640)
}

func (c *ModelCompiler) generateModelTree(path string) error {
	treeFile := filepath.Join(path, c.modelInfo.Name+".tree")
	log.Infof("Generating YANG tree '%s'", treeFile)

	yangDir := filepath.Join(path, "yang")
	args := []string{"-f", "tree", "--ignore-error=XPATH_FUNCTION", "-p", yangDir, "-o", treeFile}

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
	if err := c.generateModel(path); err != nil {
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
	return c.applyTemplate(mainTemplate, c.getTemplatePath(mainTemplate), mainFile)
}

func (c *ModelCompiler) generateModel(path string) error {
	modelDir := filepath.Join(path, "api")
	modelFile := filepath.Join(modelDir, "model.go")
	log.Infof("Generating plugin model '%s'", modelFile)
	c.createDir(modelDir)
	return c.applyTemplate(modelTemplate, c.getTemplatePath(modelTemplate), modelFile)
}

func (c *ModelCompiler) generateGoModule(path string) error {
	gomodFile := filepath.Join(path, "go.mod")
	log.Infof("Generating plugin Go module '%s'", gomodFile)
	return c.applyTemplate(gomodTemplate, c.getTemplatePath(gomodTemplate), gomodFile)
}

func (c *ModelCompiler) generateMakefile(path string) error {
	makefileFile := filepath.Join(path, "Makefile")
	log.Infof("Generating plugin Makefile '%s'", makefileFile)
	return c.applyTemplate(makefileTemplate, c.getTemplatePath(makefileTemplate), makefileFile)
}

func (c *ModelCompiler) generateDockerfile(path string) error {
	dockerfileFile := filepath.Join(path, "Dockerfile")
	log.Infof("Generating plugin Dockerfile '%s'", dockerfileFile)
	return c.applyTemplate(dockerfileTemplate, c.getTemplatePath(dockerfileTemplate), dockerfileFile)
}

// TODO we should be able to run this generated code right after we generate it,
// so that we can remove a step from `make models-images`
func (c *ModelCompiler) generateOpenApi(path string) error {
	// the Schema we need to import is generated at runtime, so we need to generate the tool
	// to import such schema and generate the OpenApi specs
	dir := filepath.Join(path, "openapi")
	openapiGenFile := filepath.Join(dir, "openapi-gen.go")
	c.createDir(dir)

	log.Infof("Generating plugin OpenApi Gen file '%s'", openapiGenFile)
	return c.applyTemplate(openapiGenTemplate, c.getTemplatePath(openapiGenTemplate), openapiGenFile)
}

//func (c *ModelCompiler) generateGnmiClientGenerator(path string) error {
//	// the Schema we need to import is generated at runtime, so we need to generate the tool
//	// to import such schema and generate the OpenApi specs
//	dir := filepath.Join(path, "gnmi-gen")
//	gnmiGen := filepath.Join(dir, "gnmi-gen.go")
//	c.createDir(dir)
//
//	log.Infof("Generating plugin GnmiGen file '%s'", gnmiGen)
//	return c.applyTemplate(gnmiGenTemplate, c.getTemplatePath(gnmiGenTemplate), gnmiGen)
//}

//func (c *ModelCompiler) generateGnmiClient(path string) error {
//	generatorPath := filepath.Join(path, "gnmi-gen/gnmi-gen.go")
//
//	args := []string{
//		"run",
//		generatorPath,
//		"--debug",
//	}
//
//	log.Infof("Executing: generator %s", path, strings.Join(args, " "))
//	cmd := exec.Command("go", args...)
//	cmd.Env = os.Environ()
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	err := cmd.Run()
//	if err != nil {
//		return err
//	}
//	return nil
//}
