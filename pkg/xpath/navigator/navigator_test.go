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

package navigator

import (
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type testStruct struct {
	A string `path:"a" module:"onf-test1"`
	B int    `path:"b" module:"onf-test1"`
}

func Test_processStruct(t *testing.T) {
	// TODO: create a more detailed struct val and a YANG Entry that can be passed to this

	entry := &yang.Entry{
		Name: "test1",
	}

	testStruct1 := &testStruct{
		A: "test1",
		B: 10,
	}

	testStructValue := reflect.ValueOf(testStruct1)
	t.Skip()

	result := processStruct(testStructValue, "test1", entry)
	assert.NotNil(t, result)

}
