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
    "github.com/openconfig/gnmi/proto/gnmi"
    "google.golang.org/grpc"
    "time"
)

type GnmiClient struct {
    client gnmi.GNMIClient
}

func NewGnmiClient(conn *grpc.ClientConn) *GnmiClient {
    gnmi_client := gnmi.NewGNMIClient(conn)
    return &GnmiClient{client: gnmi_client}
}

// GetResponseUpdate -- extract the single Update from the GetResponse
func GetResponseUpdate(gr *gnmi.GetResponse) (*gnmi.TypedValue, error) {
    if len(gr.Notification) != 1 {
        return nil, fmt.Errorf("unexpected number of GetResponse notifications %d", len(gr.Notification))
    }
    n0 := gr.Notification[0]
    if len(n0.Update) != 1 {
        return nil, fmt.Errorf("unexpected number of GetResponse notification updates %d", len(n0.Update))
    }
    u0 := n0.Update[0]
    if u0.Val == nil {
        return nil, nil
    }
    return &gnmi.TypedValue{
        Value: u0.Val.Value,
    }, nil
}


func (c *GnmiClient) GetCont1aCont2aLeaf2g(ctx context.Context, target string,
    ) (bool, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return false, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return false, err
    }

    return val.GetBoolVal(), nil
    }


func (c *GnmiClient) ListCont1aCont2aLeaf2g(ctx context.Context, target string,
    ) (bool, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return false, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return false, err
    }

    return val.GetBoolVal(), nil
    }


func (c *GnmiClient) DeleteCont1aCont2aLeaf2g(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aCont2aLeaf2g(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1aCont2aLeaf2a(ctx context.Context, target string,
    ) (uint8, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return uint8(val.GetUintVal()), nil
    }


func (c *GnmiClient) ListCont1aCont2aLeaf2a(ctx context.Context, target string,
    ) (uint8, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return uint8(val.GetUintVal()), nil
    }


func (c *GnmiClient) DeleteCont1aCont2aLeaf2a(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aCont2aLeaf2a(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1aCont2aLeaf2b(ctx context.Context, target string,
    ) (float64, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return float64(val.GetFloatVal()), nil
    }


func (c *GnmiClient) ListCont1aCont2aLeaf2b(ctx context.Context, target string,
    ) (float64, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return float64(val.GetFloatVal()), nil
    }


func (c *GnmiClient) DeleteCont1aCont2aLeaf2b(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aCont2aLeaf2b(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1aCont2aLeaf2c(ctx context.Context, target string,
    ) (string, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return "", err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    return val.GetStringVal(), nil
    }


func (c *GnmiClient) ListCont1aCont2aLeaf2c(ctx context.Context, target string,
    ) (string, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return "", err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    return val.GetStringVal(), nil
    }


func (c *GnmiClient) DeleteCont1aCont2aLeaf2c(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aCont2aLeaf2c(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1aCont2aLeaf2d(ctx context.Context, target string,
    ) (float64, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return float64(val.GetFloatVal()), nil
    }


func (c *GnmiClient) ListCont1aCont2aLeaf2d(ctx context.Context, target string,
    ) (float64, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return float64(val.GetFloatVal()), nil
    }


func (c *GnmiClient) DeleteCont1aCont2aLeaf2d(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aCont2aLeaf2d(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1aCont2aLeaf2e(ctx context.Context, target string,
    ) (int16, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return int16(val.GetIntVal()), nil
    }


func (c *GnmiClient) ListCont1aCont2aLeaf2e(ctx context.Context, target string,
    ) (int16, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return int16(val.GetIntVal()), nil
    }


func (c *GnmiClient) DeleteCont1aCont2aLeaf2e(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aCont2aLeaf2e(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1aCont2aLeaf2f(ctx context.Context, target string,
    ) ([]byte, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return []byte{}, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return []byte{}, err
    }

    return val.GetBytesVal(), nil
    }


func (c *GnmiClient) ListCont1aCont2aLeaf2f(ctx context.Context, target string,
    ) ([]byte, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return []byte{}, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return []byte{}, err
    }

    return val.GetBytesVal(), nil
    }


func (c *GnmiClient) DeleteCont1aCont2aLeaf2f(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aCont2aLeaf2f(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1aLeaf1a(ctx context.Context, target string,
    ) (string, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return "", err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    return val.GetStringVal(), nil
    }


func (c *GnmiClient) ListCont1aLeaf1a(ctx context.Context, target string,
    ) (string, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return "", err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    return val.GetStringVal(), nil
    }


func (c *GnmiClient) DeleteCont1aLeaf1a(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1aLeaf1a(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetCont1b_StateLeaf2d(ctx context.Context, target string,
    ) (uint16, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return uint16(val.GetUintVal()), nil
    }


func (c *GnmiClient) ListCont1b_StateLeaf2d(ctx context.Context, target string,
    ) (uint16, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return 0, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    return uint16(val.GetUintVal()), nil
    }


func (c *GnmiClient) DeleteCont1b_StateLeaf2d(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateCont1b_StateLeaf2d(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) GetLeafattoplevel(ctx context.Context, target string,
    ) (string, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return "", err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    return val.GetStringVal(), nil
    }


func (c *GnmiClient) ListLeafattoplevel(ctx context.Context, target string,
    ) (string, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return "", err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    return val.GetStringVal(), nil
    }


func (c *GnmiClient) DeleteLeafattoplevel(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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


func (c *GnmiClient) UpdateLeafattoplevel(ctx context.Context, target string, val *gnmi.TypedValue,
    ) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
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

