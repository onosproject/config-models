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

func NewRicGnmiClient(conn *grpc.ClientConn) *GnmiClient {
	gnmi_client := gnmi.NewGNMIClient(conn)
	return &GnmiClient{client: gnmi_client}
}

func (c *GnmiClient) Delete_Nodes_Node(ctx context.Context, target string,
	key string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
				},
				{
					Name: "node",
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

func (c *GnmiClient) Get_Nodes_Node(ctx context.Context, target string,
	key string,
) (*Xapp_Nodes_Node, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
				},
				{
					Name: "node",
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

	if reflect.ValueOf(st.Nodes).Kind() == reflect.Ptr && reflect.ValueOf(st.Nodes).IsNil() {
		return nil, status.Error(codes.NotFound, "Xapp_Nodes_Node-not-found")
	}
	if res, ok := st.Nodes.Node[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "Xapp_Nodes_Node-not-found")
}

func (c *GnmiClient) Update_Nodes_Node(ctx context.Context, target string, data Xapp_Nodes_Node,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
				},
				{
					Name: "node",
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

func (c *GnmiClient) Delete_Nodes_Node_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
				},
				{
					Name: "node",
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

func (c *GnmiClient) Get_Nodes_Node_List(ctx context.Context, target string,
) (map[string]*Xapp_Nodes_Node, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
				},
				{
					Name: "node",
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

	if reflect.ValueOf(st.Nodes).Kind() == reflect.Ptr && reflect.ValueOf(st.Nodes).IsNil() {
		return nil, status.Error(codes.NotFound, "Xapp_Nodes_Node-not-found")
	}
	if reflect.ValueOf(st.Nodes.Node).Kind() == reflect.Ptr && reflect.ValueOf(st.Nodes.Node).IsNil() {
		return nil, status.Error(codes.NotFound, "Xapp_Nodes_Node-not-found")
	}

	return st.Nodes.Node, nil
}

func (c *GnmiClient) Update_Nodes_Node_List(ctx context.Context, target string, list map[string]*Xapp_Nodes_Node,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "nodes",
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

func (c *GnmiClient) Delete_Nodes(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
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

func (c *GnmiClient) Delete_ReportPeriod(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "report_period",
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

func (c *GnmiClient) Get_Nodes(ctx context.Context, target string,
) (*Xapp_Nodes, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
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

	if reflect.ValueOf(st.Nodes).Kind() == reflect.Ptr && reflect.ValueOf(st.Nodes).IsNil() {
		return nil, status.Error(codes.NotFound, "Xapp_Nodes-not-found")
	}

	return st.Nodes, nil

}

func (c *GnmiClient) Get_ReportPeriod(ctx context.Context, target string,
) (*KpimonXapp_ReportPeriod, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "report_period",
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

	if reflect.ValueOf(st.ReportPeriod).Kind() == reflect.Ptr && reflect.ValueOf(st.ReportPeriod).IsNil() {
		return nil, status.Error(codes.NotFound, "KpimonXapp_ReportPeriod-not-found")
	}

	return st.ReportPeriod, nil

}

func (c *GnmiClient) Update_Nodes(ctx context.Context, target string, data Xapp_Nodes,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "nodes",
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

func (c *GnmiClient) Update_ReportPeriod(ctx context.Context, target string, data KpimonXapp_ReportPeriod,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "report_period",
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

func (c *GnmiClient) Delete_ReportPeriodInterval(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "report_period",
				},
				{
					Name: "interval",
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

func (c *GnmiClient) Get_ReportPeriodInterval(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "report_period",
				},
				{
					Name: "interval",
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

	if uint32(val.GetUintVal()) == 0 {
		return 0, status.Error(codes.NotFound, "ReportPeriodInterval-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Update_ReportPeriodInterval(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "report_period",
				},
				{
					Name: "interval",
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
