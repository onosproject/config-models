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
		case "/t1:cont1a/cont2a/leaf2a":
			assert.Equal(t, "1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/cont2a/leaf2b":
			assert.Equal(t, "0.432", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_DECIMAL, (&value).Type)
		case "/t1:cont1a/cont2a/leaf2d":
			assert.Equal(t, "1.540", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_DECIMAL, (&value).Type)
		case "/t1:cont1a/cont2a/leaf2e":
			assert.Equal(t, "[5 4 3 2 1] 32", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_LEAFLIST_INT, (&value).Type)
		case "/t1:cont1a/cont2a/leaf2f":
			assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_BYTES, (&value).Type)
		case "/t1:cont1a/cont2a/leaf2g":
			assert.Equal(t, "true", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_BOOL, (&value).Type)
		case "/t1:cont1a/leaf1a":
			assert.Equal(t, "leaf1aval", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/list2a[name=0]/name":
			assert.Equal(t, "l2a1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/list2a[name=0]/ref2d":
			assert.Equal(t, "1.54", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/list2a[name=0]/tx-power":
			assert.Equal(t, "5", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/list2a[name=0]/range-min":
			assert.Equal(t, "20", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/list2a[name=0]/range-max":
			assert.Equal(t, "20", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/list2a[name=1]/name":
			assert.Equal(t, "l2a2", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/list2a[name=1]/tx-power":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/list2a[name=1]/range-min":
			assert.Equal(t, "2", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/list2a[name=1]/range-max":
			assert.Equal(t, "4", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1e:list5[key1=0][key2=*]/key1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list5[key1=0][key2=*]/key2":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1e:list5[key1=0][key2=*]/leaf5a":
			assert.Equal(t, "5a five-6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list5[key1=1][key2=*]/key1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list5[key1=1][key2=*]/key2":
			assert.Equal(t, "7", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_UINT, (&value).Type)
		case "/t1:cont1a/t1e:list5[key1=1][key2=*]/leaf5a":
			assert.Equal(t, "5a five-7", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/id":
			assert.Equal(t, "l2a1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/leaf4b":
			assert.Equal(t, "this is list4-l2a1", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=0][fkey2=*]/fkey1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=0][fkey2=*]/fkey2":
			assert.Equal(t, "7", (&value).ValueToString()) // TODO should be UINT
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=0][fkey2=*]/displayname":
			assert.Equal(t, "Value l2a1-five-7", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=1][fkey2=*]/fkey1":
			assert.Equal(t, "five", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=1][fkey2=*]/fkey2":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type) // TODO should be UINT
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=1][fkey2=*]/displayname":
			assert.Equal(t, "Value l2a1-five-6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=2][fkey2=*]/fkey1":
			assert.Equal(t, "six", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=2][fkey2=*]/fkey2":
			assert.Equal(t, "6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type) // TODO should be UINT
		case "/t1:cont1a/t1e:list4[id=0]/list4a[fkey1=2][fkey2=*]/displayname":
			assert.Equal(t, "Value l2a1-six-6", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=1]/id":
			assert.Equal(t, "l2a2", (&value).ValueToString())
			assert.Equal(t, configapi.ValueType_STRING, (&value).Type)
		case "/t1:cont1a/t1e:list4[id=1]/leaf4b":
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
			pathObjStr:  `path:"/t1:cont1a/leaf1a" value_type:STRING description:"Leaf inside Container 1a" length:"5..10" AttrName:"leaf1a" `,
			pathWithIdx: `/t1:cont1a/leaf1a`,
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
		`/t1:cont1a/t1e:list4[test1]/leaf4b`: {
			pathObjStr:  `path:"/t1:cont1a/t1e:list4[id=*]/leaf4b" value_type:STRING description:"leaf 4a on list4a elements" length:"1..20" AttrName:"leaf4b" `,
			pathWithIdx: `/t1:cont1a/t1e:list4[id=test1]/leaf4b`,
			found:       true,
		},
		// TODO fix this - does not handle double indices properly
		//`/t1:cont1a/t1e:list4[test1]/list4a[k1][k2]/displayname`: {
		//	pathObjStr:  `path:"/t1:cont1a/t1e:list4[id=*]/list4a[fkey1=*][fkey2=*]/displayname" value_type:STRING description:"an optional display name attribute with 2 different length ranges" length:"1..5" length:"10..20" AttrName:"displayname" `,
		//	pathWithIdx: `/t1:cont1a/t1e:list4[id=test1]/list4a[fkey1=k1][fkey2=k2]/displayname`,
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
		`/t1:cont1a/cont2a/leaf2c`: {
			pathObjStr:  `sub_path:"/" value_type:STRING description:"Read only leaf inside Container 2a" AttrName:"leaf2c" `,
			pathWithIdx: `/t1:cont1a/cont2a/leaf2c`,
			found:       true,
		},
		`/t1:cont1b-state/leaf2d`: {
			pathObjStr:  `sub_path:"/leaf2d" value_type:UINT type_opts:16 description:"A state attribute" AttrName:"leaf2d" `,
			pathWithIdx: `/t1:cont1b-state/leaf2d`,
			found:       true,
		},
		`/t1:cont1b-state/list2b[5]/index`: {
			pathObjStr:  `sub_path:"/list2b[index=*]/index" value_type:UINT type_opts:8 description:"The list index" IsAKey:true AttrName:"index" `,
			pathWithIdx: `/t1:cont1b-state/list2b[index=5]/index`,
			found:       true,
		},
		`/t1:cont1b-state/list2b[5]/leaf3c`: {
			pathObjStr:  `sub_path:"/list2b[index=*]/leaf3c" value_type:STRING description:"A string attribute in the list" AttrName:"leaf3c" `,
			pathWithIdx: `/t1:cont1b-state/list2b[index=5]/leaf3c`,
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
			expectedPath:  `/t1:cont1a/leaf1a`,
			expectedValue: `test-string`,
			expectedType:  configapi.ValueType_STRING,
		},
		`/t1:cont1a/cont2a/leaf2a`: {
			value:         `1`,
			expectedPath:  `/t1:cont1a/cont2a/leaf2a`,
			expectedValue: `1`,
			expectedType:  configapi.ValueType_UINT,
		},
		`/t1:cont1a/list2a[2a-1]/tx-power`: {
			value:         `6`,
			expectedPath:  `/t1:cont1a/list2a[name=2a-1]/tx-power`,
			expectedValue: `6`,
			expectedType:  configapi.ValueType_UINT,
		},
		`/cont1b-state/list2b[5]/leaf3c`: {
			value:         `test-string`,
			expectedPath:  `/t1:cont1b-state/list2b[index=5]/leaf3c`,
			expectedValue: `test-string`,
			expectedType:  configapi.ValueType_STRING,
		},
		`/t1:cont1a/leaf-non-existent`: {
			errString: `unable to locate /t1:cont1a/leaf-non-existent in model`,
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
