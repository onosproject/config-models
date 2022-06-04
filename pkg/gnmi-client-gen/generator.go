/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"errors"
	"fmt"
	t "github.com/onosproject/config-models/pkg/gnmi-client-gen/template"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/genutil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"sort"
	"strings"
	"text/template"
)

const templateFile = "gnmi_client.go.tpl"

var log = logging.GetLogger("gnmi-client-gen")

type GnmiEndpoints struct {
	LeavesEndpoints    []LeavesEndpoint
	ContainerEndpoints []ContainerEndpoint
	ListEndpoints      []ListEndpoint
	PluginName         string
}

// LeavesEndpoint contains the information needed to generate the gNMI client for
// elements which are at the end of the tree (leaves) and thus are simple types.
type LeavesEndpoint struct {
	Method            string // GET or SET (NOTE: do we need an Enum?)
	MethodName        string // Get, Update and Delete
	ModelName         string
	Path              []string // A string representing the gNMI path (/ separated)
	GoType            string
	GoReturnType      string
	GoEmptyReturnType string
}

type ContainerEndpoint struct {
	ModelName       string
	ModelPath       string
	ParentModelPath string
	Method          string   // GET or SET (NOTE: do we need an Enum?)
	MethodName      string   // Get, Update and Delete + itemName
	Path            []string // A list of strings representing the gNMI path
}

type ListKey struct {
	Type      string   // the generated Key type (can be a Go type, eg: string)
	Keys      []Key    // a list of keys that we need to set in the path
	ParentKey *ListKey // a reference to a parent key, this is needed in case of nested lists
	ModelName string   // represent the model in the path, to simplify recursion
}
type Key struct {
	Name string
	Type string
	Ptr  bool // whether this is pointer or not
}

type ListEndpoint struct {
	ContainerEndpoint
	Key              ListKey
	ParentPath       []string // A list of strings representing the gNMI path to the parent model
	PluralMethodName string   // Used for the methods that apply to the entire list
}

func BuildGnmiStruct(debug bool, pluginName string, entry *yang.Entry, parentPath []string, parentKey *ListKey) (*GnmiEndpoints, error) {

	if debug {
		log.SetLevel(logging.DebugLevel)
	}

	g := &GnmiEndpoints{
		LeavesEndpoints:    []LeavesEndpoint{},
		ContainerEndpoints: []ContainerEndpoint{},
		ListEndpoints:      []ListEndpoint{},
		PluginName:         pluginName,
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
			_g, err := BuildGnmiStruct(debug, pluginName, item, itemPath, parentKey)
			if err != nil {
				return nil, err
			}
			g.LeavesEndpoints = append(g.LeavesEndpoints, _g.LeavesEndpoints...)
			g.ContainerEndpoints = append(g.ContainerEndpoints, _g.ContainerEndpoints...)
			g.ListEndpoints = append(g.ListEndpoints, _g.ListEndpoints...)

			// generate endpoints for container
			eps, err := generateGnmiEndpointsForContainer(item, itemPath)
			if err != nil {
				return nil, err
			}
			g.ContainerEndpoints = append(g.ContainerEndpoints, eps...)
		} else if item.IsList() {
			// generate the key for this list
			key, err := GetListKey(item)
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Unimplemented {
					log.Warnw("skipping-generation-for-list", "item", item.Name, "err", err.Error())
					continue
				} else {
					return nil, err
				}
			}

			// if there is parent key add it to key for this item
			if parentKey != nil {
				key.ParentKey = parentKey
			}

			// recurse down the spec
			_g, err := BuildGnmiStruct(debug, pluginName, item, itemPath, &key)
			if err != nil {
				return nil, err
			}
			g.LeavesEndpoints = append(g.LeavesEndpoints, _g.LeavesEndpoints...)
			g.ContainerEndpoints = append(g.ContainerEndpoints, _g.ContainerEndpoints...)
			g.ListEndpoints = append(g.ListEndpoints, _g.ListEndpoints...)

			// generate endpoints for lists
			eps, err := generateGnmiEndpointsForLists(item, itemPath, &key)
			if err != nil {
				return nil, err
			}
			g.ListEndpoints = append(g.ListEndpoints, eps...)
		} else {
			return nil, fmt.Errorf("unhandled item %v.Type %v", item.Name, item.Type)
		}
	}
	ep, err := sortByName(*g)
	return &ep, err
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

		t, err := yangTypeToGoType(item)
		if err != nil {
			if s, ok := status.FromError(err); ok && s.Code() == codes.Unimplemented {
				log.Warnw("skipping-generation-for-leaf", "item", item.Name, "err", err.Error())
				continue
			} else {
				return nil, err
			}
		}
		ep := LeavesEndpoint{
			Method:            m,
			ModelName:         itemName,
			MethodName:        fmt.Sprintf("%s_%s", epName, itemName),
			Path:              path,
			GoType:            t,
			GoReturnType:      yangTypeToGoReturnVal(item.Type.Kind),
			GoEmptyReturnType: yangTypeToGoEmptyReturnVal(item.Type.Kind),
		}
		eps = append(eps, ep)
	}
	log.Debugw("generating-methods-for-leaf-item",
		"name", itemName, "path", path, "endpoints", eps)
	return eps, nil
}

