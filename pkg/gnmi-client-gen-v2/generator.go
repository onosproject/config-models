/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen_v2

//{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry -}}
import (
	"fmt"
	t "github.com/onosproject/config-models/pkg/gnmi-client-gen-v2/template"
	"github.com/openconfig/goyang/pkg/yang"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strings"
	"text/template"
)

const templateFile = "gnmi_client.go.tpl"

type templateData struct {
	PluginName string
	Entry      *yang.Entry
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
	case yang.Ystring:
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
		return "", status.Error(codes.Unimplemented, "yang.Yidentityref type is not supported yet")
	case yang.Yleafref:
		v, err := findLeafRefType(entry.Type.Path, entry)
		fmt.Println("TEO", entry.Name, entry.Type.Path, v)
		return v, err
		//case yang.Yunion:

	}
	// not ideal, but for now we'll take it
	return "", status.Error(codes.Unimplemented, fmt.Sprintf("%s type is not supported yet", entry.Type.Kind.String()))
}

func Generate(pluginName string, entry *yang.Entry, output io.Writer) error {

	if entry == nil {
		return fmt.Errorf("entry-cannot-be-nil")
	}

	templateFunctions := template.FuncMap{
		"sanitize":         sanitize,
		"capitalize":       capitalize,
		"hasParent":        hasParent,
		"goReturnVal":      goReturnVal,
		"goEmptyReturnVal": goEmptyReturnVal,
		"goType":           goType,
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

	var downwardPath = []string{}
	var cur_entry = entry

	// if the path is absolute (eg: /t1:cont1a/t1:list2a/t1:name) go back to the root and then descend
	if path[0:1] == "/" {
		dp := strings.Split(path[1:], "/")

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
		// go back till the root
		for {
			if cur_entry.Parent != nil {
				cur_entry = cur_entry.Parent
			} else {
				break
			}
		}

	} else {
		// identify how many levels we have to go up the tree
		lb := strings.Count(path, "../")

		// identify the path to take once we have gone up the tree
		downwardPath = strings.Split(strings.ReplaceAll(path, "../", ""), "/")

		// this is the entry we are moving to

		for i := 0; i < lb; i++ {
			// we're going up the tree till it's needed
			cur_entry = cur_entry.Parent
		}

	}
	for _, k := range downwardPath {
		// and then descending to the leafref path
		cur_entry = cur_entry.Dir[k]
	}

	// not simply convert the type
	return goType(cur_entry)

}
