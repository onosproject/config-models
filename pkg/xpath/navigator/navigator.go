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
	"encoding/base64"
	"fmt"
	"github.com/antchfx/xpath"
	"github.com/openconfig/goyang/pkg/yang"
	"reflect"
	"sort"
	"strings"
)

// YangNodeNavigator - implements xpath.NodeNavigator
type YangNodeNavigator struct {
	root, curr     *yang.Entry
	currKeys       []string
	currListKeys   []string
	currIdx        int
	currListIdx    interface{}
	deviceRoot     interface{}
	currValueStack []interface{}
}

func NewYangNodeNavigator(root *yang.Entry, device interface{}) xpath.NodeNavigator {
	currKeys, listKeys := orderKeys(root.Dir, root.Key)

	nav := &YangNodeNavigator{
		root:           root,
		curr:           root,
		currKeys:       currKeys,
		currListKeys:   listKeys,
		deviceRoot:     device,
		currValueStack: []interface{}{device},
	}
	return nav
}

// NodeType returns the XPathNodeType of the current node.
func (x *YangNodeNavigator) NodeType() xpath.NodeType {
	if x.curr.IsLeaf() && x.curr.Parent.IsList() && strings.Contains(x.curr.Parent.Key, x.curr.Name) {
		return xpath.AttributeNode
	}
	if x.curr.IsLeaf() {
		return xpath.TextNode
	}
	if x.curr.IsContainer() || x.curr.IsLeafList() || x.curr.IsList() {
		return xpath.ElementNode
	}

	return xpath.CommentNode
}

// LocalName gets the Name of the current node.
func (x *YangNodeNavigator) LocalName() string {
	return x.curr.Name
}

// Prefix returns namespace prefix associated with the current node.
func (x *YangNodeNavigator) Prefix() string {
	if x.curr.Prefix != nil {
		return x.curr.Prefix.Name
	}
	return ""
}

// Value gets the value of current node.
func (x *YangNodeNavigator) Value() string {
	if x.curr.IsLeaf() {
		val := reflect.ValueOf(x.currValueStack[len(x.currValueStack)-1])
		switch k := val.Kind(); k {
		case reflect.Slice:
			switch val.Type().Name() {
			case "Binary":
				srcBytes := make([]byte, 0, val.Len())
				for i := 0; i < val.Len(); i++ {
					valI := val.Index(i).Interface()
					srcBytes = append(srcBytes, valI.(uint8))
				}
				return base64.StdEncoding.EncodeToString(srcBytes)
			}
			return fmt.Sprint(val.Interface())
		case reflect.Struct:
			return fmt.Sprint(val.Elem().Interface())
		case reflect.Ptr:
			return fmt.Sprint(val.Elem().Interface())
		case reflect.Int64: // For identities (enumerations)
			if x.curr.Type != nil && x.curr.Type.IdentityBase != nil &&
				len(x.curr.Type.IdentityBase.Values) > int(val.Int()-1) {
				return x.curr.Type.IdentityBase.Values[val.Int()-1].Name
			}
			return fmt.Sprintf("%d", val.Int())
		default:
			fmt.Printf("unexpected kind %s\n", k)
		}
	}
	return ""
}

// Copy does a deep copy of the YangNodeNavigator and all its components.
func (x *YangNodeNavigator) Copy() xpath.NodeNavigator {
	return nil
}

// MoveToRoot moves the YangNodeNavigator to the root node of the current node.
func (x *YangNodeNavigator) MoveToRoot() {
	x.curr = x.root
	x.currKeys, x.currListKeys = orderKeys(x.curr.Dir, x.curr.Key)
}

// MoveToParent moves the YangNodeNavigator to the parent node of the current node.
func (x *YangNodeNavigator) MoveToParent() bool {
	if x.curr.Parent != nil {
		x.curr = x.curr.Parent
		x.currKeys, x.currListKeys = orderKeys(x.curr.Parent.Dir, x.curr.Key)
		x.currIdx = findKeyIndex(x.currKeys, x.curr.Name)
		x.currValueStack = x.currValueStack[:len(x.currValueStack)-1]
		return true
	}
	return false
}

// MoveToNextAttribute moves the YangNodeNavigator to the next attribute on current node.
func (x *YangNodeNavigator) MoveToNextAttribute() bool {
	return false
}

