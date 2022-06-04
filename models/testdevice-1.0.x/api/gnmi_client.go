/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
 */

// Generated via gnmi-gen.go, do NOT edit

package api

import (
	"context"
	"fmt"
	"github.com/onosproject/config-models/pkg/gnmi-client-gen/gnmi_utils"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"time"
)

type GnmiClient struct {
	client gnmi.GNMIClient
}

func NewAetherGnmiClient(conn *grpc.ClientConn) *GnmiClient {
	gnmi_client := gnmi.NewGNMIClient(conn)
	return &GnmiClient{client: gnmi_client}
}

func (c *GnmiClient) Get_Cont1a(ctx context.Context, target string) (*OnfTest1_Cont1A, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
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
		return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A-not-found")
	}

	return st.Cont1A, nil

}

func (c *GnmiClient) Get_Cont1a_Cont2a(ctx context.Context, target string) (*OnfTest1_Cont1A_Cont2A, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
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

	if reflect.ValueOf(st.Cont1A.Cont2A).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A).IsNil() {
		return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A-not-found")
	}

	return st.Cont1A.Cont2A, nil

}

func (c *GnmiClient) Get_Cont1a_Cont2a_Leaf2a(ctx context.Context, target string) (uint8, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
				},

				{
					Name: "leaf2a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return 0, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return 0, err
	}

	if uint8(val.GetUintVal()) == 0 {
		return 0, status.Error(codes.NotFound, "leaf2a-not-found")
	}

	return uint8(val.GetUintVal()), nil
}
func (c *GnmiClient) Get_Cont1a_Cont2a_Leaf2b(ctx context.Context, target string) (float64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
				},

				{
					Name: "leaf2b",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return 0, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return 0, err
	}

	if float64(val.GetFloatVal()) == 0 {
		return 0, status.Error(codes.NotFound, "leaf2b-not-found")
	}

	return float64(val.GetFloatVal()), nil
}
func (c *GnmiClient) Get_Cont1a_Cont2a_Leaf2c(ctx context.Context, target string) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
				},

				{
					Name: "leaf2c",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return "", err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return "", err
	}

	if val.GetStringVal() == "" {
		return "", status.Error(codes.NotFound, "leaf2c-not-found")
	}

	return val.GetStringVal(), nil
}
func (c *GnmiClient) Get_Cont1a_Cont2a_Leaf2d(ctx context.Context, target string) (float64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
				},

				{
					Name: "leaf2d",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return 0, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return 0, err
	}

	if float64(val.GetFloatVal()) == 0 {
		return 0, status.Error(codes.NotFound, "leaf2d-not-found")
	}

	return float64(val.GetFloatVal()), nil
}
func (c *GnmiClient) Get_Cont1a_Cont2a_Leaf2e(ctx context.Context, target string) (int16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
				},

				{
					Name: "leaf2e",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return 0, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return 0, err
	}

	if int16(val.GetIntVal()) == 0 {
		return 0, status.Error(codes.NotFound, "leaf2e-not-found")
	}

	return int16(val.GetIntVal()), nil
}
func (c *GnmiClient) Get_Cont1a_Cont2a_Leaf2f(ctx context.Context, target string) ([]byte, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
				},

				{
					Name: "leaf2f",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
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

	if val.GetBytesVal() == nil {
		return nil, status.Error(codes.NotFound, "leaf2f-not-found")
	}

	return val.GetBytesVal(), nil
}
func (c *GnmiClient) Get_Cont1a_Cont2a_Leaf2g(ctx context.Context, target string) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "cont2a",
				},

				{
					Name: "leaf2g",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return false, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return false, err
	}

	if val.GetBoolVal() == false {
		return false, status.Error(codes.NotFound, "leaf2g-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_Cont1a_Leaf1a(ctx context.Context, target string) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1a",
				},

				{
					Name: "leaf1a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return "", err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return "", err
	}

	if val.GetStringVal() == "" {
		return "", status.Error(codes.NotFound, "leaf1a-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_Cont1bState(ctx context.Context, target string) (*OnfTest1_Cont1BState, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1b-state",
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

	if reflect.ValueOf(st.Cont1BState).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A).IsNil() {
		return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A-not-found")
	}

	return st.Cont1BState, nil

}

func (c *GnmiClient) Get_Cont1bState_Leaf2d(ctx context.Context, target string) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "cont1b-state",
				},

				{
					Name: "leaf2d",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return 0, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return 0, err
	}

	if uint16(val.GetUintVal()) == 0 {
		return 0, status.Error(codes.NotFound, "leaf2d-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_Leafattoplevel(ctx context.Context, target string) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "leafAtTopLevel",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.GetRequest{
		Encoding: gnmi.Encoding_PROTO,
		Path:     path,
	}
	res, err := c.client.Get(gnmiCtx, req)

	if err != nil {
		return "", err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return "", err
	}

	if val.GetStringVal() == "" {
		return "", status.Error(codes.NotFound, "leafAtTopLevel-not-found")
	}

	return val.GetStringVal(), nil
}
