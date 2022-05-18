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
	LeavesEndpoints    []LeavesEndpoint
	ContainerEndpoints []ContainerEndpoint
	ListEndpoints      []ListEndpoint
	BaseModel          string // the name of the yang model
}

// LeavesEndpoint contains the information needed to generate the gNMI client for
// elements which are at the end of the tree (leaves) and thus are simple types.
type LeavesEndpoint struct {
	Method            string // GET or SET (NOTE: do we need an Enum?)
	MethodName        string // Get, Update, List and Delete
	ModelName         string
	Path              []string // A string representing the gNMI path (/ separated)
	GoType            string
	GoReturnType      string
	GoEmptyReturnType string
}

type ContainerEndpoint struct {
	ModuleName      string
	ModelName       string
	ModelPath       string
	ParentModelPath string
	Method          string   // GET or SET (NOTE: do we need an Enum?)
	MethodName      string   // Get, Update, List and Delete
	Path            []string // A list of strings representing the gNMI path
}

type ListKey struct {
	Name string
	Type string
}

type ListEndpoint struct {
	ContainerEndpoint
	Key        ListKey
	ParentPath []string // A list of strings representing the gNMI path to the parent model
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

func BuildGnmiStruct(debug bool, modelName string, entry *yang.Entry, parentPath []string) (*GnmiEndpoints, error) {

	if debug {
		log.SetLevel(logging.DebugLevel)
	}

	g := &GnmiEndpoints{
		LeavesEndpoints:    []LeavesEndpoint{},
		ContainerEndpoints: []ContainerEndpoint{},
		BaseModel:          modelName,
	}

	for _, item := range entry.Dir {
		itemPath := append(parentPath, item.Name)
		log.Debugw("item", "path", itemPath)

		if item.IsLeaf() || item.IsLeafList() {
			eps, err := generateGnmiEndpointsForLeaf(item, itemPath)
			if err != nil {
				return nil, err
			}
			g.LeavesEndpoints = append(g.LeavesEndpoints, eps...)
		} else if item.Kind == yang.ChoiceEntry {
			// FIXME add support
			log.Warn("item-is-choice-entry")
		} else if item.IsContainer() {
			_g, err := BuildGnmiStruct(debug, modelName, item, itemPath)
			if err != nil {
				return nil, err
			}
			g.LeavesEndpoints = append(g.LeavesEndpoints, _g.LeavesEndpoints...)
			g.ContainerEndpoints = append(g.ContainerEndpoints, _g.ContainerEndpoints...)
			g.ListEndpoints = append(g.ListEndpoints, _g.ListEndpoints...)

			// generate endpoints for container
			eps, err := generateGnmiEndpointsForContainer(item, itemPath, modelName)
			if err != nil {
				return nil, err
			}
			g.ContainerEndpoints = append(g.ContainerEndpoints, eps...)
		} else if item.IsList() {
			eps, err := generateGnmiEndpointsForLists(item, itemPath, modelName)
			if err != nil {
				return nil, err
			}
			g.ListEndpoints = append(g.ListEndpoints, eps...)
		} else {
			return nil, fmt.Errorf("unhandled item %v.Type %v", item.Name, item.Type)
		}
	}
	return g, nil
}

const gnmiGet = "get"
const gnmiUpdate = "update"
const gnmiDelete = "delete"

// NOTE do we need a "create" or an update without an ID ==  a create?
var methods = []string{gnmiGet, gnmiDelete, gnmiUpdate}

func generateGnmiEndpointsForLeaf(item *yang.Entry, path []string) ([]LeavesEndpoint, error) {
	eps := []LeavesEndpoint{}
	itemName := PathToCamelCaseName(path)

	caser := cases.Title(language.English)
	for _, m := range methods {
		epName := caser.String(m)

		ep := LeavesEndpoint{
			Method:            m,
			ModelName:         itemName,
			MethodName:        fmt.Sprintf("%s%s", epName, itemName),
			Path:              path,
			GoType:            yangTypeToGoType(item.Type.Kind),
			GoReturnType:      yangTypeToGoReturnVal(item.Type.Kind),
			GoEmptyReturnType: yangTypeToGoEmptyReturnVal(item.Type.Kind),
		}
		eps = append(eps, ep)
	}
	log.Debugw("generating-methods-for-leaf-item",
		"name", itemName, "path", path, "endpoints", eps)
	return eps, nil
}

func generateGnmiEndpointsForContainer(item *yang.Entry, path []string, moduleName string) ([]ContainerEndpoint, error) {
	// NOTE YGOT model names are generated here: https://github.com/openconfig/ygot/blob/6f722a0cce2a47949294afa0c3f23b080d51e501/ygen/goelements.go#L193
	eps := []ContainerEndpoint{}
	// itemName is the name of the Model as generated by YGOT
	itemName := PathToYgotModelName(path)
	// pathInModel is where this particular container is nested in the YGOT device
	pathInModel := PathToYgotModelPath(path)
	// parentModelPath is the path to the parent container
	parentModelPath := PathToYgotModelPath(path[:len(path)-1])

	caser := cases.Title(language.English)
	for _, m := range methods {

		epName := caser.String(m)

		ep := ContainerEndpoint{
			Method:          m,
			MethodName:      fmt.Sprintf("%s%s", epName, itemName),
			Path:            path,
			ModelName:       itemName,
			ModuleName:      moduleName,
			ModelPath:       pathInModel,
			ParentModelPath: parentModelPath,
		}
		eps = append(eps, ep)
	}
	log.Debugw("generating-methods-for-container-item",
		"name", itemName, "path", path, "endpoints", eps)
	return eps, nil
}

// I assume that a list always contains Yang containers, we can't have a list of strings
func generateGnmiEndpointsForLists(item *yang.Entry, path []string, moduleName string) ([]ListEndpoint, error) {

	eps := []ListEndpoint{}

	// itemName is the name of the Model as generated by YGOT
	itemName := PathToYgotModelName(path)
	// pathInModel is where this particular container is nested in the YGOT device
	pathInModel := PathToYgotModelPath(path)
	// parentModelPath is the path to the parent container
	parentModelPath := PathToYgotModelPath(path[:len(path)-1])

	caser := cases.Title(language.English)
	for _, m := range methods {

		epName := caser.String(m)

		key := GetListKey(item)
		ep := ListEndpoint{
			ContainerEndpoint: ContainerEndpoint{
				Method:          m,
				MethodName:      fmt.Sprintf("%s%s", epName, itemName),
				Path:            path,
				ModelName:       itemName,
				ModuleName:      moduleName,
				ModelPath:       pathInModel,
				ParentModelPath: parentModelPath,
			},
			ParentPath: path[:len(path)-1],
			Key:        key,
		}

		if key.Type != "unknown" {
			// NOTE temporary workaround to avoid generating code
			// for lists with composite keys
			eps = append(eps, ep)
		}
	}
	log.Debugw("generating-methods-for-list-items",
		"name", itemName, "path", path, "endpoints", eps)
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
