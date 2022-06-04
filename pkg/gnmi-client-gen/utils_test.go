/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetListKey(t *testing.T) {

	type args struct {
		entry *yang.Entry
	}

	tests := []struct {
		name string
		args args
		want ListKey
	}{
		{
			name: "simple-key",
			args: args{
				entry: &yang.Entry{
					Name: "foo",
					Key:  "id",
					Dir: map[string]*yang.Entry{
						"id": {Type: &yang.YangType{Kind: yang.Ystring}},
					},
					Annotation: map[string]interface{}{
						"structname": "OnfTest1_Cont1A_List5",
					},
				},
			},
			want: ListKey{ModelName: "foo", Type: "string", Keys: []Key{{Name: "Id", Type: "string"}}}},
		{
			name: "composite-key",
			args: args{
				entry: &yang.Entry{
					Name: "list5",
					Key:  "key1 key2",
					Annotation: map[string]interface{}{
						"structname": "OnfTest1_Cont1A_List5",
					},
					Dir: map[string]*yang.Entry{
						"key1": {Type: &yang.YangType{Kind: yang.Ystring}},
						"key2": {Type: &yang.YangType{Kind: yang.Ystring}},
					},
				},
			},
			want: ListKey{
				ModelName: "list5",
				Type:      "OnfTest1_Cont1A_List5_Key",
				Keys: []Key{
					{Name: "Key1", Type: "string"},
					{Name: "Key2", Type: "string"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetListKey(tt.args.entry)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.ModelName, res.ModelName)
			assert.Equal(t, tt.want.Type, res.Type)
			assert.Equal(t, len(tt.want.Keys), len(res.Keys))
			for i, k := range res.Keys {
				w := tt.want.Keys[i]
				assert.Equal(t, w.Name, k.Name)
				assert.Equal(t, w.Type, k.Type)
			}
		})
	}
}

func TestYangTypeToGoType(t *testing.T) {

	type args struct {
		entry *yang.Entry
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"string-type", args{entry: &yang.Entry{Type: &yang.YangType{Kind: yang.Ystring}}}, "string"},
		{
			"leafref-one-level-above",
			args{
				entry: &yang.Entry{
					Parent: &yang.Entry{
						Name: "parent",
						Dir: map[string]*yang.Entry{
							"config": {
								Name: "config",
								Dir: map[string]*yang.Entry{
									"key-leafref": {
										Name: "key-leafref",
										Type: &yang.YangType{
											Kind: yang.Yuint16,
										},
									},
								},
							},
						},
					},
					Name: "key-leafref",
					Type: &yang.YangType{
						Kind: yang.Yleafref,
						Path: "../config/key-leafref",
					},
				},
			},
			"uint16",
		},
		{
			"leafref-absolute-path",
			args{
				entry: &yang.Entry{
					Name:        "id",
					Description: "Link to list2a names",
					Type: &yang.YangType{
						Kind: yang.Yleafref,
						Path: "/t1:cont1a/t1:list2a/t1:name",
					},
					Parent: &yang.Entry{
						Name: "list4",
						Parent: &yang.Entry{
							Name: "cont1a",
							Parent: &yang.Entry{
								Name:   "device",
								Parent: nil,
								Dir: map[string]*yang.Entry{
									"cont1a": {
										Name: "cont1a",
										Dir: map[string]*yang.Entry{
											"list2a": {
												Name: "list2a",
												Dir: map[string]*yang.Entry{
													"name": {
														Name: "name",
														Type: &yang.YangType{
															Kind: yang.Ystring,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := yangTypeToGoType(tt.args.entry)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestPathToYgotModelPath(t *testing.T) {
	type args struct {
		path []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{path: []string{"foo"}}, "Foo"},
		{"test1", args{path: []string{"foo", "bar"}}, "Foo.Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PathToYgotModelPath(tt.args.path), "PathToYgotModelPath(%v)", tt.args.path)
		})
	}
}
