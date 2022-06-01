/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	configapi "github.com/onosproject/onos-api/go/onos/config/v2"
	"github.com/openconfig/goyang/pkg/yang"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	slash     = "/"
	equals    = "="
	bracketsq = "["
	brktclose = "]"
	colon     = ":"
)

type indexValue struct {
	name  string
	value *configapi.TypedValue
	order int
}

// regexp to find indices in paths names
const matchOnIndex = `(\[.*?]).*?`

var rOnIndex = regexp.MustCompile(matchOnIndex)

// Native r/o and r/w path maps
var roPathMap ReadOnlyPathMap
var rwPathMap ReadWritePathMap

// ExtractPaths parse the schema entries out in to flat paths
func ExtractPaths(entries map[string]*yang.Entry) ([]*admin.ReadOnlyPath, []*admin.ReadWritePath) {

	roPaths, rwPaths, err := extractPaths(entries["Device"], yang.TSUnset, "", "")
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	return roPaths, rwPaths
}

// extractPaths - recursive function that walks the YGOT tree to extract paths
func extractPaths(deviceEntry *yang.Entry, parentState yang.TriState, parentPath string,
	subpathPrefix string) ([]*admin.ReadOnlyPath, []*admin.ReadWritePath, error) {

	readOnlyPaths := make([]*admin.ReadOnlyPath, 0)
	readWritePaths := make([]*admin.ReadWritePath, 0)

	for _, dirEntry := range deviceEntry.Dir {
		itemPath := formatName(dirEntry, false, parentPath, subpathPrefix)
		if dirEntry.IsLeaf() || dirEntry.IsLeafList() {
			// No need to recurse
			t, typeOpts, err := toValueType(dirEntry.Type, dirEntry.IsLeafList())
			if err != nil {
				return nil, nil, err
			}
			tObj := admin.ReadOnlySubPath{
				SubPath:     itemPath,
				ValueType:   t,
				TypeOpts:    typeOpts,
				Description: dirEntry.Description,
				Units:       dirEntry.Units,
				IsAKey:      false,
				AttrName:    dirEntry.Name,
			}
			var enum map[int]string
			if dirEntry.Type.Kind == yang.Yidentityref {
				enum = handleIdentity(dirEntry.Type)
			}
			fmt.Println(enum)
			// Check to see if this attribute is a key in a list
			if dirEntry.Parent.IsList() {
				keyNames := strings.Split(dirEntry.Parent.Key, " ")
				itemPathParts := strings.Split(itemPath, "/")
				attrName := itemPathParts[len(itemPathParts)-1]
				for _, k := range keyNames {
					if strings.EqualFold(attrName, k) {
						tObj.IsAKey = true
						break
					}
				}
			}
			if parentState == yang.TSFalse {
				//leafMap, ok := readOnlyPaths[parentPath]
				//if !ok {
				//	leafMap = make(ReadOnlySubPathMap)
				//	readOnlyPaths[parentPath] = leafMap
				//}
				//leafMap[strings.Replace(itemPath, parentPath, "", 1)] = tObj
			} else if dirEntry.Config == yang.TSFalse {
				//leafMap := make(ReadOnlySubPathMap)
				//leafMap["/"] = tObj
				//readOnlyPaths[itemPath] = leafMap
			} else {
				ranges := make([]string, 0)
				for _, r := range dirEntry.Type.Range {
					ranges = append(ranges, fmt.Sprintf("%v", r))
				}
				lengths := make([]string, 0)
				for _, l := range dirEntry.Type.Length {
					lengths = append(lengths, fmt.Sprintf("%v", l))
				}
				rwElem := admin.ReadWritePath{
					Path:        "",
					ValueType:   tObj.ValueType,
					Units:       tObj.Units,
					Description: tObj.Description,
					Mandatory:   dirEntry.Mandatory == yang.TSTrue,
					Default:     dirEntry.Default,
					Range:       ranges,
					Length:      lengths,
					TypeOpts:    nil,
					IsAKey:      false,
					AttrName:    "",
				}
				readWritePaths = append(readWritePaths, &rwElem)
			}
		} else if dirEntry.IsContainer() {
			if dirEntry.Config == yang.TSFalse || parentState == yang.TSFalse {
				subpathPfx := subpathPrefix
				if parentState == yang.TSFalse {
					subpathPfx = itemPath[len(parentPath):]
				}
				subPaths, _, err := extractPaths(dirEntry, yang.TSFalse, itemPath, subpathPfx)
				if err != nil {
					return nil, nil, err
				}
				subPathsMap := make([]*admin.ReadOnlySubPath, 0)
				for _, v := range subPaths {
					for _, u := range v.SubPath {
						subPathsMap = append(subPathsMap, u)
					}
				}
				continue
			}
			readOnlyPathsTemp, readWritePathTemp, err := extractPaths(dirEntry, dirEntry.Config, itemPath, "")
			if err != nil {
				return nil, nil, err
			}
			for k, v := range readOnlyPathsTemp {
				readOnlyPaths[k] = v
			}
			for k, v := range readWritePathTemp {
				readWritePaths[k] = v
			}
		} else if dirEntry.IsList() {
			itemPath = formatName(dirEntry, true, parentPath, subpathPrefix)
			if dirEntry.Config == yang.TSFalse || parentState == yang.TSFalse {
				subpathPfx := subpathPrefix
				if parentState == yang.TSFalse {
					subpathPfx = itemPath[len(parentPath):]
				}
				subPaths, _, err := extractPaths(dirEntry, yang.TSFalse, parentPath, subpathPfx)
				if err != nil {
					return nil, nil, err
				}
				subPathsMap := make([]*admin.ReadOnlySubPath, 0)
				for _, v := range subPaths {
					for _, v := range v.SubPath {
						subPathsMap = append(subPathsMap, v)
					}
				}
				//readOnlyPaths = append(readOnlyPaths, subPathsMap)
				continue
			}
			readOnlyPathsTemp, readWritePathsTemp, err := extractPaths(dirEntry, dirEntry.Config, itemPath, "")
			if err != nil {
				return nil, nil, err
			}
			for k, v := range readOnlyPathsTemp {
				readOnlyPaths[k] = v
			}
			for k, v := range readWritePathsTemp {
				readWritePaths[k] = v
				// Need to copy the index of the list across to the RO list too
				//roIdxName := k[:strings.LastIndex(k, "/")]
				//roIdxSubPath := k[strings.LastIndex(k, "/"):]
				//indices, _ := ExtractIndexNames(itemPath[strings.LastIndex(itemPath, "/"):])
				//isIdxAttr := false
				//for _, idx := range indices {
				//	if roIdxSubPath == fmt.Sprintf("/%s", idx) {
				//		isIdxAttr = true
				//	}
				//}
				//if roIdxName == itemPath && isIdxAttr {
				//	roIdx := ReadOnlyAttrib{
				//		ValueType:   v.ValueType,
				//		Description: v.Description,
				//		Units:       v.Units,
				//	}
				//	readOnlyPaths[roIdxName] = make(map[string]ReadOnlyAttrib)
				//	readOnlyPaths[roIdxName][roIdxSubPath] = roIdx
				//}
			}

		} else if dirEntry.IsChoice() || dirEntry.IsCase() {
			// Recurse down through Choice and Case
			readOnlyPathsTemp, readWritePathsTemp, err := extractPaths(dirEntry, dirEntry.Config, parentPath, "")
			if err != nil {
				return nil, nil, err
			}
			for k, v := range readOnlyPathsTemp {
				readOnlyPaths[k] = v
			}
			for k, v := range readWritePathsTemp {
				readWritePaths[k] = v
			}
		} else {
			log.Warnf("Unexpected type of leaf for %s %v", itemPath, dirEntry)
		}
	}

	return readOnlyPaths, readWritePaths, nil
}

