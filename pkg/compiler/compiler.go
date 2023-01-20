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
	yang               = "yang"
	dotYang            = ".yang"
	pyang              = "pyang"
)

const yangBaseDirectory = "/var/model-compiler/yang-base"

// NewCompiler creates a new config model compiler
func NewCompiler() *ModelCompiler {
	return &ModelCompiler{}
}

type Dictionary struct {
	Name               string
	Version            string
	ArtifactName       string
	GoPackage          string
	ModelData          []*gnmi.ModelData
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
	metaData   *MetaData
	modelInfo  *api.ModelInfo
	dictionary Dictionary
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

	// Format YANG files if the model requests formatting
	if c.metaData.FormatYang {
		err = c.formatYang(path)
		if err != nil {
			log.Errorf("YANG file formatting failed: %+v", err)
			return err
		}
	}

	// Create dictionary from metadata and model info
	c.dictionary = Dictionary{
		Name:               c.modelInfo.Name,
		Version:            c.modelInfo.Version,
		ArtifactName:       c.metaData.ArtifactName,
		GoPackage:          c.metaData.GoPackage,
		ModelData:          c.modelInfo.ModelData,
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

	return nil
}

func (c *ModelCompiler) loadModelMetaData(path string) error {
	c.metaData = &MetaData{LintModel: true, RequireHyphenated: true, FormatYang: true}
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

func (c *ModelCompiler) lintModel(path string) error {
	log.Infof("Linting YANG files")

	// Append the root YANG files to the command-line arguments
	yangDir := filepath.Join(path, "yang")
	yangDirs := []string{yangBaseDirectory, yangDir}
	args := []string{"--lint", "-W", "error", "--ignore-error=XPATH_FUNCTION", "-p", strings.Join(yangDirs, ":")}
	if c.metaData.RequireHyphenated {
		args = append(args, "--lint-ensure-hyphenated-names")
	}
	for _, module := range c.metaData.Modules {
		args = append(args, filepath.Join(yangDir, module.YangFile))
	}

	log.Infof("Executing pyang %v", args)
	cmd := exec.Command(pyang, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *ModelCompiler) formatYang(path string) error {
	log.Infof("Formatting YANG files")

	// Append the root YANG files to the command-line arguments
	yangDir := filepath.Join(path, yang)
	yangDirs := []string{yangBaseDirectory, yangDir}

	tempFile := fmt.Sprintf("%s/temp-formatted.yang", os.TempDir())
	err := filepath.Walk(yangDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || !strings.HasSuffix(info.Name(), dotYang) {
				return nil
			}

			args := []string{"-f", yang, "--ignore-error=XPATH_FUNCTION", "-p", strings.Join(yangDirs, ":"), path,
				"-o", tempFile}
			log.Infof("Formatting YANG with: pyang %v", args)
			cmd := exec.Command(pyang, args...)
			cmd.Env = os.Environ()
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
			input, err := os.ReadFile(tempFile)
			if err != nil {
				return err
			}
			err = os.WriteFile(path, input, 0644)
			if err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}

	return nil
}

func (c *ModelCompiler) generateGolangBindings(path string) error {
	apiDir := filepath.Join(path, "api")
	c.createDir(apiDir)

	apiFile := filepath.Join(apiDir, "generated.go")
	log.Infof("Generating YANG bindings '%s'", apiFile)

	args := []string{
		fmt.Sprintf("-output_file=%s", apiFile),
		"-package_name=api",
		"-generate_fakeroot",
		"--include_descriptions",
	}

	// Append all YANG files to the command-line arguments
	pathDirs := []string{yangBaseDirectory}
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
	args = append(args, fmt.Sprintf("-path=%s", strings.Join(pathDirs, ",")))
	for _, module := range c.metaData.Modules {
		args = append(args, filepath.Join(path, "yang", module.YangFile))
	}

	log.Infof("Executing: generator %v", args)
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
	yangDirs := []string{yangBaseDirectory, yangDir}
	args := []string{"-f", "tree", "--ignore-error=XPATH_FUNCTION", "-p", strings.Join(yangDirs, ":"), "-o", treeFile}

	// Append the root YANG files to the command-line arguments
	for _, module := range c.metaData.Modules {
		args = append(args, filepath.Join(yangDir, module.YangFile))
	}

	log.Infof("Executing pyang %v", args)
	cmd := exec.Command(pyang, args...)
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
