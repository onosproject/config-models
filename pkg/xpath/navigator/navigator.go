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

const (
	goStruct        = "gostruct"
	orderedAttrList = "orderedattrlist"
)

// YangNodeNavigator - implements xpath.NodeNavigator
type YangNodeNavigator struct {
	root, curr *yang.Entry
}

func NewYangNodeNavigator(root *yang.Entry, device interface{}) xpath.NodeNavigator {
	nav := &YangNodeNavigator{
		root: root,
		curr: root,
	}

	addGoStructToYangEntry(root, device)

	return nav
}

func addGoStructToYangEntry(dir *yang.Entry, yangStruct interface{}) []string {
	if dir.Annotation == nil {
		dir.Annotation = make(map[string]interface{})
	}
	// Create a new entry per list index
	if dir.IsList() {
		listKeys := make([]string, 0)
		mapIter := reflect.ValueOf(yangStruct).MapRange()
		for mapIter.Next() {
			orderedKeys := make([]string, 0, len(dir.Dir))
			newDir := deepCopyDir(dir)
			if newDir.Annotation == nil {
				newDir.Annotation = make(map[string]interface{})
			}
			for k, v := range newDir.Dir {
				processStruct(mapIter.Value(), k, v)
				v.Parent = newDir
				orderedKeys = append(orderedKeys, k)
			}
			newDir.Annotation[goStruct] = mapIter.Value().Interface()
			listKey := fmt.Sprintf("%s__%v", dir.Name, mapIter.Key().Interface())
			sort.Strings(orderedKeys)
			newDir.Annotation[orderedAttrList] = orderedKeys
			newDir.Annotation[dir.Name] = fmt.Sprint(mapIter.Key().Interface())
			dir.Parent.Dir[listKey] = newDir
			listKeys = append(listKeys, listKey)
		}
		return listKeys
	}
	// Else is a struct
	dir.Annotation[goStruct] = yangStruct
	if dir.IsLeaf() {
		return []string{dir.Name}
	}
	orderedKeys := make([]string, 0, len(dir.Dir))
	for k, v := range dir.Dir {
		structVal := reflect.ValueOf(yangStruct)
		switch structVal.Kind() {
		case reflect.Ptr:
			keysAdded := processStruct(structVal, k, v)
			orderedKeys = append(orderedKeys, keysAdded...)
		default:
			panic(fmt.Errorf("unhandled kind %s", structVal.Kind().String()))
		}
	}
	if len(orderedKeys) > 0 {
		sort.Strings(orderedKeys)
		dir.Annotation[orderedAttrList] = orderedKeys
	}
	return []string{dir.Name}
}

func processStruct(structVal reflect.Value, dirName string, dirValue *yang.Entry) []string {
	for i := 0; i < structVal.Elem().Type().NumField(); i++ {
		fieldPathName := structVal.Elem().Type().Field(i).Tag.Get("path")
		if fieldPathName != dirName {
			continue
		}
		fieldName := structVal.Elem().Type().Field(i).Name
		val := structVal.Elem().FieldByName(fieldName)
		if !val.IsZero() {
			return addGoStructToYangEntry(dirValue, val.Interface()) //Recursive
		}
		return nil
	}
	return nil
}

func deepCopyDir(dir *yang.Entry) *yang.Entry {
	newDir := &yang.Entry{
		Parent:      dir.Parent,
		Node:        dir.Node,
		Name:        dir.Name,
		Description: dir.Description,
		Default:     dir.Default,
		Units:       dir.Units,
		Errors:      dir.Errors,
		Kind:        dir.Kind,
		Config:      dir.Config,
		Prefix:      dir.Prefix,
		Mandatory:   dir.Mandatory,
		Dir:         nil,
		Key:         dir.Key,
		Type:        dir.Type,
		Exts:        dir.Exts,
		ListAttr:    dir.ListAttr,
		RPC:         dir.RPC,
		Identities:  dir.Identities,
		Augments:    dir.Augments,
		Augmented:   dir.Augmented,
		Deviations:  dir.Deviations,
		Deviate:     dir.Deviate,
		Uses:        dir.Uses,
		Extra:       dir.Extra,
		Annotation:  nil,
	}

	if dir.Dir != nil {
		newDir.Dir = make(map[string]*yang.Entry)
		for k, v := range dir.Dir {
			newDir.Dir[k] = deepCopyDir(v)
		}
	}

	if dir.Annotation != nil {
		newDir.Annotation = make(map[string]interface{})
		for k, v := range dir.Annotation {
			if k != goStruct && k != orderedAttrList {
				newDir.Annotation[k] = v
			}
		}
	}

	return newDir
}

