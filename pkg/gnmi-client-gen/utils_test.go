/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen

import (
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetListKey(t *testing.T) {

	type args struct {
		entry *yang.Entry
	}

	tests := []struct {
		name string
		args args
		want ListKey
	}{
		{name: "simple-key", args: args{entry: &yang.Entry{
			Key: "id",
			Dir: map[string]*yang.Entry{
				"id": {Type: &yang.YangType{Kind: yang.Ystring}},
			},
			Annotation: map[string]interface{}{
				"structname": "OnfTest1_Cont1A_List5",
			},
		}}, want: ListKey{Type: "string", Keys: []Key{{Name: "Id", Type: "string"}}}},
		{
			name: "composite-key",
			args: args{
				entry: &yang.Entry{
					Name: "list5",
					Key:  "key1 key2",
					Annotation: map[string]interface{}{
						"structname": "OnfTest1_Cont1A_List5",
					},
					Dir: map[string]*yang.Entry{
						"key1": {Type: &yang.YangType{Kind: yang.Ystring}},
						"key2": {Type: &yang.YangType{Kind: yang.Ystring}},
					},
				},
			},
			want: ListKey{
				Type: "OnfTest1_Cont1A_List5_Key",
				Keys: []Key{
					{Name: "Key1", Type: "string"},
					{Name: "Key2", Type: "string"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetListKey(tt.args.entry)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Type, res.Type)
			assert.Equal(t, len(tt.want.Keys), len(res.Keys))
			for i, k := range res.Keys {
				w := tt.want.Keys[i]
				assert.Equal(t, w.Name, k.Name)
				assert.Equal(t, w.Type, k.Type)
			}
		})
	}
}
