// Copyright 2020-present Open Networking Foundation.
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

package openapi_gen

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
	"sort"
	"strings"
)

var swagger openapi3.Swagger

var respGet200Desc = "GET OK 200"

func BuildOpenapi(yangSchema *ytypes.Schema) (*openapi3.Swagger, error) {
	topEntry := yangSchema.SchemaTree["Device"]
	paths, components, err := buildSchema(topEntry, yang.TSFalse, "")
	if err != nil {
		return nil, err
	}

	swagger = openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   "onos-config model plugin",
			Version: "1.0.0-oas3",
		},
		Paths:      paths,
		Components: *components,
	}

	if err := swagger.Validate(context.Background()); err != nil {
		return nil, err
	}

	swaggerLdr := openapi3.NewSwaggerLoader()
	if err = swaggerLdr.ResolveRefsIn(&swagger, nil); err != nil {
		return nil, fmt.Errorf("error on Resolving Refs %v", err)
	}

	return &swagger, nil
}

// buildSchema is a recursive function to extract a list of read only paths from a YGOT schema
func buildSchema(deviceEntry *yang.Entry, parentState yang.TriState, parentPath string) (openapi3.Paths, *openapi3.Components, error) {
	openapiPaths := make(openapi3.Paths)
	openapiComponents := openapi3.Components{
		Schemas: make(map[string]*openapi3.SchemaRef),
	}
	for _, dirEntry := range deviceEntry.Dir {
		itemPath := formatName(dirEntry, false, parentPath, "")
		if dirEntry.IsLeaf() || dirEntry.IsLeafList() {
			// No need to recurse
			var schemaVal *openapi3.Schema
			switch dirEntry.Type.Kind {
			case yang.Ystring, yang.Yunion, yang.Yleafref, yang.Yidentityref:
				schemaVal = openapi3.NewStringSchema()
			case yang.Ybool:
				schemaVal = openapi3.NewBoolSchema()
			case yang.Yuint8, yang.Yuint16, yang.Yint8, yang.Yint16:
				schemaVal = openapi3.NewIntegerSchema()
			case yang.Yuint32, yang.Yint32:
				schemaVal = openapi3.NewInt32Schema()
			case yang.Yuint64, yang.Yint64:
				schemaVal = openapi3.NewInt64Schema()
			default:
				return nil, nil, fmt.Errorf("unhandled leaf %v %s", dirEntry.Type.Kind, dirEntry.Type.Name)
			}
			schemaVal.Title = dirEntry.Name

			if dirEntry.IsLeaf() {
				openapiComponents.Schemas[toCamelCase(itemPath)] = &openapi3.SchemaRef{
					Value: schemaVal,
				}
			} else { // Leaflist
				arr := openapi3.NewSchema()
				arr.Type = "leaf-list"
				arr.Items = &openapi3.SchemaRef{
					Value: schemaVal,
				}
				arr.Title = fmt.Sprintf("leaf-list-%s", dirEntry.Name)
				openapiComponents.Schemas[toCamelCase(itemPath)] = &openapi3.SchemaRef{
					Value: arr,
				}
			}
		} else if dirEntry.IsContainer() {
			newPath := newPathItem(dirEntry, itemPath, parentPath)
			openapiPaths[itemPath] = newPath

			paths, components, err := buildSchema(dirEntry, dirEntry.Config, itemPath)
			if err != nil {
				return nil, nil, err
			}
			for k, v := range paths {
				openapiPaths[k] = v
			}

			schemaVal := openapi3.NewObjectSchema()
			schemaVal.Properties = make(map[string]*openapi3.SchemaRef)
			schemaVal.Title = toCamelCase(itemPath)
			asRef := &openapi3.SchemaRef{
				Value: schemaVal,
			}
			openapiComponents.Schemas[toCamelCase(itemPath)] = asRef
			respGet200 := openapi3.NewResponse()
			respGet200.Description = &respGet200Desc
			respGet200.Content = openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
				Ref:   fmt.Sprintf("#/components/schemas/%s", toCamelCase(itemPath)),
				Value: asRef.Value,
			})
			newPath.Get.AddResponse(200, respGet200)

			for k, v := range components.Schemas {
				switch v.Value.Type {
				case "array": // List as a child of container
					arrayObj := openapi3.NewArraySchema()
					arrayObj.Items = &openapi3.SchemaRef{
						Ref:   fmt.Sprintf("#/components/schemas/%s", k),
						Value: v.Value.Items.Value,
					}
					arrayObj.Title = v.Value.Title
					schemaVal.Properties[fmt.Sprintf("List%s", k)] = &openapi3.SchemaRef{
						Value: arrayObj,
					}
					openapiComponents.Schemas[k] = v.Value.Items
				case "object": // Container as a child of container
					openapiComponents.Schemas[k] = v
				case "string", "boolean", "integer": // leaf as a child of list
					schemaVal.Properties[v.Value.Title] = v
				case "leaf-list":
					v.Value.Type = "array"
					schemaVal.Properties[v.Value.Title] = v
				default:
					return nil, nil, fmt.Errorf("undhanled in container %s: %s", k, v.Value.Type)
				}

			}
		} else if dirEntry.IsList() {
			keys := strings.Split(dirEntry.Key, " ")
			listItemPath := itemPath
			for _, k := range keys {
				listItemPath += fmt.Sprintf("/{%s}", k)
			}
			openapiPaths[listItemPath] = newPathItem(dirEntry, itemPath, parentPath)

			paths, components, err := buildSchema(dirEntry, dirEntry.Config, itemPath)
			if err != nil {
				return nil, nil, err
			}
			for k, v := range paths {
				openapiPaths[k] = v
			}

			arr := openapi3.NewArraySchema()
			arr.Items = &openapi3.SchemaRef{
				Value: openapi3.NewObjectSchema(),
			}
			arr.Title = fmt.Sprintf("Item%s", toCamelCase(itemPath))
			openapiComponents.Schemas[toCamelCase(itemPath)] = &openapi3.SchemaRef{
				Value: arr,
			}

			respGet200 := openapi3.NewResponse()
			respGet200.Description = &respGet200Desc
			respGet200.Content = openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
				Ref:   fmt.Sprintf("#/components/schemas/%s", toCamelCase(itemPath)),
				Value: openapiComponents.Schemas[toCamelCase(itemPath)].Value,
			})
			openapiPaths[listItemPath].Get.AddResponse(200, respGet200)

			for k, v := range components.Schemas {
				switch v.Value.Type {
				case "array": // List as a child of list
					arrayObj := openapi3.NewArraySchema()
					arrayObj.Items = &openapi3.SchemaRef{
						Ref:   fmt.Sprintf("#/components/schemas/%s", k),
						Value: v.Value.Items.Value,
					}
					arrayObj.Title = fmt.Sprintf("List%s", k)
					arr.Items.Value.Properties[fmt.Sprintf("List%s", k)] = &openapi3.SchemaRef{
						Value: arrayObj,
					}
					openapiComponents.Schemas[k] = v.Value.Items
				case "object": // Container as a child of list
					if v.Value.Title != "" {
						arr.Items.Value.Properties[k] = &openapi3.SchemaRef{
							Ref:   fmt.Sprintf("#/components/schemas/%s", v.Value.Title),
							Value: v.Value,
						}
					}
					openapiComponents.Schemas[k] = v
				case "string", "boolean", "integer": // leaf as a child of list
					arr.Items.Value.Properties[v.Value.Title] = v
				default:
					return nil, nil, fmt.Errorf("unhandled in list %s: %s", k, v.Value.Type)
				}
			}
		} else {
			return nil, nil, fmt.Errorf("unhandled dirEntry.Type %v", dirEntry.Type)
		}
	}
	return openapiPaths, &openapiComponents, nil
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
		name = fmt.Sprintf("%s/%s[%s=*]", parentAndSubPath, dirEntry.Name, strings.Join(keyParts, " "))
	} else {
		name = fmt.Sprintf("%s/%s", parentAndSubPath, dirEntry.Name)
	}

	return name
}

