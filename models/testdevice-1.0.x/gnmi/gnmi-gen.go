/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// TODO this file needs to be generated via a template as it requires dynamic import (same a openapi-gen)

package main

import (
	"flag"
	"fmt"
	"github.com/onosproject/config-models/models/testdevice-1.0.x/api"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/goyang/pkg/yang"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

var log = logging.GetLogger("gnmi-gen")

func main() {
	var outputFile string
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug logs")
	flag.StringVar(&outputFile, "o", "", "Where to output generated code, stdout is default")
	flag.Parse()

	if debug {
		log.SetLevel(logging.DebugLevel)
	}

	schemaMap, err := api.Schema()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	topEntry := schemaMap.SchemaTree["Device"]
	res, err := buildGnmiStruct(topEntry, "")
	if err != nil {
		log.Errorw("failed to generate gNMI client", "err", err)
	}
	fmt.Println(res)
}

// TODO we might need/want to move the following in a separate package,
// as per pkg/openapi-gen/openapi-gen.go

// find a better name,
// keeps a flat list of gNMI methods we need to create

type GnmiEndpoints struct {
	Endpoints []GnmiEndpoint
}

type GnmiEndpoint struct {
	methodName string
	path       string
}

func buildGnmiStruct(entry *yang.Entry, parentPath string) (*GnmiEndpoints, error) {

	g := &GnmiEndpoints{
		Endpoints: []GnmiEndpoint{},
	}

	for _, item := range entry.Dir {
		itemPath := fmt.Sprintf("%s/%s", parentPath, item.Name)
		log.Infow("itme-path", "itemPath", itemPath)

		if item.IsLeaf() || item.IsLeafList() {
			log.Debug("item-is-leaf")
			eps, err := generateGnmiEndpointsForItem(item, itemPath)
			if err != nil {
				return nil, err
			}
			g.Endpoints = append(g.Endpoints, eps...)
		} else if item.Kind == yang.ChoiceEntry {
			// FIXME add support
			log.Debug("item-is-choice-entry")
		} else if item.IsContainer() {
			log.Debug("item-is-container")
			_g, err := buildGnmiStruct(item, itemPath)
			if err != nil {
				return nil, err
			}
			g.Endpoints = append(g.Endpoints, _g.Endpoints...)
		} else if item.IsList() {
			// FIXME add support
			log.Debug("item-is-list")
		} else {
			return nil, fmt.Errorf("unhandled item %v.Type %v", item.Name, item.Type)
		}
	}
	return g, nil
}

const gnmiGet = "get"
const gnmiUpdate = "update"
const gnmiList = "list"
const gnmiDelete = "delete"

func generateGnmiEndpointsForItem(item *yang.Entry, path string) ([]GnmiEndpoint, error) {
	eps := []GnmiEndpoint{}

	// NOTE do we need a "create" or an update without an ID ==  a create?
	methods := []string{gnmiGet, gnmiList, gnmiDelete, gnmiUpdate}

	caser := cases.Title(language.English)
	for _, m := range methods {
		epName := caser.String(m)
		ep := GnmiEndpoint{
			methodName: fmt.Sprintf("%s%s", epName, caser.String(item.Name)),
			path:       path,
		}
		eps = append(eps, ep)
		//switch m {
		//case gnmiGet:
		//}
	}
	return eps, nil
}
