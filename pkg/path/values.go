/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package path

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	configapi "github.com/onosproject/onos-api/go/onos/config/v2"
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
	if prefixPath == "/" {
		prefixPath = ""
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
				value, parentPath, len(roPaths), len(rwPaths), err.Error())
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
	var pathElem *admin.ReadWritePath
	var subPath *admin.ReadOnlySubPath
	var enum map[int]string
	var typeOpts []uint64
	var err error
	pathElem, modelPath, ok = findModelRwPathNoIndices(parentPath)
	if !ok {
		subPath, modelPath, ok = findModelRoPathNoIndices(parentPath)
		if !ok {
			if roPaths == nil || rwPaths == nil {
				// If RO paths was not given - then we assume this missing pathWithIdx was a RO pathWithIdx
				return nil, nil
			}
			return nil, fmt.Errorf("unable to locate %s in model", parentPath)
		}
		modeltype = subPath.ValueType
		// enum = subPath.Enum  // TODO - fix this
		if subPath.TypeOpts != nil {
			typeOpts = make([]uint64, len(subPath.TypeOpts))
			copy(typeOpts, subPath.TypeOpts)
		}
	} else {
		modeltype = pathElem.ValueType
		//enum = pathElem.Enum // TODO - fix this
		if pathElem.TypeOpts != nil {
			typeOpts = make([]uint64, len(pathElem.TypeOpts))
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
		typedValue = configapi.NewTypedValueDecimal(digits, uint8(precision))
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

func findModelRwPathNoIndices(searchpath string) (*admin.ReadWritePath, string, bool) {
	searchpath = removeDoubleSlash(searchpath)
	searchpathNoIndices := removePathIndices(searchpath)
	for _, rwPath := range rwPaths {
		if removePathIndices(rwPath.Path) == searchpathNoIndices {
			pathWithNumericalIdx, err := insertNumericalIndices(rwPath.Path, searchpath)
			if err != nil {
				return nil, fmt.Sprintf("could not replace wildcards in model pathWithIdx with numerical ids %v", err), false
			}
			return rwPath, pathWithNumericalIdx, true
		}
	}
	return nil, "", false
}

func findModelRoPathNoIndices(searchpath string) (*admin.ReadOnlySubPath, string, bool) {
	searchpathNoIndices := removePathIndices(searchpath)
	for _, roPath := range roPaths {
		for _, subpathValue := range roPath.SubPath {
			var fullpath string
			if subpathValue.SubPath == "/" {
				fullpath = roPath.Path
			} else {
				fullpath = fmt.Sprintf("%s%s", roPath.Path, subpathValue.SubPath)
			}
			if removePathIndices(fullpath) == searchpathNoIndices {
				pathWithNumericalIdx, err := insertNumericalIndices(fullpath, searchpath)
				if err != nil {
					return nil, fmt.Sprintf("could not replace wildcards in model pathWithIdx with numerical ids %v", err), false
				}
				return subpathValue, pathWithNumericalIdx, true
			}
		}
	}
	return nil, "", false
}

// For RW paths
func indicesOfPath(searchpath string) []string {
	searchpathNoIndices := removePathIndices(searchpath)
	// First search through the RW paths
	for _, p := range roPaths {
		pathNoIndices := removePathIndices(p.Path)
		// Find a short pathWithIdx
		if pathNoIndices[:strings.LastIndex(pathNoIndices, slash)] == searchpathNoIndices {
			idxNames, _ := ExtractIndexNames(p.Path)
			return idxNames
		}
	}

	// If not found then search through the RO paths
	for _, value := range roPaths {
		for _, subpath := range value.SubPath {
			var fullpath string
			if subpath.SubPath == "/" {
				fullpath = value.Path
			} else {
				fullpath = fmt.Sprintf("%s%s", value.Path, subpath)
			}
			pathNoIndices := removePathIndices(fullpath)
			// Find a short pathWithIdx
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
			modelParts[idx] = strings.Replace(modelParts[idx], "*", jsonPart[brktIdx+1:len(jsonPart)-1], 1)
		}
	}

	return strings.Join(modelParts, "/"), nil
}

func prefixLength(objPath string, parentPath string) int {
	objPathParts := strings.Split(objPath, "/")
	parentPathParts := strings.Split(parentPath, "/")
	return len(strings.Join(objPathParts[:len(parentPathParts)], "/"))
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

// for a pathWithIdx like
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

// ExtractIndexNames - get an ordered array of index names and index values
func ExtractIndexNames(path string) ([]string, []string) {
	indexNames := make([]string, 0)
	indexValues := make([]string, 0)
	jsonMatches := rOnIndex.FindAllStringSubmatch(path, -1)
	for _, m := range jsonMatches {
		idxName := m[1][1:strings.LastIndex(m[1], "=")]
		indexNames = append(indexNames, idxName)
		idxValue := m[1][strings.LastIndex(m[1], "=")+1 : len(m[1])-1]
		indexValues = append(indexValues, idxValue)
	}
	return indexNames, indexValues
}
