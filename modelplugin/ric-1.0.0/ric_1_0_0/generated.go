// Code generated by YGOT. DO NOT EDIT.
/*
Package ric_1_0_0 is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /Users/adibrastegarnia/go/pkg/mod/github.com/openconfig/ygot@v0.8.12/genutil/names.go
using the following YANG input files:
	- test1@2020-11-18.yang
	- xapp@2020-11-24.yang
Imported modules were sourced from:
	- yang/...
*/
package ric_1_0_0

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
)

// Binary is a type that is used for fields that have a YANG type of
// binary. It is used such that binary fields can be distinguished from
// leaf-lists of uint8s (which are mapped to []uint8, equivalent to
// []byte in reflection).
type Binary []byte

// YANGEmpty is a type that is used for fields that have a YANG type of
// empty. It is used such that empty fields can be distinguished from boolean fields
// in the generated code.
type YANGEmpty bool

var (
	SchemaTree map[string]*yang.Entry
)

func init() {
	var err error
	if SchemaTree, err = UnzipSchema(); err != nil {
		panic("schema error: " +  err.Error())
	}
}

// Schema returns the details of the generated schema.
func Schema() (*ytypes.Schema, error) {
	uzp, err := UnzipSchema()
	if err != nil {
		return nil, fmt.Errorf("cannot unzip schema, %v", err)
	}

	return &ytypes.Schema{
		Root: &Device{},
		SchemaTree: uzp,
		Unmarshal: Unmarshal,
	}, nil
}

// UnzipSchema unzips the zipped schema and returns a map of yang.Entry nodes,
// keyed by the name of the struct that the yang.Entry describes the schema for.
func UnzipSchema() (map[string]*yang.Entry, error) {
	var schemaTree map[string]*yang.Entry
	var err error
	if schemaTree, err = ygot.GzipToSchema(ySchema); err != nil {
		return nil, fmt.Errorf("could not unzip the schema; %v", err)
	}
	return schemaTree, nil
}

// Unmarshal unmarshals data, which must be RFC7951 JSON format, into
// destStruct, which must be non-nil and the correct GoStruct type. It returns
// an error if the destStruct is not found in the schema or the data cannot be
// unmarshaled. The supplied options (opts) are used to control the behaviour
// of the unmarshal function - for example, determining whether errors are
// thrown for unknown fields in the input JSON.
func Unmarshal(data []byte, destStruct ygot.GoStruct, opts ...ytypes.UnmarshalOpt) error {
	tn := reflect.TypeOf(destStruct).Elem().Name()
	schema, ok := SchemaTree[tn]
	if !ok {
		return fmt.Errorf("could not find schema for type %s", tn )
	}
	var jsonTree interface{}
	if err := json.Unmarshal([]byte(data), &jsonTree); err != nil {
		return err
	}
	return ytypes.Unmarshal(schema, destStruct, jsonTree, opts...)
}