func GetPathValues(prefixPath string, genericJSON []byte) ([]*configapi.PathValue, error) {
	var f interface{}
	err := json.Unmarshal(genericJSON, &f)
	if err != nil {
		return nil, err
	}
	if fAsMap, ok := f.(map[string]interface{}); ok {
		if fResult, isResult := fAsMap["result"]; isResult {
			if fResultAsMap, ok := fResult.([]interface{}); ok {
				f = fResultAsMap[0]
			}
		}
	}
	values, err := extractValuesWithPaths(f, removeIndexNames(prefixPath))
	if err != nil {
		return nil, fmt.Errorf("error decomposing JSON %v", err)
	}
	return values, nil
}

// extractValuesIntermediate recursively walks a JSON tree to create a flat set
// of paths and values.
func extractValuesWithPaths(f interface{}, parentPath string) ([]*configapi.PathValue, error) {
	changes := make([]*configapi.PathValue, 0)

	switch value := f.(type) {
	case map[string]interface{}:
		mapChanges, err := handleMap(value, parentPath)
		if err != nil {
			return nil, err
		}
		changes = append(changes, mapChanges...)

	case []interface{}:
		indexNames := indicesOfPath(parentPath)
		// Iterate through to look for indexes first
		for idx, v := range value {
			indices := make([]indexValue, 0)
			nonIndexPaths := make([]string, 0)
			objs, err := extractValuesWithPaths(v, fmt.Sprintf("%s[%d]", parentPath, idx))
			if err != nil {
				return nil, err
			}
			for _, obj := range objs {
				isIndex := false
				for i, idxName := range indexNames {
					if removePathIndices(obj.Path) == fmt.Sprintf("%s/%s", removePathIndices(parentPath), idxName) {
						indices = append(indices, indexValue{name: idxName, value: &obj.Value, order: i})
						isIndex = true
						break
					}
				}
				if !isIndex {
					nonIndexPaths = append(nonIndexPaths, obj.Path)
				}
			}
			sort.Slice(indices, func(i, j int) bool {
				return indices[i].order < indices[j].order
			})
			// Now we have indices, need to go through again
			for _, obj := range objs {
				for _, nonIdxPath := range nonIndexPaths {
					if obj.Path == nonIdxPath {
						suffixLen := prefixLength(obj.Path, parentPath)
						obj.Path, err = replaceIndices(obj.Path, suffixLen, indices)
						if err != nil {
							return nil, fmt.Errorf("error replacing indices in %s %v", obj.Path, err)
						}
						changes = append(changes, obj)
					}
				}
			}
		}
	default:
		attr, err := handleAttribute(value, parentPath)
		if err != nil {
			return nil, fmt.Errorf("error handling json attribute value %v. Parent %s. #RO:%d #RW:%d %s",
				value, parentPath, len(roPathMap), len(rwPathMap), err.Error())
		}
		if attr != nil {
			changes = append(changes, attr)
		}
	}

	return changes, nil
}

