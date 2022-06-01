// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	configapi "github.com/onosproject/onos-api/go/onos/config/v2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/goyang/pkg/yang"
)

var log = logging.GetLogger("utils", "path")

// MatchOnIndex - regexp to find indices in paths names
const MatchOnIndex = `(\[.*?]).*?`

// validPathRegexp - permissible values in paths
const validPathRegexp = `(/[a-zA-Z0-9:=\-\._[\]]+)+`

// IndexAllowedChars - regexp to restrict characters in index names
const IndexAllowedChars = `^([a-zA-Z0-9\*\-\._])+$`

// ReadOnlyAttrib is the known metadata about a Read Only leaf
// Deprecated - remove me
type ReadOnlyAttrib struct {
	ValueType   configapi.ValueType
	TypeOpts    []uint8
	Description string
	Units       string
	Enum        map[int]string
	IsAKey      bool
	AttrName    string
}

// ReadOnlySubPathMap abstracts the read only subpath
// Deprecated - remove me
type ReadOnlySubPathMap map[string]ReadOnlyAttrib

// ReadOnlyPathMap abstracts the read only path
// Deprecated - remove me
type ReadOnlyPathMap map[string]ReadOnlySubPathMap

// Deprecated - remove me
var rIndexAllowedChars = regexp.MustCompile(IndexAllowedChars)

// JustPaths extracts keys from a read only path map
// Deprecated - remove me
func (ro ReadOnlyPathMap) JustPaths() []string {
	keys := make([]string, 0)
	for k, subPaths := range ro {
		for k1 := range subPaths {
			if k1 == "/" {
				keys = append(keys, k)
			} else {
				keys = append(keys, k+k1)
			}
		}
	}
	return keys
}

// ReadWritePathElem holds data about a leaf or container
// Deprecated - remove me
type ReadWritePathElem struct {
	ReadOnlyAttrib
	Mandatory bool
	Default   []string
	Range     []string
	Length    []string
}

// ReadWritePathMap is a map of ReadWrite paths their metadata
// Deprecated - remove me
type ReadWritePathMap map[string]ReadWritePathElem

// ExtractIndexNames - get an ordered array of index names and index values
// Deprecated - remove me
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

// Deprecated - remove me
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

//Paths extract the read only path up to the first read only container
// Deprecated - remove me
func Paths(readOnly ReadOnlyPathMap) []string {
	keys := make([]string, 0, len(readOnly))
	for k := range readOnly {
		keys = append(keys, k)
	}
	return keys
}

func toValueType(entry *yang.YangType, isLeafList bool) (configapi.ValueType, []uint64, error) {
	switch entry.Kind.String() {
	case "int8", "int16", "int32", "int64":
		width := extractIntegerWidth(entry.Kind.String())
		if isLeafList {
			return configapi.ValueType_LEAFLIST_INT, []uint64{uint64(width)}, nil
		}
		return configapi.ValueType_INT, []uint64{uint64(width)}, nil
	case "uint8", "uint16", "uint32", "uint64":
		width := extractIntegerWidth(entry.Kind.String())
		if isLeafList {
			return configapi.ValueType_LEAFLIST_UINT, []uint64{uint64(width)}, nil
		}
		return configapi.ValueType_UINT, []uint64{uint64(width)}, nil
	case "decimal64":
		if isLeafList {
			return configapi.ValueType_LEAFLIST_DECIMAL, []uint64{uint64(entry.FractionDigits)}, nil
		}
		return configapi.ValueType_DECIMAL, []uint64{uint64(entry.FractionDigits)}, nil
	case "string", "enumeration", "leafref", "identityref", "union", "instance-identifier":
		if isLeafList {
			return configapi.ValueType_LEAFLIST_STRING, nil, nil
		}
		return configapi.ValueType_STRING, nil, nil
	case "boolean":
		if isLeafList {
			return configapi.ValueType_LEAFLIST_BOOL, nil, nil
		}
		return configapi.ValueType_BOOL, nil, nil
	case "bits", "binary":
		if isLeafList {
			return configapi.ValueType_LEAFLIST_BYTES, nil, nil
		}
		return configapi.ValueType_BYTES, nil, nil
	case "empty":
		return configapi.ValueType_EMPTY, nil, nil
	default:
		return configapi.ValueType_EMPTY, nil,
			errors.NewInvalid("unhandled type in ModelPlugin %s %s %s",
				entry.Name, entry.Kind.String(), entry.Type)
	}
}

func extractIntegerWidth(typeName string) configapi.Width {
	switch typeName {
	case "int8", "uint8":
		return configapi.WidthEight
	case "int16", "uint16":
		return configapi.WidthSixteen
	case "int32", "uint32":
		return configapi.WidthThirtyTwo
	case "int64", "uint64", "counter64":
		return configapi.WidthSixtyFour
	default:
		return configapi.WidthThirtyTwo
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
