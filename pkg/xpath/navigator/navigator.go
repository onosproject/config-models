// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package navigator

import (
	"encoding/base64"
	"fmt"
	"github.com/SeanCondon/xpath"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"reflect"
	"sort"
	"strings"
)

const (
	goStruct        = "gostruct"
	orderedAttrList = "orderedattrlist"
)

type XpathSelect struct {
	Name     string
	Path     string
	Expected []string
}

type XpathEvaluate struct {
	Name     string
	Path     string
	Expected interface{}
}

// YangNodeNavigator - implements xpath.NodeNavigator
type YangNodeNavigator struct {
	root, curr *yang.Entry
}

var log = logging.GetLogger("config-model", "navigator")

func NewYangNodeNavigator(root *yang.Entry, device ygot.ValidatedGoStruct) xpath.NodeNavigator {
	nav := &YangNodeNavigator{
		root: root,
		curr: root,
	}

	addGoStructToYangEntry(root, device)

	return nav
}

// addGoStructToYangEntry - recursive function that walks the Abstract Syntax
// Tree and matches up the GoStruct
// Also extracts the "must" statements in to XPath queries
func addGoStructToYangEntry(dir *yang.Entry, yangStruct interface{}) map[string]*yang.Entry {
	resultMap := make(map[string]*yang.Entry)

	if dir.Annotation == nil {
		dir.Annotation = make(map[string]interface{})
	}

	mustStmnt, ok := dir.Extra["must"]
	if ok {
		dir.Annotation["must"] = extractMust(mustStmnt)
	}

	// Create a new entry per list index
	if dir.IsList() {
		//mapKeysAsValues := reflect.ValueOf(yangStruct).MapKeys()
		mapIter := reflect.ValueOf(yangStruct).MapRange()
		//sort.SliceStable(mapKeysAsValues, func(i, j int) bool {
		//	return mapKeysAsValues[i].String() < mapKeysAsValues[j].String()
		//})
		for {
			if !mapIter.Next() {
				break
			}
			orderedKeys := make([]string, 0, len(dir.Dir))
			newDir := deepCopyDir(dir)
			if newDir.Annotation == nil {
				newDir.Annotation = make(map[string]interface{})
			}
			newListChildren := make(map[string]*yang.Entry)
			for k, v := range newDir.Dir {
				v.Parent = newDir
				for childKey, childValue := range processStruct(mapIter.Value(), k, v) {
					// Copy each one in to childEntries map
					newListChildren[childKey] = childValue
				}
			}
			for newListKey, newListValue := range newListChildren {
				newDir.Dir[newListKey] = newListValue
				orderedKeys = append(orderedKeys, newListKey)
			}
			sort.Strings(orderedKeys)
			newDir.Annotation[orderedAttrList] = orderedKeys
			newDir.Annotation[dir.Name] = fmt.Sprint(mapIter.Key().Interface())
			newDir.Annotation[goStruct] = mapIter.Value().Interface()
			listKey := fmt.Sprintf("%s__%v", dir.Name, mapIter.Key().Interface())
			resultMap[listKey] = newDir
		}
		return resultMap
	}
	// Else is a struct
	dir.Annotation[goStruct] = yangStruct
	if dir.IsLeaf() {
		resultMap[dir.Name] = dir
		return resultMap
	}
	childMap := make(map[string]*yang.Entry)
	for k, v := range dir.Dir {
		structVal := reflect.ValueOf(yangStruct)
		switch structVal.Kind() {
		case reflect.Ptr:
			for childKey, childValue := range processStruct(structVal, k, v) {
				childMap[childKey] = childValue
			}
		default:
			panic(fmt.Errorf("unhandled kind %s", structVal.Kind().String()))
		}
	}
	if len(childMap) > 0 {
		orderedKeys := make([]string, 0, len(childMap))
		for k, v := range childMap {
			v.Parent = dir
			dir.Dir[k] = v
			orderedKeys = append(orderedKeys, k)
		}
		sort.Strings(orderedKeys)
		dir.Annotation[orderedAttrList] = orderedKeys
	}
	resultMap[dir.Name] = dir
	return resultMap
}

