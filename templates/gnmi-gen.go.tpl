/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"flag"
	"fmt"
	"{{ .GoPackage }}/api"
	gnmi_client_gen "github.com/onosproject/config-models/pkg/gnmi-client-gen"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"os"
	"path/filepath"
)

const outputFolder = "api"
var outputFile string
var log = logging.GetLogger("{{ .Name }}-gnmi-gen")

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug logs")
	flag.StringVar(&outputFile, "o", "gnmi_client.go", "Where to output generated code")
	flag.Parse()

	if debug {
		log.SetLevel(logging.DebugLevel)
	}

	schemaMap, err := api.Schema()
	if err != nil {
	fmt.Println(err)
		os.Exit(-1)
	}

	modelName := gnmi_client_gen.CapitalizeModelName(api.ModelData())

	topEntry := schemaMap.SchemaTree["Device"]
	res, err := gnmi_client_gen.BuildGnmiStruct(debug, modelName, topEntry, []string{})
	if err != nil {
		log.Errorw("failed to generate gNMI Endpoint list", "err", err)
	}

	outPath := filepath.Join(outputFolder, outputFile)
	err = gnmi_client_gen.ApplyTemplate(res, outPath)
	if err != nil {
		log.Errorw("failed to generate Go code for gNMI client", "err", err)
		os.Exit(1)
	}
	log.Infow("code-generated", "output", outPath)
}
