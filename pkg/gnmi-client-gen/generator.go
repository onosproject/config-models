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
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/genutil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strings"
	"text/template"
)

var log = logging.GetLogger("gnmi-client-gen")

const templateFile = "gnmi_client.go.tpl"

type templateData struct {
	PluginName string
	Entry      *yang.Entry
}

func isContainer(e *yang.Entry) bool {
	return e.IsContainer()
}

func isList(e *yang.Entry) bool {
	return e.IsList()
}

// takes tuple of parameters and returns a map
// it's used to pass multiple values to a template
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// remove all chars that are not valid in variables/methods names
func sanitize(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

func capitalize(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func hasParent(entry *yang.Entry) bool {
	return entry.Parent != nil
}

// returns the correct format for an existing value
func goReturnVal(val *yang.Entry) (string, error) {
	switch val.Type.Kind {
	case yang.Yint8:
		return "int8(val.GetIntVal())", nil
	case yang.Yint16:
		return "int16(val.GetIntVal())", nil
	case yang.Yint32:
		return "int32(val.GetIntVal())", nil
	case yang.Yint64:
		return "int64(val.GetIntVal())", nil
	case yang.Yuint8:
		return "uint8(val.GetUintVal())", nil
	case yang.Yuint16:
		return "uint16(val.GetUintVal())", nil
	case yang.Yuint32:
		return "uint32(val.GetUintVal())", nil
	case yang.Yuint64:
		return "uint64(val.GetUintVal())", nil
	case yang.Ybool, yang.Yempty:
		return "val.GetBoolVal()", nil
	case yang.Ystring, yang.Yunion:
		return "val.GetStringVal()", nil
	case yang.Ydecimal64:
		return "float64(val.GetFloatVal())", nil
	case yang.Ybinary:
		return "val.GetBytesVal()", nil
	case yang.Yenum, yang.Yidentityref:
		return "int64(val.GetIntVal())", nil
	case yang.Yleafref:
		t, err := findLeafRefType(val.Type.Path, val)
		if err != nil {
			return "", err
		}

		if t == "float64" {
			t = "decimal64"
		}

		yt, ok := yang.TypeKindFromName[t]
		if !ok {
			return "", status.Errorf(codes.Internal, "type %s is not a valid yang kind", t)
		}
		return goReturnVal(&yang.Entry{Type: &yang.YangType{Kind: yt}})
	}

	return "", status.Errorf(codes.Unimplemented, "%T type is not supported yet", val.Type.Kind)
}

// returns the correct format to return in case of error
func goEmptyReturnVal(val *yang.Entry) (string, error) {
	switch val.Type.Kind {
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
		return "0", nil
	case yang.Ybool, yang.Yempty:
		return "false", nil
	case yang.Ystring:
		return "\"\"", nil
	case yang.Ybinary:
		return "nil", nil
	case yang.Yleafref:
		t, err := findLeafRefType(val.Type.Path, val)
		if err != nil {
			return "", err
		}

		if t == "float64" {
			t = "decimal64"
		}

		yt, ok := yang.TypeKindFromName[t]
		if !ok {
			return "", status.Errorf(codes.Internal, "type %s is not a valid yang kind", t)
		}
		return goEmptyReturnVal(&yang.Entry{Type: &yang.YangType{Kind: yt}})
	case yang.Yunion:
		return "\"\"", nil
	}

	return "", status.Errorf(codes.Unimplemented, "%T type is not supported yet", val.Type.Kind)
}

func goType(entry *yang.Entry) (string, error) {
	// NOTE inspired by https://github.com/openconfig/ygot/blob/master/ytypes/util_types.go#L353
	switch entry.Type.Kind {
	case yang.Yint8:
		return "int8", nil
	case yang.Yint16:
		return "int16", nil
	case yang.Yint32:
		return "int32", nil
	case yang.Yint64:
		return "int64", nil
	case yang.Yuint8:
		return "uint8", nil
	case yang.Yuint16:
		return "uint16", nil
	case yang.Yuint32:
		return "uint32", nil
	case yang.Yuint64:
		return "uint64", nil
	case yang.Ybool, yang.Yempty:
		return "bool", nil
	case yang.Ystring:
		return "string", nil
	case yang.Ydecimal64:
		return "float64", nil
	case yang.Ybinary:
		return "[]byte", nil
	case yang.Yenum:
		return "int64", nil
	case yang.Yidentityref:
		// FIXME this is not enough, eg:
		// SYSLOG_FACILITY is generated via YGOT as `type E_OpenconfigSystemLogging_SYSLOG_FACILITY int64`
		// NOTE that PrefixedName returns :SYSLOG_FACILITY
		//return entry.Type.IdentityBase.PrefixedName()
		// not ideal, but for now we'll take it
		log.Warnw("type is not supported yet", "kind", entry.Type.Kind.String(), "entry-name", entry.Name)
		return "int64", nil
	case yang.Yleafref:
		v, err := findLeafRefType(entry.Type.Path, entry)
		return v, err
	case yang.Yunion:
		return "string", nil
	}
	// not ideal, but for now we'll take it
	log.Warnw("type is not supported yet", "kind", entry.Type.Kind.String(), "entry-name", entry.Name)
	return "interface{}", nil
}

// extracts the structname from a yang entry
func structName(entry *yang.Entry, forList bool) (string, error) {
	name, ok := entry.Annotation["structname"]
	if !ok {
		return "", status.Errorf(codes.NotFound, "structname not found in annotations")
	}
	if entry.IsList() && forList {
		keys, err := listKeys(entry)
		if err != nil {
			return "", err
		}
		// if the entry is a list and we are targeting the container
		// NOTE we still need to support composite keys
		return fmt.Sprintf("map[%s]*%s", keys[0].Gotype, name), nil
	}
	return fmt.Sprintf("*%s", name), nil
}

// returns true if the entry is the root of the tree
func isRoot(entry *yang.Entry) bool {
	if entry.Annotation != nil {
		if v, ok := entry.Annotation["isFakeRoot"]; ok {
			return v == true
		}
	}
	return false
}

// return the path to a model from the root
func devicePath(entry *yang.Entry, forParent bool) (string, error) {

	name := genutil.EntryCamelCaseName(entry)
	if entry.IsList() {
		var keyName string
		// if the item is a list, we need to return the element in the list based on the key
		keys := strings.Split(entry.Key, " ")
		if len(keys) > 1 {
			keyName = "key"
		} else {
			keyName = entry.Key
		}
		if !forParent {
			name = sanitize(fmt.Sprintf("%s[%s]", name, fmt.Sprintf("%s_%s", entry.Name, keyName)))
		}
	}

	if entry.Parent != nil && !isRoot(entry.Parent) {
		// forParent applies only for the last level,
		// in a structure like item1 -> item2 -> item3
		// we return:
		// - forParent = false => Item1.Item2.Item3
		// - forParent = true => Item1.Item2
		parentName, err := devicePath(entry.Parent, false)
		if err != nil {
			return "", err
		}
		if !forParent || entry.IsList() {
			return fmt.Sprintf("%s.%s", parentName, name), nil
		}
		return parentName, nil
	}

	return name, nil
}

// extract the list keys from a yang entry
// and returns the tuples
type YangKey struct {
	// Name represents the variable name in the method arguments,
	// it's prefixed with the entry name to avoid duplicates, eg: application_id and device_id
	Name string
	// Key is the name of the key in the Yang path
	Key string
	// Gotype is the Go type that corresponds to a key, used to build the method parameters list
	Gotype string
}

func listKeys(entry *yang.Entry) ([]YangKey, error) {
	list := []YangKey{}
	if !entry.IsList() {
		// if this is not a list, return an empty list
		return list, nil
	}
	for _, k := range strings.Split(entry.Key, " ") {
		// find the child entry the key refers to get the type
		if e, ok := entry.Dir[k]; ok {
			// the only error case is for unsupported types, ignore for now
			if t, err := goType(e); err == nil {
				list = append(list, YangKey{
					Name:   fmt.Sprintf("%s_%s", entry.Name, k),
					Key:    k,
					Gotype: t,
				})
			}

		}
	}
	return list, nil
}

func Generate(pluginName string, entry *yang.Entry, output io.Writer) error {

	if entry == nil {
		return fmt.Errorf("entry-cannot-be-nil")
	}

	templateFunctions := template.FuncMap{
		"dict":             dict,
		"isContainer":      isContainer,
		"isList":           isList,
		"sanitize":         sanitize,
		"capitalize":       capitalize,
		"hasParent":        hasParent,
		"goReturnVal":      goReturnVal,
		"goEmptyReturnVal": goEmptyReturnVal,
		"goType":           goType,
		"structName":       structName,
		"devicePath":       devicePath,
		"listKeys":         listKeys,
		"isRoot":           isRoot,
	}

	t, err := template.New(templateFile).
		Funcs(templateFunctions).
		ParseFS(t.GnmiGenTemplate, "*.go.tpl")
	if err != nil {
		return err
	}

	return t.Execute(output, templateData{
		PluginName: pluginName,
		Entry:      entry,
	})
}

func findLeafRefType(path string, entry *yang.Entry) (string, error) {

	var dp []string
	var curEntry = entry

	// if the path is absolute (eg: /t1:cont1a/t1:list2a/t1:name) go back to the root and then descend
	if path[0:1] == "/" {
		dp = strings.Split(path[1:], "/")

		// go back till the root
		for {
			if curEntry.Parent != nil {
				curEntry = curEntry.Parent
			} else {
				break
			}
		}

	} else {
		// identify how many levels we have to go up the tree
		lb := strings.Count(path, "../")

		// identify the path to take once we have gone up the tree
		dp = strings.Split(strings.ReplaceAll(path, "../", ""), "/")

		// this is the entry we are moving to

		for i := 0; i < lb; i++ {
			// we're going up the tree till it's needed
			curEntry = curEntry.Parent
		}

	}

	var downwardPath = []string{}
	// remove the prefix from the pieces
	for _, p := range dp {
		// a piece of the path can be ta:cont1a or cont1a
		k := strings.Split(p, ":")
		if len(k) > 1 {
			downwardPath = append(downwardPath, k[1])
		} else {
			downwardPath = append(downwardPath, p)
		}
	}

	for _, k := range downwardPath {
		// and then descending to the leafref path
		curEntry = curEntry.Dir[k]
	}

	// not simply convert the type
	return goType(curEntry)

}
