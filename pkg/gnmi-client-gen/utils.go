/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"fmt"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/genutil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func PathToYgotModelName(path []string) string {
	name := []string{}
	for _, p := range path {
		name = append(name, genutil.EntryCamelCaseName(&yang.Entry{Name: p}))
	}
	return strings.Join(name, "_")
}

func PathToCamelCaseName(path []string) string {
	name := ""
	for _, p := range path {
		name += genutil.EntryCamelCaseName(&yang.Entry{Name: p})
	}
	return name
}

func PathToYgotModelPath(path []string) string {
	name := []string{}
	for _, p := range path {
		name = append(name, genutil.EntryCamelCaseName(&yang.Entry{Name: p}))
	}
	return strings.Join(name, ".")
}

// given a yang entry with a key, return the appropriate type for the key
func GetListKey(entry *yang.Entry) (ListKey, error) {
	keys := strings.Split(entry.Key, " ")
	modelName := fmt.Sprintf("%s", entry.Annotation["structname"])
	caser := cases.Title(language.English)

	if len(keys) > 1 {
		key := ListKey{
			Type: fmt.Sprintf("%s_Key", modelName),
			Keys: []Key{},
		}

		for _, k := range keys {
			if kv, ok := entry.Dir[k]; ok {
				t, err := yangTypeToGoType(kv)
				if err != nil {
					return ListKey{}, err
				}
				key.Keys = append(key.Keys, Key{
					Name: caser.String(k),
					Type: t,
				})
			}
		}
		return key, nil
	} else {
		if kv, ok := entry.Dir[keys[0]]; ok {
			t, err := yangTypeToGoType(kv)
			if err != nil {
				return ListKey{}, err
			}
			return ListKey{
				Type: t,
				Keys: []Key{
					{
						Name: caser.String(keys[0]),
						Type: t,
					},
				},
			}, nil
		}

	}
	return ListKey{}, fmt.Errorf("cannot-generate-key-from-%s", keys)
}

func yangTypeToGoType(entry *yang.Entry) (string, error) {
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
		return findLeafRefType(entry.Type.Path, entry)
		//case yang.Yunion:

	}
	// not ideal, but for now we'll take it
	return "", status.Error(codes.Unimplemented, fmt.Sprintf("%s type is not supported yet", entry.Type.Kind.String()))
}

// returns the correct format for an existing value
func yangTypeToGoReturnVal(val yang.TypeKind) string {
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
	return "val.GetValue()"
}

// returns the correct format to return in case of error
func yangTypeToGoEmptyReturnVal(val yang.TypeKind) string {
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
		return "nil"
	}
	// not ideal, but for now we'll take it
	return "nil"
}

func findLeafRefType(path string, entry *yang.Entry) (string, error) {

	var downwardPath = []string{}
	var cur_entry = entry

	// if the path is absolute (eg: /t1:cont1a/t1:list2a/t1:name) go back to the root and then descend
	if path[0:1] == "/" {
		dp := strings.Split(path[1:], "/")

		// remove the prefix from the pieces
		for _, p := range dp {
			k := strings.Split(p, ":")
			downwardPath = append(downwardPath, k[1])
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
	return yangTypeToGoType(cur_entry)

}