func handleMap(value map[string]interface{}, parentPath string) ([]*configapi.PathValue, error) {
	changes := make([]*configapi.PathValue, 0)

	for key, v := range value {
		objs, err := extractValuesWithPaths(v, fmt.Sprintf("%s/%s", parentPath, stripNamespace(key)))
		if err != nil {
			return nil, err
		}
		if len(objs) > 0 {
			firstType := (objs[0].Value).Type
			matching := true
			for _, o := range objs {
				// In a leaf list all value types have to match
				if o.Value.Type != firstType {
					// Not a leaf list
					matching = false
					break
				}
			}
			if !matching {
				changes = append(changes, objs...)
			} else {
				switch (objs[0].Value).Type {
				case configapi.ValueType_LEAFLIST_INT:
					llVals := make([]int64, 0)
					var width configapi.Width
					for _, obj := range objs {
						var llI []int64
						llI, width = (*configapi.TypedLeafListInt)(&obj.Value).List()
						llVals = append(llVals, llI...)
					}
					newCv := configapi.PathValue{Path: objs[0].Path, Value: *configapi.NewLeafListIntTv(llVals, width)}
					changes = append(changes, &newCv)
				case configapi.ValueType_LEAFLIST_STRING:
					llVals := make([]string, 0)
					for _, obj := range objs {
						llI := (*configapi.TypedLeafListString)(&obj.Value)
						llVals = append(llVals, llI.List()...)
					}
					newCv := configapi.PathValue{Path: objs[0].Path, Value: *configapi.NewLeafListStringTv(llVals)}
					changes = append(changes, &newCv)
				case configapi.ValueType_LEAFLIST_UINT:
					llVals := make([]uint64, 0)
					var width configapi.Width
					for _, obj := range objs {
						var llI []uint64
						llI, width = (*configapi.TypedLeafListUint)(&obj.Value).List()
						llVals = append(llVals, llI...)
					}
					newCv := configapi.PathValue{Path: objs[0].Path, Value: *configapi.NewLeafListUintTv(llVals, width)}
					changes = append(changes, &newCv)
				case configapi.ValueType_LEAFLIST_BOOL:
					llVals := make([]bool, 0)
					for _, obj := range objs {
						llI := (*configapi.TypedLeafListBool)(&obj.Value)
						llVals = append(llVals, llI.List()...)
					}
					newCv := configapi.PathValue{Path: objs[0].Path, Value: *configapi.NewLeafListBoolTv(llVals)}
					changes = append(changes, &newCv)
				case configapi.ValueType_LEAFLIST_BYTES:
					llVals := make([][]byte, 0)
					for _, obj := range objs {
						llI := (*configapi.TypedLeafListBytes)(&obj.Value)
						llVals = append(llVals, llI.List()...)
					}
					newCv := configapi.PathValue{Path: objs[0].Path, Value: *configapi.NewLeafListBytesTv(llVals)}
					changes = append(changes, &newCv)
				case configapi.ValueType_LEAFLIST_DECIMAL:
					llDigits := make([]int64, 0)
					var llPrecision uint8
					for _, obj := range objs {
						var digitsList []int64
						digitsList, llPrecision = (*configapi.TypedLeafListDecimal)(&obj.Value).List()
						llDigits = append(llDigits, digitsList...)
					}
					newCv := configapi.PathValue{Path: objs[0].Path, Value: *configapi.NewLeafListDecimalTv(llDigits, llPrecision)}
					changes = append(changes, &newCv)
				case configapi.ValueType_LEAFLIST_FLOAT:
					llVals := make([]float32, 0)
					for _, obj := range objs {
						llI := (*configapi.TypedLeafListFloat)(&obj.Value)
						llVals = append(llVals, llI.List()...)
					}
					newCv := configapi.PathValue{Path: objs[0].Path, Value: *configapi.NewLeafListFloatTv(llVals)}
					changes = append(changes, &newCv)
				default:
					// Not a leaf list
					changes = append(changes, objs...)
				}
			}
		}
	}
	return changes, nil
}

