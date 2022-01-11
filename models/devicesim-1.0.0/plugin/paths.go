// Copyright 2022-present Open Networking Foundation.
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

package main

import (
	"fmt"
	"github.com/onosproject/config-models/models/devicesim-1.0.0/api"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	devicechange "github.com/onosproject/onos-api/go/onos/config/change/device"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/openconfig/goyang/pkg/yang"
	"regexp"
	"sort"
	"strings"
)

const rootPath = "Device"

// regexp to find indices in paths names
const matchOnIndex = `(\[.*?]).*?`

var rOnIndex = regexp.MustCompile(matchOnIndex)

// ReadOnlyAttrib is the known metadata about a Read Only leaf
type ReadOnlyAttrib struct {
	ValueType   devicechange.ValueType
	TypeOpts    []uint8
	Description string
	Units       string
	Enum        map[int]string
	IsAKey      bool
	AttrName    string
}

// ReadOnlySubPathMap abstracts the read only subpath
type ReadOnlySubPathMap map[string]ReadOnlyAttrib

// ReadOnlyPathMap abstracts the read only path
type ReadOnlyPathMap map[string]ReadOnlySubPathMap

// ReadWritePathElem holds data about a leaf or container
type ReadWritePathElem struct {
	ReadOnlyAttrib
	Mandatory bool
	Default   string
	Range     []string
	Length    []string
}

// ReadWritePathMap is a map of ReadWrite paths a their metadata
type ReadWritePathMap map[string]ReadWritePathElem

func ExtractPaths() ([]*admin.ReadOnlyPath, []*admin.ReadWritePath) {
	entries, err := api.UnzipSchema()
	if err != nil {
		log.Errorf("Unable to extract model schema: %+v", err)
		return nil, nil
	}

	ro, rw := extractPaths(entries[rootPath], yang.TSUnset, "", "")

	roPaths = make([]*admin.ReadOnlyPath, 0, len(ro))
	for k, p := range ro {
		roPaths = append(roPaths, convertRoPath(k, p))
	}
	log.Info("Read only paths are: %+v", roPaths)

	rwPaths = make([]*admin.ReadWritePath, 0, len(rw))
	for k, p := range rw {
		rwPaths = append(rwPaths, convertRwPath(k, p))
	}
	log.Info("Read write paths are: %+v", rwPaths)

	return roPaths, rwPaths
}

func convertRoPath(path string, psm ReadOnlySubPathMap) *admin.ReadOnlyPath {
	sm := make([]*admin.ReadOnlySubPath, 0, 0)

	for k, s := range psm {
		sm = append(sm, &admin.ReadOnlySubPath{
			SubPath:              k,
			ValueType:            s.ValueType,
			// TODO: add other ReadOnlyAttrib fields to the API
		})
	}

	return &admin.ReadOnlyPath{
		Path:                 path,
		SubPath:              sm,
	}
}

func convertRwPath(path string, pe ReadWritePathElem) *admin.ReadWritePath {
	return &admin.ReadWritePath{
		Path:                 path,
		ValueType:            pe.ValueType,
		Units:                pe.Units,
		Description:          pe.Description,
		Mandatory:            pe.Mandatory,
		Default:              pe.Default,
		Range:                pe.Range,
		Length:               pe.Length,
		// TODO: add other ReadOnlyAttrib fields to the API
	}
}

