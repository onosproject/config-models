// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package compiler

import (
	"fmt"
	"github.com/spf13/viper"
)

// MetaData plugin meta-data
type MetaData struct {
	Name               string   `mapstructure:"name" yaml:"name"`
	Version            string   `mapstructure:"version" yaml:"version"`
	Modules            []Module `mapstructure:"modules" yaml:"modules"`
	GetStateMode       uint32   `mapstructure:"getStateMode" yaml:"getStateMode"`
	LintModel          bool     `mapstructure:"lintModel" yaml:"lintModel"`
	GenOpenAPI         bool     `mapstructure:"genOpenAPI" yaml:"genOpenAPI"`
	OpenAPITargetAlias string   `mapstructure:"openAPITargetAlias" yaml:"openAPITargetAlias"`
	GoPackage          string   `mapstructure:"goPackage" yaml:"goPackage"`
	ArtifactName       string   `mapstructure:"artifactName" yaml:"artifactName"`
	ContactName        string   `mapstructure:"contactName" yaml:"contactName"`
	ContactUrl         string   `mapstructure:"contactUrl" yaml:"contactUrl"`
	ContactEmail       string   `mapstructure:"contactEmail" yaml:"contactEmail"`
	LicenseName        string   `mapstructure:"licenseName" yaml:"licenseName"`
	LicenseUrl         string   `mapstructure:"licenseUrl" yaml:"licenseUrl"`
}

type Module struct {
	Name         string `mapstructure:"name" yaml:"name"`
	Revision     string `mapstructure:"revision" yaml:"revision"`
	Organization string `mapstructure:"organization" yaml:"organization"`
	YangFile     string `mapstructure:"file" yaml:"file"`
}

// LoadMetaData loads the metadata.yaml file
func LoadMetaData(path string, configFile string, metaData *MetaData) error {
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFile)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(metaData)
}

// ValidateMetaData checks that required attributes are set
func ValidateMetaData(metaData *MetaData) error {
	if metaData.Name == "" {
		return fmt.Errorf("name is mandatory")
	}
	if metaData.Version == "" {
		return fmt.Errorf("version is mandatory")
	}
	if metaData.ArtifactName == "" {
		return fmt.Errorf("artifactName is mandatory")
	}
	if metaData.GoPackage == "" {
		return fmt.Errorf("goPackage is mandatory")
	}
	if metaData.Modules == nil || len(metaData.Modules) == 0 {
		return fmt.Errorf("no modules are listed")
	}
	if metaData.ContactName == "" {
		return fmt.Errorf("ContactName is mandatory")
	}
	if metaData.LicenseName == "" {
		return fmt.Errorf("licenseName is mandatory")
	}
	return nil
}