func handleAttribute(value interface{}, parentPath string) (*configapi.PathValue, error) {
	var modeltype configapi.ValueType
	var modelPath string
	var ok bool
	var pathElem *ReadWritePathElem
	var subPath *ReadOnlyAttrib
	var enum map[int]string
	var typeOpts []uint8
	var err error
	pathElem, modelPath, ok = findModelRwPathNoIndices(parentPath)
	if !ok {
		subPath, modelPath, ok = findModelRoPathNoIndices(parentPath)
		if !ok {
			if roPathMap == nil || rwPathMap == nil {
				// If RO paths was not given - then we assume this missing path was a RO path
				return nil, nil
			}
			return nil, fmt.Errorf("unable to locate %s in model", parentPath)
		}
		modeltype = subPath.ValueType
		enum = subPath.Enum
		if subPath.TypeOpts != nil {
			typeOpts = make([]uint8, len(subPath.TypeOpts))
			copy(typeOpts, subPath.TypeOpts)
		}
	} else {
		modeltype = pathElem.ValueType
		enum = pathElem.Enum
		if pathElem.TypeOpts != nil {
			typeOpts = make([]uint8, len(pathElem.TypeOpts))
			copy(typeOpts, pathElem.TypeOpts)
		}
	}
	var typedValue *configapi.TypedValue
	switch modeltype {
	case configapi.ValueType_STRING:
		var stringVal string
		switch valueTyped := value.(type) {
		case string:
			if len(enum) > 0 {
				stringVal, err = convertEnumIdx(valueTyped, enum, parentPath)
				if err != nil {
					return nil, err
				}
			} else {
				stringVal = valueTyped
			}
		case float64:
			if len(enum) > 0 {
				stringVal, err = convertEnumIdx(fmt.Sprintf("%g", valueTyped), enum, parentPath)
				if err != nil {
					return nil, err
				}
			} else {
				stringVal = fmt.Sprintf("%g", valueTyped)
			}
		case bool:
			stringVal = fmt.Sprintf("%v", value)
		}
		typedValue = configapi.NewTypedValueString(stringVal)
	case configapi.ValueType_BOOL:
		typedValue = configapi.NewTypedValueBool(value.(bool))
	case configapi.ValueType_INT:
		var intVal int
		switch valueTyped := value.(type) {
		case string:
			intVal64, err := strconv.ParseInt(valueTyped, 10, int(typeOpts[0]))
			if err != nil {
				return nil, fmt.Errorf("error converting to %v %s", modeltype, valueTyped)
			}
			intVal = int(intVal64)
		case float64:
			intVal = int(valueTyped)
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		if len(typeOpts) == 0 {
			return nil, fmt.Errorf("expected INT to have a field width e.g. 8, 16, 32, 64")
		}
		typedValue = configapi.NewTypedValueInt(intVal, configapi.Width(typeOpts[0]))
	case configapi.ValueType_UINT:
		var uintVal uint
		switch valueTyped := value.(type) {
		case string:
			intVal, err := strconv.ParseInt(valueTyped, 10, int(typeOpts[0]))
			if err != nil {
				return nil, fmt.Errorf("error converting to %v %s", modeltype, valueTyped)
			}
			uintVal = uint(intVal)
		case float64:
			uintVal = uint(valueTyped)
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		if len(typeOpts) == 0 {
			return nil, fmt.Errorf("expected UINT to have a field width e.g. 8, 16, 32, 64")
		}
		typedValue = configapi.NewTypedValueUint(uintVal, configapi.Width(typeOpts[0]))
	case configapi.ValueType_DECIMAL:
		var digits int64
		precision := typeOpts[0]
		switch valueTyped := value.(type) {
		case float64:
			digits = int64(valueTyped * math.Pow(10, float64(precision)))
		case string:
			floatVal, err := strconv.ParseFloat(valueTyped, 64)
			if err != nil {
				return nil, fmt.Errorf("error converting string to float %v", err)
			}
			digits = int64(floatVal * math.Pow(10, float64(precision)))
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		if len(typeOpts) == 0 {
			return nil, fmt.Errorf("expected DECIMAL to have a precision")
		}
		typedValue = configapi.NewTypedValueDecimal(digits, precision)
	case configapi.ValueType_BYTES:
		var dstBytes []byte
		switch valueTyped := value.(type) {
		case string:
			// Values should be base64
			dstBytes, err = base64.StdEncoding.DecodeString(valueTyped)
			if err != nil {
				return nil, fmt.Errorf("expected binary value as base64. error decoding %s as base64 %v", valueTyped, err)
			}
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		typedValue = configapi.NewTypedValueBytes(dstBytes)
	default:
		typedValue, err = handleAttributeLeafList(modeltype, value)
		if err != nil {
			return nil, err
		}
	}
	return &configapi.PathValue{Path: modelPath, Value: *typedValue}, nil
}

// A continuation of handle attribute above
func handleAttributeLeafList(modeltype configapi.ValueType,
	value interface{}) (*configapi.TypedValue, error) {

	var typedValue *configapi.TypedValue

	switch modeltype {
	case configapi.ValueType_LEAFLIST_INT:
		var leafvalue int64
		switch valueTyped := value.(type) {
		case float64:
			leafvalue = int64(valueTyped)
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		typedValue = configapi.NewLeafListIntTv([]int64{leafvalue}, configapi.WidthThirtyTwo) // TODO: lookup width from model
	case configapi.ValueType_LEAFLIST_UINT:
		var leafvalue uint64
		switch valueTyped := value.(type) {
		case float64:
			leafvalue = uint64(valueTyped)
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		typedValue = configapi.NewLeafListUintTv([]uint64{leafvalue}, configapi.WidthThirtyTwo) // TODO: lookup width from model
	case configapi.ValueType_LEAFLIST_FLOAT:
		var leafvalue float32
		switch valueTyped := value.(type) {
		case float64:
			leafvalue = float32(valueTyped)
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		typedValue = configapi.NewLeafListFloatTv([]float32{leafvalue})
	case configapi.ValueType_LEAFLIST_STRING:
		var leafvalue string
		switch valueTyped := value.(type) {
		case string:
			leafvalue = valueTyped
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		typedValue = configapi.NewLeafListStringTv([]string{leafvalue})
	case configapi.ValueType_LEAFLIST_BOOL:
		var leafvalue bool
		switch valueTyped := value.(type) {
		case bool:
			leafvalue = valueTyped
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		typedValue = configapi.NewLeafListBoolTv([]bool{leafvalue})
	case configapi.ValueType_LEAFLIST_BYTES:
		var leafvalue []byte
		var err error
		switch valueTyped := value.(type) {
		case string:
			// Values should be base64
			leafvalue, err = base64.StdEncoding.DecodeString(valueTyped)
			if err != nil {
				return nil, fmt.Errorf("expected binary value as base64. error decoding %s as base64 %v", valueTyped, err)
			}
		default:
			return nil, fmt.Errorf("unhandled conversion to %v %s", modeltype, valueTyped)
		}
		typedValue = configapi.NewLeafListBytesTv([][]byte{leafvalue})
	default:
		return nil, fmt.Errorf("unhandled conversion to %v", modeltype)
	}
	return typedValue, nil
}

func findModelRwPathNoIndices(searchpath string) (*ReadWritePathElem, string, bool) {
	searchpath = removeDoubleSlash(searchpath)
	searchpathNoIndices := removePathIndices(searchpath)
	for path, value := range rwPathMap {
		if removePathIndices(path) == searchpathNoIndices {
			pathWithNumericalIdx, err := insertNumericalIndices(path, searchpath)
			if err != nil {
				return nil, fmt.Sprintf("could not replace wildcards in model path with numerical ids %v", err), false
			}
			return &value, pathWithNumericalIdx, true
		}
	}
	return nil, "", false
}

func findModelRoPathNoIndices(searchpath string) (*ReadOnlyAttrib, string, bool) {
	searchpathNoIndices := removePathIndices(searchpath)
	for path, value := range roPathMap {
		for subpath, subpathValue := range value {
			var fullpath string
			if subpath == "/" {
				fullpath = path
			} else {
				fullpath = fmt.Sprintf("%s%s", path, subpath)
			}
			if removePathIndices(fullpath) == searchpathNoIndices {
				pathWithNumericalIdx, err := insertNumericalIndices(fullpath, searchpath)
				if err != nil {
					return nil, fmt.Sprintf("could not replace wildcards in model path with numerical ids %v", err), false
				}
				return &subpathValue, pathWithNumericalIdx, true
			}
		}
	}
	return nil, "", false
}

// For RW paths
func indicesOfPath(searchpath string) []string {
	searchpathNoIndices := removePathIndices(searchpath)
	// First search through the RW paths
	for p := range rwPathMap {
		pathNoIndices := removePathIndices(p)
		// Find a short path
		if pathNoIndices[:strings.LastIndex(pathNoIndices, slash)] == searchpathNoIndices {
			idxNames, _ := ExtractIndexNames(p)
			return idxNames
		}
	}

	// If not found then search through the RO paths
	for p, value := range roPathMap {
		for subpath := range value {
			var fullpath string
			if subpath == "/" {
				fullpath = p
			} else {
				fullpath = fmt.Sprintf("%s%s", p, subpath)
			}
			pathNoIndices := removePathIndices(fullpath)
			// Find a short path
			if pathNoIndices[:strings.LastIndex(pathNoIndices, slash)] == searchpathNoIndices {
				idxNames, _ := ExtractIndexNames(fullpath)
				return idxNames
			}
		}
	}

	return []string{}
}

// YGOT does not handle namespaces, so there is no point in us maintaining them
// They may come from the southbound or northbound in a JSON payload though, so
// we have to be able to deal with them
func stripNamespace(path string) string {
	pathParts := strings.Split(path, "/")
	for idx, pathPart := range pathParts {
		colonPos := strings.Index(pathPart, colon)
		if colonPos > 0 {
			pathParts[idx] = pathPart[colonPos+1:]
		}
	}
	return strings.Join(pathParts, "/")
}

func insertNumericalIndices(modelPath string, jsonPath string) (string, error) {
	jsonParts := strings.Split(jsonPath, slash)
	modelParts := strings.Split(modelPath, slash)
	if len(modelParts) != len(jsonParts) {
		return "", fmt.Errorf("strings must have the same number of / characters %d!=%d", len(modelParts), len(jsonParts))
	}
	for idx, jsonPart := range jsonParts {
		brktIdx := strings.LastIndex(jsonPart, bracketsq)
		if brktIdx > 0 {
			modelParts[idx] = strings.ReplaceAll(modelParts[idx], "*", jsonPart[brktIdx+1:len(jsonPart)-1])
		}
	}

	return strings.Join(modelParts, "/"), nil
}

func prefixLength(objPath string, parentPath string) int {
	objPathParts := strings.Split(objPath, "/")
	parentPathParts := strings.Split(parentPath, "/")
	return len(strings.Join(objPathParts[:len(parentPathParts)], "/"))
}

func ConvertRoPath(path string, psm ReadOnlySubPathMap) *admin.ReadOnlyPath {
	sm := make([]*admin.ReadOnlySubPath, 0)

	for k, s := range psm {
		tos := make([]uint64, 0, len(s.TypeOpts))
		for _, to := range s.TypeOpts {
			tos = append(tos, uint64(to))
		}

		sm = append(sm, &admin.ReadOnlySubPath{
			SubPath:     k,
			ValueType:   s.ValueType,
			Description: s.Description,
			TypeOpts:    tos,
			IsAKey:      s.IsAKey,
			AttrName:    s.AttrName,
		})
	}

	return &admin.ReadOnlyPath{
		Path:    path,
		SubPath: sm,
	}
}

func ConvertRwPath(path string, pe ReadWritePathElem) *admin.ReadWritePath {
	tos := make([]uint64, 0, len(pe.TypeOpts))
	for _, to := range pe.TypeOpts {
		tos = append(tos, uint64(to))
	}
	var firstDefault string // TODO convert the Default to []string in onos-api
	if len(pe.Default) > 0 {
		firstDefault = pe.Default[0]
	}
	return &admin.ReadWritePath{
		Path:        path,
		ValueType:   pe.ValueType,
		Units:       pe.Units,
		Description: pe.Description,
		Mandatory:   pe.Mandatory,
		Default:     firstDefault,
		Range:       pe.Range,
		Length:      pe.Length,
		TypeOpts:    tos,
		IsAKey:      pe.IsAKey,
		AttrName:    pe.AttrName,
	}
}

// There might not be an index for everything
func replaceIndices(path string, ignoreAfter int, indices []indexValue) (string, error) {
	ignored := path[ignoreAfter:]
	pathParts := strings.Split(path[:ignoreAfter], bracketsq)
	idxOffset := len(pathParts) - len(indices) - 1

	// Range in reverse
	for i := len(pathParts) - 1; i > 0; i-- {
		pathPart := pathParts[i]
		eqIdx := strings.LastIndex(pathPart, equals)
		if eqIdx > 0 {
			closeIdx := strings.LastIndex(pathPart, brktclose)
			idxName := pathPart[:eqIdx]
			var actualValue string
			if i-idxOffset-1 < 0 {
				continue
			}
			index := indices[i-idxOffset-1]
			if index.name != idxName {
				//continue
				return "", fmt.Errorf("unexpected index name %s", index.name)
			}
			switch index.value.Type {
			case configapi.ValueType_STRING:
				actualValue = string(index.value.Bytes)
			case configapi.ValueType_UINT, configapi.ValueType_INT:
				actualValue = fmt.Sprintf("%d", (*configapi.TypedUint)(index.value).Uint())
			}
			pathParts[i] = fmt.Sprintf("%s=%s%s", idxName, actualValue, pathPart[closeIdx:])
		}
	}

	return fmt.Sprintf("%s%s", strings.Join(pathParts, bracketsq), ignored), nil
}

func convertEnumIdx(valueTyped string, enum map[int]string,
	parentPath string) (string, error) {
	var stringVal string
	for k, v := range enum {
		if v == valueTyped {
			stringVal = valueTyped
			break
		} else if fmt.Sprintf("%d", k) == valueTyped {
			stringVal = v
			break
		}
	}
	if stringVal == "" {
		enumOpts := make([]string, len(enum)*2)
		i := 0
		for k, v := range enum {
			enumOpts[i*2] = fmt.Sprintf("%d", k)
			enumOpts[i*2+1] = v
			i++
		}
		return "", fmt.Errorf("value %s for %s does not match any enumerated value %s",
			valueTyped, parentPath, strings.Join(enumOpts, ";"))
	}
	return stringVal, nil
}

// for a path like
// "/interfaces/interface[name=eth1]/subinterfaces/subinterface[index=120]/config/description",
// Remove the "name=" and "index="
func removeIndexNames(prefixPath string) string {
	splitPath := strings.Split(prefixPath, equals)
	for i, pathPart := range splitPath {
		if i < len(splitPath)-1 {
			lastBrktPos := strings.LastIndex(pathPart, bracketsq)
			splitPath[i] = pathPart[:lastBrktPos] + "["
		}
	}

	return strings.Join(splitPath, "")
}

func removePathIndices(path string) string {
	indices := rOnIndex.FindAllStringSubmatch(path, -1)
	for _, i := range indices {
		path = strings.Replace(path, i[0], "", 1)
	}
	return path
}

func removeDoubleSlash(path string) string {
	if strings.HasPrefix(path, "//") {
		return path[1:]
	}
	return path
}