func generateGnmiEndpointsForContainer(item *yang.Entry, path []string) ([]ContainerEndpoint, error) {
	// NOTE YGOT model names are generated here: https://github.com/openconfig/ygot/blob/6f722a0cce2a47949294afa0c3f23b080d51e501/ygen/goelements.go#L193
	eps := []ContainerEndpoint{}
	// itemName is the name of the Model as generated by YGOT (without module prefix)
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
			MethodName:      fmt.Sprintf("%s_%s", epName, itemName),
			Path:            path,
			ModelName:       fmt.Sprintf("%s", item.Annotation["structname"]),
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
func generateGnmiEndpointsForLists(item *yang.Entry, path []string, key *ListKey) ([]ListEndpoint, error) {

	eps := []ListEndpoint{}

	// itemName is the name of the Model as generated by YGOT (without module prefix)
	itemName := PathToYgotModelName(path)

	// pathInModel is where this particular container is nested in the YGOT device
	pathInModel := PathToYgotModelPath(path)
	// parentModelPath is the path to the parent container
	parentModelPath := PathToYgotModelPath(path[:len(path)-1])

	caser := cases.Title(language.English)
	for _, m := range methods {

		epName := caser.String(m)

		ep := ListEndpoint{
			ContainerEndpoint: ContainerEndpoint{
				Method:          m,
				MethodName:      fmt.Sprintf("%s_%s", epName, itemName),
				Path:            path,
				ModelName:       fmt.Sprintf("%s", item.Annotation["structname"]),
				ModelPath:       pathInModel,
				ParentModelPath: parentModelPath,
			},
			PluralMethodName: fmt.Sprintf("%s_%s_List", epName, itemName),
			ParentPath:       path[:len(path)-1],
			Key:              *key,
		}
		eps = append(eps, ep)
	}
	log.Debugw("generating-methods-for-list-items",
		"name", itemName, "path", path, "endpoints", eps)
	return eps, nil
}

// sortByName will order the lists based on the MethodName
// a possible future improvement is to order them by ModelName while retaining the order Get, Update, Delete
// but the main reason to have this is to make it easier to compare diffs in PRs for an easier review process
func sortByName(endpoints GnmiEndpoints) (GnmiEndpoints, error) {
	sort.Sort(leavesByMethodName(endpoints.LeavesEndpoints))
	sort.Sort(containersByMethodName(endpoints.ContainerEndpoints))
	sort.Sort(listsByMethodName(endpoints.ListEndpoints))
	return endpoints, nil
}

func ApplyTemplate(epList *GnmiEndpoints, outPath string) error {

	var funcs template.FuncMap = map[string]interface{}{
		"quote": func(value interface{}) string {
			return fmt.Sprintf("\"%s\"", value)
		},
		"replace": func(search, replace string, value interface{}) string {
			return strings.ReplaceAll(fmt.Sprint(value), search, replace)
		},
		"sub": func(a int, b int) int {
			return a - b
		},
		"lower": strings.ToLower,
		"toName": func(s string) string {
			return genutil.EntryCamelCaseName(&yang.Entry{Name: s})
		},
		"hasParentKey": func(key ListKey) bool {
			return key.ParentKey != nil
		},
		// takes tuple of parameters and returns a map
		// it's used to pass multiple values to a template
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	}

	t, err := template.New(templateFile).
		Funcs(funcs).
		ParseFS(t.GnmiGenTemplate, "*.go.tpl")
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

// helpers to the sortByName method
func less(s []interface{}, i, j int) bool {
	ok1 := false
	ok2 := false

	// cast to LeavesEndpoint
	leaf1, ok1 := s[i].(LeavesEndpoint)
	leaf2, ok2 := s[j].(LeavesEndpoint)
	if ok1 && ok2 {
		return leaf1.MethodName < leaf2.MethodName
	}

	cont1, ok1 := s[i].(ContainerEndpoint)
	cont2, ok2 := s[j].(ContainerEndpoint)
	if ok1 && ok2 {
		return cont1.MethodName < cont2.MethodName
	}

	list1, ok1 := s[i].(ListEndpoint)
	list2, ok2 := s[j].(ListEndpoint)
	if ok1 && ok2 {
		return list1.MethodName < list2.MethodName
	}
	panic(fmt.Sprintf("type-%T-not-supported-for-ordering", s[i]))
}

type leavesByMethodName []LeavesEndpoint

func (s leavesByMethodName) Len() int      { return len(s) }
func (s leavesByMethodName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s leavesByMethodName) Less(i, j int) bool {
	// we have to manually convert to a list of interfaces
	b := make([]interface{}, len(s))
	for i := range s {
		b[i] = s[i]
	}
	return less(b, i, j)
}

type containersByMethodName []ContainerEndpoint

func (s containersByMethodName) Len() int      { return len(s) }
func (s containersByMethodName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s containersByMethodName) Less(i, j int) bool {
	// we have to manually convert to a list of interfaces
	b := make([]interface{}, len(s))
	for i := range s {
		b[i] = s[i]
	}
	return less(b, i, j)
}

type listsByMethodName []ListEndpoint

func (s listsByMethodName) Len() int      { return len(s) }
func (s listsByMethodName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s listsByMethodName) Less(i, j int) bool {
	// we have to manually convert to a list of interfaces
	b := make([]interface{}, len(s))
	for i := range s {
		b[i] = s[i]
	}
	return less(b, i, j)
}
