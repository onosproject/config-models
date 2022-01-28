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
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	AdditionalPropertyTarget    = "AdditionalPropertyTarget"
	AdditionalPropertyUnchanged = "AdditionalPropertyUnchanged"
)

var swagger openapi3.Swagger

var respGet200Desc = "GET OK 200"
var pathPrefix string
var targetParameter *openapi3.ParameterRef

type ApiGenSettings struct {
	ModelType    string
	ModelVersion string
	Title        string
	Description  string
}

func (settings *ApiGenSettings) ApplyDefaults() {
	if settings.ModelType == "" {
		panic("ModelType not specified")
	}

	// Fill in defaults for any unset settings
	if settings.ModelVersion == "" {
		settings.ModelVersion = "0.0.1"
	}
	if settings.Title == "" {
		settings.Title = fmt.Sprintf("%s onos-config model plugin", settings.ModelType)
	}
	if settings.Description == "" {
		settings.Description = fmt.Sprintf("This OpenAPI 3 specification is generated from"+
			"%s onos-config model plugin", settings.ModelType)
	}
}

func BuildOpenapi(yangSchema *ytypes.Schema, settings *ApiGenSettings) (*openapi3.Swagger, error) {
	settings.ApplyDefaults()

	pathPrefix = fmt.Sprintf("/%s/v%s/{target}", strings.ToLower(settings.ModelType), settings.ModelVersion)
	targetParameter = targetParam()

	topEntry := yangSchema.SchemaTree["Device"]
	paths, components, err := buildSchema(topEntry, yang.TSFalse, "")
	if err != nil {
		return nil, err
	}

	components.Parameters = make(map[string]*openapi3.ParameterRef)
	components.Parameters["target"] = targetParameter

	swagger = openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:          settings.Title,
			Version:        settings.ModelVersion,
			TermsOfService: "https://opennetworking.org/wp-content/uploads/2019/02/ONF-Licensing-and-IPR-FAQ-2020-06.pdf",
			Contact: &openapi3.Contact{
				Name:  "Open Networking Foundation",
				URL:   "https://opennetworking.org",
				Email: "info@opennetworking.org",
			},
			License: &openapi3.License{
				Name: "LicenseRef-ONF-Member-1.0",
				URL:  "https://opennetworking.org/wp-content/uploads/2020/06/ONF-Member-Only-Software-License-v1.0.pdf",
			},
			Description: settings.Description,
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
			Content:     stringContent,
		},
	}

	return &targetParam
}

// add AdditionalProperties reference to target to a particular schema
func addAdditionalProperties(schemaVal *openapi3.Schema, name string) {
	schemaValAdditionalRef := openapi3.NewObjectSchema()
	schemaValAdditionalRef.Properties = make(map[string]*openapi3.SchemaRef)
	schemaValAdditionalRef.Title = "ref"
	schemaVal.AdditionalProperties = &openapi3.SchemaRef{
		Value: schemaValAdditionalRef,
		Ref:   fmt.Sprintf("#/components/schemas/%s", name),
	}
}