// Device represents the /device YANG schema element.
type Device struct {
	Cont1A	*Test1_Cont1A	`path:"cont1a" module:"test1"`
	Nodes	*Xapp_Nodes	`path:"nodes" module:"xapp"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Device"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Device) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A represents the /test1/cont1a YANG schema element.
type Test1_Cont1A struct {
	Leaf1A	*string	`path:"leaf1a" module:"test1"`
	Leaf2A	*string	`path:"leaf2a" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Xapp_Nodes represents the /xapp/nodes YANG schema element.
type Xapp_Nodes struct {
	Node	map[string]*Xapp_Nodes_Node	`path:"node" module:"xapp"`
}

// IsYANGGoStruct ensures that Xapp_Nodes implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Xapp_Nodes) IsYANGGoStruct() {}

// NewNode creates a new entry in the Node list of the
// Xapp_Nodes struct. The keys of the list are populated from the input
// arguments.
func (t *Xapp_Nodes) NewNode(Id string) (*Xapp_Nodes_Node, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.Node == nil {
		t.Node = make(map[string]*Xapp_Nodes_Node)
	}

	key := Id

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.Node[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list Node", key)
	}

	t.Node[key] = &Xapp_Nodes_Node{
		Id: &Id,
	}

	return t.Node[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Xapp_Nodes) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Xapp_Nodes"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Xapp_Nodes) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Xapp_Nodes_Node represents the /xapp/nodes/node YANG schema element.
type Xapp_Nodes_Node struct {
	Id	*string	`path:"id" module:"xapp"`
	Ip	*string	`path:"ip" module:"xapp"`
	PlmnId	*string	`path:"plmn-id" module:"xapp"`
	Port	*uint16	`path:"port" module:"xapp"`
}

// IsYANGGoStruct ensures that Xapp_Nodes_Node implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Xapp_Nodes_Node) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the Xapp_Nodes_Node struct, which is a YANG list entry.
func (t *Xapp_Nodes_Node) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Id == nil {
		return nil, fmt.Errorf("nil value for key Id")
	}

	return map[string]interface{}{
		"id": *t.Id,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Xapp_Nodes_Node) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Xapp_Nodes_Node"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Xapp_Nodes_Node) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }



var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5a, 0x6b, 0x6f, 0xb3, 0x36,
		0x14, 0xfe, 0x9e, 0x5f, 0x81, 0x2c, 0xbd, 0x12, 0xf4, 0x85, 0x04, 0xd2, 0x90, 0x36, 0x7c, 0x89,
		0xb2, 0xb7, 0xab, 0x26, 0xf5, 0xb2, 0xa9, 0xab, 0xa6, 0x49, 0x84, 0x55, 0x5e, 0x70, 0x52, 0x6b,
		0x89, 0x41, 0xe0, 0x74, 0xad, 0x12, 0xfe, 0xfb, 0xc4, 0x25, 0x69, 0xb8, 0xdb, 0xac, 0xe9, 0x45,
		0xe2, 0x0b, 0xa5, 0x3e, 0x8f, 0xf1, 0xf1, 0x79, 0x1e, 0x7c, 0x4e, 0x0f, 0xdd, 0x74, 0x04, 0x41,
		0x10, 0xc0, 0x2d, 0x5c, 0x21, 0x60, 0x08, 0xc0, 0x46, 0x4f, 0x78, 0x86, 0x80, 0x1c, 0x8f, 0x5e,
		0x61, 0x62, 0x03, 0x43, 0xd0, 0x92, 0x5f, 0x7f, 0x38, 0x64, 0x8e, 0x17, 0xc0, 0x10, 0xd4, 0x64,
		0xe0, 0x02, 0x7b, 0xc0, 0x10, 0xe2, 0x47, 0x44, 0x03, 0x33, 0x87, 0x50, 0x0d, 0xa6, 0xc6, 0x52,
		0x8f, 0x4f, 0xec, 0x72, 0xda, 0x9a, 0x5e, 0x66, 0x3f, 0x9c, 0x5d, 0x6e, 0x6f, 0xf8, 0xcd, 0x43,
		0x73, 0xfc, 0x9c, 0x5b, 0x25, 0xb5, 0x12, 0xd5, 0x32, 0xab, 0x44, 0xd6, 0xdf, 0x9d, 0xb5, 0x37,
		0x43, 0x85, 0x33, 0x63, 0x4f, 0xd0, 0xcb, 0xbf, 0x8e, 0x17, 0x3a, 0x03, 0xdc, 0x78, 0x11, 0xb9,
		0x18, 0xf8, 0x0b, 0xf4, 0x27, 0xde, 0x62, 0xbd, 0x42, 0x84, 0x02, 0x43, 0xa0, 0xde, 0x1a, 0x95,
		0x00, 0x0f, 0x50, 0xa1, 0x4f, 0x39, 0x50, 0x90, 0x1a, 0x09, 0x32, 0x3b, 0xcd, 0x06, 0x78, 0x6f,
		0x58, 0x22, 0x38, 0x2f, 0x08, 0x74, 0x2e, 0x0c, 0x09, 0xae, 0xc4, 0xb9, 0x24, 0xf0, 0x6a, 0x89,
		0xb9, 0x8c, 0x00, 0x16, 0x22, 0xd8, 0x08, 0x61, 0x25, 0x86, 0x9b, 0x20, 0x6e, 0xa2, 0x98, 0x09,
		0x2b, 0x26, 0xae, 0x84, 0xc0, 0xfd, 0x53, 0xef, 0x5f, 0x5c, 0xc4, 0x16, 0x27, 0x9f, 0x7a, 0x98,
		0x2c, 0xaa, 0x62, 0xb5, 0x7b, 0x5d, 0xce, 0x2b, 0x30, 0xd7, 0x88, 0x2c, 0xe8, 0x23, 0x30, 0x04,
		0xb3, 0x72, 0xb7, 0xd5, 0xd1, 0x8e, 0x9e, 0x74, 0x83, 0x49, 0x2d, 0x2d, 0x8c, 0x82, 0xca, 0xc1,
		0xff, 0x80, 0xcb, 0x35, 0xca, 0xbf, 0xf9, 0xa5, 0xf8, 0x4b, 0x0f, 0xce, 0x28, 0x76, 0xc8, 0x05,
		0x5e, 0x60, 0xea, 0x87, 0x0b, 0xd5, 0xce, 0x0b, 0x64, 0x86, 0x2d, 0xc2, 0xe7, 0xa3, 0x6f, 0xf1,
		0x5c, 0x3d, 0xe2, 0x1e, 0x3b, 0xcd, 0xac, 0x56, 0x87, 0x0d, 0x5f, 0x10, 0xc3, 0xe8, 0x60, 0xe9,
		0x33, 0x1e, 0x40, 0xfd, 0xf6, 0x00, 0x6a, 0x0f, 0xa0, 0xf6, 0x00, 0xfa, 0xd8, 0x03, 0xa8, 0xaf,
		0xeb, 0x5f, 0xf8, 0x04, 0xaa, 0xac, 0x92, 0x26, 0x84, 0x38, 0x14, 0x86, 0x2e, 0x17, 0x17, 0x4b,
		0xfe, 0xec, 0x11, 0xad, 0xa0, 0x0b, 0x23, 0x3d, 0x82, 0x1e, 0x45, 0x3e, 0xd5, 0x7a, 0x85, 0xa5,
		0x68, 0x0c, 0xa7, 0xde, 0x7a, 0x46, 0x49, 0xf2, 0x3a, 0xdc, 0x87, 0xe8, 0x87, 0x1f, 0x21, 0x7a,
		0x92, 0x7e, 0x19, 0x5f, 0x9d, 0x3a, 0x70, 0x08, 0x10, 0xc7, 0x46, 0x7e, 0x79, 0x05, 0x1c, 0x9b,
		0xdf, 0xa1, 0x00, 0x7e, 0x86, 0xae, 0xfb, 0xf9, 0x4a, 0xe0, 0xc8, 0xab, 0xb7, 0x2a, 0x82, 0xc3,
		0x58, 0xd6, 0x67, 0xa0, 0x08, 0x55, 0x9d, 0x7f, 0xb4, 0xf7, 0xc8, 0x3f, 0x25, 0x84, 0x7c, 0xf2,
		0x0c, 0x54, 0x4c, 0x58, 0xb3, 0x1c, 0x54, 0x46, 0xe4, 0x1e, 0x80, 0xed, 0xfa, 0xed, 0xef, 0xc2,
		0x89, 0xed, 0xba, 0x7d, 0xb3, 0x1d, 0x94, 0xb5, 0x24, 0xf3, 0x90, 0xcd, 0x4b, 0x3a, 0x2f, 0xf9,
		0x8d, 0x45, 0xd0, 0x58, 0x0c, 0x0d, 0x44, 0xc1, 0x98, 0x14, 0x6a, 0xa2, 0x5d, 0x5b, 0xb0, 0xf0,
		0x17, 0x2e, 0xf9, 0x02, 0xa6, 0x61, 0xda, 0xaa, 0xf0, 0x1d, 0x60, 0x97, 0x43, 0xc3, 0x6e, 0xab,
		0xe1, 0x56, 0xc3, 0x69, 0x45, 0x28, 0xd0, 0xb6, 0x3d, 0xe4, 0xfb, 0x3c, 0x3a, 0x1e, 0x31, 0x60,
		0x13, 0x5f, 0x4c, 0xa6, 0x48, 0xb1, 0x31, 0x98, 0xf1, 0xfc, 0x69, 0xc0, 0xe1, 0x3b, 0xcf, 0x1f,
		0x13, 0x79, 0x15, 0x43, 0x4a, 0x91, 0x47, 0x98, 0xb7, 0xb3, 0x9f, 0x28, 0x8a, 0xa6, 0xaa, 0x8c,
		0xac, 0xad, 0xa9, 0x29, 0x23, 0x2b, 0xbe, 0xd5, 0xa2, 0x1f, 0xf1, 0x7d, 0xdf, 0x54, 0x95, 0xc1,
		0xee, 0x5e, 0x37, 0x55, 0x45, 0xb7, 0xa4, 0xe9, 0xb4, 0x2b, 0x6d, 0x4e, 0x03, 0xfe, 0x89, 0xe2,
		0x37, 0x73, 0x3a, 0x75, 0x37, 0xb7, 0x41, 0x78, 0xbd, 0x0e, 0xac, 0xef, 0xd2, 0x18, 0x30, 0x7b,
		0x6b, 0x31, 0x21, 0x03, 0xf9, 0x88, 0x6c, 0x0e, 0xbf, 0x00, 0x9b, 0xc6, 0x36, 0x8c, 0x39, 0x54,
		0xe6, 0x13, 0xe5, 0xd2, 0xda, 0xa8, 0xf2, 0x20, 0x90, 0x0c, 0x49, 0xcc, 0x8e, 0x19, 0xd2, 0x46,
		0x95, 0xf5, 0x40, 0x14, 0x0b, 0x2c, 0xe3, 0xa2, 0x67, 0x48, 0x5b, 0x51, 0x14, 0x13, 0x1e, 0x53,
		0xdc, 0x9a, 0xaa, 0x66, 0x8d, 0xa3, 0xdb, 0xf8, 0xba, 0x57, 0x07, 0x13, 0x58, 0x2a, 0xd4, 0x84,
		0xcc, 0x2d, 0xe1, 0xbf, 0x0c, 0xeb, 0xbb, 0x21, 0x6d, 0x86, 0xc1, 0xee, 0x3e, 0xba, 0x4a, 0x5b,
		0xb1, 0x7b, 0x32, 0x9d, 0x76, 0xbb, 0x27, 0x52, 0xbc, 0x81, 0x04, 0x77, 0x12, 0x5b, 0xc7, 0x86,
		0x91, 0x1b, 0x92, 0xc4, 0x6f, 0xdd, 0x63, 0xc8, 0xb2, 0xf3, 0xff, 0x9e, 0xd3, 0x2c, 0xf1, 0xba,
		0xcb, 0x15, 0x51, 0x78, 0x2a, 0xc8, 0xdd, 0x84, 0x36, 0x05, 0xb7, 0x29, 0xf8, 0xd3, 0x95, 0x91,
		0xae, 0xe3, 0x51, 0x0e, 0x29, 0x87, 0xe8, 0x56, 0xc7, 0xad, 0x8e, 0x33, 0x9a, 0x50, 0xc8, 0x7a,
		0xf5, 0x37, 0xf2, 0x38, 0xc4, 0x3c, 0x64, 0x80, 0xde, 0x41, 0xb2, 0x38, 0x4a, 0x2d, 0xc9, 0xd3,
		0xea, 0x6d, 0xd8, 0x0f, 0xcd, 0xf5, 0x45, 0x79, 0xe7, 0x35, 0xe8, 0x8e, 0x72, 0x96, 0x6b, 0xdc,
		0x2d, 0xe1, 0xb7, 0x0a, 0xc5, 0x50, 0xd7, 0x4f, 0xf5, 0x77, 0x0c, 0xc7, 0x47, 0x17, 0x13, 0x5c,
		0xcd, 0xab, 0x2b, 0xf4, 0x52, 0xd9, 0x76, 0x02, 0xd7, 0xd8, 0xa7, 0x13, 0x4a, 0x6b, 0x5a, 0x5c,
		0x37, 0x98, 0xfc, 0xbc, 0x44, 0xe1, 0xc9, 0xe1, 0x57, 0x33, 0x15, 0x2a, 0xe0, 0x00, 0xa9, 0x9d,
		0x0f, 0x06, 0xc3, 0xb3, 0xc1, 0x40, 0x3d, 0x3b, 0x3d, 0x53, 0x47, 0xba, 0xae, 0x0d, 0xb5, 0x0a,
		0xaa, 0xc0, 0xaf, 0x9e, 0x8d, 0x3c, 0x64, 0xff, 0x14, 0xfa, 0x4c, 0xd6, 0xcb, 0x25, 0xd7, 0x56,
		0x6b, 0xfa, 0xe9, 0x65, 0x7d, 0xf5, 0xf0, 0x18, 0xec, 0x45, 0xfd, 0xed, 0x5e, 0x45, 0xcf, 0xb5,
		0xa0, 0xc5, 0xfe, 0x27, 0x74, 0xdd, 0x87, 0xdb, 0x70, 0x62, 0x74, 0x05, 0x1f, 0xf1, 0x51, 0xe0,
		0xd5, 0xf9, 0xfa, 0x4f, 0x02, 0xaf, 0xfe, 0x96, 0x7e, 0x10, 0xe8, 0x1c, 0xb8, 0x54, 0xe6, 0x0a,
		0xc0, 0xfe, 0x25, 0xfc, 0x07, 0xdd, 0x39, 0x4e, 0x3e, 0xf7, 0x64, 0xdd, 0x03, 0x87, 0xa6, 0x94,
		0x33, 0x17, 0xf1, 0x7f, 0xec, 0xc4, 0x0b, 0x76, 0x82, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01,
		0x00, 0x00, 0xff, 0xff, 0x65, 0x0f, 0x50, 0xbf, 0xd0, 0x23, 0x00, 0x00,
	}
)


// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
var ΛEnumTypes = map[string][]reflect.Type{
}

