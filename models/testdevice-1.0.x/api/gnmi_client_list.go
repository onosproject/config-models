///*
// * SPDX-FileCopyrightText: 2022-present Intel Corporation
// *
// * SPDX-License-Identifier: Apache-2.0
// */
//
package api

import (
	"context"
	"github.com/onosproject/config-models/pkg/gnmi-client-gen/gnmi_utils"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"time"
)

// TODO autogenerate these methods
func (c *GnmiClient) UpdateCont1AList2AItem(ctx context.Context, target string, data OnfTest1_Cont1A_List2A,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},
				{
					Name: "list2a",
					Key: map[string]string{
						"name": *data.Name,
					},
				},
			},
			Target: target,
		},
	}

	req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
	if err != nil {
		return nil, err
	}

	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1AList2AItem(ctx context.Context, target string, key string,
) (*OnfTest1_Cont1A_List2A, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},
				{
					Name: "list2a",
					Key: map[string]string{
						"name": key,
					},
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_JSON,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return nil, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return nil, err
	}

	json := val.GetJsonVal()
	st := Device{}
	Unmarshal(json, &st)

	if reflect.ValueOf(st.Cont1A).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A).IsNil() {
		return nil, status.Error(codes.NotFound, "Cont1A-not-found")
	}

	if res, ok := st.Cont1A.List2A[key]; ok {
		return res, nil
	}
	return nil, status.Error(codes.NotFound, "item-not-found")
}