// MoveToChild moves the YangNodeNavigator to the first child node of the current node.
func (x *YangNodeNavigator) MoveToChild() bool {
	if x.curr.IsList() {
		x.currKeys, x.currListKeys = orderKeys(x.curr.Dir, x.curr.Key)
		x.currIdx = -len(x.currListKeys)
		currValue, currListIdx := listEntryFromPathTag(x.currValueStack, x.currListIdx, false)
		x.currListIdx = currListIdx
		x.currValueStack[len(x.currValueStack)-1] = currValue // Replace last entry
		x.currValueStack = append(x.currValueStack, fieldFromPathTag(x.currValueStack, x.currListKeys[0]))
		x.curr = x.curr.Dir[x.currListKeys[0]] // TODO: don't use fixed value
		return true
	}
	if !x.curr.IsLeaf() && x.curr.Dir != nil {
		currKeys, _ := orderKeys(x.curr.Dir, x.curr.Key)
		for i := 0; i < len(x.currKeys); i++ {
			curr, ok := x.curr.Dir[currKeys[i]]
			if !ok {
				continue
			}
			valueObject := fieldFromPathTag(x.currValueStack, curr.Name)
			if valueObject == nil {
				continue
			}
			x.curr = curr
			if !x.curr.IsLeaf() {
				x.currKeys, x.currListKeys = orderKeys(x.curr.Dir, x.curr.Key)
			}
			x.currValueStack = append(x.currValueStack, valueObject)
			return true
		}
	}
	return false
}

// MoveToFirst moves the YangNodeNavigator to the first sibling node of the current node.
func (x *YangNodeNavigator) MoveToFirst() bool {
	for i := 0; i < len(x.currKeys); i++ {
		// The first item might not have a value
		currIdx := i
		curr := x.curr.Parent.Dir[x.currKeys[currIdx]]
		// First check it has a value
		currValueStack := x.currValueStack[:len(x.currValueStack)-1]
		newVal := fieldFromPathTag(currValueStack, curr.Name)
		if newVal != nil {
			x.currIdx = currIdx
			x.curr = curr
			x.currValueStack = x.currValueStack[:len(x.currValueStack)-1]
			x.currValueStack = append(x.currValueStack, newVal)
			return true
		}
	}
	return false
}

// MoveToNext moves the YangNodeNavigator to the next sibling node of the current node.
func (x *YangNodeNavigator) MoveToNext() bool {
	if x.curr.IsList() { // This could be a continuation from a previous list entry
		currValueStack := x.currValueStack[:len(x.currValueStack)-1]
		listParent := fieldFromPathTag(currValueStack, x.curr.Name)
		currValueStack = append(currValueStack, listParent)
		newVal, nextIdx := listEntryFromPathTag(currValueStack, x.currListIdx, true)
		if newVal != nil {
			x.currValueStack[len(x.currValueStack)-1] = listParent
			x.currListIdx = nextIdx
			x.currKeys, _ = orderKeys(x.curr.Dir, "")
			x.currIdx = -len(x.currListKeys)
			return true
		}
	}
	// handle any extra list keys values if they exist
	if x.currIdx < -1 {
		x.currIdx++
		listIdx := x.currIdx + len(x.currListKeys)
		x.curr = x.curr.Parent.Dir[x.currListKeys[listIdx]]
		// First check it has a value
		currValueStack := x.currValueStack[:len(x.currValueStack)-1]
		nextIdx := fieldFromPathTag(currValueStack, x.curr.Name)
		currValueStack = append(currValueStack, nextIdx)
		x.currValueStack = currValueStack
		return true
	}

	for i := 0; i < len(x.currKeys); i++ {
		currIdx := x.currIdx + i + 1
		if currIdx > len(x.currKeys)-1 {
			return false
		}
		curr := x.curr.Parent.Dir[x.currKeys[currIdx]]
		// First check it has a value
		currValueStack := x.currValueStack[:len(x.currValueStack)-1]

		// Look first for a next entry in a list
		if curr.IsList() {
			listParent := fieldFromPathTag(currValueStack, curr.Name)
			currValueStack = append(currValueStack, listParent)
			newVal, nextIdx := listEntryFromPathTag(currValueStack, x.currListIdx, true)
			if newVal != nil {
				x.curr = curr
				x.currValueStack[len(x.currValueStack)-1] = listParent // Replace last entry
				x.currListIdx = nextIdx
				x.currKeys, _ = orderKeys(x.curr.Dir, "")
				x.currIdx = -len(x.currListKeys)
				return true
			} else {
				return false
			}
		}
		newVal := fieldFromPathTag(currValueStack, curr.Name)
		if newVal != nil {
			x.currIdx = currIdx
			x.curr = curr
			x.currValueStack = x.currValueStack[:len(x.currValueStack)-1]
			x.currValueStack = append(x.currValueStack, newVal)
			return true
		}
	}
	return false
}

