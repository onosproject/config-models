package main

import (
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/onosproject/config-models/models/testdevice-2.0.x/api"
	openapi_gen "github.com/onosproject/config-models/pkg/openapi-gen"
	"io/ioutil"
	"os"
)

func main() {
	var outputFile string
	flag.StringVar(&outputFile, "o", "", "Where to output generated code, stdout is default")
	flag.Parse()

	schemaMap, err := api.Schema()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// description, err := ioutil.ReadFile("description.md") // From https://docs.aetherproject.org/master/_sources/developer/roc-api.rst.txt

	description := "TODO add a description field in Metadata.yaml?"
	settings := openapi_gen.ApiGenSettings{
		ModelType:    "testdevice-2.0.x",
		ModelVersion: "2.0.0",
		Title:        "testdevice-2.0.x-2.0.0",
		Description:  string(description),
	}

	schema, err := openapi_gen.BuildOpenapi(schemaMap, &settings)
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
