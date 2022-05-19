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
	"strings"
)

func PathToCamelCaseName(path []string) string {
	name := ""
	for _, p := range path {
		name += genutil.EntryCamelCaseName(&yang.Entry{Name: p})
	}
	return name
}

// deprecated: the model name is in item.Annotation["structname"]
func PathToYgotModelName(path []string) string {
	name := []string{}
	for _, p := range path {
		name = append(name, genutil.EntryCamelCaseName(&yang.Entry{Name: p}))
	}
	return strings.Join(name, "_")
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
				key.Keys = append(key.Keys, Key{
					Name: caser.String(k),
					Type: yangTypeToGoType(kv.Type.Kind),
				})
			}
		}
		return key, nil
	} else {
		if kv, ok := entry.Dir[keys[0]]; ok {
			return ListKey{
				Type: yangTypeToGoType(kv.Type.Kind),
				Keys: []Key{
					{
						Name: caser.String(keys[0]),
						Type: yangTypeToGoType(kv.Type.Kind),
					},
				},
			}, nil
		}

	}
	return ListKey{}, fmt.Errorf("cannot-generate-key-from-%s", keys)
}

func yangTypeToGoType(val yang.TypeKind) string {
	// NOTE inspired by https://github.com/openconfig/ygot/blob/master/ytypes/util_types.go#L353
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
	case yang.Yleafref:
		return "string"
	}
	// not ideal, but for now we'll take it
	return "interface{}"
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