// MoveToPrevious moves the YangNodeNavigator to the previous sibling node of the current node.
func (x *YangNodeNavigator) MoveToPrevious() bool {
	for i := 0; i < len(x.currKeys); i++ {
		currIdx := x.currIdx - i - 1
		if currIdx < 0 {
			return false
		}
		curr := x.curr.Parent.Dir[x.currKeys[currIdx]]
		// First check it has a value
		currValueStack := x.currValueStack[:len(x.currValueStack)-1]
		newVal := fieldFromPathTag(currValueStack, curr.Name)
		if newVal != nil {
			x.currIdx = currIdx
			x.curr = curr
			x.currValueStack = x.currValueStack[:len(x.currValueStack)-1]
			x.currValueStack = append(x.currValueStack, newVal)
			return true
		}
	}
	return false
}

// MoveTo moves the YangNodeNavigator to the same position as the specified YangNodeNavigator.
func (x *YangNodeNavigator) MoveTo(dest xpath.NodeNavigator) bool {
	return false
}

func orderKeys(dir map[string]*yang.Entry, listKeysStr string) ([]string, []string) {
	currKeys := make([]string, 0, len(dir))
	var listKeys []string
	if listKeysStr == "" {
		listKeys = make([]string, 0)
	} else {
		listKeys = strings.Split(listKeysStr, " ")
	}
outerLoop:
	for k := range dir {
		for _, l := range listKeys {
			if l == k {
				continue outerLoop
			}
		}
		currKeys = append(currKeys, k)
	}
	sort.Strings(currKeys)

	return currKeys, listKeys
}

func findKeyIndex(keys []string, name string) int {
	for i, k := range keys {
		if k == name {
			return i
		}
	}
	return -1
}

func fieldFromPathTag(currValue []interface{}, tag string) interface{} {
	topCurrValue := currValue[len(currValue)-1]
	if topCurrValue == nil {
		return nil
	}
	currValueType := reflect.TypeOf(topCurrValue).Elem()
	for i := 0; i < currValueType.NumField(); i++ {
		if currValueType.Field(i).Tag.Get("path") == tag {
			fieldName := currValueType.Field(i).Name
			val := reflect.ValueOf(topCurrValue).Elem().FieldByName(fieldName)
			if val.IsValid() {
				if !val.IsZero() {
					return val.Interface()
				}
			}
			return nil
		}
	}
	return nil
}

func listEntryFromPathTag(currValueStack []interface{}, index interface{}, next bool) (interface{}, interface{}) {
	topCurrValue := currValueStack[len(currValueStack)-1]
	currValueType := reflect.TypeOf(topCurrValue)
	topCurrValueType := reflect.ValueOf(topCurrValue)
	switch currValueType.Kind() {
	case reflect.Map:
		mapKeys := topCurrValueType.MapKeys()
		sort.Slice(mapKeys, func(i, j int) bool {
			return mapKeys[i].String() < mapKeys[j].String()
		})
		gotIndex := false
		for _, k := range mapKeys {
			if index != nil && reflect.TypeOf(index).Kind() == k.Kind() {
				if k.Interface() == index {
					gotIndex = true
					if next {
						continue
					} else {
						return topCurrValueType.MapIndex(k).Interface(), k.Interface()
					}
				}
				if gotIndex {
					return topCurrValueType.MapIndex(k).Interface(), k.Interface()
				}
			} else {
				return topCurrValueType.MapIndex(k).Interface(), k.Interface()
			}
		}
	case reflect.Ptr:
		topParentValue := currValueStack[len(currValueStack)-2]
		fmt.Println(topParentValue)

		// Need to build up a new Key object for the next instance
		// This happens where there is more than one key in a list
		// Then rather than using the native type(s) a struct is created
		// Here we know what that key struct type is from prior "index"
		indexType := reflect.TypeOf(index)
		newKeyValue := reflect.New(indexType) // Ptr
		for i := 0; i < indexType.NumField(); i++ {
			sfKey := indexType.FieldByIndex([]int{i})
			switch sfKey.Type.Kind() {
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
				fieldI := topCurrValueType.Elem().FieldByIndex([]int{i}).Elem().Uint()
				newKeyValue.Elem().FieldByIndex([]int{i}).SetUint(fieldI)
			default:
				panic(fmt.Errorf("unhandled type %v", sfKey))
			}
		}
		fmt.Printf("new key %v (old %v)\n", newKeyValue.Interface(), index)
	default:
		fmt.Printf("unhandled type %s\n", currValueType.Kind().String())
	}
	return nil, nil
}