// processStruct - part of the recursive function addGoStructToYangEntry
func processStruct(structVal reflect.Value, dirName string, dirValue *yang.Entry) map[string]*yang.Entry {
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

// extractMust - this is necessary since the Must statement is not
// yet a first class citizen of the yang.Entry - for the moment it
// is crammed in to the Extra field
func extractMust(mustStmnt []interface{}) *yang.Must {
	mustStruct := new(yang.Must)
	for _, s := range mustStmnt {
		sMap, mapOK := s.(map[string]interface{})
		if mapOK {
			mustStruct.Name = sMap["Name"].(string)
			desc, descOK := sMap["Description"]
			if descOK {
				descMap, descMapOK := desc.(map[string]interface{})
				if descMapOK {
					mustStruct.Description = &yang.Value{
						Name: descMap["Name"].(string),
					}
				}
			}
			err, errOK := sMap["ErrorMessage"]
			if errOK {
				errMap, errMapOK := err.(map[string]interface{})
				if errMapOK {
					mustStruct.ErrorMessage = &yang.Value{
						Name: errMap["Name"].(string),
					}
				}
			}
			errAppTag, errAppTagOK := sMap["ErrorAppTag"]
			if errAppTagOK {
				errAppTagMap, errMapAppTagOK := errAppTag.(map[string]interface{})
				if errMapAppTagOK {
					mustStruct.ErrorAppTag = &yang.Value{
						Name: errAppTagMap["Name"].(string),
					}
				}
			}
		}
	}
	return mustStruct
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
			if strings.Contains(k, "__") {
				continue
			}
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

// WalkAndValidateMust - walk through the YNN and validate any Must statements
// This goes down first and then across
func (x *YangNodeNavigator) WalkAndValidateMust() error {
	for {
		if x.MoveToChild() ||
			x.MoveToNext() ||
			(x.MoveToParent() && x.MoveToNext()) ||
			(x.MoveToParent() && x.MoveToNext()) ||
			(x.MoveToParent() && x.MoveToNext()) ||
			(x.MoveToParent() && x.MoveToNext()) ||
			(x.MoveToParent() && x.MoveToNext()) {

			mustIf, ok := x.curr.Annotation["must"]
			if ok {
				mustStruct, okMustStruct := mustIf.(*yang.Must)
				if okMustStruct {
					mustExpr, err := xpath.Compile(mustStruct.Name)
					if err != nil {
						return err
					}
					x1 := x.Copy().(*YangNodeNavigator)
					result := mustExpr.Evaluate(x1)
					resultBool, resultOk := result.(bool)
					if !resultOk {
						return fmt.Errorf("result of %s cannot be evaluated as bool %v",
							mustExpr.String(), result)
					}
					if !resultBool {
						items := x1.generateMustError("@*")
						if len(items) == 0 {
							items = x1.generateMustError("*")
						}
						return fmt.Errorf("%s. Must statement '%v' to true. Container(s): %v",
							mustStruct.ErrorMessage.Name,
							mustStruct.Name, items)
					}
					log.Infof("Checking Must rule %s: %v", mustExpr.String(), resultBool)
				}
			}
			continue
		}
		return nil
	}
}

func (x *YangNodeNavigator) generateMustError(expr string) []string {
	currentExpr, currentErr := xpath.Compile(expr)
	if currentErr != nil {
		return nil
	}
	currentIter := currentExpr.Select(x)
	items := make([]string, 0)
	for currentIter.MoveNext() {
		items = append(items, fmt.Sprintf("%s=%s", currentIter.Current().LocalName(), currentIter.Current().Value()))
	}
	return items
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
	ynnCopy := YangNodeNavigator{
		root: x.root,
		curr: x.curr,
	}

	return &ynnCopy
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
	if x.curr.IsList() && x.curr.Key != "" {
		keys := strings.Split(x.curr.Key, " ")
		x.curr = x.curr.Dir[keys[0]]
		return true
	} else if x.curr.Parent != nil && x.curr.Parent.IsList() {
		keys := strings.Split(x.curr.Parent.Key, " ")
		for i, k := range keys {
			if x.curr.Name == k && i < len(keys)-1 {
				x.curr = x.curr.Parent.Dir[keys[i+1]]
				return true
			}
		}
	}
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
			valueObject := getGoStruct(curr.Annotation)
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
func (x *YangNodeNavigator) MoveTo(other xpath.NodeNavigator) bool {
	node, ok := other.(*YangNodeNavigator)
	if !ok || node.root != x.root {
		return false
	}

	x.curr = node.curr
	return true
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
