/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"fmt"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup(t *testing.T) {
	// Make sure the default is restored as some tests change it
	methods = []string{gnmiGet, gnmiDelete, gnmiUpdate}
}

func Test_generateGnmiEndpointsForLists(t *testing.T) {
	setup(t)
	// only generate endpoinds for the GET method
	methods = []string{gnmiGet}

	type args struct {
		item *yang.Entry
		path []string
	}
	tests := []struct {
		name    string
		args    args
		want    []ListEndpoint
		wantErr assert.ErrorAssertionFunc
	}{
		{"list-endpoint-plural",
			args{
				item: &yang.Entry{
					Name: "application",
					Key:  "id",
					Annotation: map[string]interface{}{
						"structname": "Test_Application",
					},
					Dir: map[string]*yang.Entry{
						"id": {Type: &yang.YangType{Kind: yang.Ystring}},
					},
				},
				path: []string{"application"},
			},
			[]ListEndpoint{
				{
					ContainerEndpoint{
						ModelName:       "Test_Application",
						ModelPath:       "Application",
						ParentModelPath: "",
						Method:          "get",
						MethodName:      "Get_Application",
						Path:            []string{"application"},
					},
					ListKey{
						Type: "string",
						Keys: []Key{
							{Name: "Id", Type: "string", Ptr: false},
						},
					},
					[]string{},
					"Get_Application_List",
				},
			},
			assert.NoError,
		},
		{"list-endpoint-plural-s",
			args{
				item: &yang.Entry{
					Name: "vcs",
					Key:  "id",
					Annotation: map[string]interface{}{
						"structname": "Test_Vcs",
					},
					Dir: map[string]*yang.Entry{
						"id": {Type: &yang.YangType{Kind: yang.Ystring}},
					},
				},
				path: []string{"vcs"},
			},
			[]ListEndpoint{
				{
					ContainerEndpoint{
						ModelName:       "Test_Vcs",
						ModelPath:       "Vcs",
						ParentModelPath: "",
						Method:          "get",
						MethodName:      "Get_Vcs",
						Path:            []string{"vcs"},
					},
					ListKey{
						Type: "string",
						Keys: []Key{
							{Name: "Id", Type: "string", Ptr: false},
						},
					},
					[]string{},
					"Get_Vcs_List",
				},
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateGnmiEndpointsForLists(tt.args.item, tt.args.path, nil)
			if !tt.wantErr(t, err, fmt.Sprintf("generateGnmiEndpointsForLists(%v, %v)", tt.args.item, tt.args.path)) {
				return
			}
			assert.Equal(t, len(methods), len(got), "more than expected endpoints have been generated")
			assert.Equalf(t, tt.want, got, "generateGnmiEndpointsForLists(%v, %v)", tt.args.item, tt.args.path)
		})
	}
}

