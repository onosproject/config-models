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
	"os"
	"sort"
	"strings"
)

var swagger openapi3.Swagger

var respGet200Desc = "GET OK 200"
var pathPrefix string
var targetParameter *openapi3.ParameterRef

func BuildOpenapi(yangSchema *ytypes.Schema, modelType string, modelVersion string) (*openapi3.Swagger, error) {
	pathPrefix = fmt.Sprintf("/%s/v%s/{target}", strings.ToLower(modelType), modelVersion)
	targetParameter = targetParam()

	topEntry := yangSchema.SchemaTree["Device"]
	paths, components, err := buildSchema(topEntry, yang.TSFalse, pathPrefix)
	if err != nil {
		return nil, err
	}

	components.Parameters = make(map[string]*openapi3.ParameterRef)
	components.Parameters["target"] = targetParameter

	swagger = openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   fmt.Sprintf("%s onos-config model plugin", modelType),
			Version: modelVersion,
		},
		Paths:      paths,
		Components: *components,
	}

	if err := swagger.Validate(context.Background()); err != nil {
		return nil, err
	}

	swaggerLdr := openapi3.NewSwaggerLoader()
	if err = swaggerLdr.ResolveRefsIn(&swagger, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error on Resolving Refs %v\n", err)
	}

	return &swagger, nil
}

func targetParam() *openapi3.ParameterRef {

	stringContent := openapi3.NewContent()
	mt := openapi3.NewMediaType()
	mt.Schema = &openapi3.SchemaRef{
		Value: openapi3.NewStringSchema(),
	}
	stringContent["text/plain; charset=utf-8"] = mt

	targetParam := openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			Name:        "target",
			In:          "path",
			Description: "target (device in onos-config)",
			Required:    true,
			Example:     "internal",
			Content:     stringContent,
		},
	}

	return &targetParam
}

// buildSchema is a recursive function to extract a list of read only paths from a YGOT schema
func buildSchema(deviceEntry *yang.Entry, parentState yang.TriState, parentPath string) (openapi3.Paths, *openapi3.Components, error) {
	openapiPaths := make(openapi3.Paths)
	openapiComponents := openapi3.Components{
		Schemas:       make(map[string]*openapi3.SchemaRef),
		RequestBodies: make(map[string]*openapi3.RequestBodyRef),
	}
	for _, dirEntry := range deviceEntry.Dir {
		itemPath := formatName(dirEntry, false, parentPath)
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
		} else if dirEntry.IsContainer() || dirEntry.Kind == yang.ChoiceEntry || dirEntry.Kind == yang.CaseEntry {
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

			rbRef := &openapi3.RequestBodyRef{
				Value: openapi3.NewRequestBody().WithContent(
					openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Value: asRef.Value,
						Ref:   fmt.Sprintf("#/components/schemas/%s", toCamelCase(itemPath)),
					}),
				),
			}
			openapiComponents.RequestBodies[fmt.Sprintf("%sRequestBody", toCamelCase(itemPath))] = rbRef

			if newPath.Post.RequestBody.Ref != "" {
				newPath.Post.RequestBody.Value = rbRef.Value
			}

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

			for k, v := range components.RequestBodies {
				openapiComponents.RequestBodies[k] = v
			}
		} else if dirEntry.IsList() {
			keys := strings.Split(dirEntry.Key, " ")
			listItemPath := itemPath
			for _, k := range keys {
				listItemPath += fmt.Sprintf("/{%s}", k)
			}
			openapiPaths[listItemPath] = newPathItem(dirEntry, itemPath, listItemPath)

			paths, components, err := buildSchema(dirEntry, dirEntry.Config, listItemPath)
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
			asRef := &openapi3.SchemaRef{
				Value: arr,
			}
			openapiComponents.Schemas[toCamelCase(itemPath)] = asRef

			rbRef := &openapi3.RequestBodyRef{
				Value: openapi3.NewRequestBody().WithContent(
					openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Value: asRef.Value,
						Ref:   fmt.Sprintf("#/components/schemas/%s", toCamelCase(itemPath)),
					}),
				),
			}
			openapiComponents.RequestBodies[fmt.Sprintf("%sRequestBody", toCamelCase(itemPath))] = rbRef

			if openapiPaths[listItemPath].Post.RequestBody.Ref != "" {
				openapiPaths[listItemPath].Post.RequestBody.Value = rbRef.Value
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
			for k, v := range components.RequestBodies {
				openapiComponents.RequestBodies[k] = v
			}
		} else {
			return nil, nil, fmt.Errorf("unhandled dirEntry %v.Type %v", dirEntry.Name, dirEntry.Type)
		}
	}
	return openapiPaths, &openapiComponents, nil
}

func formatName(dirEntry *yang.Entry, isList bool, parentPath string) string {
	parentAndSubPath := parentPath
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
	getOp.Responses = make(openapi3.Responses)

	parameters := make(openapi3.Parameters, 0)
	pathKeys := strings.Split(parentPath, "/")
	firstTarget := true
	for _, pathKey := range pathKeys {
		if pathKey == "{target}" && firstTarget {
			targetParameterRef := openapi3.ParameterRef{
				Ref:   fmt.Sprintf("#/components/parameters/target"),
				Value: targetParameter.Value,
			}
			parameters = append(parameters, &targetParameterRef)
			firstTarget = false

		} else if strings.HasPrefix(pathKey, "{") && strings.HasSuffix(pathKey, "}") {
			k := pathKey[1 : len(pathKey)-1]
			p := openapi3.ParameterRef{
				//Ref: k,
				Value: openapi3.NewPathParameter(k),
			}
			p.Value.Description = fmt.Sprintf("key %s", pathKey)
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
	if dirEntry.Kind == yang.ChoiceEntry {
		newPath.Description = fmt.Sprintf("YANG Choice: %s", dirEntry.Name)
	} else if dirEntry.Kind == yang.CaseEntry {
		newPath.Description = fmt.Sprintf("YANG Choice Case: %s", dirEntry.Name)
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
		postOp.Responses = make(openapi3.Responses)
		postOp.Responses["201"] = &openapi3.ResponseRef{Value: openapi3.NewResponse().WithDescription("created")}
		postOp.RequestBody = &openapi3.RequestBodyRef{
			Ref: fmt.Sprintf("#/components/requestBodies/%sRequestBody", toCamelCase(itemPath)),
			// Value is filled in later
		}
		newPath.Post = postOp
	}

	return &newPath
}

func toCamelCase(itemPath string) string {
	temp := strings.ReplaceAll(itemPath, "/", "_")
	return strcase.ToCamel(temp)
}
