// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi_gen

import (
	"github.com/openconfig/goyang/pkg/yang"
	"gotest.tools/assert"
	"testing"
)

// Test the range min..10 | 20..100
func Test_yangRangeDouble(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Kind:  yang.MinNumber,
			Value: 0,
		},
		Max: yang.Number{
			Kind:  yang.Positive,
			Value: 10,
		},
	})
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Kind:  yang.Positive,
			Value: 20,
		},
		Max: yang.Number{
			Kind:  yang.Positive,
			Value: 100,
		},
	})

	min, max, err := yangRange(testRange1)
	assert.NilError(t, err)
	assert.Assert(t, min == nil)
	assert.Assert(t, max != nil)
	if max != nil {
		assert.Equal(t, 100.0, *max)
	}
}

// Test the range -0.02..0.002
func Test_yangRangeDecimal(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Kind:           yang.Negative,
			Value:          201,
			FractionDigits: 2,
		},
		Max: yang.Number{
			Kind:           yang.Positive,
			Value:          2005,
			FractionDigits: 2,
		},
	})

	min, max, err := yangRange(testRange1)
	assert.NilError(t, err)
	assert.Assert(t, min != nil)
	if min != nil {
		assert.Equal(t, 2.0100000000000002, *min)
	}
	assert.Assert(t, max != nil)
	if max != nil {
		assert.Equal(t, 20.05, *max)
	}
}

func Test_pathToSchemaName(t *testing.T) {
	schName1 := pathToSchemaName("/qos-profile/qos-profile/{id}/arp")
	assert.Equal(t, "/qos-profile/qos-profile/arp", schName1)

	schName2 := pathToSchemaName("/subscriber/ue/{id}/profiles/access-profile/{access-profile}")
	assert.Equal(t, "/subscriber/ue/profiles/access-profile", schName2)
}

func Test_newPathItem(t *testing.T) {

	targetParameter = targetParam("test-target")

	testDirEntry := yang.Entry{
		Config: yang.TSTrue,
		Parent: &yang.Entry{},
		Type: &yang.YangType{
			Name: "Test1",
		},
	}

	pathItem := newPathItem(&testDirEntry, "/test-1/test-2/{id}/test-3/{id}/test-4",
		"/parent-1/{parent1-name}/parent-2/{parent2-name}")
	assert.Assert(t, pathItem != nil)
	if pathItem != nil {
		g := pathItem.Get
		assert.Assert(t, g != nil)
		if g != nil {
			assert.Equal(t, "GET /test-1/test-2/{id}/test-3/{id}/test-4 Generated from YANG model", g.Summary)
			assert.DeepEqual(t, []string{"Parent-1_Parent-2"}, g.Tags)
		}
		assert.Assert(t, pathItem.Post != nil)
		assert.Equal(t, "POST Generated from YANG model", pathItem.Post.Summary)
		assert.Equal(t, "postTest-1_Test-2_Test-3_Test-4", pathItem.Post.OperationID)

		assert.Assert(t, pathItem.Delete != nil)
		assert.Equal(t, "DELETE Generated from YANG model", pathItem.Delete.Summary)
		assert.Equal(t, "deleteTest-1_Test-2_Test-3_Test-4", pathItem.Delete.OperationID)

		assert.Equal(t, 3, len(pathItem.Parameters))
		for _, p := range pathItem.Parameters {
			switch name := p.Value.Name; name {
			case "target":
				assert.Equal(t, "target (device in onos-config)", p.Value.Description)
			case "parent1-name":
				assert.Equal(t, "key {parent1-name}", p.Value.Description)
			case "parent2-name":
				assert.Equal(t, "key {parent2-name}", p.Value.Description)
			default:
				t.Errorf("unexpected parameter %s", name)
			}
		}
	}
}

func Test_buildSchemaIntegerLeaf(t *testing.T) {

	targetParameter = targetParam("test-target")

	testLeaf1 := yang.Entry{
		Name:   "Leaf1",
		Config: yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yint16,
		},
	}

	testDirEntry := yang.Entry{
		Name:   "Test1",
		Parent: &yang.Entry{},
		Kind:   yang.DirectoryEntry,
		Config: yang.TSTrue,
		Dir:    make(map[string]*yang.Entry),
		Type: &yang.YangType{
			Name: "Test1",
		},
	}
	testLeaf1.Parent = &testDirEntry
	testDirEntry.Dir["leaf1"] = &testLeaf1

	paths, components, err := buildSchema(&testDirEntry, yang.TSUnset, "/test")
	assert.NilError(t, err)
	assert.Equal(t, len(paths), 0)
	assert.Equal(t, len(components.Schemas), 1)
	s, ok := components.Schemas["Test_Leaf1"]
	assert.Assert(t, ok, "expecting Test_Leaf1")
	assert.Equal(t, "Leaf1", s.Value.Title)
	assert.Equal(t, "integer", s.Value.Type)
}