func TestBuildGnmiStruct_ordering(t *testing.T) {
	setup(t)
	type args struct {
		debug      bool
		pluginName string
		entry      *yang.Entry
		parentPath []string
	}
	tests := []struct {
		name    string
		args    args
		want    *GnmiEndpoints
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"leaves-order",
			args{
				debug:      false,
				pluginName: "Test",
				entry: &yang.Entry{
					Name: "Device",
					Dir: map[string]*yang.Entry{
						"b_leaf": {Name: "b_leaf", Kind: yang.LeafEntry, Type: &yang.YangType{Kind: yang.Ystring}},
						"a_leaf": {Name: "a_leaf", Kind: yang.LeafEntry, Type: &yang.YangType{Kind: yang.Ystring}},
					},
				},
				parentPath: []string{},
			},
			&GnmiEndpoints{
				LeavesEndpoints: []LeavesEndpoint{
					{Method: gnmiDelete, MethodName: "Delete_ALeaf", ModelName: "ALeaf", Path: []string{"a_leaf"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
					{Method: gnmiDelete, MethodName: "Delete_BLeaf", ModelName: "BLeaf", Path: []string{"b_leaf"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
					{Method: gnmiGet, MethodName: "Get_ALeaf", ModelName: "ALeaf", Path: []string{"a_leaf"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
					{Method: gnmiGet, MethodName: "Get_BLeaf", ModelName: "BLeaf", Path: []string{"b_leaf"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
					{Method: gnmiUpdate, MethodName: "Update_ALeaf", ModelName: "ALeaf", Path: []string{"a_leaf"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
					{Method: gnmiUpdate, MethodName: "Update_BLeaf", ModelName: "BLeaf", Path: []string{"b_leaf"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
				},
				ContainerEndpoints: []ContainerEndpoint{},
				ListEndpoints:      []ListEndpoint{},
				PluginName:         "Test",
			},
			assert.NoError,
		},
		{
			"container-order",
			args{
				debug:      false,
				pluginName: "Test",
				entry: &yang.Entry{
					Name: "Device",
					Dir: map[string]*yang.Entry{
						"b_cont": {Name: "b_cont", Kind: yang.DirectoryEntry, Annotation: map[string]interface{}{"structname": "Bcont"}},
						"a_cont": {Name: "a_cont", Kind: yang.DirectoryEntry, Annotation: map[string]interface{}{"structname": "Acont"}},
					},
				},
				parentPath: []string{},
			},
			&GnmiEndpoints{
				LeavesEndpoints: []LeavesEndpoint{},
				ContainerEndpoints: []ContainerEndpoint{
					{Method: gnmiDelete, MethodName: "Delete_ACont", ModelName: "Acont", Path: []string{"a_cont"}, ModelPath: "ACont"},
					{Method: gnmiDelete, MethodName: "Delete_BCont", ModelName: "Bcont", Path: []string{"b_cont"}, ModelPath: "BCont"},
					{Method: gnmiGet, MethodName: "Get_ACont", ModelName: "Acont", Path: []string{"a_cont"}, ModelPath: "ACont"},
					{Method: gnmiGet, MethodName: "Get_BCont", ModelName: "Bcont", Path: []string{"b_cont"}, ModelPath: "BCont"},
					{Method: gnmiUpdate, MethodName: "Update_ACont", ModelName: "Acont", Path: []string{"a_cont"}, ModelPath: "ACont"},
					{Method: gnmiUpdate, MethodName: "Update_BCont", ModelName: "Bcont", Path: []string{"b_cont"}, ModelPath: "BCont"},
				},
				ListEndpoints: []ListEndpoint{},
				PluginName:    "Test",
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// NOTE we need to repeat the test multiple times as the order is not guaranteed and can randomly be correct on a single run
			count := 10

			for i := 1; i <= count; i++ {
				got, err := BuildGnmiStruct(tt.args.debug, tt.args.pluginName, tt.args.entry, tt.args.parentPath, nil)
				if !tt.wantErr(t, err, fmt.Sprintf("BuildGnmiStruct(%v, %v, %v, %v)", tt.args.debug, tt.args.pluginName, tt.args.entry, tt.args.parentPath)) {
					return
				}
				assert.Equalf(t, tt.want, got, "gnmi-endpoint-were-not-correctly-generated-on-run-%d: BuildGnmiStruct(%v, %v, %v, %v)", i, tt.args.debug, tt.args.pluginName, tt.args.entry, tt.args.parentPath)
			}
		})
	}
}

func TestBuildGnmiStruct_nested_lists(t *testing.T) {

	setup(t)
	// only generate endpoinds for the GET method
	methods = []string{gnmiGet}

	listAttr := &yang.ListAttr{
		MinElements: 0,
		MaxElements: 10,
	}

	type args struct {
		debug      bool
		pluginName string
		entry      *yang.Entry
		parentPath []string
	}
	tests := []struct {
		name    string
		args    args
		want    *GnmiEndpoints
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"nested-lists-single-keyed",
			args{
				debug:      false,
				pluginName: "Test",
				entry: &yang.Entry{
					Name: "Device",
					Dir: map[string]*yang.Entry{
						"rootlist": {
							Name: "rootlist",
							Annotation: map[string]interface{}{
								"structname": "Test_RootList",
							},
							Key: "rootlist_id",
							Dir: map[string]*yang.Entry{
								"rootlist_id": {
									Name: "rootlist_id",
									Type: &yang.YangType{Kind: yang.Ystring},
								},
								"childlist": {
									Name: "childlist",
									Annotation: map[string]interface{}{
										"structname": "Test_ChildList",
									},
									Key: "childlist_id",
									Dir: map[string]*yang.Entry{
										"childlist_id": {
											Name: "childlist_id",
											Type: &yang.YangType{Kind: yang.Ystring},
										},
									},
									ListAttr: listAttr,
								},
							},
							ListAttr: listAttr,
						},
					},
				},
				parentPath: []string{},
			},
			&GnmiEndpoints{
				LeavesEndpoints: []LeavesEndpoint{
					{Method: gnmiGet, MethodName: "Get_RootlistChildlistChildlistId", ModelName: "RootlistChildlistChildlistId", Path: []string{"rootlist", "childlist", "childlist_id"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
					{Method: gnmiGet, MethodName: "Get_RootlistRootlistId", ModelName: "RootlistRootlistId", Path: []string{"rootlist", "rootlist_id"}, GoType: "string", GoReturnType: "val.GetStringVal()", GoEmptyReturnType: "\"\""},
				},
				ContainerEndpoints: []ContainerEndpoint{},
				ListEndpoints: []ListEndpoint{
					{
						ContainerEndpoint{
							ModelName:  "Test_RootList",
							ModelPath:  "Rootlist",
							Method:     gnmiGet,
							MethodName: "Get_Rootlist",
							Path:       []string{"rootlist"},
						},
						ListKey{
							Type: "string",
							Keys: []Key{
								{Name: "Rootlist_id", Type: "string", Ptr: false},
							},
						},
						[]string{},
						"Get_Rootlist_List",
					},
					{
						ContainerEndpoint{
							ModelName:       "Test_ChildList",
							ModelPath:       "Rootlist.Childlist",
							ParentModelPath: "Rootlist",
							Method:          gnmiGet,
							MethodName:      "Get_Rootlist_Childlist",
							Path:            []string{"rootlist", "childlist"},
						},
						ListKey{
							Type: "string",
							Keys: []Key{
								{Name: "Childlist_id", Type: "string", Ptr: false},
							},
							ParentKey: &ListKey{
								Type: "string",
								Keys: []Key{
									{Name: "Rootlist_id", Type: "string", Ptr: false},
								},
							},
						},
						[]string{"rootlist"},
						"Get_Rootlist_Childlist_List",
					},
				},
				PluginName: "Test",
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildGnmiStruct(tt.args.debug, tt.args.pluginName, tt.args.entry, tt.args.parentPath, nil)
			if !tt.wantErr(t, err, fmt.Sprintf("BuildGnmiStruct(%v, %v, %v, %v)", tt.args.debug, tt.args.pluginName, tt.args.entry, tt.args.parentPath)) {
				return
			}
			assert.Equal(t, len(tt.want.LeavesEndpoints), len(got.LeavesEndpoints), "wrong-number-of-leaves-endpoints-generated")
			assert.Equal(t, len(tt.want.ContainerEndpoints), len(got.ContainerEndpoints), "wrong-number-of-leaves-endpoints-generated")
			assert.Equal(t, len(tt.want.ListEndpoints), len(got.ListEndpoints), "wrong-number-of-leaves-endpoints-generated")

			for i, leafEp := range tt.want.LeavesEndpoints {
				assert.Equalf(t, leafEp, got.LeavesEndpoints[i], "leaf-enpoint-for-%s-was-incorrectly generated", leafEp.ModelName)
			}

			for i, containerEp := range tt.want.ContainerEndpoints {
				assert.Equalf(t, containerEp, got.ContainerEndpoints[i], "container-enpoint-for-%s-was-incorrectly generated", containerEp.ModelName)
			}

			for i, listEp := range tt.want.ListEndpoints {
				assert.Equalf(t, listEp, got.ListEndpoints[i], "list-enpoint-for-%s-was-incorrectly generated", listEp.ModelName)
			}

		})
	}
}
