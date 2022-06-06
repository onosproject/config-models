/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	aether_2_1_x "github.com/onosproject/aether-models/models/aether-2.1.x/api"
	testdevice "github.com/onosproject/config-models/models/testdevice-1.0.x/api"
	"github.com/onosproject/config-models/pkg/gnmi-client-gen-v2/testdata"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExtractSchema(t *testing.T) {
	// uncomment this to manually run the test to regenerate the sample schema
	// not needed otherwise
	t.Skip()
	folder := "sampleSchemas"

	testdevice, err := testdevice.Schema()
	assert.NoError(t, err)

	a21x, err := aether_2_1_x.Schema()
	assert.NoError(t, err)

	schemas := map[string]*ytypes.Schema{
		"testdevice-1": testdevice,
		"aether-2.1.x": a21x,
	}

	for name, schema := range schemas {
		empJSON, err := json.MarshalIndent(schema.SchemaTree["Device"], "", "  ")
		assert.NoError(t, err)

		file, err := os.Create(fmt.Sprintf("%s/%s.json", folder, name))
		if err != nil {
			t.Fail()
		}
		_, err = file.WriteString(string(empJSON))
		assert.NoError(t, err)
		defer file.Close()
	}

}

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
			"List1[list1_key]",
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
			"List1[list1_key].List2[list2_key]",
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