// extractPaths is a recursive function to extract a list of read only paths from a YGOT schema
func extractPaths(deviceEntry *yang.Entry, parentState yang.TriState, parentPath string, subpathPrefix string) (ReadOnlyPathMap, ReadWritePathMap) {
	readOnlyPaths := make(ReadOnlyPathMap)
	readWritePaths := make(ReadWritePathMap)
	for _, dirEntry := range deviceEntry.Dir {
		itemPath := formatName(dirEntry, false, parentPath, subpathPrefix)
		if dirEntry.IsLeaf() || dirEntry.IsLeafList() {
			// No need to recurse
			t, typeOpts, err := toValueType(dirEntry.Type, dirEntry.IsLeafList())
			tObj := ReadOnlyAttrib{
				ValueType:   t,
				TypeOpts:    typeOpts,
				Description: dirEntry.Description,
				Units:       dirEntry.Units,
				AttrName:    dirEntry.Name,
			}
			if err != nil {
				log.Errorf(err.Error())
			}
			var enum map[int]string
			if dirEntry.Type.Kind == yang.Yidentityref {
				enum = handleIdentity(dirEntry.Type)
			}
			tObj.Enum = enum
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
				leafMap, ok := readOnlyPaths[parentPath]
				if !ok {
					leafMap = make(ReadOnlySubPathMap)
					readOnlyPaths[parentPath] = leafMap
				}
				leafMap[strings.Replace(itemPath, parentPath, "", 1)] = tObj
			} else if dirEntry.Config == yang.TSFalse {
				leafMap := make(ReadOnlySubPathMap)
				leafMap["/"] = tObj
				readOnlyPaths[itemPath] = leafMap
			} else {
				ranges := make([]string, 0)
				for _, r := range dirEntry.Type.Range {
					ranges = append(ranges, fmt.Sprintf("%v", r))
				}
				lengths := make([]string, 0)
				for _, l := range dirEntry.Type.Length {
					lengths = append(lengths, fmt.Sprintf("%v", l))
				}
				rwElem := ReadWritePathElem{
					ReadOnlyAttrib: tObj,
					Mandatory:      dirEntry.Mandatory == yang.TSTrue,
					Default:        dirEntry.Default,
					Range:          ranges,
					Length:         lengths,
				}
				readWritePaths[itemPath] = rwElem
			}
		} else if dirEntry.IsContainer() {
			if dirEntry.Config == yang.TSFalse || parentState == yang.TSFalse {
				subpathPfx := subpathPrefix
				if parentState == yang.TSFalse {
					subpathPfx = itemPath[len(parentPath):]
				}
				subPaths, _ := extractPaths(dirEntry, yang.TSFalse, itemPath, subpathPfx)
				subPathsMap := make(ReadOnlySubPathMap)
				for _, v := range subPaths {
					for k, u := range v {
						subPathsMap[k] = u
					}
				}
				readOnlyPaths[itemPath] = subPathsMap
				continue
			}
			readOnlyPathsTemp, readWritePathTemp := extractPaths(dirEntry, dirEntry.Config, itemPath, "")
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
				subPaths, _ := extractPaths(dirEntry, yang.TSFalse, parentPath, subpathPfx)
				subPathsMap := make(ReadOnlySubPathMap)
				for _, v := range subPaths {
					for k, u := range v {
						subPathsMap[k] = u
					}
				}
				readOnlyPaths[itemPath] = subPathsMap
				continue
			}
			readOnlyPathsTemp, readWritePathsTemp := extractPaths(dirEntry, dirEntry.Config, itemPath, "")
			for k, v := range readOnlyPathsTemp {
				readOnlyPaths[k] = v
			}
			for k, v := range readWritePathsTemp {
				readWritePaths[k] = v
				// Need to copy the index of the list across to the RO list too
				roIdxName := k[:strings.LastIndex(k, "/")]
				roIdxSubPath := k[strings.LastIndex(k, "/"):]
				indices, _ := extractIndexNames(itemPath[strings.LastIndex(itemPath, "/"):])
				isIdxAttr := false
				for _, idx := range indices {
					if roIdxSubPath == fmt.Sprintf("/%s", idx) {
						isIdxAttr = true
					}
				}
				if roIdxName == itemPath && isIdxAttr {
					roIdx := ReadOnlyAttrib{
						ValueType:   v.ValueType,
						Description: v.Description,
						Units:       v.Units,
					}
					readOnlyPaths[roIdxName] = make(map[string]ReadOnlyAttrib)
					readOnlyPaths[roIdxName][roIdxSubPath] = roIdx
				}
			}

		} else if dirEntry.IsChoice() || dirEntry.IsCase() {
			// Recurse down through Choice and Case
			readOnlyPathsTemp, readWritePathsTemp := extractPaths(dirEntry, dirEntry.Config, parentPath, "")
			for k, v := range readOnlyPathsTemp {
				readOnlyPaths[k] = v
			}
			for k, v := range readWritePathsTemp {
				readWritePaths[k] = v
			}
		} else {
			log.Errorf("Unexpected type of leaf for %s %v", itemPath, dirEntry)
		}
	}
	return readOnlyPaths, readWritePaths
}