// buildSchema is a recursive function to extract a list of read only paths from a YGOT schema
func buildSchema(deviceEntry *yang.Entry, parentState yang.TriState, parentPath string) (openapi3.Paths, *openapi3.Components, error) {
	openapiPaths := make(openapi3.Paths)
	openapiComponents := openapi3.Components{
		Schemas:       make(map[string]*openapi3.SchemaRef),
		RequestBodies: make(map[string]*openapi3.RequestBodyRef),
	}

	// At the root of the API, add in the definition of "additionalPropertyTarget"
	if parentPath == "" {
		schemaValTarget := openapi3.NewObjectSchema()
		schemaValTarget.Title = "target"
		schemaValTarget.Type = "string"
		schemaValTarget.Description = "an override of the target (device)"
		schemaValTargetRef := &openapi3.SchemaRef{
			Value: schemaValTarget,
		}

		schemaValAddTarget := openapi3.NewObjectSchema()
		schemaValAddTarget.Properties = make(map[string]*openapi3.SchemaRef)
		schemaValAddTarget.Title = AdditionalPropertyTarget
		schemaValAddTarget.Description = "Optionally specify a target other than the default (only on PATCH method)"
		schemaValAddTargetRef := &openapi3.SchemaRef{
			Value: schemaValAddTarget,
		}
		schemaValAddTarget.Properties["target"] = schemaValTargetRef
		openapiComponents.Schemas[AdditionalPropertyTarget] = schemaValAddTargetRef

		schemaValUnchanged := openapi3.NewObjectSchema()
		schemaValUnchanged.Title = "unchanged"
		schemaValUnchanged.Type = "string"
		schemaValUnchanged.Description = "A comma seperated list of unchanged mandatory attribute names"
		schemaValUnchangedRef := &openapi3.SchemaRef{
			Value: schemaValUnchanged,
		}

		schemaValAddUnchanged := openapi3.NewObjectSchema()
		schemaValAddUnchanged.Properties = make(map[string]*openapi3.SchemaRef)
		schemaValAddUnchanged.Title = AdditionalPropertyUnchanged
		schemaValAddUnchanged.Description = "To optionally omit 'required' properties, add them to 'unchanged' list"
		schemaValAddUnchangedRef := &openapi3.SchemaRef{
			Value: schemaValAddUnchanged,
		}
		schemaValAddUnchanged.Properties["unchanged"] = schemaValUnchangedRef
		openapiComponents.Schemas[AdditionalPropertyUnchanged] = schemaValAddUnchangedRef
	}

	for _, dirEntry := range deviceEntry.Dir {
		itemPath := fmt.Sprintf("%s/%s", parentPath, dirEntry.Name)
		if dirEntry.IsLeaf() || dirEntry.IsLeafList() {
			// No need to recurse
			var schemaVal *openapi3.Schema
			switch dirEntry.Type.Kind {
			case yang.Ystring:
				schemaVal = openapi3.NewStringSchema()
				if dirEntry.Type.Length != nil {
					min, max, err := yangRange(dirEntry.Type.Length)
					if err != nil {
						return nil, nil, err
					}
					if min != nil {
						schemaVal.MinLength = uint64(*min)
					}
					if max != nil {
						v := uint64(*max)
						schemaVal.MaxLength = &v
					}
				}
				if dirEntry.Type.Pattern != nil && len(dirEntry.Type.Pattern) > 0 {
					// All we can do is take the first one
					schemaVal.Pattern = dirEntry.Type.Pattern[0]
				}
				if dirEntry.Type.Default != "" {
					schemaVal.Default = dirEntry.Type.Default
				}
			case yang.Yunion:
				schemaVal = openapi3.NewStringSchema()
				if dirEntry.Type.Default != "" {
					schemaVal.Default = dirEntry.Type.Default
				}
			case yang.Yleafref:
				schemaVal = openapi3.NewStringSchema()
				if dirEntry.Type.Default != "" {
					schemaVal.Default = dirEntry.Type.Default
				}
				if schemaVal.Extensions == nil {
					schemaVal.Extensions = make(map[string]interface{})
				}
				schemaVal.Extensions["x-leafref"] = dirEntry.Type.Path
			case yang.Yidentityref, yang.Yenum:
				schemaVal = openapi3.NewStringSchema()
				if dirEntry.Type.IdentityBase != nil {
					schemaVal.Enum = make([]interface{}, 0)
					for _, val := range dirEntry.Type.IdentityBase.Values {
						schemaVal.Enum = append(schemaVal.Enum, val.Name)
					}
				}
			case yang.Ybool:
				schemaVal = openapi3.NewBoolSchema()
				if dirEntry.Type.Default == "true" || dirEntry.Default == "true" {
					schemaVal.Default = true
				} else if dirEntry.Type.Default == "false" || dirEntry.Default == "false" {
					schemaVal.Default = false
				}
			case yang.Yuint8, yang.Yuint16, yang.Yint8, yang.Yint16, yang.Yuint32, yang.Yint32, yang.Yuint64, yang.Yint64, yang.Ydecimal64:
				switch dirEntry.Type.Kind {
				case yang.Yuint32, yang.Yint32:
					schemaVal = openapi3.NewInt32Schema()
				case yang.Yuint64, yang.Yint64:
					schemaVal = openapi3.NewInt64Schema()
				case yang.Ydecimal64:
					schemaVal = openapi3.NewFloat64Schema()
				default:
					schemaVal = openapi3.NewIntegerSchema()
				}
				def, err := yangDefault(dirEntry)
				if err != nil {
					return nil, nil, err
				}
				schemaVal.Default = def
				if dirEntry.Type.Range != nil {
					start, end, err := yangRange(dirEntry.Type.Range)
					if err != nil {
						return nil, nil, err
					}
					if start != nil {
						startFloat := float64(*start)
						schemaVal.Min = &startFloat
					}
					if end != nil {
						endFloat := (float64)(*end)
						schemaVal.Max = &endFloat
					}
				}
			case yang.Ybinary:
				schemaVal = openapi3.NewBytesSchema()
				if dirEntry.Type.Length != nil {
					min, max, err := yangRange(dirEntry.Type.Length)
					if err != nil {
						return nil, nil, err
					}
					if min != nil {
						schemaVal.MinLength = uint64(*min)
					}
					if max != nil {
						v := uint64(*max)
						schemaVal.MaxLength = &v
					}
				}
				if dirEntry.Type.Default != "" {
					schemaVal.Default = dirEntry.Type.Default
				}
			case yang.Yempty:
				schemaVal = openapi3.NewStringSchema()
				var emptylen uint64 = 0
				schemaVal.MaxLength = &emptylen
			default:
				return nil, nil, fmt.Errorf("unhandled leaf %v %s", dirEntry.Type.Kind, dirEntry.Type.Name)
			}
			schemaVal.Title = dirEntry.Name
			schemaVal.Description = dirEntry.Description
			if dirEntry.Mandatory.Value() {
				schemaVal.Required = append(schemaVal.Required, dirEntry.Name)
			} else if strings.Contains(dirEntry.Parent.Key, dirEntry.Name) {
				schemaVal.Required = append(schemaVal.Required, dirEntry.Name)
			}

			if dirEntry.IsLeaf() {
				openapiComponents.Schemas[toUnderScore(itemPath)] = &openapi3.SchemaRef{
					Value: schemaVal,
				}
			} else { // Leaflist
				arr := openapi3.NewSchema()
				arr.Type = "leaf-list"
				arr.Items = &openapi3.SchemaRef{
					Value: schemaVal,
				}
				arr.Title = fmt.Sprintf("leaf-list-%s", dirEntry.Name)
				openapiComponents.Schemas[toUnderScore(itemPath)] = &openapi3.SchemaRef{
					Value: arr,
				}
			}
		} else if dirEntry.Kind == yang.ChoiceEntry {
			for name, dir := range dirEntry.Dir {
				_, components, err := buildSchema(dir, dir.Config, parentPath)
				if err != nil {
					return nil, nil, err
				}
				for k, v := range components.Schemas {
					v.Value.Description = fmt.Sprintf("For choice %s:%s", dirEntry.Name, name)
					openapiComponents.Schemas[toUnderScore(k)] = v
				}
			}

		} else if dirEntry.IsContainer() {
			newPath := newPathItem(dirEntry, itemPath, parentPath)
			openapiPaths[pathWithPrefix(itemPath)] = newPath

			paths, components, err := buildSchema(dirEntry, dirEntry.Config, itemPath)
			if err != nil {
				return nil, nil, err
			}
			for k, v := range paths {
				openapiPaths[k] = v
			}

			schemaVal := openapi3.NewObjectSchema()
			schemaVal.Properties = make(map[string]*openapi3.SchemaRef)
			schemaVal.Title = toUnderScore(itemPath)
			schemaVal.Description = dirEntry.Description
			if len(strings.Split(itemPath, "/")) <= 2 {
				addAdditionalProperties(schemaVal, AdditionalPropertyTarget)
			}
			asRef := &openapi3.SchemaRef{
				Value: schemaVal,
			}
			openapiComponents.Schemas[toUnderScore(itemPath)] = asRef

			rbRef := &openapi3.RequestBodyRef{
				Value: openapi3.NewRequestBody().WithContent(
					openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Value: asRef.Value,
						Ref:   fmt.Sprintf("#/components/schemas/%s", toUnderScore(itemPath)),
					}),
				),
			}
			openapiComponents.RequestBodies[fmt.Sprintf("RequestBody_%s", toUnderScore(itemPath))] = rbRef

			if newPath.Post != nil && newPath.Post.RequestBody.Ref != "" {
				newPath.Post.RequestBody.Value = rbRef.Value
			}

			respGet200 := openapi3.NewResponse()
			respGet200.Description = &respGet200Desc
			respGet200.Content = openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
				Ref:   fmt.Sprintf("#/components/schemas/%s", toUnderScore(itemPath)),
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
					arrayObj.MinItems = v.Value.MinItems
					arrayObj.MaxItems = v.Value.MaxItems
					arrayObj.UniqueItems = v.Value.UniqueItems
					arrayObj.Extensions = v.Value.Extensions
					arrayObj.Description = v.Value.Description
					schemaVal.Properties[strings.ToLower(lastPartOf(k))] = &openapi3.SchemaRef{
						Value: arrayObj,
					}
					openapiComponents.Schemas[k] = v.Value.Items
				case "object": // Container as a child of container
					schemaPath := pathToSchemaName(itemPath)
					root := k[len(schemaPath):]
					if v.Value.Title != "" && !strings.Contains(root, "_") {
						schemaVal.Properties[strings.ToLower(lastPartOf(k))] = &openapi3.SchemaRef{
							Ref:   fmt.Sprintf("#/components/schemas/%s", v.Value.Title),
							Value: v.Value,
						}
					}
					openapiComponents.Schemas[k] = v
				case "string", "boolean", "integer", "number": // leaf as a child of list
					if v.Value.Required != nil {
						schemaVal.Required = append(schemaVal.Required, v.Value.Required...)
						sort.Strings(schemaVal.Required)
						v.Value.Required = nil
					}
					schemaVal.Properties[v.Value.Title] = v
				case "leaf-list":
					v.Value.Type = "array"
					schemaVal.Properties[v.Value.Title] = v
				default:
					return nil, nil, fmt.Errorf("undhanled in container %s: %s", k, v.Value.Type)
				}
			}
			if len(schemaVal.Required) > 0 {
				addAdditionalProperties(schemaVal, AdditionalPropertyUnchanged)
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
			openapiPaths[pathWithPrefix(listItemPath)] = newPathItem(dirEntry, itemPath, listItemPath)

			paths, components, err := buildSchema(dirEntry, dirEntry.Config, listItemPath)
			if err != nil {
				return nil, nil, err
			}
			for k, v := range paths {
				openapiPaths[k] = v
			}

			arr := openapi3.NewArraySchema()
			arr.Extensions = make(map[string]interface{})
			arr.Extensions["x-keys"] = keys
			arr.Items = &openapi3.SchemaRef{
				Value: openapi3.NewObjectSchema(),
			}
			arr.MinItems = dirEntry.ListAttr.MinElements
			if dirEntry.ListAttr.MaxElements != math.MaxUint64 {
				arr.MaxItems = &dirEntry.ListAttr.MaxElements
			}
			arr.UniqueItems = true
			arr.Title = fmt.Sprintf("Item%s", toUnderScore(itemPath))
			arr.Description = dirEntry.Description

			if dirEntry.Extra != nil {
				mustArgs := make([]yang.Must, 0)
				for k, v := range dirEntry.Extra {
					switch k {
					case "must":
						for _, e := range v {
							emap, ok := e.(map[string]interface{})
							if ok {
								m := yang.Must{}
								if name, ok := emap["Name"]; ok {
									m.Name = name.(string)
								} else {
									continue
								}
								if errMsg, ok := emap["ErrorMessage"]; ok {
									if errMsgMap, ok := errMsg.(map[string]interface{}); ok {
										if errMsgName, ok := errMsgMap["Name"]; ok {
											m.ErrorMessage = &yang.Value{
												Name: errMsgName.(string),
											}
										}
									}
								}
								mustArgs = append(mustArgs, m)
							}
						}
					}
				}
				arr.Extensions["x-must"] = mustArgs
			}

			asRef := &openapi3.SchemaRef{
				Value: arr,
			}
			openapiComponents.Schemas[toUnderScore(itemPath)] = asRef

			rbRef := &openapi3.RequestBodyRef{
				Value: openapi3.NewRequestBody().WithContent(
					openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
						Value: asRef.Value,
						Ref:   fmt.Sprintf("#/components/schemas/%s", toUnderScore(itemPath)),
					}),
				),
			}
			openapiComponents.RequestBodies[fmt.Sprintf("RequestBody_%s", toUnderScore(itemPath))] = rbRef

			if openapiPaths[pathWithPrefix(listItemPath)] != nil &&
				openapiPaths[pathWithPrefix(listItemPath)].Post != nil &&
				openapiPaths[pathWithPrefix(listItemPath)].Post.RequestBody.Ref != "" {
				openapiPaths[pathWithPrefix(listItemPath)].Post.RequestBody.Value = rbRef.Value
			}

			respGet200 := openapi3.NewResponse()
			respGet200.Description = &respGet200Desc
			respGet200.Content = openapi3.NewContentWithJSONSchemaRef(&openapi3.SchemaRef{
				Ref:   fmt.Sprintf("#/components/schemas/%s", toUnderScore(itemPath)),
				Value: openapiComponents.Schemas[toUnderScore(itemPath)].Value,
			})
			openapiPaths[pathWithPrefix(listItemPath)].Get.AddResponse(200, respGet200)
			if len(strings.Split(itemPath, "/")) <= 2 {
				addAdditionalProperties(arr.Items.Value, AdditionalPropertyTarget)
			}
			for k, v := range components.Schemas {
				switch v.Value.Type {
				case "array": // List as a child of list
					arrayObj := openapi3.NewArraySchema()
					arrayObj.Items = &openapi3.SchemaRef{
						Ref:   fmt.Sprintf("#/components/schemas/%s", k),
						Value: v.Value.Items.Value,
					}
					arrayObj.UniqueItems = v.Value.UniqueItems
					arrayObj.MinItems = v.Value.MinItems
					arrayObj.MaxItems = v.Value.MaxItems
					arrayObj.Extensions = v.Value.Extensions
					arrayObj.Description = v.Value.Description
					arrayObj.Title = lastPartOf(k)
					arr.Items.Value.Properties[strings.ToLower(lastPartOf(k))] = &openapi3.SchemaRef{
						Value: arrayObj,
					}
					openapiComponents.Schemas[k] = v.Value.Items
				case "object": // Container as a child of list
					schemaPath := pathToSchemaName(itemPath)
					root := k[len(schemaPath):]
					if v.Value.Title != "" && !strings.Contains(root, "_") {
						arr.Items.Value.Properties[strings.ToLower(lastPartOf(k))] = &openapi3.SchemaRef{
							Ref:   fmt.Sprintf("#/components/schemas/%s", v.Value.Title),
							Value: v.Value,
						}
					}
					openapiComponents.Schemas[k] = v
				case "string", "boolean", "integer", "number": // leaf as a child of list
					if v.Value.Required != nil {
						arr.Items.Value.Required = append(arr.Items.Value.Required, v.Value.Required...)
						sort.Strings(arr.Items.Value.Required)
						v.Value.Required = nil
					}
					arr.Items.Value.Properties[v.Value.Title] = v
				default:
					return nil, nil, fmt.Errorf("unhandled in list %s: %s", k, v.Value.Type)
				}
			}
			if len(arr.Items.Value.Required) > 1 {
				addAdditionalProperties(arr.Items.Value, AdditionalPropertyUnchanged)
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

func newPathItem(dirEntry *yang.Entry, itemPath string, parentPath string) *openapi3.PathItem {
	getOp := openapi3.NewOperation()
	getOp.Summary = fmt.Sprintf("GET %s", itemPath)
	getOp.OperationID = fmt.Sprintf("get%s", toUnderScore(itemPath))
	getOp.Tags = []string{toUnderScore(parentPath)}
	getOp.Responses = make(openapi3.Responses)

	parameters := make(openapi3.Parameters, 0)
	pathKeys := strings.Split(parentPath, "/")
	targetParameterRef := openapi3.ParameterRef{
		Ref:   "#/components/parameters/target",
		Value: targetParameter.Value,
	}
	parameters = append(parameters, &targetParameterRef)

	parameterNames := make(map[string]interface{})

	for _, pathKey := range pathKeys {
		if strings.HasPrefix(pathKey, "{") && strings.HasSuffix(pathKey, "}") {
			k := pathKey[1 : len(pathKey)-1]
			pathKey := pathKey // pinning
			if _, alreadyUsed := parameterNames[k]; alreadyUsed {
				newK := fmt.Sprintf("%s_%d", k, len(parameterNames))
				pathKey = fmt.Sprintf("{%s}", newK)
				// TODO: add the changed parameter name back in to the stored path
				k = newK
			}
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
			parameterNames[k] = struct{}{}
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
		deleteOp.Summary = fmt.Sprintf("DELETE %s", itemPath)
		deleteOp.OperationID = fmt.Sprintf("delete%s", toUnderScore(itemPath))
		deleteOp.Responses = openapi3.NewResponses()
		newPath.Delete = deleteOp

		postOp := openapi3.NewOperation()
		postOp.Summary = fmt.Sprintf("POST %s", itemPath)
		postOp.OperationID = fmt.Sprintf("post%s", toUnderScore(itemPath))
		postOp.Responses = make(openapi3.Responses)
		postOp.Responses["201"] = &openapi3.ResponseRef{Value: openapi3.NewResponse().WithDescription("created")}
		postOp.RequestBody = &openapi3.RequestBodyRef{
			Ref: fmt.Sprintf("#/components/requestBodies/RequestBody_%s", toUnderScore(itemPath)),
			// Value is filled in later
		}
		newPath.Post = postOp
	}

	return &newPath
}

func toUnderScore(itemPath string) string {
	pathParts := make([]string, 0)
	for _, pathPart := range strings.Split(itemPath, "/") {
		if pathPart == "" || strings.HasPrefix(pathPart, "{") || strings.HasSuffix(pathPart, "}") {
			continue
		}
		pathParts = append(pathParts, uppercaseFirstCharacter(pathPart))
	}

	return strings.Join(pathParts, "_")
}

func lastPartOf(path string) string {
	pathParts := strings.Split(path, "_")
	return pathParts[len(pathParts)-1]
}

// Uppercase the first character in a string. This assumes UTF-8, so we have
// to be careful with unicode, don't treat it as a byte array.
func uppercaseFirstCharacter(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func pathWithPrefix(itemPath string) string {
	return fmt.Sprintf("%s%s", pathPrefix, itemPath)
}

// Removes any indices
func pathToSchemaName(itemPath string) string {
	parts := strings.Split(itemPath, "/")
	partsWoIdx := make([]string, 0)
	for _, p := range parts {
		if !strings.Contains(p, "{") {
			partsWoIdx = append(partsWoIdx, p)
		}
	}
	return strings.ToLower(strings.Join(partsWoIdx, "/"))
}

// If there is more than 1 range - try to find the overall min and max
// If YANG uses min and max, then leave out any statement in OpenAPI 3
// Leave it to the implementation handle the min and max for the type
func yangRange(yangRange yang.YangRange) (*float64, *float64, error) {
	var minVal *float64
	var maxVal *float64
	var hasMinMin, hasMaxMax bool
	if yangRange.Len() == 0 {
		return nil, nil, fmt.Errorf("unexpected nil range")
	}
	for i := 0; i < yangRange.Len(); i++ {
		if yangRange[i].Min.Kind == yang.MinNumber {
			minVal = nil
			hasMinMin = true
		} else if m := floatFromYnumber(yangRange[i].Min); (!hasMinMin && minVal == nil) || (minVal != nil && *m < *minVal) {
			minVal = m
		}

		if yangRange[i].Max.Kind == yang.MaxNumber {
			maxVal = nil
			hasMaxMax = true
		} else if m := floatFromYnumber(yangRange[i].Max); (!hasMaxMax && maxVal == nil) || (maxVal != nil && *m > *maxVal) {
			maxVal = m
		}
	}
	return minVal, maxVal, nil
}

func yangDefault(leaf *yang.Entry) (interface{}, error) {
	if leaf.Type.Default != "" {
		switch leaf.Type.Kind {
		case yang.Yint8:
			intValue, err := strconv.ParseInt(leaf.Type.Default, 10, 8)
			return int8(intValue), err
		case yang.Yuint8:
			intValue, err := strconv.ParseUint(leaf.Type.Default, 10, 8)
			return uint8(intValue), err
		case yang.Yint16:
			intValue, err := strconv.ParseInt(leaf.Type.Default, 10, 16)
			return int16(intValue), err
		case yang.Yuint16:
			intValue, err := strconv.ParseUint(leaf.Type.Default, 10, 16)
			return uint16(intValue), err
		case yang.Yint32:
			intValue, err := strconv.ParseInt(leaf.Type.Default, 10, 32)
			return int32(intValue), err
		case yang.Yuint32:
			intValue, err := strconv.ParseUint(leaf.Type.Default, 10, 32)
			return uint32(intValue), err
		case yang.Yint64:
			intValue, err := strconv.ParseInt(leaf.Type.Default, 10, 64)
			return int64(intValue), err
		case yang.Yuint64:
			intValue, err := strconv.ParseUint(leaf.Type.Default, 10, 64)
			return uint64(intValue), err
		case yang.Ydecimal64:
			return strconv.ParseFloat(leaf.Type.Default, 64)
		}
	}
	return nil, nil
}

func floatFromYnumber(ynumber yang.Number) *float64 {
	neg := 1.0
	if ynumber.Kind == yang.Negative {
		neg = -1.0
	}
	if !ynumber.IsDecimal() {
		v := float64(ynumber.Value) * neg
		return &v
	}
	v := float64(ynumber.Value) * math.Pow(10, -1.0*float64(ynumber.FractionDigits))
	return &v
}
