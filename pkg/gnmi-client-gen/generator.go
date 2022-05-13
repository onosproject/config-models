/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"fmt"
	t "github.com/onosproject/config-models/pkg/gnmi-client-gen/template"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/goyang/pkg/yang"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const templateFolder = "template"
const templateFile = "gnmi_client.go.tpl"

var log = logging.GetLogger("gnmi-client-gen")

type GnmiEndpoints struct {
	Endpoints []GnmiEndpoint
	BaseModel string // the name of the yang model
}

type GnmiEndpoint struct {
	Method     string // GET or SET (NOTE: do we need an Enum?)
	MethodName string // Get, Update, List and Delete
	Path       string // A string representing the gNMI path (/ separated)
	ValueType  yang.TypeKind
}

func CapitalizeModelName(model []*gnmi.ModelData) string {
	caser := cases.Title(language.English)

	// onf-test1 -> OnfTest1
	modelName := model[0].Name
	pieces := strings.Split(modelName, "-")
	capitalized := ""
	for _, p := range pieces {
		capitalized = capitalized + caser.String(p)
	}
	return capitalized
}

func BuildGnmiStruct(modelName string, entry *yang.Entry, parentPath string, parentName string) (*GnmiEndpoints, error) {
	caser := cases.Title(language.English)
	g := &GnmiEndpoints{
		Endpoints: []GnmiEndpoint{},
		BaseModel: modelName,
	}

	for _, item := range entry.Dir {
		itemPath := fmt.Sprintf("%s/%s", parentPath, item.Name)
		itemName := fmt.Sprintf("%s%s", parentName, strings.ReplaceAll(caser.String(item.Name), "-", "_"))
		log.Debugw("item", "path", itemPath, "name", itemName)

		if item.IsLeaf() || item.IsLeafList() {
			eps, err := generateGnmiEndpointsForItem(item, itemPath, itemName)
			if err != nil {
				return nil, err
			}
			g.Endpoints = append(g.Endpoints, eps...)
		} else if item.Kind == yang.ChoiceEntry {
			// FIXME add support
			log.Debug("item-is-choice-entry")
		} else if item.IsContainer() {
			log.Debug("item-is-container")
			_g, err := BuildGnmiStruct(modelName, item, itemPath, itemName)
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

func generateGnmiEndpointsForItem(item *yang.Entry, path string, name string) ([]GnmiEndpoint, error) {
	log.Debugw("generating-methods-for-item", "name", name, "path", path)
	eps := []GnmiEndpoint{}

	// NOTE do we need a "create" or an update without an ID ==  a create?
	methods := []string{gnmiGet, gnmiList, gnmiDelete, gnmiUpdate}

	caser := cases.Title(language.English)
	for _, m := range methods {
		epName := caser.String(m)

		var gnmiMethod string
		if m == gnmiGet || m == gnmiList {
			gnmiMethod = "GET"
		} else {
			gnmiMethod = "SET"
		}

		ep := GnmiEndpoint{
			Method:     gnmiMethod,
			MethodName: fmt.Sprintf("%s%s", epName, name),
			Path:       path,
			ValueType:  item.Type.Kind,
		}
		eps = append(eps, ep)
		//switch m {
		//case gnmiGet:
		//}
	}
	return eps, nil
}

func ApplyTemplate(epList *GnmiEndpoints, outPath string) error {

	var funcs template.FuncMap = map[string]interface{}{
		"quote": func(value interface{}) string {
			return fmt.Sprintf("\"%s\"", value)
		},
		"replace": func(search, replace string, value interface{}) string {
			return strings.ReplaceAll(fmt.Sprint(value), search, replace)
		},
		"splitPath": func(val string) []string {
			// NOTE the first element of the list is empty as it starts with /
			// so just drop it
			return strings.Split(strings.TrimSpace(val), "/")[1:]
		},
		// NOTE inspired by https://github.com/openconfig/ygot/blob/master/ytypes/util_types.go#L353
		"yang2goType": func(val yang.TypeKind) string {
			switch val {
			case yang.Yint8:
				return "int8"
			case yang.Yint16:
				return "int16"
			case yang.Yint32:
				return "int32"
			case yang.Yint64:
				return "int64"
			case yang.Yuint8:
				return "uint8"
			case yang.Yuint16:
				return "uint16"
			case yang.Yuint32:
				return "uint32"
			case yang.Yuint64:
				return "uint64"
			case yang.Ybool, yang.Yempty:
				return "bool"
			case yang.Ystring:
				return "string"
			case yang.Ydecimal64:
				return "float64"
			case yang.Ybinary:
				return "[]byte"
			case yang.Yenum, yang.Yidentityref:
				return "int64"
			}
			// not ideal, but for now we'll take it
			return "interface{}"
		},
		"yang2returnVal": func(val yang.TypeKind) string {
			switch val {
			case yang.Yint8:
				return "int8(val.GetIntVal())"
			case yang.Yint16:
				return "int16(val.GetIntVal())"
			case yang.Yint32:
				return "int32(val.GetIntVal())"
			case yang.Yint64:
				return "int64(val.GetIntVal())"
			case yang.Yuint8:
				return "uint8(val.GetUintVal())"
			case yang.Yuint16:
				return "uint16(val.GetUintVal())"
			case yang.Yuint32:
				return "uint32(val.GetUintVal())"
			case yang.Yuint64:
				return "uint64(val.GetUintVal())"
			case yang.Ybool, yang.Yempty:
				return "val.GetBoolVal()"
			case yang.Ystring:
				return "val.GetStringVal()"
			case yang.Ydecimal64:
				return "float64(val.GetFloatVal())"
			case yang.Ybinary:
				return "val.GetBytesVal()"
			case yang.Yenum, yang.Yidentityref:
				return "int64(val.GetIntVal())"
			}
			// not ideal, but for now we'll take it
			return "GetValue()"
		},
		// given a Yang TypeKind return the appropriate value in case of error
		"emptyReturnValue": func(val yang.TypeKind) string {
			switch val {
			case yang.Yint8,
				yang.Yint16,
				yang.Yint32,
				yang.Yint64,
				yang.Yuint8,
				yang.Yuint16,
				yang.Yuint32,
				yang.Yuint64,
				yang.Ydecimal64,
				yang.Yenum, yang.Yidentityref:
				return "0"
			case yang.Ybool, yang.Yempty:
				return "false"
			case yang.Ystring:
				return "\"\""
			case yang.Ybinary:
				return "[]byte{}"
			}
			// not ideal, but for now we'll take it
			return "interface{}"
		},
	}

	t, err := template.New(templateFile).
		Funcs(funcs).
		ParseFS(t.GnmiGenTemplate, templateFile)
	if err != nil {
		return err
	}

	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, epList)
}

func getTemplatePath(name string) string {
	return filepath.Join(templateFolder, name)
}