func toValueType(entry *yang.YangType, isLeafList bool) (devicechange.ValueType, []uint8, error) {
	switch entry.Kind.String() {
	case "int8", "int16", "int32", "int64":
		width := extractIntegerWidth(entry.Kind.String())
		if isLeafList {
			return devicechange.ValueType_LEAFLIST_INT, []uint8{uint8(width)}, nil
		}
		return devicechange.ValueType_INT, []uint8{uint8(width)}, nil
	case "uint8", "uint16", "uint32", "uint64":
		width := extractIntegerWidth(entry.Kind.String())
		if isLeafList {
			return devicechange.ValueType_LEAFLIST_UINT, []uint8{uint8(width)}, nil
		}
		return devicechange.ValueType_UINT, []uint8{uint8(width)}, nil
	case "decimal64":
		if isLeafList {
			return devicechange.ValueType_LEAFLIST_DECIMAL, []uint8{uint8(entry.FractionDigits)}, nil
		}
		return devicechange.ValueType_DECIMAL, []uint8{uint8(entry.FractionDigits)}, nil
	case "string", "enumeration", "leafref", "identityref", "union", "instance-identifier":
		if isLeafList {
			return devicechange.ValueType_LEAFLIST_STRING, nil, nil
		}
		return devicechange.ValueType_STRING, nil, nil
	case "boolean":
		if isLeafList {
			return devicechange.ValueType_LEAFLIST_BOOL, nil, nil
		}
		return devicechange.ValueType_BOOL, nil, nil
	case "bits", "binary":
		if isLeafList {
			return devicechange.ValueType_LEAFLIST_BYTES, nil, nil
		}
		return devicechange.ValueType_BYTES, nil, nil
	case "empty":
		return devicechange.ValueType_EMPTY, nil, nil
	default:
		return devicechange.ValueType_EMPTY, nil,
			errors.NewInvalid("unhandled type in ModelPlugin %s %s %s",
				entry.Name, entry.Kind.String(), entry.Type)
	}
}

func extractIntegerWidth(typeName string) devicechange.Width {
	switch typeName {
	case "int8", "uint8":
		return devicechange.WidthEight
	case "int16", "uint16":
		return devicechange.WidthSixteen
	case "int32", "uint32":
		return devicechange.WidthThirtyTwo
	case "int64", "uint64", "counter64":
		return devicechange.WidthSixtyFour
	default:
		return devicechange.WidthThirtyTwo
	}
}

func handleIdentity(yangType *yang.YangType) map[int]string {
	identityMap := make(map[int]string)
	identityMap[0] = "UNSET"
	for i, val := range yangType.IdentityBase.Values {
		identityMap[i+1] = val.Name
	}
	return identityMap
}

func extractIndexNames(path string) ([]string, []string) {
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

func formatName(dirEntry *yang.Entry, isList bool, parentPath string, subpathPrefix string) string {
	parentAndSubPath := parentPath
	if subpathPrefix != "/" {
		parentAndSubPath = fmt.Sprintf("%s%s", parentPath, subpathPrefix)
	}

	var name string
	if isList {
		//have to ensure index order is consistent where there's more than one
		keyParts := strings.Split(dirEntry.Key, " ")
		sort.Slice(keyParts, func(i, j int) bool {
			return keyParts[i] < keyParts[j]
		})
		name = fmt.Sprintf("%s/%s", parentAndSubPath, dirEntry.Name)
		for _, k := range keyParts {
			name += fmt.Sprintf("[%s=*]", k)
		}
	} else {
		name = fmt.Sprintf("%s/%s", parentAndSubPath, dirEntry.Name)
	}

	return name
}
