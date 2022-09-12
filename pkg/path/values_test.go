/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package path

import (
	configapi "github.com/onosproject/onos-api/go/onos/config/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type findIdxTestRwTest struct {
	pathObjStr  string
	pathWithIdx string
	found       bool
}

type handleAttrTest struct {
	value         string
	expectedPath  string
	expectedValue string
	expectedType  configapi.ValueType
	errString     string
}

func Test_GetPathValues(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("testdata/sample-testdevice-1-config.json")
	assert.NoError(t, err)

	pathValues, err := GetPathValues("", sampleConfig)
	assert.NoError(t, err)
	assert.Equal(t, 35, len(pathValues))

	for _, pathValue := range pathValues {
		value := pathValue.GetValue()
		switch path := pathValue.Path; path {
		case "/t1:cont1a/t1:cont2a/t1:leaf2a":
			assert.Equal(t, "1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1:cont2a/t1:leaf2b":
			assert.Equal(t, "0.432", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_DECIMAL, (&value).Type)
		case "/t1:cont1a/t1:cont2a/t1:leaf2d":
			assert.Equal(t, "1.540", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_DECIMAL, (&value).Type)
		case "/t1:cont1a/t1:cont2a/t1:leaf2e":
			assert.Equal(t, "[5 4 3 2 1] 32", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_LEAFLIST_INT, (&value).Type)
		case "/t1:cont1a/t1:cont2a/t1:leaf2f":
			assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_BYTES, (&value).Type)
		case "/t1:cont1a/t1:cont2a/t1:leaf2g":
			assert.Equal(t, "true", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_BOOL, (&value).Type)
		case "/t1:cont1a/t1:leaf1a":
			assert.Equal(t, "leaf1aval", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=0]/t1:name":
			assert.Equal(t, "l2a1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=0]/t1:ref2d":
			assert.Equal(t, "1.54", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=0]/t1:tx-power":
			assert.Equal(t, "5", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=0]/t1:range-min":
			assert.Equal(t, "20", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=0]/t1:range-max":
			assert.Equal(t, "20", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=1]/t1:name":
			assert.Equal(t, "l2a2", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=1]/t1:tx-power":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=1]/t1:range-min":
			assert.Equal(t, "2", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1:list2a[t1:name=1]/t1:range-max":
			assert.Equal(t, "4", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1e:list5[t1e:key1=0][t1e:key2=*]/t1e:key1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list5[t1e:key1=0][t1e:key2=*]/t1e:key2":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1e:list5[t1e:key1=0][t1e:key2=*]/t1e:leaf5a":
			assert.Equal(t, "5a five-6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list5[t1e:key1=1][t1e:key2=*]/t1e:key1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list5[t1e:key1=1][t1e:key2=*]/t1e:key2":
			assert.Equal(t, "7", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1e:list5[t1e:key1=1][t1e:key2=*]/t1e:leaf5a":
			assert.Equal(t, "5a five-7", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:id":
			assert.Equal(t, "l2a1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:leaf4b":
			assert.Equal(t, "this is list4-l2a1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=0][t1e:fkey2=*]/t1e:fkey1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=0][t1e:fkey2=*]/t1e:fkey2":
			assert.Equal(t, "7", (&value).ValueToString()) // TODO should be UINT
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=0][t1e:fkey2=*]/t1e:displayname":
			assert.Equal(t, "Value l2a1-five-7", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=1][t1e:fkey2=*]/t1e:fkey1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=1][t1e:fkey2=*]/t1e:fkey2":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type) // TODO should be UINT
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=1][t1e:fkey2=*]/t1e:displayname":
			assert.Equal(t, "Value l2a1-five-6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=2][t1e:fkey2=*]/t1e:fkey1":
			assert.Equal(t, "six", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=2][t1e:fkey2=*]/t1e:fkey2":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type) // TODO should be UINT
		case "/t1:cont1a/t1e:list4[t1e:id=0]/t1e:list4a[t1e:fkey1=2][t1e:fkey2=*]/t1e:displayname":
			assert.Equal(t, "Value l2a1-six-6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=1]/t1e:id":
			assert.Equal(t, "l2a2", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[t1e:id=1]/t1e:leaf4b":
			assert.Equal(t, "this is list4-l2a2", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

}

func Test_findModelRwPathNoIndicesNew(t *testing.T) {
	tests := map[string]*findIdxTestRwTest{
		`/t1:cont1a/leaf1a`: {
			pathObjStr:  `path:"/t1:cont1a/t1:leaf1a" value_type:STRING description:"Leaf inside Container 1a" length:"5..10" AttrName:"leaf1a" `,
			pathWithIdx: `/t1:cont1a/t1:leaf1a`,
			found:       true,
		},
		`/t1:leafAtTopLevel`: {
			pathObjStr:  `path:"/t1:leafAtTopLevel" value_type:STRING description:"A leaf at the top level (not recommended but must be supported)" AttrName:"leafAtTopLevel" `,
			pathWithIdx: `/t1:leafAtTopLevel`,
			found:       true,
		},
		`//t1:leafAtTopLevel`: { // With double slash
			pathObjStr:  `path:"/t1:leafAtTopLevel" value_type:STRING description:"A leaf at the top level (not recommended but must be supported)" AttrName:"leafAtTopLevel" `,
			pathWithIdx: `/t1:leafAtTopLevel`,
			found:       true,
		},
		`/t1:cont1a/t1e:list4[test1]/t1:leaf4b`: {
			pathObjStr:  `path:"/t1:cont1a/t1e:list4[t1e:id=*]/t1e:leaf4b" value_type:STRING description:"leaf 4a on list4a elements" length:"1..20" AttrName:"leaf4b" `,
			pathWithIdx: `/t1:cont1a/t1e:list4[t1e:id=test1]/t1e:leaf4b`,
			found:       true,
		},
		// TODO fix this - does not handle double indices properly
		//`/t1:cont1a/t1e:list4[test1]/t1e:list4a[k1][k2]/t1e:displayname`: {
		//	pathObjStr:  `path:"/t1:cont1a/t1e:list4[t1e:id=*]/t1e:list4a[t1e:fkey1=*][t1e:fkey2=*]/t1e:displayname" value_type:STRING description:"an optional display name attribute with 2 different length ranges" length:"1..5" length:"10..20" AttrName:"displayname" `,
		//	pathWithIdx: `/t1:cont1a/t1e:list4[t1e:id=test1]/t1e:list4a[t1e:fkey1=k1][t1e:fkey2=k2]/t1e:displayname`,
		//	found:       true,
		//},
	}

	for searchPath, result := range tests {
		pathObj, withNumIdx, found := findModelRwPathNoIndices(searchPath)
		assert.Equal(t, result.pathObjStr, pathObj.String())
		assert.Equal(t, result.found, found)
		assert.Equal(t, result.pathWithIdx, withNumIdx)
	}
}

func Test_findModelRoPathNoIndicesNew(t *testing.T) {
	tests := map[string]*findIdxTestRwTest{
		`/t1:cont1a/t1:cont2a/t1:leaf2c`: {
			pathObjStr:  `sub_path:"/" value_type:STRING description:"Read only leaf inside Container 2a" AttrName:"leaf2c" `,
			pathWithIdx: `/t1:cont1a/t1:cont2a/t1:leaf2c`,
			found:       true,
		},
		`/t1:cont1b-state/t1:leaf2d`: {
			pathObjStr:  `sub_path:"/t1:leaf2d" value_type:UINT type_opts:16 description:"A state attribute" AttrName:"leaf2d" `,
			pathWithIdx: `/t1:cont1b-state/t1:leaf2d`,
			found:       true,
		},
		`/t1:cont1b-state/t1:list2b[5]/t1:index`: {
			pathObjStr:  `sub_path:"/t1:list2b[t1:index=*]/t1:index" value_type:UINT type_opts:8 description:"The list index" IsAKey:true AttrName:"index" `,
			pathWithIdx: `/t1:cont1b-state/t1:list2b[t1:index=5]/t1:index`,
			found:       true,
		},
		`/t1:cont1b-state/t1:list2b[5]/t1:leaf3c`: {
			pathObjStr:  `sub_path:"/t1:list2b[t1:index=*]/t1:leaf3c" value_type:STRING description:"A string attribute in the list" AttrName:"leaf3c" `,
			pathWithIdx: `/t1:cont1b-state/t1:list2b[t1:index=5]/t1:leaf3c`,
			found:       true,
		},
	}

	for searchPath, result := range tests {
		pathObj, withNumIdx, found := findModelRoPathNoIndices(searchPath)
		assert.Equal(t, result.pathObjStr, pathObj.String())
		assert.Equal(t, result.found, found)
		assert.Equal(t, result.pathWithIdx, withNumIdx)
	}
}

func Test_handleAttribute(t *testing.T) {

	tests := map[string]handleAttrTest{
		`/t1:cont1a/leaf1a`: {
			value:         `test-string`,
			expectedPath:  `/t1:cont1a/t1:leaf1a`,
			expectedValue: `test-string`,
			expectedType:  configapi.ValueType_STRING,
		},
		`/t1:cont1a/t1:cont2a/leaf2a`: {
			value:         `1`,
			expectedPath:  `/t1:cont1a/t1:cont2a/t1:leaf2a`,
			expectedValue: `1`,
			expectedType:  configapi.ValueType_UINT,
		},
		`/t1:cont1a/list2a[2a-1]/tx-power`: {
			value:         `6`,
			expectedPath:  `/t1:cont1a/t1:list2a[t1:name=2a-1]/t1:tx-power`,
			expectedValue: `6`,
			expectedType:  configapi.ValueType_UINT,
		},
		`/cont1b-state/list2b[5]/leaf3c`: {
			value:         `test-string`,
			expectedPath:  `/t1:cont1b-state/t1:list2b[t1:index=5]/t1:leaf3c`,
			expectedValue: `test-string`,
			expectedType:  configapi.ValueType_STRING,
		},
		`/t1:cont1a/leaf-non-existent`: {
			errString: `unable to locate /t1:cont1a/t1:leaf-non-existent in model`,
		},
	}

	for parentPath, tt := range tests {
		pathValue, err := handleAttribute(tt.value, parentPath)
		if tt.errString != "" {
			assert.Errorf(t, err, tt.errString)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, pathValue)
			assert.Equal(t, tt.expectedPath, pathValue.GetPath())
			value := pathValue.GetValue()
			assert.Equal(t, tt.expectedValue, (&value).ValueToString())
			assert.Equal(t, tt.expectedType, value.GetType())
		}
	}

}
