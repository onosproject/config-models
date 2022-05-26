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

func NewTestdeviceGnmiClient(conn *grpc.ClientConn) *GnmiClient {
    gnmi_client := gnmi.NewGNMIClient(conn)
    return &GnmiClient{client: gnmi_client}
}









func (c *GnmiClient) GetCont1A_List2A_Item(ctx context.Context, target string, 
key string,
) (*OnfTest1_Cont1A_List2A, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "cont1a",
                },
            {
                    Name: "list2a",
                    Key: map[string]string{
                        
                        "name": fmt.Sprint(key),
                        
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
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List2A-not-found")
    }
    if res, ok := st.Cont1A.List2A[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List2A-not-found")
    }








func (c *GnmiClient) DeleteCont1A_List2A_Item(ctx context.Context, target string, 
key string,
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
                    Name: "list2a",
                    Key: map[string]string{
                        
                        "name": fmt.Sprint(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateCont1A_List2A_Item(ctx context.Context, target string,  data OnfTest1_Cont1A_List2A,
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
                    Name: "list2a",
                    Key: map[string]string{
                        "name": fmt.Sprint(*data.Name),
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








func (c *GnmiClient) GetCont1A_List4_Item(ctx context.Context, target string, 
key string,
) (*OnfTest1_Cont1A_List4, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "cont1a",
                },
            {
                    Name: "list4",
                    Key: map[string]string{
                        
                        "id": fmt.Sprint(key),
                        
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
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List4-not-found")
    }
    if res, ok := st.Cont1A.List4[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List4-not-found")
    }








func (c *GnmiClient) DeleteCont1A_List4_Item(ctx context.Context, target string, 
key string,
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
                    Name: "list4",
                    Key: map[string]string{
                        
                        "id": fmt.Sprint(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateCont1A_List4_Item(ctx context.Context, target string,  data OnfTest1_Cont1A_List4,
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
                    Name: "list4",
                    Key: map[string]string{
                        "id": fmt.Sprint(*data.Id),
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








func (c *GnmiClient) GetCont1A_List5_Item(ctx context.Context, target string, 
key OnfTest1_Cont1A_List5_Key,
) (*OnfTest1_Cont1A_List5, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "cont1a",
                },
            {
                    Name: "list5",
                    Key: map[string]string{
                        
                        "key1": fmt.Sprint(key.Key1),
                        
                        
                        "key2": fmt.Sprint(key.Key2),
                        
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
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List5-not-found")
    }
    if res, ok := st.Cont1A.List5[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List5-not-found")
    }








func (c *GnmiClient) DeleteCont1A_List5_Item(ctx context.Context, target string, 
key OnfTest1_Cont1A_List5_Key,
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
                    Name: "list5",
                    Key: map[string]string{
                        
                        "key1": fmt.Sprint(key.Key1),
                        
                        
                        "key2": fmt.Sprint(key.Key2),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateCont1A_List5_Item(ctx context.Context, target string,  data OnfTest1_Cont1A_List5,
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
                    Name: "list5",
                    Key: map[string]string{
                        "key1": fmt.Sprint(*data.Key1),
                        "key2": fmt.Sprint(*data.Key2),
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








func (c *GnmiClient) GetCont1BState_List2B_Item(ctx context.Context, target string, 
key uint8,
) (*OnfTest1_Cont1BState_List2B, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "cont1b-state",
                },
            {
                    Name: "list2b",
                    Key: map[string]string{
                        
                        "index": fmt.Sprint(key),
                        
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

    if reflect.ValueOf(st.Cont1BState).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1BState).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1BState_List2B-not-found")
    }
    if res, ok := st.Cont1BState.List2B[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTest1_Cont1BState_List2B-not-found")
    }








func (c *GnmiClient) DeleteCont1BState_List2B_Item(ctx context.Context, target string, 
key uint8,
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
                    Name: "list2b",
                    Key: map[string]string{
                        
                        "index": fmt.Sprint(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateCont1BState_List2B_Item(ctx context.Context, target string,  data OnfTest1_Cont1BState_List2B,
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
                    Name: "list2b",
                    Key: map[string]string{
                        "index": fmt.Sprint(*data.Index),
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





func (c *GnmiClient) GetCont1A_List2A(ctx context.Context, target string, 
) (map[string]*OnfTest1_Cont1A_List2A, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "cont1a",
            },
        {
                Name: "list2a",
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
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List2A-not-found")
    }
    if reflect.ValueOf(st.Cont1A.List2A).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A.List2A).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List2A-not-found")
    }

    return st.Cont1A.List2A, nil
}





func (c *GnmiClient) DeleteCont1A_List2A(ctx context.Context, target string, 
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
                Name: "list2a",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateCont1A_List2A(ctx context.Context, target string,  list map[string]*OnfTest1_Cont1A_List2A,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "cont1a",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "name": fmt.Sprint(*item.Name),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetCont1A_List4(ctx context.Context, target string, 
) (map[string]*OnfTest1_Cont1A_List4, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "cont1a",
            },
        {
                Name: "list4",
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
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List4-not-found")
    }
    if reflect.ValueOf(st.Cont1A.List4).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A.List4).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List4-not-found")
    }

    return st.Cont1A.List4, nil
}





func (c *GnmiClient) DeleteCont1A_List4(ctx context.Context, target string, 
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
                Name: "list4",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateCont1A_List4(ctx context.Context, target string,  list map[string]*OnfTest1_Cont1A_List4,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "cont1a",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": fmt.Sprint(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetCont1A_List5(ctx context.Context, target string, 
) (map[OnfTest1_Cont1A_List5_Key]*OnfTest1_Cont1A_List5, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "cont1a",
            },
        {
                Name: "list5",
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
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List5-not-found")
    }
    if reflect.ValueOf(st.Cont1A.List5).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A.List5).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_List5-not-found")
    }

    return st.Cont1A.List5, nil
}





func (c *GnmiClient) DeleteCont1A_List5(ctx context.Context, target string, 
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
                Name: "list5",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateCont1A_List5(ctx context.Context, target string,  list map[OnfTest1_Cont1A_List5_Key]*OnfTest1_Cont1A_List5,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "cont1a",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "key1": fmt.Sprint(*item.Key1),
                    "key2": fmt.Sprint(*item.Key2),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetCont1BState_List2B(ctx context.Context, target string, 
) (map[uint8]*OnfTest1_Cont1BState_List2B, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "cont1b-state",
            },
        {
                Name: "list2b",
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

    if reflect.ValueOf(st.Cont1BState).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1BState).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1BState_List2B-not-found")
    }
    if reflect.ValueOf(st.Cont1BState.List2B).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1BState.List2B).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1BState_List2B-not-found")
    }

    return st.Cont1BState.List2B, nil
}





func (c *GnmiClient) DeleteCont1BState_List2B(ctx context.Context, target string, 
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
                Name: "list2b",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateCont1BState_List2B(ctx context.Context, target string,  list map[uint8]*OnfTest1_Cont1BState_List2B,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "cont1b-state",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "index": fmt.Sprint(*item.Index),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetCont1A_Cont2A(ctx context.Context, target string, 
) (*OnfTest1_Cont1A_Cont2A, error) {
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
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_Cont2A-not-found")
    }
    if reflect.ValueOf(st.Cont1A.Cont2A).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A.Cont2A).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A_Cont2A-not-found")
    }

    return st.Cont1A.Cont2A, nil

}



func (c *GnmiClient) DeleteCont1A_Cont2A(ctx context.Context, target string, 
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
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateCont1A_Cont2A(ctx context.Context, target string,  data OnfTest1_Cont1A_Cont2A,
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



path :=  []*gnmi.Path{
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



func (c *GnmiClient) DeleteCont1A(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "cont1a",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateCont1A(ctx context.Context, target string,  data OnfTest1_Cont1A,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
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



path :=  []*gnmi.Path{
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

    if reflect.ValueOf(st.Cont1BState).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1BState).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfTest1_Cont1BState-not-found")
    }

    return st.Cont1BState, nil

}



func (c *GnmiClient) DeleteCont1BState(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "cont1b-state",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateCont1BState(ctx context.Context, target string,  data OnfTest1_Cont1BState,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    if val.GetStringVal() ==  "" {
    return "", status.Error(codes.NotFound, "LeafAtTopLevel-not-found")
    }

    return val.GetStringVal(), nil
}



func (c *GnmiClient) DeleteLeafAtTopLevel(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    if float64(val.GetFloatVal()) ==  0 {
    return 0, status.Error(codes.NotFound, "Cont1ACont2ALeaf2D-not-found")
    }

    return float64(val.GetFloatVal()), nil
}



func (c *GnmiClient) DeleteCont1ACont2ALeaf2D(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    if int16(val.GetIntVal()) ==  0 {
    return 0, status.Error(codes.NotFound, "Cont1ACont2ALeaf2E-not-found")
    }

    return int16(val.GetIntVal()), nil
}



func (c *GnmiClient) DeleteCont1ACont2ALeaf2E(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    if val.GetBytesVal() ==  nil {
    return nil, status.Error(codes.NotFound, "Cont1ACont2ALeaf2F-not-found")
    }

    return val.GetBytesVal(), nil
}



func (c *GnmiClient) DeleteCont1ACont2ALeaf2F(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return false, err
    }

    if val.GetBoolVal() ==  false {
    return false, status.Error(codes.NotFound, "Cont1ACont2ALeaf2G-not-found")
    }

    return val.GetBoolVal(), nil
}



func (c *GnmiClient) DeleteCont1ACont2ALeaf2G(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    if uint8(val.GetUintVal()) ==  0 {
    return 0, status.Error(codes.NotFound, "Cont1ACont2ALeaf2A-not-found")
    }

    return uint8(val.GetUintVal()), nil
}



func (c *GnmiClient) DeleteCont1ACont2ALeaf2A(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    if float64(val.GetFloatVal()) ==  0 {
    return 0, status.Error(codes.NotFound, "Cont1ACont2ALeaf2B-not-found")
    }

    return float64(val.GetFloatVal()), nil
}



func (c *GnmiClient) DeleteCont1ACont2ALeaf2B(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    if val.GetStringVal() ==  "" {
    return "", status.Error(codes.NotFound, "Cont1ACont2ALeaf2C-not-found")
    }

    return val.GetStringVal(), nil
}



func (c *GnmiClient) DeleteCont1ACont2ALeaf2C(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return "", err
    }

    if val.GetStringVal() ==  "" {
    return "", status.Error(codes.NotFound, "Cont1ALeaf1A-not-found")
    }

    return val.GetStringVal(), nil
}



func (c *GnmiClient) DeleteCont1ALeaf1A(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return 0, err
    }

    if uint16(val.GetUintVal()) ==  0 {
    return 0, status.Error(codes.NotFound, "Cont1BStateLeaf2D-not-found")
    }

    return uint16(val.GetUintVal()), nil
}



func (c *GnmiClient) DeleteCont1BStateLeaf2D(ctx context.Context, target string,
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
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
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
