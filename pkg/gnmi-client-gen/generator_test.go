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
