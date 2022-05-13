/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/genutil"
	"strings"
)

func PathToCamelCaseName(path []string) string {
	name := ""
	for _, p := range path {
		name += genutil.EntryCamelCaseName(&yang.Entry{Name: p})
	}
	return name
}

func PathToYgotModelName(path []string, prefix string) string {
	name := []string{prefix}
	for _, p := range path {
		name = append(name, genutil.EntryCamelCaseName(&yang.Entry{Name: p}))
	}
	return strings.Join(name, "_")
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
	return "GetValue()"
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
		return "[]byte{}"
	}
	// not ideal, but for now we'll take it
	return "interface{}"
}
