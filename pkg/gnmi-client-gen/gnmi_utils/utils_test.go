/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_utils

import (
	"context"
	"fmt"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type nestedStruct struct {
	string *string `path:"ns_string"`
}

type TestStruct struct {
	String          *string       `path:"string"`
	Uint8           *uint8        `path:"uint8"`
	NestedStruct    nestedStruct  `path:"nested_struct"`
	NestedStructPtr *nestedStruct `path:"nested_struct_ptr"`
}

type FieldIndex int

func (f FieldIndex) Int() int {
	return int(f)
}

const (
	strFieldIndex             FieldIndex = 0
	uint8FieldIndex           FieldIndex = 1
	nestedStructFieldIndex    FieldIndex = 2
	nestedStructPtrFieldIndex FieldIndex = 3
)

var (
	str string = "str"
	u8  uint8  = 2
)

func TestHasNonZeroField(t *testing.T) {

	type args struct {
		data      TestStruct
		fieldIdx  FieldIndex // identifies the position of the field we are checking in the struct
		fieldName string     // name of the field
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// check that we return the correct value for basic fields
		{"stringFieldHasValue", args{data: TestStruct{String: &str}, fieldIdx: strFieldIndex}, true},
		{"stringFieldNoValue", args{data: TestStruct{}, fieldIdx: strFieldIndex}, false},
		{"uint8FieldHasValue", args{data: TestStruct{Uint8: &u8}, fieldIdx: uint8FieldIndex}, true},
		{"uint8FieldNoValue", args{data: TestStruct{}, fieldIdx: uint8FieldIndex}, false},
		// and for nested structures
		{"nestedStructFieldHasValue", args{data: TestStruct{NestedStruct: nestedStruct{string: &str}}, fieldIdx: nestedStructFieldIndex}, true},
		{"nestedStructFieldHasEmptyStruct", args{data: TestStruct{NestedStruct: nestedStruct{}}, fieldIdx: nestedStructFieldIndex}, false},
		{"nestedStructFieldNoValue", args{data: TestStruct{}, fieldIdx: nestedStructFieldIndex}, false},
		// and for pointers to nested structures
		{"nestedStructPtrFieldHasValue", args{data: TestStruct{NestedStructPtr: &nestedStruct{string: &str}}, fieldIdx: nestedStructPtrFieldIndex}, true},
		{"nestedStructPtrFieldHasPointerToEmptyStruct", args{data: TestStruct{NestedStructPtr: &nestedStruct{}}, fieldIdx: nestedStructPtrFieldIndex}, true}, // a pointer to an empty structure is still a valid pointer
		{"nestedStructPtrFieldNoValue", args{data: TestStruct{}, fieldIdx: nestedStructPtrFieldIndex}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_v := reflect.ValueOf(tt.args.data).Field(tt.args.fieldIdx.Int())
			_t := reflect.ValueOf(tt.args.data).Type().Field(tt.args.fieldIdx.Int())
			res := hasNonZeroField(_t, _v)
			logMsg := fmt.Sprintf("field '%s'", _t.Name)
			if tt.want {
				logMsg += " expected to have a value, but it didn't."
			} else {
				logMsg += " NOT expected to have a value, but it did."
			}
			assert.Equal(t, tt.want, res, fmt.Sprintf("Test %s failed checking field '%s': %s", tt.name, tt.args.fieldName, logMsg))
		})
	}
}

func TestCreateGnmiSetForContainer(t *testing.T) {
	ctx := context.Background()
	target := "test-target"
	basePath := &gnmi.Path{
		Elem: []*gnmi.PathElem{
			{
				Name: "cont1a",
			},
			{
				Name: "cont2a",
			},
		},
		Target: target,
	}

	tests := []struct {
		name string
		args TestStruct
		want *gnmi.SetRequest
	}{
		// we have 2 fields in the struct, we should be getting 2 updates
		{"setReqNoNestedFields", TestStruct{String: &str, Uint8: &u8}, &gnmi.SetRequest{
			Update: []*gnmi.Update{
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: str}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "string"})},
				},
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_UintVal{UintVal: uint64(u8)}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "uint8"})},
				},
			},
		}},
		// we have a single field in the structs, we should be getting 1 update
		{"setReqNoNestedFields2", TestStruct{String: &str}, &gnmi.SetRequest{
			Update: []*gnmi.Update{
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: str}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "string"})},
				},
			},
		}},
		// we have a nested structure with a field, we should be getting one update with the appropriate path
		{"setReqNestedSingle", TestStruct{NestedStruct: nestedStruct{string: &str}}, &gnmi.SetRequest{
			Update: []*gnmi.Update{
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: str}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "nested_struct"}, &gnmi.PathElem{Name: "ns_string"})},
				},
			},
		}},
		// we have two values and a nested structure with a field, we should be getting three updates with the appropriate path
		{"setReqNested", TestStruct{String: &str, Uint8: &u8, NestedStruct: nestedStruct{string: &str}}, &gnmi.SetRequest{
			Update: []*gnmi.Update{
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: str}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "string"})},
				},
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_UintVal{UintVal: uint64(u8)}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "uint8"})},
				},
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: str}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "nested_struct"}, &gnmi.PathElem{Name: "ns_string"})},
				},
			},
		}},
		// we have a pointer to a nested structure with a field, we should be getting one update with the appropriate path
		{"setReqNestedPtrSingle", TestStruct{NestedStructPtr: &nestedStruct{string: &str}}, &gnmi.SetRequest{
			Update: []*gnmi.Update{
				{
					Val:  &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: str}},
					Path: &gnmi.Path{Target: target, Elem: append(basePath.Elem, &gnmi.PathElem{Name: "nested_struct_ptr"}, &gnmi.PathElem{Name: "ns_string"})},
				},
			},
		}},
		// we have a pointer to an empty nested structure with a field, we should be getting no updates
		{"setReqNestedPtrEmpty", TestStruct{NestedStructPtr: &nestedStruct{}}, &gnmi.SetRequest{
			Update: []*gnmi.Update{},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := CreateGnmiSetForContainer(ctx, tt.args, basePath, target)
			assert.Nil(t, err, fmt.Sprintf("Test %s failed: %s", tt.name, err))
			assert.Equal(t, tt.want, req, fmt.Sprintf("Test %s failed: expected and generated SetRequests are different", tt.name))
		})
	}
}
