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
	"testing"
)

func Test_orderKeysWithListKeys(t *testing.T) {
	entriesMap := map[string]*yang.Entry{
		"key1": {
			Name: "e1",
		},
		"key2": {
			Name: "e2",
		},
		"key10": {
			Name: "e10",
		},
		"key0": {
			Name: "e0",
		},
	}
	listKeysStr := "key1 key2"
	otherKeys, listKeys := orderKeys(entriesMap, listKeysStr)
	assert.ElementsMatch(t, []string{"key0", "key10"}, otherKeys)
	assert.ElementsMatch(t, []string{"key1", "key2"}, listKeys)
}

func Test_orderKeysNoList(t *testing.T) {
	entriesMap := map[string]*yang.Entry{
		"key0": {
			Name: "e0",
		},
		"key1": {
			Name: "e1",
		},
		"key2": {
			Name: "e2",
		},
		"key10": {
			Name: "e10",
		},
	}
	listKeysStr := ""
	otherKeys, listKeys := orderKeys(entriesMap, listKeysStr)
	assert.ElementsMatch(t, []string{"key0", "key1", "key2", "key10"}, otherKeys)
	assert.ElementsMatch(t, []string{}, listKeys)
}

func Test_orderKeysNoKeysButListIsError(t *testing.T) {
	entriesMap := map[string]*yang.Entry{}
	listKeysStr := "key1"
	otherKeys, listKeys := orderKeys(entriesMap, listKeysStr)
	assert.ElementsMatch(t, []string{}, otherKeys)
	assert.ElementsMatch(t, []string{"key1"}, listKeys)
}
