// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"{{ .GoPackage }}/api"
	openapi_gen "github.com/onosproject/config-models/pkg/openapi-gen"
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

	settings := openapi_gen.ApiGenSettings{
		ModelType:    "{{ .Name }}",
		ModelVersion: "{{ .Version }}",
		Title:        "{{ .Name }}-{{ .Version }}",
		TargetAlias:  "{{ .OpenAPITargetAlias }}",
        Contact:      &openapi3.Contact{
            Name:           {{.ContactName | quote}},
            URL:            {{.ContactUrl | quote}},
            Email:          {{.ContactEmail | quote}},
        },
        License:      &openapi3.License{
            Name:           {{.LicenseName | quote}},
            URL:            {{.LicenseUrl | quote}},
        },
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

	license := []byte(`# SPDX-FileCopyrightText: 2022-present Intel Corporation
# SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0
`)

	licensedYaml := append(license, yaml...)

	if outputFile != "" {
		err = os.WriteFile(outputFile, licensedYaml, 0644)
		if err != nil {
			fmt.Printf("error writing generated code to file: %s\n", err)
			os.Exit(-1)
		}
	} else {
		fmt.Println(string(licensedYaml))
	}
}
