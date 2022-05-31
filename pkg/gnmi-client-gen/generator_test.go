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

func Test_generateGnmiEndpointsForLists(t *testing.T) {

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
			got, err := generateGnmiEndpointsForLists(tt.args.item, tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("generateGnmiEndpointsForLists(%v, %v)", tt.args.item, tt.args.path)) {
				return
			}
			assert.Equal(t, len(methods), len(got), "more than expected endpoints have been generated")
			assert.Equalf(t, tt.want, got, "generateGnmiEndpointsForLists(%v, %v)", tt.args.item, tt.args.path)
		})
	}
}
