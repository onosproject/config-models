/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package testdata

import (
	"fmt"
	"github.com/openconfig/goyang/pkg/yang"
)

// this file contains the sample schemas to be used in the tests
// they are contained in a map indexed by the test name

// returns a yang.Entry to be used to generated the client in the tests
func GetSchema(testname string) (*yang.Entry, error) {
	// simple-leaves
	var leaf1 = &yang.Entry{
		Name: "leaf1",
		Kind: yang.LeafEntry,
		Type: &yang.YangType{
			Kind: yang.Ystring,
			Name: "string",
		},
	}

	var leaf2 = &yang.Entry{
		Name: "leaf2",
		Kind: yang.LeafEntry,
		Type: &yang.YangType{
			Kind: yang.Yuint16,
			Name: "uint16",
		},
	}

	// basic-container
	var cont1leaf1 = &yang.Entry{
		Name: "cont1leaf1",
		Kind: yang.LeafEntry,
		Type: &yang.YangType{
			Kind: yang.Ystring,
			Name: "string",
		},
	}

	var cont1leaf2 = &yang.Entry{
		Name: "cont1leaf2",
		Kind: yang.LeafEntry,
		Type: &yang.YangType{
			Kind: yang.Yuint16,
			Name: "uint16",
		},
	}

	var cont1 = &yang.Entry{
		Name: "cont1",
		Kind: yang.DirectoryEntry,
		Dir: map[string]*yang.Entry{
			"cont1leaf1": cont1leaf1,
			"cont1leaf2": cont1leaf2,
		},
		Annotation: map[string]interface{}{
			"structname": "Test_Cont1A",
		},
	}
	cont1leaf1.Parent = cont1
	cont1leaf2.Parent = cont1

	// basic list
	var list1leaf1 = &yang.Entry{
		Name: "list1leaf1",
		Kind: yang.LeafEntry,
		Type: &yang.YangType{
			Kind: yang.Ystring,
			Name: "string",
		},
	}

	var list1leaf2 = &yang.Entry{
		Name: "list1leaf2",
		Kind: yang.LeafEntry,
		Type: &yang.YangType{
			Kind: yang.Yuint16,
			Name: "uint16",
		},
	}
	var list1 = &yang.Entry{
		Name:     "list1",
		Kind:     yang.DirectoryEntry,
		ListAttr: &yang.ListAttr{MinElements: 0}, // only needed to match isList, not relevant for code generation
		Annotation: map[string]interface{}{
			"structname": "Test_List1",
		},
		Key: "list1leaf2", // use list1leaf2 as key so we test the conversion between int and string
		Dir: map[string]*yang.Entry{
			"list1leaf1": list1leaf1,
			"list1leaf2": list1leaf2,
		},
	}
	list1leaf1.Parent = list1
	list1leaf2.Parent = list1

	var TestYangSchemas = map[string]*yang.Entry{
		"empty-entry": nil,
		"simple-leaves": {
			Name: "Device",
			Dir: map[string]*yang.Entry{
				"leaf1": leaf1,
				"leaf2": leaf2,
			},
		},
		"basic-container": {
			Name: "Device",
			Dir: map[string]*yang.Entry{
				"cont1": cont1,
			},
		},
		"basic-list": {
			Name: "Device",
			Dir: map[string]*yang.Entry{
				"list1": list1,
			},
		},
	}

	entry, ok := TestYangSchemas[testname]
	if !ok {
		return nil, fmt.Errorf("Cannot load yang entry for test: %s", testname)
	}
	return entry, nil
}
