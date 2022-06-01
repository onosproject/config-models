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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"time"
)

type GnmiClient struct {
	client gnmi.GNMIClient
}

func NewE2nodeGnmiClient(conn *grpc.ClientConn) *GnmiClient {
	gnmi_client := gnmi.NewGNMIClient(conn)
	return &GnmiClient{client: gnmi_client}
}

func (c *GnmiClient) Get_E2Node_Intervals(ctx context.Context, target string,
) (*E2Node_E2Node_Intervals, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
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

	if reflect.ValueOf(st.E2Node).Kind() == reflect.Ptr && reflect.ValueOf(st.E2Node).IsNil() {
		return nil, status.Error(codes.NotFound, "E2Node_E2Node_Intervals-not-found")
	}
	if reflect.ValueOf(st.E2Node.Intervals).Kind() == reflect.Ptr && reflect.ValueOf(st.E2Node.Intervals).IsNil() {
		return nil, status.Error(codes.NotFound, "E2Node_E2Node_Intervals-not-found")
	}

	return st.E2Node.Intervals, nil

}

func (c *GnmiClient) Delete_E2Node_Intervals(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
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

func (c *GnmiClient) Update_E2Node_Intervals(ctx context.Context, target string, data E2Node_E2Node_Intervals,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
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

func (c *GnmiClient) Get_E2Node(ctx context.Context, target string,
) (*E2Node_E2Node, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
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

	if reflect.ValueOf(st.E2Node).Kind() == reflect.Ptr && reflect.ValueOf(st.E2Node).IsNil() {
		return nil, status.Error(codes.NotFound, "E2Node_E2Node-not-found")
	}

	return st.E2Node, nil

}

func (c *GnmiClient) Delete_E2Node(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
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

func (c *GnmiClient) Update_E2Node(ctx context.Context, target string, data E2Node_E2Node,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
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

func (c *GnmiClient) Get_E2NodeIntervalsSchedMeasReportPerUe(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "SchedMeasReportPerUe",
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
		return 0, status.Error(codes.NotFound, "E2NodeIntervalsSchedMeasReportPerUe-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Delete_E2NodeIntervalsSchedMeasReportPerUe(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "SchedMeasReportPerUe",
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

func (c *GnmiClient) Update_E2NodeIntervalsSchedMeasReportPerUe(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "SchedMeasReportPerUe",
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

func (c *GnmiClient) Get_E2NodeIntervalsPdcpMeasReportPerUe(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "PdcpMeasReportPerUe",
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
		return 0, status.Error(codes.NotFound, "E2NodeIntervalsPdcpMeasReportPerUe-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Delete_E2NodeIntervalsPdcpMeasReportPerUe(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "PdcpMeasReportPerUe",
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

func (c *GnmiClient) Update_E2NodeIntervalsPdcpMeasReportPerUe(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "PdcpMeasReportPerUe",
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

func (c *GnmiClient) Get_E2NodeIntervalsRadioMeasReportPerCell(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "RadioMeasReportPerCell",
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
		return 0, status.Error(codes.NotFound, "E2NodeIntervalsRadioMeasReportPerCell-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Delete_E2NodeIntervalsRadioMeasReportPerCell(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "RadioMeasReportPerCell",
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

func (c *GnmiClient) Update_E2NodeIntervalsRadioMeasReportPerCell(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "RadioMeasReportPerCell",
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

func (c *GnmiClient) Get_E2NodeIntervalsRadioMeasReportPerUe(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "RadioMeasReportPerUe",
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
		return 0, status.Error(codes.NotFound, "E2NodeIntervalsRadioMeasReportPerUe-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Delete_E2NodeIntervalsRadioMeasReportPerUe(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "RadioMeasReportPerUe",
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

func (c *GnmiClient) Update_E2NodeIntervalsRadioMeasReportPerUe(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "RadioMeasReportPerUe",
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

func (c *GnmiClient) Get_E2NodeIntervalsSchedMeasReportPerCell(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "SchedMeasReportPerCell",
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
		return 0, status.Error(codes.NotFound, "E2NodeIntervalsSchedMeasReportPerCell-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Delete_E2NodeIntervalsSchedMeasReportPerCell(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "SchedMeasReportPerCell",
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

func (c *GnmiClient) Update_E2NodeIntervalsSchedMeasReportPerCell(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "e2node",
				},
				{
					Name: "intervals",
				},
				{
					Name: "SchedMeasReportPerCell",
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