func newPathItem(dirEntry *yang.Entry, itemPath string, parentPath string) *openapi3.PathItem {
	getOp := openapi3.NewOperation()
	getOp.Summary = fmt.Sprintf("GET %s Generated from YANG model", itemPath)
	getOp.OperationID = fmt.Sprintf("get%s", toCamelCase(itemPath))
	getOp.Tags = []string{toCamelCase(parentPath)}
	getOp.Responses = openapi3.NewResponses()

	parameters := make(openapi3.Parameters, 0)
	if dirEntry.IsList() {
		keys := strings.Split(dirEntry.Key, " ")
		for _, k := range keys {
			p := openapi3.ParameterRef{
				//Ref: k,
				Value: openapi3.NewPathParameter(k),
			}
			p.Value.Description = fmt.Sprintf("key for %s", dirEntry.Name)
			p.Value.Content = openapi3.NewContent()
			mt := openapi3.NewMediaType()
			mt.Schema = &openapi3.SchemaRef{
				Value: openapi3.NewStringSchema(),
			}
			p.Value.Content["text/plain; charset=utf-8"] = mt
			parameters = append(parameters, &p)
		}
	}

	newPath := openapi3.PathItem{
		Get:         getOp,
		Description: dirEntry.Description,
		Parameters:  parameters,
	}

	if dirEntry.Config != yang.TSFalse && dirEntry.Parent.Config != yang.TSFalse {
		deleteOp := openapi3.NewOperation()
		deleteOp.Summary = fmt.Sprintf("DELETE Generated from YANG model")
		deleteOp.OperationID = fmt.Sprintf("delete%s", toCamelCase(itemPath))
		deleteOp.Responses = openapi3.NewResponses()
		newPath.Delete = deleteOp

		postOp := openapi3.NewOperation()
		postOp.Summary = fmt.Sprintf("POST Generated from YANG model")
		postOp.OperationID = fmt.Sprintf("post%s", toCamelCase(itemPath))
		postOp.Responses = openapi3.NewResponses()
		newPath.Post = postOp
	}

	return &newPath
}

func toCamelCase(itemPath string) string {
	temp := strings.ReplaceAll(itemPath, "/", "_")
	return strcase.ToCamel(temp)
}
