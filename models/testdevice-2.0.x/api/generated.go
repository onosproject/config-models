// Code generated by YGOT. DO NOTEDIT.
/*
Package api is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /go/pkg/mod/github.com/openconfig/ygot@v0.26.2/genutil/names.go
using the following YANG input files:
	- /config-model/yang/onf-test1@2019-06-10.yang
	- /config-model/yang/onf-test1-augmented@2020-02-29.yang
Imported modules were sourced from:
	- /var/model-compiler/yang-base/...
	- /config-model/yang/...
*/
package api

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
	ΛEnumTypes map[string][]reflect.Type
)

func init() {
	var err error
	initΛEnumTypes()
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
	Cont1A	*OnfTest1_Cont1A	`path:"cont1a" module:"onf-test1"`
	Cont1BState	*OnfTest1_Cont1BState	`path:"cont1b-state" module:"onf-test1"`
	LeafAtTopLevel	*string	`path:"leaf-at-top-level" module:"onf-test1"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Device"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Device) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Device.
func (*Device) ΛBelongingModule() string {
	return ""
}


// OnfTest1_Cont1A represents the /onf-test1/cont1a YANG schema element.
type OnfTest1_Cont1A struct {
	Cont2A	*OnfTest1_Cont1A_Cont2A	`path:"cont2a" module:"onf-test1"`
	Cont2D	*OnfTest1_Cont1A_Cont2D	`path:"cont2d" module:"onf-test1-augmented"`
	Leaf1A	*string	`path:"leaf1a" module:"onf-test1"`
	List2A	map[string]*OnfTest1_Cont1A_List2A	`path:"list2a" module:"onf-test1"`
}

// IsYANGGoStruct ensures that OnfTest1_Cont1A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OnfTest1_Cont1A) IsYANGGoStruct() {}

// NewList2A creates a new entry in the List2A list of the
// OnfTest1_Cont1A struct. The keys of the list are populated from the input
// arguments.
func (t *OnfTest1_Cont1A) NewList2A(Name string) (*OnfTest1_Cont1A_List2A, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2A == nil {
		t.List2A = make(map[string]*OnfTest1_Cont1A_List2A)
	}

	key := Name

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2A[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2A", key)
	}

	t.List2A[key] = &OnfTest1_Cont1A_List2A{
		Name: &Name,
	}

	return t.List2A[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OnfTest1_Cont1A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OnfTest1_Cont1A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OnfTest1_Cont1A.
func (*OnfTest1_Cont1A) ΛBelongingModule() string {
	return "onf-test1"
}


// OnfTest1_Cont1A_Cont2A represents the /onf-test1/cont1a/cont2a YANG schema element.
type OnfTest1_Cont1A_Cont2A struct {
	Leaf2A	*uint8	`path:"leaf2a" module:"onf-test1"`
	Leaf2B	*float64	`path:"leaf2b" module:"onf-test1"`
	Leaf2C	*string	`path:"leaf2c" module:"onf-test1"`
	Leaf2D	*float64	`path:"leaf2d" module:"onf-test1"`
	Leaf2E	[]int16	`path:"leaf2e" module:"onf-test1"`
	Leaf2F	Binary	`path:"leaf2f" module:"onf-test1"`
	Leaf2G	*bool	`path:"leaf2g" module:"onf-test1"`
}

// IsYANGGoStruct ensures that OnfTest1_Cont1A_Cont2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OnfTest1_Cont1A_Cont2A) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A_Cont2A) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OnfTest1_Cont1A_Cont2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A_Cont2A) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OnfTest1_Cont1A_Cont2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OnfTest1_Cont1A_Cont2A.
func (*OnfTest1_Cont1A_Cont2A) ΛBelongingModule() string {
	return "onf-test1"
}


// OnfTest1_Cont1A_Cont2D represents the /onf-test1/cont1a/cont2d YANG schema element.
type OnfTest1_Cont1A_Cont2D struct {
	Beer	YANGEmpty	`path:"beer" module:"onf-test1-augmented"`
	Chocolate	E_OnfTest1_Cont1A_Cont2D_Chocolate	`path:"chocolate" module:"onf-test1-augmented"`
	Leaf2D3C	*string	`path:"leaf2d3c" module:"onf-test1-augmented"`
	Pretzel	YANGEmpty	`path:"pretzel" module:"onf-test1-augmented"`
}

// IsYANGGoStruct ensures that OnfTest1_Cont1A_Cont2D implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OnfTest1_Cont1A_Cont2D) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A_Cont2D) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OnfTest1_Cont1A_Cont2D"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A_Cont2D) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OnfTest1_Cont1A_Cont2D) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OnfTest1_Cont1A_Cont2D.
func (*OnfTest1_Cont1A_Cont2D) ΛBelongingModule() string {
	return "onf-test1-augmented"
}


// OnfTest1_Cont1A_List2A represents the /onf-test1/cont1a/list2a YANG schema element.
type OnfTest1_Cont1A_List2A struct {
	Name	*string	`path:"name" module:"onf-test1"`
	RxPower	*uint16	`path:"rx-power" module:"onf-test1"`
	TxPower	*uint16	`path:"tx-power" module:"onf-test1"`
}

// IsYANGGoStruct ensures that OnfTest1_Cont1A_List2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OnfTest1_Cont1A_List2A) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the OnfTest1_Cont1A_List2A struct, which is a YANG list entry.
func (t *OnfTest1_Cont1A_List2A) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A_List2A) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OnfTest1_Cont1A_List2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1A_List2A) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OnfTest1_Cont1A_List2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OnfTest1_Cont1A_List2A.
func (*OnfTest1_Cont1A_List2A) ΛBelongingModule() string {
	return "onf-test1"
}


// OnfTest1_Cont1BState represents the /onf-test1/cont1b-state YANG schema element.
type OnfTest1_Cont1BState struct {
	Cont2C	*OnfTest1_Cont1BState_Cont2C	`path:"cont2c" module:"onf-test1"`
	Leaf2D	*uint16	`path:"leaf2d" module:"onf-test1"`
	List2B	map[OnfTest1_Cont1BState_List2B_Key]*OnfTest1_Cont1BState_List2B	`path:"list2b" module:"onf-test1"`
}

// IsYANGGoStruct ensures that OnfTest1_Cont1BState implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OnfTest1_Cont1BState) IsYANGGoStruct() {}

// OnfTest1_Cont1BState_List2B_Key represents the key for list List2B of element /onf-test1/cont1b-state.
type OnfTest1_Cont1BState_List2B_Key struct {
	Index1	uint8	`path:"index1"`
	Index2	uint8	`path:"index2"`
}

// IsYANGGoKeyStruct ensures that OnfTest1_Cont1BState_List2B_Key partially implements the
// yang.GoKeyStruct interface. This allows functions that need to
// handle this key struct to identify it as being generated by gogen.
func (OnfTest1_Cont1BState_List2B_Key) IsYANGGoKeyStruct() {}

// ΛListKeyMap returns the values of the OnfTest1_Cont1BState_List2B_Key key struct.
func (t OnfTest1_Cont1BState_List2B_Key) ΛListKeyMap() (map[string]interface{}, error) {
	return map[string]interface{}{
		"index1": t.Index1,
		"index2": t.Index2,
	}, nil
}

// NewList2B creates a new entry in the List2B list of the
// OnfTest1_Cont1BState struct. The keys of the list are populated from the input
// arguments.
func (t *OnfTest1_Cont1BState) NewList2B(Index1 uint8, Index2 uint8) (*OnfTest1_Cont1BState_List2B, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2B == nil {
		t.List2B = make(map[OnfTest1_Cont1BState_List2B_Key]*OnfTest1_Cont1BState_List2B)
	}

	key := OnfTest1_Cont1BState_List2B_Key{
		Index1: Index1,
		Index2: Index2,
	}

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2B[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2B", key)
	}

	t.List2B[key] = &OnfTest1_Cont1BState_List2B{
		Index1: &Index1,
		Index2: &Index2,
	}

	return t.List2B[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1BState) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OnfTest1_Cont1BState"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1BState) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OnfTest1_Cont1BState) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OnfTest1_Cont1BState.
func (*OnfTest1_Cont1BState) ΛBelongingModule() string {
	return "onf-test1"
}


// OnfTest1_Cont1BState_Cont2C represents the /onf-test1/cont1b-state/cont2c YANG schema element.
type OnfTest1_Cont1BState_Cont2C struct {
	Leaf3A	*bool	`path:"leaf3a" module:"onf-test1"`
	Leaf3B	*string	`path:"leaf3b" module:"onf-test1"`
}

// IsYANGGoStruct ensures that OnfTest1_Cont1BState_Cont2C implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OnfTest1_Cont1BState_Cont2C) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1BState_Cont2C) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OnfTest1_Cont1BState_Cont2C"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1BState_Cont2C) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OnfTest1_Cont1BState_Cont2C) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OnfTest1_Cont1BState_Cont2C.
func (*OnfTest1_Cont1BState_Cont2C) ΛBelongingModule() string {
	return "onf-test1"
}


// OnfTest1_Cont1BState_List2B represents the /onf-test1/cont1b-state/list2b YANG schema element.
type OnfTest1_Cont1BState_List2B struct {
	Index1	*uint8	`path:"index1" module:"onf-test1"`
	Index2	*uint8	`path:"index2" module:"onf-test1"`
	Leaf3C	*string	`path:"leaf3c" module:"onf-test1"`
	Leaf3D	E_OnfTest1Identities_MYBASE	`path:"leaf3d" module:"onf-test1"`
}

// IsYANGGoStruct ensures that OnfTest1_Cont1BState_List2B implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OnfTest1_Cont1BState_List2B) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the OnfTest1_Cont1BState_List2B struct, which is a YANG list entry.
func (t *OnfTest1_Cont1BState_List2B) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Index1 == nil {
		return nil, fmt.Errorf("nil value for key Index1")
	}

	if t.Index2 == nil {
		return nil, fmt.Errorf("nil value for key Index2")
	}

	return map[string]interface{}{
		"index1": *t.Index1,
		"index2": *t.Index2,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1BState_List2B) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OnfTest1_Cont1BState_List2B"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OnfTest1_Cont1BState_List2B) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OnfTest1_Cont1BState_List2B) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OnfTest1_Cont1BState_List2B.
func (*OnfTest1_Cont1BState_List2B) ΛBelongingModule() string {
	return "onf-test1"
}


// E_OnfTest1Identities_MYBASE is a derived int64 type which is used to represent
// the enumerated node OnfTest1Identities_MYBASE. An additional value named
// OnfTest1Identities_MYBASE_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_OnfTest1Identities_MYBASE int64

// IsYANGGoEnum ensures that OnfTest1Identities_MYBASE implements the yang.GoEnum
// interface. This ensures that OnfTest1Identities_MYBASE can be identified as a
// mapped type for a YANG enumeration.
func (E_OnfTest1Identities_MYBASE) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  OnfTest1Identities_MYBASE.
func (E_OnfTest1Identities_MYBASE) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum; }

// String returns a logging-friendly string for E_OnfTest1Identities_MYBASE.
func (e E_OnfTest1Identities_MYBASE) String() string {
	return ygot.EnumLogString(e, int64(e), "E_OnfTest1Identities_MYBASE")
}

const (
	// OnfTest1Identities_MYBASE_UNSET corresponds to the value UNSET of OnfTest1Identities_MYBASE
	OnfTest1Identities_MYBASE_UNSET E_OnfTest1Identities_MYBASE = 0
	// OnfTest1Identities_MYBASE_IDTYPE1 corresponds to the value IDTYPE1 of OnfTest1Identities_MYBASE
	OnfTest1Identities_MYBASE_IDTYPE1 E_OnfTest1Identities_MYBASE = 1
	// OnfTest1Identities_MYBASE_IDTYPE2 corresponds to the value IDTYPE2 of OnfTest1Identities_MYBASE
	OnfTest1Identities_MYBASE_IDTYPE2 E_OnfTest1Identities_MYBASE = 2
)


// E_OnfTest1_Cont1A_Cont2D_Chocolate is a derived int64 type which is used to represent
// the enumerated node OnfTest1_Cont1A_Cont2D_Chocolate. An additional value named
// OnfTest1_Cont1A_Cont2D_Chocolate_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_OnfTest1_Cont1A_Cont2D_Chocolate int64

// IsYANGGoEnum ensures that OnfTest1_Cont1A_Cont2D_Chocolate implements the yang.GoEnum
// interface. This ensures that OnfTest1_Cont1A_Cont2D_Chocolate can be identified as a
// mapped type for a YANG enumeration.
func (E_OnfTest1_Cont1A_Cont2D_Chocolate) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  OnfTest1_Cont1A_Cont2D_Chocolate.
func (E_OnfTest1_Cont1A_Cont2D_Chocolate) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum; }

// String returns a logging-friendly string for E_OnfTest1_Cont1A_Cont2D_Chocolate.
func (e E_OnfTest1_Cont1A_Cont2D_Chocolate) String() string {
	return ygot.EnumLogString(e, int64(e), "E_OnfTest1_Cont1A_Cont2D_Chocolate")
}

const (
	// OnfTest1_Cont1A_Cont2D_Chocolate_UNSET corresponds to the value UNSET of OnfTest1_Cont1A_Cont2D_Chocolate
	OnfTest1_Cont1A_Cont2D_Chocolate_UNSET E_OnfTest1_Cont1A_Cont2D_Chocolate = 0
	// OnfTest1_Cont1A_Cont2D_Chocolate_dark corresponds to the value dark of OnfTest1_Cont1A_Cont2D_Chocolate
	OnfTest1_Cont1A_Cont2D_Chocolate_dark E_OnfTest1_Cont1A_Cont2D_Chocolate = 1
	// OnfTest1_Cont1A_Cont2D_Chocolate_milk corresponds to the value milk of OnfTest1_Cont1A_Cont2D_Chocolate
	OnfTest1_Cont1A_Cont2D_Chocolate_milk E_OnfTest1_Cont1A_Cont2D_Chocolate = 2
	// OnfTest1_Cont1A_Cont2D_Chocolate_first_available corresponds to the value first_available of OnfTest1_Cont1A_Cont2D_Chocolate
	OnfTest1_Cont1A_Cont2D_Chocolate_first_available E_OnfTest1_Cont1A_Cont2D_Chocolate = 3
)


// ΛEnum is a map, keyed by the name of the type defined for each enum in the
// generated Go code, which provides a mapping between the constant int64 value
// of each value of the enumeration, and the string that is used to represent it
// in the YANG schema. The map is named ΛEnum in order to avoid clash with any
// valid YANG identifier.
var ΛEnum = map[string]map[int64]ygot.EnumDefinition{
	"E_OnfTest1Identities_MYBASE": {
		1: {Name: "IDTYPE1", DefiningModule: "onf-test1-identities"},
		2: {Name: "IDTYPE2", DefiningModule: "onf-test1-identities"},
	},
	"E_OnfTest1_Cont1A_Cont2D_Chocolate": {
		1: {Name: "dark"},
		2: {Name: "milk"},
		3: {Name: "first-available"},
	},
}


var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5d, 0x5f, 0x6f, 0xdb, 0x38,
		0x12, 0x7f, 0xef, 0xa7, 0x18, 0xf8, 0x25, 0xdb, 0x43, 0xdc, 0xc8, 0x72, 0x92, 0xa6, 0x01, 0xee,
		0xc1, 0x69, 0x52, 0x5c, 0xb1, 0xed, 0x6e, 0xd1, 0x06, 0x7b, 0xd8, 0xeb, 0x05, 0x07, 0x5a, 0x1a,
		0xdb, 0x44, 0x65, 0xd2, 0x47, 0x52, 0x6e, 0x7c, 0x8b, 0x7c, 0xf7, 0x03, 0x25, 0xd9, 0xf1, 0x5f,
		0x89, 0x94, 0x64, 0xc7, 0x4e, 0xb8, 0x2f, 0xdb, 0xd8, 0x24, 0x4d, 0x0e, 0x67, 0x7e, 0x33, 0x1c,
		0xce, 0x0c, 0xff, 0x7a, 0x05, 0x00, 0xd0, 0xf8, 0x8d, 0x0c, 0xb1, 0x71, 0x09, 0x8d, 0x10, 0xc7,
		0x34, 0xc0, 0xc6, 0x71, 0xfa, 0xe9, 0xaf, 0x94, 0x85, 0x8d, 0x4b, 0x68, 0x65, 0x7f, 0xbe, 0xe7,
		0xac, 0x47, 0xfb, 0x8d, 0x4b, 0xf0, 0xb2, 0x0f, 0xae, 0xa9, 0x68, 0x5c, 0x42, 0x3a, 0x44, 0xf2,
		0x41, 0xc0, 0x99, 0x6a, 0x91, 0x85, 0xcf, 0x16, 0x86, 0xcf, 0xbe, 0x3f, 0x5e, 0xfc, 0xf6, 0x1a,
		0x65, 0x20, 0xe8, 0x48, 0x51, 0xce, 0x74, 0xa3, 0xdb, 0x01, 0x82, 0xe2, 0x23, 0x88, 0x70, 0x8c,
		0x11, 0xe8, 0x2e, 0x84, 0x32, 0x14, 0xcb, 0xbd, 0x16, 0x27, 0x37, 0xfb, 0x78, 0x79, 0x92, 0xb3,
		0x2f, 0xbe, 0x08, 0xec, 0xd1, 0xfb, 0x95, 0xb9, 0x2d, 0xcc, 0x4f, 0xb5, 0x96, 0x7e, 0x25, 0xf9,
		0xf6, 0x1b, 0x8f, 0x45, 0x80, 0x6b, 0x7b, 0xa6, 0x33, 0xc1, 0xc9, 0x4f, 0x2e, 0xf4, 0x64, 0x1a,
		0xa3, 0xf4, 0x47, 0x8e, 0xd7, 0x37, 0xfc, 0x07, 0x91, 0x1d, 0xd1, 0x8f, 0x87, 0xc8, 0x54, 0xe3,
		0x12, 0x94, 0x88, 0x71, 0x43, 0xc3, 0xb9, 0x56, 0x7a, 0x4e, 0x2b, 0x8d, 0x1e, 0x16, 0x3e, 0x79,
		0x58, 0xa6, 0xe7, 0xd2, 0xb6, 0x2c, 0x6c, 0x8f, 0x4f, 0x36, 0x2f, 0x64, 0x7e, 0x9b, 0x7c, 0xb2,
		0x69, 0x15, 0x6b, 0xb6, 0xcb, 0x67, 0x61, 0xc1, 0x76, 0x15, 0x6c, 0x5b, 0xe1, 0xf6, 0x99, 0x6c,
		0xa3, 0xd9, 0x76, 0x9a, 0x6e, 0xab, 0xf5, 0xf6, 0x5a, 0x6f, 0xb3, 0xf1, 0x76, 0xaf, 0xdf, 0xf6,
		0x0d, 0xdb, 0x5f, 0xc8, 0x06, 0xb3, 0x06, 0x11, 0x92, 0x5e, 0x0e, 0x3b, 0xac, 0x90, 0x33, 0x6b,
		0x5f, 0xb0, 0x98, 0x25, 0xf6, 0xf8, 0x2d, 0x1e, 0xa2, 0xa0, 0x01, 0xe8, 0xce, 0x40, 0x99, 0xa4,
		0x21, 0xc2, 0xfb, 0x29, 0x93, 0x80, 0xc9, 0x70, 0x3d, 0x12, 0x47, 0x9a, 0x34, 0xdf, 0x73, 0x1b,
		0x26, 0x8d, 0xfd, 0x46, 0x6e, 0x9b, 0xbb, 0x82, 0xdf, 0xca, 0x78, 0xd3, 0x2b, 0x68, 0x56, 0xc4,
		0xa3, 0x36, 0xbc, 0x6a, 0xc7, 0xb3, 0xb6, 0xbc, 0x5b, 0x9a, 0x87, 0x4b, 0xf3, 0xb2, 0x35, 0x4f,
		0xe7, 0xf3, 0x76, 0x01, 0x8f, 0xcf, 0x7e, 0xed, 0x76, 0x32, 0x42, 0x3b, 0x3a, 0xc7, 0x94, 0xa9,
		0x0b, 0x13, 0x52, 0x67, 0x4c, 0x71, 0x66, 0xd0, 0xf4, 0x2b, 0x61, 0x7d, 0x34, 0xe2, 0x54, 0xfd,
		0x9f, 0xd9, 0xd6, 0x25, 0x03, 0x7f, 0xa6, 0xcc, 0x78, 0xaf, 0x67, 0x9d, 0xfe, 0x20, 0x51, 0x8c,
		0x9b, 0xa1, 0x76, 0x63, 0xbf, 0x0f, 0x82, 0x04, 0x5a, 0x7a, 0xaf, 0x69, 0x9f, 0x2a, 0x59, 0xcc,
		0xe6, 0xab, 0x24, 0xc6, 0x3e, 0x51, 0x74, 0xac, 0x7f, 0xbb, 0x47, 0x22, 0x89, 0xc6, 0xbd, 0x1f,
		0x8e, 0x2d, 0x48, 0x42, 0xee, 0xcb, 0x93, 0xa4, 0x7d, 0x38, 0x24, 0x79, 0x55, 0x23, 0xe1, 0x76,
		0xc6, 0x71, 0x8e, 0xe5, 0x56, 0x69, 0xf2, 0xec, 0x78, 0xae, 0xb0, 0xd5, 0x5d, 0x25, 0x48, 0xc7,
		0x7b, 0x25, 0x48, 0x33, 0x66, 0x52, 0x91, 0x6e, 0x64, 0x08, 0xee, 0x02, 0x7b, 0x28, 0x90, 0x05,
		0x5b, 0x01, 0xe1, 0xa9, 0xe6, 0xf8, 0xfa, 0xe1, 0x3d, 0x9c, 0x7b, 0xa7, 0x5e, 0xc3, 0x82, 0x75,
		0x2c, 0xf5, 0xf5, 0x3a, 0xbd, 0xfd, 0xb8, 0x36, 0x4b, 0x3e, 0x28, 0xab, 0xc2, 0xd7, 0xaa, 0xf2,
		0xd9, 0xe2, 0xf7, 0x8d, 0x9b, 0x5e, 0x95, 0xe0, 0xb3, 0xd4, 0xa2, 0xed, 0x5a, 0x5a, 0xc0, 0x5d,
		0x4b, 0x0b, 0xf8, 0x0f, 0x1e, 0x29, 0xd2, 0xc7, 0xd2, 0x16, 0xb0, 0xb3, 0x4a, 0x0f, 0xd5, 0x2a,
		0xfd, 0x4c, 0x58, 0x48, 0x14, 0x17, 0x93, 0x62, 0x2b, 0xac, 0x84, 0x05, 0x1b, 0x62, 0x40, 0x87,
		0x24, 0x3a, 0x3f, 0xb5, 0xb0, 0x62, 0x5b, 0xbe, 0x41, 0xdb, 0x15, 0xcd, 0xd3, 0x7e, 0xb9, 0xb6,
		0x6f, 0xbb, 0xbc, 0xd2, 0xd5, 0x6c, 0xb9, 0x7f, 0x76, 0x88, 0xef, 0x79, 0xde, 0x13, 0x12, 0x65,
		0x2f, 0x2d, 0x91, 0xf2, 0xba, 0x23, 0xb0, 0xd4, 0x1d, 0x81, 0xa5, 0xee, 0xf8, 0x8a, 0x24, 0x04,
		0xce, 0xa2, 0xc9, 0xce, 0xb4, 0x87, 0xef, 0xb4, 0xc7, 0xc1, 0xfa, 0x34, 0xa4, 0x12, 0x94, 0xf5,
		0x6d, 0xd4, 0xc1, 0xc5, 0xb6, 0x04, 0x23, 0xb4, 0x14, 0x8c, 0xd0, 0x52, 0x30, 0x3a, 0x8c, 0xab,
		0x01, 0x0a, 0xc8, 0x94, 0xa0, 0xb3, 0xab, 0x9c, 0x64, 0x38, 0x5b, 0xe9, 0x70, 0x6d, 0xa5, 0x3d,
		0x75, 0xda, 0x38, 0x63, 0xa9, 0x4e, 0x9d, 0x80, 0x96, 0x3a, 0x01, 0x2d, 0x75, 0x42, 0x62, 0x22,
		0x45, 0x54, 0x2a, 0xa7, 0x0d, 0x9c, 0x36, 0xc8, 0xa3, 0x33, 0x65, 0xaa, 0x75, 0x6e, 0xa1, 0x09,
		0xfc, 0xc3, 0xc5, 0xf4, 0xea, 0xf8, 0xe5, 0x3d, 0xc3, 0x13, 0xb0, 0x73, 0xc5, 0x5b, 0x49, 0xd8,
		0x27, 0x2a, 0x55, 0x47, 0x29, 0x61, 0x26, 0x65, 0x9f, 0x29, 0xbb, 0x89, 0x50, 0xcb, 0xbf, 0x21,
		0xa9, 0xf4, 0x76, 0xce, 0xf5, 0x68, 0x5d, 0x9c, 0x9e, 0x9e, 0xbf, 0x3d, 0x3d, 0xf5, 0xde, 0xb6,
		0xdf, 0x7a, 0xef, 0xce, 0xce, 0x5a, 0xe7, 0x2d, 0x93, 0xcb, 0xd7, 0xdf, 0x45, 0x88, 0x02, 0xc3,
		0xab, 0x49, 0xe3, 0x12, 0x58, 0x1c, 0x45, 0xdb, 0xd2, 0x62, 0x3d, 0x4b, 0x2d, 0xd6, 0xb3, 0xd4,
		0x62, 0x5d, 0xca, 0x88, 0x98, 0x38, 0x6f, 0xb1, 0xd3, 0x63, 0x85, 0x74, 0x4e, 0x59, 0xc5, 0x42,
		0x91, 0xbd, 0x33, 0x68, 0xfa, 0x09, 0x59, 0x5f, 0x0d, 0xf6, 0x4e, 0x93, 0xf9, 0x9e, 0xbb, 0x53,
		0x3e, 0x64, 0x9a, 0xec, 0xfb, 0xe1, 0xa4, 0x6f, 0x09, 0xeb, 0x7d, 0x4b, 0x58, 0xbf, 0xe2, 0x3c,
		0x42, 0xc2, 0x1c, 0xae, 0x3b, 0x5c, 0x2f, 0xc6, 0xf5, 0x94, 0x57, 0x6c, 0x7c, 0x55, 0xad, 0xb2,
		0x72, 0x61, 0x15, 0x52, 0x6a, 0x11, 0x93, 0xd1, 0x18, 0xc6, 0xb2, 0x38, 0x6a, 0xd3, 0xda, 0xb9,
		0xfd, 0x4b, 0x2a, 0x7c, 0xaf, 0xe1, 0xef, 0x70, 0xa4, 0x77, 0xfc, 0x08, 0xb8, 0x00, 0x16, 0x0f,
		0xbb, 0x28, 0x7e, 0x79, 0x73, 0x92, 0x46, 0xa8, 0xbe, 0x86, 0x7f, 0xc7, 0x9e, 0xd7, 0x0e, 0xc0,
		0xc8, 0xdb, 0x77, 0x23, 0x04, 0x17, 0x9f, 0x51, 0x4a, 0xd2, 0xb7, 0x60, 0xda, 0xe9, 0xac, 0x3e,
		0xf6, 0xe0, 0x93, 0x96, 0x68, 0x9f, 0x00, 0x95, 0x40, 0xba, 0x7c, 0x8c, 0x70, 0x0a, 0x6a, 0x80,
		0xa9, 0xa4, 0xfb, 0x7d, 0xd0, 0x74, 0x80, 0x2e, 0x26, 0xfc, 0x09, 0x3d, 0x2e, 0xf4, 0x97, 0x30,
		0x26, 0x11, 0x0d, 0x89, 0xc6, 0x06, 0x50, 0x1c, 0x46, 0x44, 0x4a, 0x53, 0xa6, 0x2f, 0x11, 0x49,
		0x32, 0x2f, 0x61, 0xa8, 0x57, 0xdb, 0x1c, 0x66, 0xcb, 0xb5, 0x50, 0x54, 0x55, 0xa2, 0x48, 0x16,
		0x04, 0xae, 0x3e, 0x8a, 0xd5, 0xa4, 0x52, 0x1e, 0x4a, 0x8a, 0xce, 0x9d, 0x95, 0xe8, 0x74, 0x18,
		0xe3, 0x8a, 0x64, 0xda, 0x20, 0x47, 0x6c, 0x64, 0x30, 0xc0, 0x21, 0x19, 0x91, 0xc4, 0x04, 0x6b,
		0x9c, 0x70, 0xd6, 0x6b, 0x2a, 0x94, 0xaa, 0x75, 0x92, 0xe6, 0x4e, 0x9c, 0xe4, 0xc6, 0xe6, 0xa7,
		0x23, 0x28, 0x11, 0x07, 0x8a, 0x65, 0x0c, 0xfa, 0x3b, 0xeb, 0xdd, 0xea, 0xfe, 0xff, 0xd1, 0xca,
		0xa6, 0xd5, 0x49, 0xfe, 0xe7, 0x77, 0xd6, 0xd3, 0x6e, 0x75, 0xad, 0x6b, 0x56, 0x93, 0x66, 0x07,
		0x84, 0x86, 0x59, 0x04, 0xa1, 0x61, 0x16, 0xc1, 0x9c, 0x2a, 0x0c, 0xf7, 0x23, 0x75, 0x80, 0x1c,
		0x62, 0xee, 0x00, 0xd9, 0x6d, 0xf2, 0x40, 0xd8, 0xb6, 0xbd, 0x00, 0xd7, 0x3d, 0xec, 0x0c, 0xa7,
		0x14, 0xf7, 0x53, 0xbb, 0x89, 0xb3, 0x79, 0x9b, 0x29, 0x04, 0x12, 0xf7, 0xf5, 0xda, 0x31, 0xd4,
		0x98, 0xb0, 0x36, 0xb9, 0x68, 0x4f, 0xed, 0x28, 0xf2, 0x1c, 0x0d, 0x29, 0xe2, 0x6e, 0xc4, 0x37,
		0x20, 0x32, 0x23, 0xc1, 0x0f, 0x73, 0x41, 0x49, 0x9b, 0x5b, 0x4b, 0x09, 0xa1, 0xfd, 0x81, 0x82,
		0x9e, 0xe0, 0x43, 0xf8, 0xfa, 0xe1, 0x7d, 0xf3, 0xdc, 0xf3, 0x3d, 0x43, 0x61, 0x38, 0x73, 0xc2,
		0x70, 0x78, 0xc2, 0x50, 0x84, 0xd0, 0x8f, 0x48, 0x4d, 0x14, 0x36, 0x99, 0x66, 0x0e, 0x7b, 0xf3,
		0x76, 0xae, 0xaf, 0x21, 0x15, 0x32, 0x96, 0x3a, 0x35, 0x6c, 0x6e, 0xca, 0x5a, 0x65, 0x58, 0xac,
		0x24, 0xab, 0x55, 0x31, 0xb5, 0x2b, 0xb1, 0x5e, 0x2d, 0xb6, 0x76, 0x39, 0x56, 0xb4, 0x74, 0xd3,
		0x18, 0xee, 0x95, 0x29, 0x8b, 0x3e, 0x9a, 0x96, 0x03, 0x1e, 0x70, 0xcd, 0x73, 0xf6, 0x44, 0x9f,
		0x59, 0x9d, 0xb3, 0x21, 0x2c, 0x69, 0xb6, 0x04, 0xa8, 0xc8, 0xe2, 0x21, 0x0a, 0xa2, 0x6d, 0x8b,
		0xcc, 0x65, 0x03, 0x01, 0x91, 0x08, 0x8f, 0x12, 0x51, 0x83, 0x35, 0x52, 0xd2, 0x3a, 0xa9, 0x2c,
		0x45, 0x55, 0xa4, 0xa9, 0xa2, 0x54, 0x55, 0x95, 0xae, 0xda, 0xa4, 0xac, 0x36, 0x69, 0xab, 0x2e,
		0x75, 0x76, 0xd2, 0x57, 0xc2, 0xcd, 0x6c, 0x67, 0x3d, 0x6d, 0xdc, 0xe9, 0xa9, 0x4c, 0x68, 0x11,
		0x29, 0xb1, 0xe3, 0x53, 0x13, 0xeb, 0xb4, 0x44, 0xdf, 0x1b, 0x16, 0x0f, 0xcb, 0xf3, 0xca, 0x2d,
		0xff, 0x96, 0x1a, 0x82, 0x65, 0x47, 0x48, 0x46, 0xf1, 0x92, 0x48, 0x32, 0x22, 0x7e, 0x94, 0xe4,
		0xb4, 0x64, 0x90, 0x96, 0x1e, 0x64, 0x48, 0xa3, 0x4a, 0x83, 0xf8, 0x7a, 0x90, 0x1e, 0x15, 0x52,
		0x35, 0xc9, 0x98, 0xd0, 0x28, 0xf1, 0xc6, 0x95, 0x1a, 0xee, 0xe1, 0xb8, 0x2c, 0x45, 0x3f, 0x32,
		0x55, 0x8d, 0x9c, 0x09, 0x25, 0xad, 0x21, 0x6b, 0x61, 0x88, 0x65, 0x12, 0x18, 0x85, 0x6a, 0x6c,
		0xf6, 0x58, 0xea, 0x5d, 0xb9, 0x84, 0x56, 0x39, 0x42, 0x6e, 0x5b, 0xda, 0x9f, 0x48, 0x87, 0x1b,
		0xfa, 0xad, 0xca, 0xf9, 0xb1, 0xc2, 0x93, 0xe4, 0xac, 0x73, 0x32, 0x67, 0x66, 0xd6, 0xe5, 0xd4,
		0x33, 0x38, 0x52, 0xc8, 0x11, 0x17, 0x4a, 0x36, 0x89, 0x40, 0x46, 0xec, 0x4d, 0xe3, 0x85, 0xde,
		0xce, 0x38, 0x76, 0xc6, 0xf1, 0xb6, 0x8d, 0xe3, 0x2e, 0xa2, 0x28, 0x6f, 0x17, 0x27, 0xbd, 0xab,
		0x99, 0xc4, 0x3e, 0x0b, 0x01, 0x87, 0x23, 0x35, 0x59, 0xb4, 0x88, 0xe7, 0x05, 0xc1, 0xd9, 0xc4,
		0xce, 0x26, 0x7e, 0x61, 0x36, 0xb1, 0x16, 0x88, 0x2a, 0xd6, 0x70, 0x7b, 0x5b, 0x66, 0x80, 0x05,
		0x98, 0x8f, 0x04, 0xaa, 0xff, 0x61, 0x54, 0x1e, 0x5d, 0xa6, 0x03, 0x54, 0x3c, 0x73, 0x3b, 0x70,
		0x71, 0xe0, 0xe2, 0xc0, 0xe5, 0x10, 0xc0, 0xe5, 0xf0, 0xce, 0x18, 0x0b, 0xf6, 0xfa, 0x6e, 0x42,
		0x07, 0x8a, 0xae, 0x6e, 0xed, 0x96, 0x6d, 0xb5, 0xdc, 0xc6, 0x4e, 0xe2, 0x81, 0x6a, 0x0d, 0x6a,
		0x08, 0x2b, 0x06, 0x35, 0x5c, 0x57, 0x09, 0x6a, 0xd0, 0x4a, 0xa7, 0x65, 0x50, 0x1a, 0x31, 0x6b,
		0x67, 0x16, 0xd4, 0xf0, 0x69, 0x6d, 0xac, 0xdf, 0xe6, 0xee, 0xf9, 0xca, 0xc8, 0xd5, 0x46, 0xac,
		0xaf, 0x36, 0x62, 0x21, 0x34, 0x9b, 0xdf, 0x20, 0x3f, 0xde, 0x1c, 0xe7, 0xb4, 0x31, 0x0c, 0xaa,
		0x36, 0xcb, 0xae, 0x30, 0xf7, 0x55, 0x4c, 0x03, 0x85, 0xcf, 0x0c, 0x81, 0xb8, 0x6c, 0x7c, 0xb0,
		0x7d, 0x5c, 0xf0, 0x83, 0x59, 0x5a, 0x88, 0xfd, 0x52, 0x5b, 0xde, 0xfe, 0xad, 0xb5, 0xde, 0x00,
		0x33, 0x33, 0x3c, 0xa3, 0xd2, 0xa8, 0xd4, 0x6b, 0xd6, 0xce, 0x0c, 0xcf, 0x3a, 0x20, 0xe9, 0x70,
		0x14, 0x61, 0x9a, 0x64, 0xc9, 0x7b, 0xda, 0xe6, 0xee, 0xd1, 0x7e, 0x9c, 0x5e, 0x09, 0x00, 0x55,
		0x38, 0x94, 0xae, 0xee, 0xeb, 0xde, 0xd7, 0x7d, 0xcd, 0xb4, 0xa8, 0x61, 0x34, 0x4a, 0xd2, 0xda,
		0x2e, 0x18, 0xe5, 0x76, 0x90, 0xb1, 0x08, 0x95, 0xf0, 0x03, 0x27, 0x18, 0x42, 0x77, 0x02, 0x26,
		0xe3, 0xb8, 0x20, 0xf7, 0xda, 0x0e, 0x4b, 0xcf, 0xb1, 0x58, 0xc9, 0xe1, 0x66, 0x2f, 0x9d, 0xba,
		0xe4, 0xa5, 0x65, 0x92, 0x5c, 0xb8, 0xdc, 0x25, 0xd3, 0x33, 0x59, 0x8e, 0x32, 0x13, 0xf7, 0xcd,
		0x11, 0xff, 0x69, 0x70, 0x31, 0x30, 0x93, 0xbb, 0x59, 0x0f, 0xdb, 0x4a, 0x54, 0x01, 0xd2, 0x31,
		0x82, 0x51, 0x5f, 0x87, 0xe3, 0x2f, 0x07, 0xc7, 0x63, 0xdb, 0x6a, 0x0a, 0xe7, 0x07, 0x5b, 0x4d,
		0xc1, 0xe5, 0xa0, 0xae, 0xd2, 0xa4, 0xed, 0x72, 0x50, 0xeb, 0xc0, 0x71, 0x65, 0x8d, 0xe3, 0xaa,
		0x1c, 0x8e, 0xdf, 0x0a, 0xc2, 0xe4, 0x90, 0x2a, 0x07, 0xe4, 0x0e, 0xc8, 0x5f, 0x2c, 0x90, 0xbb,
		0xfa, 0xf4, 0x07, 0xad, 0xdb, 0x9e, 0x1a, 0xc7, 0xad, 0xdc, 0x31, 0xbf, 0xe2, 0xa4, 0xc0, 0x8d,
		0x62, 0x56, 0x5f, 0xc7, 0xbc, 0xae, 0xce, 0x52, 0x3d, 0x9d, 0x9c, 0xb3, 0xa7, 0x59, 0xd1, 0x9c,
		0x7d, 0xcc, 0x06, 0x9f, 0xa5, 0x7c, 0x4f, 0xd5, 0xe0, 0x2c, 0xe9, 0x7b, 0xf6, 0xcd, 0xf4, 0xa0,
		0xf3, 0xda, 0x04, 0xd2, 0x16, 0xb5, 0xa4, 0x65, 0x28, 0xe0, 0x74, 0x0e, 0xb3, 0xcc, 0xe5, 0x08,
		0xa5, 0x04, 0x35, 0x20, 0x0c, 0xa6, 0x93, 0x80, 0x66, 0x92, 0xc5, 0x9c, 0x94, 0xf2, 0x1d, 0x6b,
		0x91, 0xd3, 0x7f, 0x4e, 0x20, 0x20, 0x0c, 0xe4, 0x80, 0x08, 0x04, 0x2a, 0xc1, 0xf7, 0x76, 0x94,
		0x04, 0x1e, 0xce, 0x2d, 0xf6, 0x29, 0x52, 0xc0, 0xeb, 0xa3, 0xd7, 0x0e, 0xa3, 0x45, 0xab, 0xd5,
		0x09, 0x98, 0x2d, 0x99, 0x4a, 0xe8, 0x0b, 0x24, 0x0a, 0x45, 0xba, 0x60, 0x2e, 0x00, 0xff, 0x1b,
		0x93, 0x08, 0x14, 0x07, 0xc3, 0xa3, 0x79, 0x5d, 0x6c, 0xf0, 0xf4, 0xb5, 0x00, 0xec, 0xa8, 0xf2,
		0x22, 0xf3, 0xfd, 0x73, 0x2f, 0x68, 0xa0, 0xf8, 0x6a, 0x5c, 0xeb, 0x16, 0x8b, 0x7c, 0xff, 0xdc,
		0x87, 0x05, 0x3b, 0xd3, 0x78, 0xab, 0xb5, 0x48, 0x5e, 0x70, 0xd3, 0x74, 0xa2, 0x5a, 0x97, 0xb9,
		0x31, 0x5a, 0xcb, 0x27, 0x95, 0x9b, 0x7b, 0x85, 0x2c, 0xa4, 0xac, 0x6f, 0x73, 0x77, 0xee, 0x8a,
		0x03, 0x6c, 0x3e, 0x72, 0xec, 0xae, 0x38, 0x40, 0x41, 0x89, 0x88, 0x15, 0x82, 0x16, 0xc6, 0x7f,
		0x40, 0xb9, 0x92, 0x11, 0x86, 0xdc, 0xb1, 0xab, 0x73, 0xac, 0x4b, 0x74, 0xae, 0xc5, 0xd0, 0x2e,
		0x40, 0xad, 0xa5, 0x57, 0x1e, 0x4d, 0xcc, 0x54, 0x4d, 0x39, 0x99, 0xfb, 0x6e, 0x94, 0x01, 0x2a,
		0xdc, 0xce, 0xde, 0xac, 0xa5, 0x12, 0xd2, 0x01, 0x55, 0xfd, 0x30, 0x21, 0x0d, 0x9e, 0x80, 0xaa,
		0x06, 0x14, 0x6b, 0xd7, 0x51, 0x79, 0x8f, 0xee, 0xf2, 0x35, 0x4b, 0xbe, 0xb6, 0x2c, 0xd2, 0x92,
		0xeb, 0x5e, 0xed, 0xcd, 0x55, 0x8b, 0x8b, 0xeb, 0x79, 0x9c, 0xed, 0xdc, 0xbc, 0xd2, 0xc7, 0x8a,
		0xbb, 0x4d, 0xa9, 0xd6, 0xe5, 0x23, 0x2f, 0x3e, 0x69, 0x3c, 0x6d, 0x95, 0xff, 0xb0, 0x71, 0x07,
		0x24, 0x06, 0x9c, 0x85, 0xeb, 0x5e, 0x37, 0x4e, 0xec, 0x5d, 0x2a, 0x81, 0xb3, 0xb4, 0xe4, 0x51,
		0x32, 0x1e, 0x10, 0xa5, 0x04, 0xed, 0xc6, 0x0a, 0xe5, 0x1b, 0xb8, 0x09, 0xa9, 0x02, 0x39, 0x19,
		0x76, 0x79, 0x04, 0x72, 0xc0, 0xe3, 0x28, 0x04, 0xc6, 0x13, 0xeb, 0x79, 0x4c, 0x25, 0xd5, 0x9c,
		0x6d, 0xf9, 0x40, 0xb2, 0xef, 0x1e, 0x48, 0x5e, 0xab, 0xb7, 0x02, 0xc3, 0xd2, 0x46, 0x81, 0xf1,
		0x03, 0xc9, 0x54, 0x26, 0xa5, 0xae, 0x80, 0xe1, 0xcf, 0xb9, 0x1d, 0x27, 0x61, 0x88, 0x21, 0x50,
		0x96, 0x1c, 0x74, 0x7c, 0xaf, 0xf5, 0x0e, 0xc6, 0x28, 0x24, 0xe5, 0xec, 0x0d, 0xc0, 0x3f, 0x11,
		0x42, 0xce, 0x8e, 0x14, 0x0c, 0xc8, 0x18, 0xb5, 0x11, 0x2c, 0xc9, 0x04, 0xa8, 0x3a, 0x92, 0x70,
		0x94, 0xc6, 0xdc, 0xa4, 0x2e, 0x93, 0x23, 0x68, 0x02, 0x1d, 0x8e, 0x22, 0x9a, 0x86, 0x58, 0x8c,
		0x88, 0xd8, 0x8c, 0x39, 0x2e, 0x02, 0xa7, 0x82, 0x4f, 0x76, 0x0b, 0xc5, 0x93, 0xda, 0x96, 0x2f,
		0x2f, 0xb7, 0x49, 0xc9, 0x27, 0x52, 0xb2, 0x6a, 0x82, 0x30, 0x26, 0x82, 0x92, 0x55, 0x9c, 0x70,
		0x3e, 0xff, 0x17, 0xec, 0xf3, 0xdf, 0x65, 0xa5, 0xc9, 0x82, 0x0a, 0xac, 0x6d, 0xcb, 0x77, 0x18,
		0xdb, 0xdd, 0x92, 0xf2, 0x90, 0x15, 0x14, 0x73, 0xe2, 0xe0, 0xc4, 0x61, 0x99, 0xce, 0x2f, 0x2a,
		0x26, 0xed, 0xcc, 0x5d, 0x82, 0xad, 0xdc, 0x0b, 0xba, 0x4b, 0xb0, 0xed, 0x5c, 0x82, 0xd5, 0xe3,
		0x0d, 0xcd, 0x8e, 0x59, 0x27, 0xb9, 0xe6, 0x77, 0xe1, 0xe1, 0xef, 0xea, 0x9b, 0x1e, 0x24, 0xcd,
		0x19, 0x7a, 0x5f, 0x35, 0x67, 0xc8, 0xa4, 0x10, 0x6a, 0xee, 0x03, 0x77, 0x6b, 0x0e, 0x89, 0x8b,
		0x47, 0xbf, 0xaa, 0xc9, 0x42, 0xbe, 0x33, 0xe7, 0x2b, 0x32, 0xaf, 0x79, 0xb2, 0x50, 0x61, 0x08,
		0x85, 0x41, 0xe8, 0x84, 0x61, 0xc8, 0xc4, 0xb6, 0x52, 0x85, 0x5a, 0x9e, 0xf1, 0x9b, 0x67, 0xcf,
		0x21, 0x5d, 0xc8, 0xdf, 0xcf, 0xe5, 0x3e, 0x55, 0xc6, 0x50, 0xd7, 0x30, 0x63, 0xa8, 0x5b, 0x36,
		0x63, 0x28, 0x05, 0xb7, 0x24, 0x53, 0xe8, 0x18, 0x7e, 0x52, 0x35, 0x00, 0x02, 0x21, 0x8f, 0xbb,
		0x11, 0xc2, 0x0f, 0x9c, 0x54, 0x75, 0x5d, 0x38, 0xac, 0xdb, 0xba, 0xeb, 0x82, 0xb2, 0x10, 0xef,
		0x5b, 0xe6, 0x47, 0xb5, 0xac, 0x7d, 0xd9, 0x04, 0x22, 0xdd, 0xdb, 0x9d, 0xd0, 0xdc, 0x09, 0x6d,
		0x5e, 0xc3, 0x5e, 0x58, 0x1c, 0xd0, 0xce, 0x0e, 0x36, 0x46, 0xd1, 0xc5, 0x9a, 0xaf, 0x2a, 0xeb,
		0xb3, 0x33, 0x77, 0x3e, 0xab, 0xc1, 0xdd, 0x96, 0xa0, 0xaa, 0x6f, 0x89, 0xe1, 0xbe, 0xc3, 0x70,
		0x87, 0xe1, 0x0e, 0xc3, 0x1d, 0x86, 0x3b, 0x0c, 0xdf, 0x07, 0x0c, 0x4f, 0xae, 0x40, 0x2c, 0x5f,
		0x5f, 0xb1, 0x7e, 0x7b, 0xa5, 0x33, 0xbd, 0x2c, 0x99, 0x79, 0xa3, 0xa6, 0xf7, 0xd5, 0x1a, 0xd9,
		0xeb, 0xc6, 0x74, 0xdf, 0x61, 0xfa, 0xc1, 0x62, 0xfa, 0x8b, 0xba, 0x39, 0x71, 0xe9, 0x43, 0x6b,
		0xfc, 0x68, 0x0e, 0xd5, 0xeb, 0x42, 0xf5, 0xd0, 0x12, 0xd5, 0x43, 0xeb, 0x8b, 0x70, 0xa0, 0x21,
		0x32, 0x45, 0xd5, 0x44, 0x60, 0x6f, 0x37, 0xd8, 0xee, 0xec, 0xf5, 0xc3, 0xc5, 0xf6, 0x39, 0x66,
		0xb1, 0x01, 0x78, 0x13, 0xb3, 0xfd, 0x63, 0x36, 0xf4, 0x15, 0x91, 0x25, 0x92, 0x6a, 0x3e, 0xff,
		0x79, 0xd5, 0xf9, 0x76, 0x63, 0xba, 0x3f, 0x09, 0x54, 0x49, 0x63, 0x65, 0x62, 0xa7, 0x50, 0x16,
		0xe6, 0xf5, 0xf1, 0xfa, 0xf6, 0xcf, 0x2f, 0x37, 0xad, 0xc6, 0x36, 0x60, 0xba, 0xd2, 0x94, 0xfc,
		0xba, 0x8b, 0x9d, 0xdf, 0xed, 0x3a, 0xf6, 0x3b, 0x3f, 0xc9, 0x32, 0x75, 0x36, 0x43, 0xae, 0xbf,
		0x62, 0xcb, 0xd9, 0x96, 0x17, 0xdb, 0xca, 0xb6, 0xac, 0xf7, 0x0a, 0x3d, 0xf7, 0x16, 0x07, 0x0c,
		0xaf, 0xd0, 0x93, 0xdc, 0xa2, 0xab, 0x9a, 0x72, 0x8b, 0x2a, 0x45, 0x80, 0xaf, 0x0f, 0xc0, 0x36,
		0x5c, 0x87, 0x49, 0x30, 0xb8, 0x56, 0xb4, 0x4d, 0xa2, 0x9a, 0x8a, 0x8f, 0x9a, 0x49, 0xe8, 0xf6,
		0xe6, 0x88, 0xf0, 0xd5, 0xa6, 0x45, 0x61, 0xe1, 0x49, 0xc5, 0x6b, 0xa2, 0x12, 0xe5, 0xfb, 0x18,
		0x1b, 0xfe, 0x0b, 0xe3, 0x0a, 0x04, 0x06, 0x7c, 0x38, 0x44, 0x16, 0x62, 0x08, 0xdd, 0x58, 0xcd,
		0x32, 0x26, 0x65, 0x3c, 0x1a, 0x71, 0xa1, 0x30, 0x7c, 0xbd, 0x21, 0xec, 0xdb, 0xdb, 0x14, 0xf6,
		0xed, 0xbd, 0xd8, 0xb0, 0xef, 0x8d, 0x9a, 0xaf, 0xf8, 0x14, 0x93, 0x77, 0x6a, 0x69, 0x7c, 0x21,
		0x4a, 0xa1, 0x60, 0x1b, 0x35, 0x4b, 0xe3, 0x7b, 0xa7, 0xf9, 0xaf, 0xbb, 0xbf, 0xda, 0x0f, 0xcd,
		0xef, 0x5e, 0xf3, 0xdd, 0xdd, 0xdf, 0x1a, 0x45, 0xc9, 0x11, 0xaf, 0x16, 0xff, 0x95, 0xad, 0x63,
		0x93, 0x88, 0x34, 0xa8, 0xfc, 0x40, 0x7e, 0xe0, 0x57, 0xce, 0x57, 0xa9, 0xb7, 0x2c, 0x36, 0x8d,
		0xf9, 0xaf, 0x16, 0x04, 0xe3, 0x1a, 0xc7, 0x34, 0xc8, 0x44, 0xe1, 0xe1, 0xd5, 0xc3, 0xff, 0x01,
		0x00, 0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x3b, 0x70, 0xe1, 0x18, 0xaa, 0x00,
		0x00,
	}
)


// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
func initΛEnumTypes(){
  ΛEnumTypes = map[string][]reflect.Type{
	"/cont1a/cont2d/snack/late-night/chocolate": []reflect.Type{
		reflect.TypeOf((E_OnfTest1_Cont1A_Cont2D_Chocolate)(0)),
	},
	"/cont1b-state/list2b/leaf3d": []reflect.Type{
		reflect.TypeOf((E_OnfTest1Identities_MYBASE)(0)),
	},
  }
}