// NodeType returns the XPathNodeType of the current node.
func (x *YangNodeNavigator) NodeType() xpath.NodeType {
	if x.curr.Parent != nil && x.curr.Parent.IsList() && strings.Contains(x.curr.Parent.Key, x.curr.Name) {
		return xpath.AttributeNode
	}
	if x.curr.IsLeaf() {
		return xpath.ElementNode
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
		value, ok := x.curr.Annotation[goStruct]
		if ok {
			switch vt := value.(type) {
			case *uint:
				return fmt.Sprint(*(vt))
			case *uint8:
				return fmt.Sprint(*(vt))
			case *uint16:
				return fmt.Sprint(*(vt))
			case *uint32:
				return fmt.Sprint(*(vt))
			case *uint64:
				return fmt.Sprint(*(vt))
			case *int:
				return fmt.Sprint(*(vt))
			case *int8:
				return fmt.Sprint(*(vt))
			case *int16:
				return fmt.Sprint(*(vt))
			case *int32:
				return fmt.Sprint(*(vt))
			case *int64:
				return fmt.Sprint(*(vt))
			case *float64:
				return fmt.Sprint(*(vt))
			case *bool:
				return fmt.Sprint(*(vt))
			case *string:
				return *(vt)
			default:
				valueReflected := reflect.ValueOf(value)
				switch valueReflected.Kind() {
				case reflect.Slice: // Most likely Binary
					bytes := valueReflected.Bytes()
					return base64.StdEncoding.EncodeToString(bytes)
				case reflect.Int64: // Most likely a YANG Identity
					ytype := x.curr.Type
					if ytype != nil {
						base := ytype.IdentityBase
						if base != nil && len(base.Values) > 0 {
							return base.Values[valueReflected.Int()-1].Name
						}
					}
					return fmt.Sprint(valueReflected.Int())
				default:
					panic(fmt.Errorf("unhandled value type %s", valueReflected.Kind()))
				}
			}
		}
	} else if x.curr.IsLeafList() {
		value, ok := x.curr.Annotation[goStruct]
		if ok {
			switch value.(type) {
			case []int16:
				return fmt.Sprint(value)
			default:
				panic(fmt.Errorf("unhandled leaf list type"))
			}
		}
	}

	return fmt.Sprintf("value of %s", x.curr.Name)
}

// Copy does a deep copy of the YangNodeNavigator and all its components.
func (x *YangNodeNavigator) Copy() xpath.NodeNavigator {
	copy := YangNodeNavigator{
		root: x.root,
		curr: x.curr,
	}

	return &copy
}

// MoveToRoot moves the YangNodeNavigator to the root node of the current node.
func (x *YangNodeNavigator) MoveToRoot() {
	x.curr = x.root
}

// MoveToParent moves the YangNodeNavigator to the parent node of the current node.
func (x *YangNodeNavigator) MoveToParent() bool {
	if x.curr.Parent != nil {
		x.curr = x.curr.Parent
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
	if !x.curr.IsLeaf() && x.curr.Dir != nil {
		currKeys := getOrderedKeys(x.curr.Annotation)
		for i := 0; i < len(currKeys); i++ {
			curr, ok := x.curr.Dir[currKeys[i]]
			if !ok {
				continue
			}
			valueObject := getGoStruct(x.curr.Annotation)
			if valueObject == nil {
				continue
			}
			x.curr = curr
			return true
		}
	}
	return false
}

// MoveToFirst moves the YangNodeNavigator to the first sibling node of the current node.
func (x *YangNodeNavigator) MoveToFirst() bool {
	parent := x.curr.Parent
	if parent != nil {
		orderedKeys := getOrderedKeys(x.curr.Parent.Annotation)
		curr, ok := parent.Dir[orderedKeys[0]]
		if ok {
			if _, structok := curr.Annotation[goStruct]; structok {
				x.curr = curr
				return true
			}
		}
	}
	return false
}

// MoveToNext moves the YangNodeNavigator to the next sibling node of the current node.
func (x *YangNodeNavigator) MoveToNext() bool {
	parent := x.curr.Parent
	if parent != nil {
		currName := x.curr.Name
		nextKey := getNextKey(x.curr.Annotation, parent.Annotation, currName)
		for ; nextKey != ""; nextKey = getNextKey(x.curr.Annotation, parent.Annotation, currName) {
			curr, ok := parent.Dir[nextKey]
			if ok {
				if _, structok := curr.Annotation[goStruct]; structok {
					x.curr = curr
					return true
				} else {
					currName = nextKey
				}
			}
		}
	}
	return false
}

// MoveToPrevious moves the YangNodeNavigator to the previous sibling node of the current node.
func (x *YangNodeNavigator) MoveToPrevious() bool {
	parent := x.curr.Parent
	if parent != nil {
		currName := x.curr.Name
		prevKey := getPreviousKey(parent.Annotation, currName)
		for ; prevKey != ""; prevKey = getPreviousKey(parent.Annotation, currName) {
			curr, ok := parent.Dir[prevKey]
			if ok {
				if _, structok := curr.Annotation[goStruct]; structok {
					x.curr = curr
					return true
				} else {
					currName = prevKey
				}
			}
		}
	}
	return false
}

// MoveTo moves the YangNodeNavigator to the same position as the specified YangNodeNavigator.
func (x *YangNodeNavigator) MoveTo(dest xpath.NodeNavigator) bool {
	return false
}

func getOrderedKeys(annotation map[string]interface{}) []string {
	orderedKeys, ok := annotation[orderedAttrList]
	if ok {
		strArray, ok := orderedKeys.([]string)
		if ok {
			return strArray
		}
		panic(fmt.Errorf("unexpected type for orderedKeys: %v", orderedKeys))
	}
	return nil
}

func getNextKey(selfAnnot map[string]interface{}, parentAnnot map[string]interface{}, key string) string {
	listEntryName, ok := selfAnnot[key]
	orderedKeys := getOrderedKeys(parentAnnot)
	for i, k := range orderedKeys {
		if k == key {
			if i < len(orderedKeys)-1 {
				return orderedKeys[i+1]
			}
		} else if ok && k == fmt.Sprintf("%s__%s", key, listEntryName) {
			if i < len(orderedKeys)-1 {
				return orderedKeys[i+1]
			}
		}
	}
	return ""
}

func getPreviousKey(annotation map[string]interface{}, key string) string {
	orderedKeys := getOrderedKeys(annotation)
	for i, k := range orderedKeys {
		if k == key {
			if i > 0 {
				return orderedKeys[i-1]
			}
		}
	}
	return ""
}

func getGoStruct(annotation map[string]interface{}) interface{} {
	gostruct, ok := annotation[goStruct]
	if ok {
		return gostruct
	}
	return nil
}
