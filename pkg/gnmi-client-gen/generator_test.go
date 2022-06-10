/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"bytes"
	"fmt"
	"github.com/onosproject/config-models/pkg/gnmi-client-gen/testdata"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	type arguments struct {
		pluginName string
		entry      *yang.Entry
	}
	tests := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty-entry",
			func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.Equal(t, "entry-cannot-be-nil", err.Error())
			},
		},
		{
			"simple-leaves",
			assert.NoError,
		},
		{
			"basic-container",
			assert.NoError,
		},
		{
			"basic-list",
			assert.NoError,
		},
		{
			"nested-list",
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			schema, err := testdata.GetSchema(tt.name)
			if err != nil {
				t.Fatalf(err.Error())
			}

			wantOutput := testdata.GetTestResult(t, tt.name)

			args := arguments{
				pluginName: "Test",
				entry:      schema,
			}

			output := &bytes.Buffer{}
			err = Generate(args.pluginName, args.entry, output)
			if !tt.wantErr(t, err, fmt.Sprintf("Generate(%v, %v, %v)", args.pluginName, args.entry, output)) {
				return
			}

			debug := os.Getenv("DEBUG")

			if debug == "true" {
				// when debugging keep whitespaces, they might fail the test
				// but it's more readable
				fmt.Println(output)
				assert.Equal(t, wantOutput, output.String())
			} else {
				// when not debugging strip whitespaces as they're not relevant
				assert.Equalf(t, testdata.RemoveAllWhitespaces(wantOutput), testdata.RemoveAllWhitespaces(output.String()), "Failed to generate template for test: %s", tt.name)
			}

		})
	}
}

func Test_devicePath(t *testing.T) {
	type args struct {
		entry     *yang.Entry
		forParent bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"leaf-at-top-level",
			args{
				entry: &yang.Entry{
					Name: "leaf1",
					Kind: yang.LeafEntry,
					Parent: &yang.Entry{
						Name: "device",
						Annotation: map[string]interface{}{
							"isFakeRoot": true,
						},
					},
				},
				forParent: false,
			},
			"Leaf1",
			assert.NoError,
		},
		{
			"path-for-parent-model",
			args{
				entry: &yang.Entry{
					Name: "leaf3",
					Kind: yang.LeafEntry,
					Parent: &yang.Entry{
						Name: "leaf2",
						Kind: yang.DirectoryEntry,
						Parent: &yang.Entry{
							Name: "leaf1",
							Kind: yang.DirectoryEntry,
						},
					},
				},
				forParent: true,
			},
			"Leaf1.Leaf2",
			assert.NoError,
		},
		{
			"basic-list",
			args{
				entry: &yang.Entry{
					Name:     "list1",
					Kind:     yang.DirectoryEntry,
					ListAttr: &yang.ListAttr{MinElements: 1},
					Key:      "id",
					Dir: map[string]*yang.Entry{
						"id": {},
					},
				},
				forParent: false,
			},
			"List1[list1_id]",
			assert.NoError,
		},
		{
			"double-keyed-list",
			args{
				entry: &yang.Entry{
					Name:     "list1",
					Kind:     yang.DirectoryEntry,
					ListAttr: &yang.ListAttr{MinElements: 1},
					Key:      "foo bar",
					Dir: map[string]*yang.Entry{
						"foo": {},
						"bar": {},
					},
				},
				forParent: false,
			},
			"List1[list1_key]",
			assert.NoError,
		},
		{
			"nested-list",
			args{
				entry: &yang.Entry{
					Name:     "list2",
					Kind:     yang.DirectoryEntry,
					ListAttr: &yang.ListAttr{MinElements: 1},
					Key:      "foo bar",
					Dir: map[string]*yang.Entry{
						"foo": {},
						"bar": {},
					},
					Parent: &yang.Entry{
						Name:     "list1",
						Kind:     yang.DirectoryEntry,
						ListAttr: &yang.ListAttr{MinElements: 1},
						Key:      "id",
						Dir: map[string]*yang.Entry{
							"id":    {},
							"list2": {},
						},
					},
				},
				forParent: false,
			},
			"List1[list1_id].List2[list2_key]",
			assert.NoError,
		},
		{
			"nested-list-for-parent",
			args{
				entry: &yang.Entry{
					Name:     "list2",
					Kind:     yang.DirectoryEntry,
					ListAttr: &yang.ListAttr{MinElements: 1},
					Key:      "foo",
					Dir: map[string]*yang.Entry{
						"foo": {},
					},
					Parent: &yang.Entry{
						Name:     "list1",
						Kind:     yang.DirectoryEntry,
						ListAttr: &yang.ListAttr{MinElements: 1},
						Key:      "id",
						Dir: map[string]*yang.Entry{
							"id":    {},
							"list2": {},
						},
					},
				},
				forParent: true,
			},
			"List1[list1_id].List2",
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := devicePath(tt.args.entry, tt.args.forParent)
			if !tt.wantErr(t, err, fmt.Sprintf("devicePath(%v)", tt.args.entry)) {
				return
			}
			assert.Equalf(t, tt.want, got, "devicePath(%v)", tt.args.entry)
		})
	}
}

func Test_findLeafRefType(t *testing.T) {
	type args struct {
		entry *yang.Entry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
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
			assert.NoError,
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
			assert.NoError,
		},
		{
			"leafref-relative-with-prefix",
			args{
				entry: &yang.Entry{
					Name: "leaf",
					Type: &yang.YangType{
						Kind: yang.Yleafref,
						Path: "../sw:vlan/sw:vlan-id",
					},
					Parent: &yang.Entry{
						Name: "parent",
						Dir: map[string]*yang.Entry{
							"vlan": {
								Name: "vlan",
								Dir: map[string]*yang.Entry{
									"vlan-id": {
										Name: "vlan-id",
										Type: &yang.YangType{
											Kind: yang.Yuint16,
										},
									},
								},
							},
						},
					},
				},
			},
			"uint16",
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findLeafRefType(tt.args.entry.Type.Path, tt.args.entry)
			if !tt.wantErr(t, err, fmt.Sprintf("findLeafRefType(%v, %v)", tt.args.entry.Type.Path, tt.args.entry)) {
				return
			}
			assert.Equalf(t, tt.want, got, "findLeafRefType(%v, %v)", tt.args.entry.Type.Path, tt.args.entry)
		})
	}
}
