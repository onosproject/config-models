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

func NewOnfTest1GnmiClient(conn *grpc.ClientConn) *GnmiClient {
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


// {GET GetLeafAtTopLevel [leafAtTopLevel] string val.GetStringVal() ""}

// {GET ListLeafAtTopLevel [leafAtTopLevel] string val.GetStringVal() ""}

// {SET DeleteLeafAtTopLevel [leafAtTopLevel] string val.GetStringVal() ""}

// {SET UpdateLeafAtTopLevel [leafAtTopLevel] string val.GetStringVal() ""}

// {GET GetCont1ACont2ALeaf2D [cont1a cont2a leaf2d] float64 float64(val.GetFloatVal()) 0}

// {GET ListCont1ACont2ALeaf2D [cont1a cont2a leaf2d] float64 float64(val.GetFloatVal()) 0}

// {SET DeleteCont1ACont2ALeaf2D [cont1a cont2a leaf2d] float64 float64(val.GetFloatVal()) 0}

// {SET UpdateCont1ACont2ALeaf2D [cont1a cont2a leaf2d] float64 float64(val.GetFloatVal()) 0}

// {GET GetCont1ACont2ALeaf2E [cont1a cont2a leaf2e] int16 int16(val.GetIntVal()) 0}

// {GET ListCont1ACont2ALeaf2E [cont1a cont2a leaf2e] int16 int16(val.GetIntVal()) 0}

// {SET DeleteCont1ACont2ALeaf2E [cont1a cont2a leaf2e] int16 int16(val.GetIntVal()) 0}

// {SET UpdateCont1ACont2ALeaf2E [cont1a cont2a leaf2e] int16 int16(val.GetIntVal()) 0}

// {GET GetCont1ACont2ALeaf2F [cont1a cont2a leaf2f] []byte val.GetBytesVal() []byte{}}

// {GET ListCont1ACont2ALeaf2F [cont1a cont2a leaf2f] []byte val.GetBytesVal() []byte{}}

// {SET DeleteCont1ACont2ALeaf2F [cont1a cont2a leaf2f] []byte val.GetBytesVal() []byte{}}

// {SET UpdateCont1ACont2ALeaf2F [cont1a cont2a leaf2f] []byte val.GetBytesVal() []byte{}}

// {GET GetCont1ACont2ALeaf2G [cont1a cont2a leaf2g] bool val.GetBoolVal() false}

// {GET ListCont1ACont2ALeaf2G [cont1a cont2a leaf2g] bool val.GetBoolVal() false}

// {SET DeleteCont1ACont2ALeaf2G [cont1a cont2a leaf2g] bool val.GetBoolVal() false}

// {SET UpdateCont1ACont2ALeaf2G [cont1a cont2a leaf2g] bool val.GetBoolVal() false}

// {GET GetCont1ACont2ALeaf2A [cont1a cont2a leaf2a] uint8 uint8(val.GetUintVal()) 0}

// {GET ListCont1ACont2ALeaf2A [cont1a cont2a leaf2a] uint8 uint8(val.GetUintVal()) 0}

// {SET DeleteCont1ACont2ALeaf2A [cont1a cont2a leaf2a] uint8 uint8(val.GetUintVal()) 0}

// {SET UpdateCont1ACont2ALeaf2A [cont1a cont2a leaf2a] uint8 uint8(val.GetUintVal()) 0}

// {GET GetCont1ACont2ALeaf2B [cont1a cont2a leaf2b] float64 float64(val.GetFloatVal()) 0}

// {GET ListCont1ACont2ALeaf2B [cont1a cont2a leaf2b] float64 float64(val.GetFloatVal()) 0}

// {SET DeleteCont1ACont2ALeaf2B [cont1a cont2a leaf2b] float64 float64(val.GetFloatVal()) 0}

// {SET UpdateCont1ACont2ALeaf2B [cont1a cont2a leaf2b] float64 float64(val.GetFloatVal()) 0}

// {GET GetCont1ACont2ALeaf2C [cont1a cont2a leaf2c] string val.GetStringVal() ""}

// {GET ListCont1ACont2ALeaf2C [cont1a cont2a leaf2c] string val.GetStringVal() ""}

// {SET DeleteCont1ACont2ALeaf2C [cont1a cont2a leaf2c] string val.GetStringVal() ""}

// {SET UpdateCont1ACont2ALeaf2C [cont1a cont2a leaf2c] string val.GetStringVal() ""}

// {GET GetCont1ALeaf1A [cont1a leaf1a] string val.GetStringVal() ""}

// {GET ListCont1ALeaf1A [cont1a leaf1a] string val.GetStringVal() ""}

// {SET DeleteCont1ALeaf1A [cont1a leaf1a] string val.GetStringVal() ""}

// {SET UpdateCont1ALeaf1A [cont1a leaf1a] string val.GetStringVal() ""}

// {GET GetCont1BStateLeaf2D [cont1b-state leaf2d] uint16 uint16(val.GetUintVal()) 0}

// {GET ListCont1BStateLeaf2D [cont1b-state leaf2d] uint16 uint16(val.GetUintVal()) 0}

// {SET DeleteCont1BStateLeaf2D [cont1b-state leaf2d] uint16 uint16(val.GetUintVal()) 0}

// {SET UpdateCont1BStateLeaf2D [cont1b-state leaf2d] uint16 uint16(val.GetUintVal()) 0}



func (c *GnmiClient) GetLeafAtTopLevel(ctx context.Context, target string,
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


func (c *GnmiClient) ListLeafAtTopLevel(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteLeafAtTopLevel(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateLeafAtTopLevel(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ACont2ALeaf2D(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ACont2ALeaf2D(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ACont2ALeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ACont2ALeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ACont2ALeaf2E(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ACont2ALeaf2E(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ACont2ALeaf2E(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ACont2ALeaf2E(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ACont2ALeaf2F(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ACont2ALeaf2F(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ACont2ALeaf2F(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ACont2ALeaf2F(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ACont2ALeaf2G(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ACont2ALeaf2G(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ACont2ALeaf2G(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ACont2ALeaf2G(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ACont2ALeaf2A(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ACont2ALeaf2A(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ACont2ALeaf2A(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ACont2ALeaf2A(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ACont2ALeaf2B(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ACont2ALeaf2B(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ACont2ALeaf2B(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ACont2ALeaf2B(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ACont2ALeaf2C(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ACont2ALeaf2C(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ACont2ALeaf2C(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ACont2ALeaf2C(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1ALeaf1A(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1ALeaf1A(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1ALeaf1A(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1ALeaf1A(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) GetCont1BStateLeaf2D(ctx context.Context, target string,
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


func (c *GnmiClient) ListCont1BStateLeaf2D(ctx context.Context, target string,
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


func (c *GnmiClient) DeleteCont1BStateLeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
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


func (c *GnmiClient) UpdateCont1BStateLeaf2D(ctx context.Context, target string, val *gnmi.TypedValue,
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

