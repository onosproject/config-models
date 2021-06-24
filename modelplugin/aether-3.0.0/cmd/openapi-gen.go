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

package main

import (
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/onosproject/config-models/modelplugin/aether-3.0.0/aether_3_0_0"
	openapi_gen "github.com/onosproject/config-models/pkg/openapi-gen"
	"io/ioutil"
	"os"
)

func main() {
	var outputFile string
	flag.StringVar(&outputFile, "o", "", "Where to output generated code, stdout is default")
	flag.Parse()

	schemaMap, err := aether_3_0_0.Schema()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	settings := openapi_gen.ApiGenSettings{
		ModelType:    "Aether",
		ModelVersion: "3.0.0",
		Example:      "connectivity-service-v3",
		Title:        "Aether 3.0.0",
	}

	schema, err := openapi_gen.BuildOpenapi(schemaMap, settings)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	yaml, err := yaml.Marshal(schema)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	license := []byte(`# SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
`)

	licensedYaml := append(license, yaml...)

	if outputFile != "" {
		err = ioutil.WriteFile(outputFile, licensedYaml, 0644)
		if err != nil {
			fmt.Printf("error writing generated code to file: %s\n", err)
			os.Exit(-1)
		}
	} else {
		fmt.Println(string(licensedYaml))
	}
}
