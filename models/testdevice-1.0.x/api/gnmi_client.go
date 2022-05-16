/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
 */

// Generated via gnmi-gen.go, do NOT edit

package api

import (
	"context"
	"github.com/onosproject/config-models/pkg/gnmi-client-gen/gnmi_utils"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"time"
)

type GnmiClient struct {
	client gnmi.GNMIClient
}

func NewOnfTest1GnmiClient(conn *grpc.ClientConn) *GnmiClient {
	gnmi_client := gnmi.NewGNMIClient(conn)
	return &GnmiClient{client: gnmi_client}
}

func (c *GnmiClient) GetCont1A_Cont2A(ctx context.Context, target string,
) (*OnfTest1_Cont1A_Cont2A, error) {
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

	return st.Cont1A.Cont2A, nil

}

func (c *GnmiClient) UpdateCont1A_Cont2A(ctx context.Context, target string, data OnfTest1_Cont1A_Cont2A,
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
					Name: "cont2a",
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

func (c *GnmiClient) GetCont1A(ctx context.Context, target string,
) (*OnfTest1_Cont1A, error) {
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

	return st.Cont1A, nil

}

func (c *GnmiClient) UpdateCont1A(ctx context.Context, target string, data OnfTest1_Cont1A,
) (*gnmi.SetResponse, error) {
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

	req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
	if err != nil {
		return nil, err
	}

	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1BState(ctx context.Context, target string,
) (*OnfTest1_Cont1BState, error) {
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

	return st.Cont1BState, nil

}

func (c *GnmiClient) UpdateCont1BState(ctx context.Context, target string, data OnfTest1_Cont1BState,
) (*gnmi.SetResponse, error) {
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

	req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
	if err != nil {
		return nil, err
	}

	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ACont2ALeaf2E(ctx context.Context, target string,
) (int16, error) {
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

	return int16(val.GetIntVal()), nil
}

func (c *GnmiClient) ListCont1ACont2ALeaf2E(ctx context.Context, target string,
) (int16, error) {
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

	return int16(val.GetIntVal()), nil
}

func (c *GnmiClient) DeleteCont1ACont2ALeaf2E(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2e",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ACont2ALeaf2E(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2e",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ACont2ALeaf2F(ctx context.Context, target string,
) ([]byte, error) {
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
		return []byte{}, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return []byte{}, err
	}

	return val.GetBytesVal(), nil
}

func (c *GnmiClient) ListCont1ACont2ALeaf2F(ctx context.Context, target string,
) ([]byte, error) {
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
		return []byte{}, err
	}

	val, err := gnmi_utils.GetResponseUpdate(res)

	if err != nil {
		return []byte{}, err
	}

	return val.GetBytesVal(), nil
}

func (c *GnmiClient) DeleteCont1ACont2ALeaf2F(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2f",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ACont2ALeaf2F(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2f",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ACont2ALeaf2G(ctx context.Context, target string,
) (bool, error) {
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

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) ListCont1ACont2ALeaf2G(ctx context.Context, target string,
) (bool, error) {
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

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) DeleteCont1ACont2ALeaf2G(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2g",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ACont2ALeaf2G(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2g",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ACont2ALeaf2A(ctx context.Context, target string,
) (uint8, error) {
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

	return uint8(val.GetUintVal()), nil
}

func (c *GnmiClient) ListCont1ACont2ALeaf2A(ctx context.Context, target string,
) (uint8, error) {
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

	return uint8(val.GetUintVal()), nil
}

func (c *GnmiClient) DeleteCont1ACont2ALeaf2A(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ACont2ALeaf2A(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ACont2ALeaf2B(ctx context.Context, target string,
) (float64, error) {
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

	return float64(val.GetFloatVal()), nil
}

func (c *GnmiClient) ListCont1ACont2ALeaf2B(ctx context.Context, target string,
) (float64, error) {
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

	return float64(val.GetFloatVal()), nil
}

func (c *GnmiClient) DeleteCont1ACont2ALeaf2B(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2b",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ACont2ALeaf2B(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2b",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ACont2ALeaf2C(ctx context.Context, target string,
) (string, error) {
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

	return val.GetStringVal(), nil
}

func (c *GnmiClient) ListCont1ACont2ALeaf2C(ctx context.Context, target string,
) (string, error) {
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

	return val.GetStringVal(), nil
}

func (c *GnmiClient) DeleteCont1ACont2ALeaf2C(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2c",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ACont2ALeaf2C(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2c",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ACont2ALeaf2D(ctx context.Context, target string,
) (float64, error) {
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

	return float64(val.GetFloatVal()), nil
}

func (c *GnmiClient) ListCont1ACont2ALeaf2D(ctx context.Context, target string,
) (float64, error) {
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

	return float64(val.GetFloatVal()), nil
}

func (c *GnmiClient) DeleteCont1ACont2ALeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2d",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ACont2ALeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "cont2a",
				},
				{
					Name: "leaf2d",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1ALeaf1A(ctx context.Context, target string,
) (string, error) {
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

	return val.GetStringVal(), nil
}

func (c *GnmiClient) ListCont1ALeaf1A(ctx context.Context, target string,
) (string, error) {
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

	return val.GetStringVal(), nil
}

func (c *GnmiClient) DeleteCont1ALeaf1A(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "leaf1a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1ALeaf1A(ctx context.Context, target string, val *gnmi.TypedValue,
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
					Name: "leaf1a",
				},
			},
			Target: target,
		},
	}

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetCont1BStateLeaf2D(ctx context.Context, target string,
) (uint16, error) {
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

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) ListCont1BStateLeaf2D(ctx context.Context, target string,
) (uint16, error) {
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

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) DeleteCont1BStateLeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
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

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateCont1BStateLeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
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

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) GetLeafAtTopLevel(ctx context.Context, target string,
) (string, error) {
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

	return val.GetStringVal(), nil
}

func (c *GnmiClient) ListLeafAtTopLevel(ctx context.Context, target string,
) (string, error) {
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

	return val.GetStringVal(), nil
}

func (c *GnmiClient) DeleteLeafAtTopLevel(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
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

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) UpdateLeafAtTopLevel(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
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

	req := &gnmi.SetRequest{
		Update: []*gnmi.Update{
			{
				Path: path[0],
				Val:  val,
			},
		},
	}
	return c.client.Set(gnmiCtx, req)
}
