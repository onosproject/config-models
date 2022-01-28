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

package compiler

import (
	"github.com/spf13/viper"
)

// MetaData plugin meta-data
type MetaData struct {
	Name         string   `mapstructure:"name" yaml:"name"`
	Version      string   `mapstructure:"version" yaml:"version"`
	Modules      []Module `mapstructure:"modules" yaml:"modules"`
	GetStateMode uint32   `mapstructure:"getStateMode" yaml:"getStateMode"`
	LintModel    bool     `mapstructure:"lintModel" yaml:"lintModel"`
	GenOpenAPI   bool     `mapstructure:"genOpenAPI" yaml:"genOpenAPI"`
	GoPackage    string   `mapstructure:"goPackage" yaml:"goPackage"`
}

type Module struct {
	Name         string `mapstructure:"name" yaml:"name"`
	Revision     string `mapstructure:"revision" yaml:"revision"`
	Organization string `mapstructure:"organization" yaml:"organization"`
	YangFile     string `mapstructure:"file" yaml:"file"`
}

// LoadMetaData loads the metadata.yaml file
func LoadMetaData(path string, metaData *MetaData) error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("metadata")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(metaData)
}
