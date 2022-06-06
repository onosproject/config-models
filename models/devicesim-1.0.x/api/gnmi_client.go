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

func NewDevicesimGnmiClient(conn *grpc.ClientConn) *GnmiClient {
	gnmi_client := gnmi.NewGNMIClient(conn)
	return &GnmiClient{client: gnmi_client}
}

func (c *GnmiClient) Delete_Components_Component(ctx context.Context, target string,
	key string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
				},
				{
					Name: "component",
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

func (c *GnmiClient) Delete_Interfaces_Interface(ctx context.Context, target string,
	key string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
				},
				{
					Name: "interface",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_Users_User(ctx context.Context, target string,
	key string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
				},
				{
					Name: "user",
					Key: map[string]string{

						"username": fmt.Sprint(key),
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

func (c *GnmiClient) Delete_System_Aaa_ServerGroups_ServerGroup(ctx context.Context, target string,
	key string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
				},
				{
					Name: "server-group",
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

func (c *GnmiClient) Delete_System_Dns_HostEntries_HostEntry(ctx context.Context, target string,
	key string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
				},
				{
					Name: "host-entry",
					Key: map[string]string{

						"hostname": fmt.Sprint(key),
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

func (c *GnmiClient) Delete_System_Ntp_NtpKeys_NtpKey(ctx context.Context, target string,
	key uint16,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
				},
				{
					Name: "ntp-key",
					Key: map[string]string{

						"key-id": fmt.Sprint(key),
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

func (c *GnmiClient) Delete_System_Openflow_Controllers_Controller(ctx context.Context, target string,
	key string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
				},
				{
					Name: "controller",
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

func (c *GnmiClient) Delete_System_Processes_Process(ctx context.Context, target string,
	key uint64,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
				},
				{
					Name: "process",
					Key: map[string]string{

						"pid": fmt.Sprint(key),
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

func (c *GnmiClient) Get_Components_Component(ctx context.Context, target string,
	key string,
) (*OpenconfigPlatform_Components_Component, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
				},
				{
					Name: "component",
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

	if reflect.ValueOf(st.Components).Kind() == reflect.Ptr && reflect.ValueOf(st.Components).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigPlatform_Components_Component-not-found")
	}
	if res, ok := st.Components.Component[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigPlatform_Components_Component-not-found")
}

func (c *GnmiClient) Get_Interfaces_Interface(ctx context.Context, target string,
	key string,
) (*OpenconfigInterfaces_Interfaces_Interface, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
				},
				{
					Name: "interface",
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

	if reflect.ValueOf(st.Interfaces).Kind() == reflect.Ptr && reflect.ValueOf(st.Interfaces).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigInterfaces_Interfaces_Interface-not-found")
	}
	if res, ok := st.Interfaces.Interface[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigInterfaces_Interfaces_Interface-not-found")
}

func (c *GnmiClient) Get_System_Aaa_Authentication_Users_User(ctx context.Context, target string,
	key string,
) (*OpenconfigSystem_System_Aaa_Authentication_Users_User, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
				},
				{
					Name: "user",
					Key: map[string]string{

						"username": fmt.Sprint(key),
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

	if reflect.ValueOf(st.System.Aaa.Authentication.Users).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.Users).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Users_User-not-found")
	}
	if res, ok := st.System.Aaa.Authentication.Users.User[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Users_User-not-found")
}

func (c *GnmiClient) Get_System_Aaa_ServerGroups_ServerGroup(ctx context.Context, target string,
	key string,
) (*OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
				},
				{
					Name: "server-group",
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

	if reflect.ValueOf(st.System.Aaa.ServerGroups).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.ServerGroups).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup-not-found")
	}
	if res, ok := st.System.Aaa.ServerGroups.ServerGroup[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup-not-found")
}

func (c *GnmiClient) Get_System_Dns_HostEntries_HostEntry(ctx context.Context, target string,
	key string,
) (*OpenconfigSystem_System_Dns_HostEntries_HostEntry, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
				},
				{
					Name: "host-entry",
					Key: map[string]string{

						"hostname": fmt.Sprint(key),
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

	if reflect.ValueOf(st.System.Dns.HostEntries).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns.HostEntries).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_HostEntries_HostEntry-not-found")
	}
	if res, ok := st.System.Dns.HostEntries.HostEntry[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_HostEntries_HostEntry-not-found")
}

func (c *GnmiClient) Get_System_Ntp_NtpKeys_NtpKey(ctx context.Context, target string,
	key uint16,
) (*OpenconfigSystem_System_Ntp_NtpKeys_NtpKey, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
				},
				{
					Name: "ntp-key",
					Key: map[string]string{

						"key-id": fmt.Sprint(key),
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

	if reflect.ValueOf(st.System.Ntp.NtpKeys).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp.NtpKeys).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_NtpKeys_NtpKey-not-found")
	}
	if res, ok := st.System.Ntp.NtpKeys.NtpKey[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_NtpKeys_NtpKey-not-found")
}

func (c *GnmiClient) Get_System_Openflow_Controllers_Controller(ctx context.Context, target string,
	key string,
) (*OpenconfigSystem_System_Openflow_Controllers_Controller, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
				},
				{
					Name: "controller",
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

	if reflect.ValueOf(st.System.Openflow.Controllers).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Controllers).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Controllers_Controller-not-found")
	}
	if res, ok := st.System.Openflow.Controllers.Controller[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Controllers_Controller-not-found")
}

func (c *GnmiClient) Get_System_Processes_Process(ctx context.Context, target string,
	key uint64,
) (*OpenconfigSystem_System_Processes_Process, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
				},
				{
					Name: "process",
					Key: map[string]string{

						"pid": fmt.Sprint(key),
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

	if reflect.ValueOf(st.System.Processes).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Processes).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Processes_Process-not-found")
	}
	if res, ok := st.System.Processes.Process[key]; ok {
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Processes_Process-not-found")
}

func (c *GnmiClient) Update_Components_Component(ctx context.Context, target string, data OpenconfigPlatform_Components_Component,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
				},
				{
					Name: "component",
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

func (c *GnmiClient) Update_Interfaces_Interface(ctx context.Context, target string, data OpenconfigInterfaces_Interfaces_Interface,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
				},
				{
					Name: "interface",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_Users_User(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication_Users_User,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
				},
				{
					Name: "user",
					Key: map[string]string{
						"username": fmt.Sprint(*data.Username),
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

func (c *GnmiClient) Update_System_Aaa_ServerGroups_ServerGroup(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
				},
				{
					Name: "server-group",
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

func (c *GnmiClient) Update_System_Dns_HostEntries_HostEntry(ctx context.Context, target string, data OpenconfigSystem_System_Dns_HostEntries_HostEntry,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
				},
				{
					Name: "host-entry",
					Key: map[string]string{
						"hostname": fmt.Sprint(*data.Hostname),
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

func (c *GnmiClient) Update_System_Ntp_NtpKeys_NtpKey(ctx context.Context, target string, data OpenconfigSystem_System_Ntp_NtpKeys_NtpKey,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
				},
				{
					Name: "ntp-key",
					Key: map[string]string{
						"key-id": fmt.Sprint(*data.KeyId),
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

func (c *GnmiClient) Update_System_Openflow_Controllers_Controller(ctx context.Context, target string, data OpenconfigSystem_System_Openflow_Controllers_Controller,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
				},
				{
					Name: "controller",
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

func (c *GnmiClient) Update_System_Processes_Process(ctx context.Context, target string, data OpenconfigSystem_System_Processes_Process,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
				},
				{
					Name: "process",
					Key: map[string]string{
						"pid": fmt.Sprint(*data.Pid),
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

func (c *GnmiClient) Delete_Components_Component_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
				},
				{
					Name: "component",
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

func (c *GnmiClient) Delete_Interfaces_Interface_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
				},
				{
					Name: "interface",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_Users_User_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
				},
				{
					Name: "user",
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

func (c *GnmiClient) Delete_System_Aaa_ServerGroups_ServerGroup_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
				},
				{
					Name: "server-group",
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

func (c *GnmiClient) Delete_System_Dns_HostEntries_HostEntry_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
				},
				{
					Name: "host-entry",
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

func (c *GnmiClient) Delete_System_Ntp_NtpKeys_NtpKey_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
				},
				{
					Name: "ntp-key",
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

func (c *GnmiClient) Delete_System_Openflow_Controllers_Controller_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
				},
				{
					Name: "controller",
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

func (c *GnmiClient) Delete_System_Processes_Process_List(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
				},
				{
					Name: "process",
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

func (c *GnmiClient) Get_Components_Component_List(ctx context.Context, target string,
) (map[string]*OpenconfigPlatform_Components_Component, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
				},
				{
					Name: "component",
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

	if reflect.ValueOf(st.Components).Kind() == reflect.Ptr && reflect.ValueOf(st.Components).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigPlatform_Components_Component-not-found")
	}
	if reflect.ValueOf(st.Components.Component).Kind() == reflect.Ptr && reflect.ValueOf(st.Components.Component).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigPlatform_Components_Component-not-found")
	}

	return st.Components.Component, nil
}

func (c *GnmiClient) Get_Interfaces_Interface_List(ctx context.Context, target string,
) (map[string]*OpenconfigInterfaces_Interfaces_Interface, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
				},
				{
					Name: "interface",
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

	if reflect.ValueOf(st.Interfaces).Kind() == reflect.Ptr && reflect.ValueOf(st.Interfaces).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigInterfaces_Interfaces_Interface-not-found")
	}
	if reflect.ValueOf(st.Interfaces.Interface).Kind() == reflect.Ptr && reflect.ValueOf(st.Interfaces.Interface).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigInterfaces_Interfaces_Interface-not-found")
	}

	return st.Interfaces.Interface, nil
}

func (c *GnmiClient) Get_System_Aaa_Authentication_Users_User_List(ctx context.Context, target string,
) (map[string]*OpenconfigSystem_System_Aaa_Authentication_Users_User, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
				},
				{
					Name: "user",
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

	if reflect.ValueOf(st.System.Aaa.Authentication.Users).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.Users).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Users_User-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication.Users.User).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.Users.User).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Users_User-not-found")
	}

	return st.System.Aaa.Authentication.Users.User, nil
}

func (c *GnmiClient) Get_System_Aaa_ServerGroups_ServerGroup_List(ctx context.Context, target string,
) (map[string]*OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
				},
				{
					Name: "server-group",
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

	if reflect.ValueOf(st.System.Aaa.ServerGroups).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.ServerGroups).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.ServerGroups.ServerGroup).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.ServerGroups.ServerGroup).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup-not-found")
	}

	return st.System.Aaa.ServerGroups.ServerGroup, nil
}

func (c *GnmiClient) Get_System_Dns_HostEntries_HostEntry_List(ctx context.Context, target string,
) (map[string]*OpenconfigSystem_System_Dns_HostEntries_HostEntry, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
				},
				{
					Name: "host-entry",
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

	if reflect.ValueOf(st.System.Dns.HostEntries).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns.HostEntries).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_HostEntries_HostEntry-not-found")
	}
	if reflect.ValueOf(st.System.Dns.HostEntries.HostEntry).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns.HostEntries.HostEntry).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_HostEntries_HostEntry-not-found")
	}

	return st.System.Dns.HostEntries.HostEntry, nil
}

func (c *GnmiClient) Get_System_Ntp_NtpKeys_NtpKey_List(ctx context.Context, target string,
) (map[uint16]*OpenconfigSystem_System_Ntp_NtpKeys_NtpKey, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
				},
				{
					Name: "ntp-key",
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

	if reflect.ValueOf(st.System.Ntp.NtpKeys).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp.NtpKeys).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_NtpKeys_NtpKey-not-found")
	}
	if reflect.ValueOf(st.System.Ntp.NtpKeys.NtpKey).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp.NtpKeys.NtpKey).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_NtpKeys_NtpKey-not-found")
	}

	return st.System.Ntp.NtpKeys.NtpKey, nil
}

func (c *GnmiClient) Get_System_Openflow_Controllers_Controller_List(ctx context.Context, target string,
) (map[string]*OpenconfigSystem_System_Openflow_Controllers_Controller, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
				},
				{
					Name: "controller",
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

	if reflect.ValueOf(st.System.Openflow.Controllers).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Controllers).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Controllers_Controller-not-found")
	}
	if reflect.ValueOf(st.System.Openflow.Controllers.Controller).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Controllers.Controller).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Controllers_Controller-not-found")
	}

	return st.System.Openflow.Controllers.Controller, nil
}

func (c *GnmiClient) Get_System_Processes_Process_List(ctx context.Context, target string,
) (map[uint64]*OpenconfigSystem_System_Processes_Process, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
				},
				{
					Name: "process",
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

	if reflect.ValueOf(st.System.Processes).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Processes).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Processes_Process-not-found")
	}
	if reflect.ValueOf(st.System.Processes.Process).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Processes.Process).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Processes_Process-not-found")
	}

	return st.System.Processes.Process, nil
}

func (c *GnmiClient) Update_Components_Component_List(ctx context.Context, target string, list map[string]*OpenconfigPlatform_Components_Component,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "components",
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

func (c *GnmiClient) Update_Interfaces_Interface_List(ctx context.Context, target string, list map[string]*OpenconfigInterfaces_Interfaces_Interface,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "interfaces",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_Users_User_List(ctx context.Context, target string, list map[string]*OpenconfigSystem_System_Aaa_Authentication_Users_User,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "system",
		},
		{
			Name: "aaa",
		},
		{
			Name: "authentication",
		},
		{
			Name: "users",
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
					"username": fmt.Sprint(*item.Username),
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

func (c *GnmiClient) Update_System_Aaa_ServerGroups_ServerGroup_List(ctx context.Context, target string, list map[string]*OpenconfigSystem_System_Aaa_ServerGroups_ServerGroup,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "system",
		},
		{
			Name: "aaa",
		},
		{
			Name: "server-groups",
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

func (c *GnmiClient) Update_System_Dns_HostEntries_HostEntry_List(ctx context.Context, target string, list map[string]*OpenconfigSystem_System_Dns_HostEntries_HostEntry,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "system",
		},
		{
			Name: "dns",
		},
		{
			Name: "host-entries",
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
					"hostname": fmt.Sprint(*item.Hostname),
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

func (c *GnmiClient) Update_System_Ntp_NtpKeys_NtpKey_List(ctx context.Context, target string, list map[uint16]*OpenconfigSystem_System_Ntp_NtpKeys_NtpKey,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "system",
		},
		{
			Name: "ntp",
		},
		{
			Name: "ntp-keys",
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
					"key-id": fmt.Sprint(*item.KeyId),
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

func (c *GnmiClient) Update_System_Openflow_Controllers_Controller_List(ctx context.Context, target string, list map[string]*OpenconfigSystem_System_Openflow_Controllers_Controller,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "system",
		},
		{
			Name: "openflow",
		},
		{
			Name: "controllers",
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

func (c *GnmiClient) Update_System_Processes_Process_List(ctx context.Context, target string, list map[uint64]*OpenconfigSystem_System_Processes_Process,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	basePathElems := []*gnmi.PathElem{
		{
			Name: "system",
		},
		{
			Name: "processes",
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
					"pid": fmt.Sprint(*item.Pid),
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

func (c *GnmiClient) Delete_Components(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
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

func (c *GnmiClient) Delete_Interfaces(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
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

func (c *GnmiClient) Delete_System(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
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

func (c *GnmiClient) Delete_System_Aaa(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
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

func (c *GnmiClient) Delete_System_Aaa_Accounting(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
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

func (c *GnmiClient) Delete_System_Aaa_Accounting_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Aaa_Accounting_Events(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Aaa_Accounting_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_AdminUser(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_AdminUser_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_AdminUser_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Delete_System_Aaa_Authentication_Users(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Delete_System_Aaa_Authorization(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
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

func (c *GnmiClient) Delete_System_Aaa_Authorization_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Aaa_Authorization_Events(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Aaa_Authorization_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Aaa_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Aaa_ServerGroups(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
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

func (c *GnmiClient) Delete_System_Aaa_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Clock(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
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

func (c *GnmiClient) Delete_System_Clock_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Clock_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Dns(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
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

func (c *GnmiClient) Delete_System_Dns_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Dns_HostEntries(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
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

func (c *GnmiClient) Delete_System_Dns_Servers(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "servers",
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

func (c *GnmiClient) Delete_System_Dns_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Logging(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
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

func (c *GnmiClient) Delete_System_Logging_Console(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
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

func (c *GnmiClient) Delete_System_Logging_Console_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Logging_Console_Selectors(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Logging_Console_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Logging_RemoteServers(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "remote-servers",
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

func (c *GnmiClient) Delete_System_Memory(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
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

func (c *GnmiClient) Delete_System_Memory_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Memory_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Ntp(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
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

func (c *GnmiClient) Delete_System_Ntp_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_Ntp_NtpKeys(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
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

func (c *GnmiClient) Delete_System_Ntp_Servers(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "servers",
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

func (c *GnmiClient) Delete_System_Ntp_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Openflow(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
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

func (c *GnmiClient) Delete_System_Openflow_Agent(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
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

func (c *GnmiClient) Delete_System_Openflow_Agent_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Openflow_Agent_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_Openflow_Controllers(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
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

func (c *GnmiClient) Delete_System_Processes(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
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

func (c *GnmiClient) Delete_System_SshServer(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
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

func (c *GnmiClient) Delete_System_SshServer_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_SshServer_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_System_TelnetServer(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
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

func (c *GnmiClient) Delete_System_TelnetServer_Config(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Delete_System_TelnetServer_State(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Get_Components(ctx context.Context, target string,
) (*OpenconfigPlatform_Components, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
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

	if reflect.ValueOf(st.Components).Kind() == reflect.Ptr && reflect.ValueOf(st.Components).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigPlatform_Components-not-found")
	}

	return st.Components, nil

}

func (c *GnmiClient) Get_Interfaces(ctx context.Context, target string,
) (*OpenconfigInterfaces_Interfaces, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
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

	if reflect.ValueOf(st.Interfaces).Kind() == reflect.Ptr && reflect.ValueOf(st.Interfaces).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigInterfaces_Interfaces-not-found")
	}

	return st.Interfaces, nil

}

func (c *GnmiClient) Get_System(ctx context.Context, target string,
) (*OpenconfigSystem_System, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System-not-found")
	}

	return st.System, nil

}

func (c *GnmiClient) Get_System_Aaa(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa-not-found")
	}
	if reflect.ValueOf(st.System.Aaa).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa-not-found")
	}

	return st.System.Aaa, nil

}

func (c *GnmiClient) Get_System_Aaa_Accounting(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Accounting, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
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

	if reflect.ValueOf(st.System.Aaa).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Accounting).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Accounting).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting-not-found")
	}

	return st.System.Aaa.Accounting, nil

}

func (c *GnmiClient) Get_System_Aaa_Accounting_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Accounting_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Aaa.Accounting).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Accounting).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting_Config-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Accounting.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Accounting.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting_Config-not-found")
	}

	return st.System.Aaa.Accounting.Config, nil

}

func (c *GnmiClient) Get_System_Aaa_Accounting_Events(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Accounting_Events, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Aaa.Accounting).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Accounting).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting_Events-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Accounting.Events).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Accounting.Events).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting_Events-not-found")
	}

	return st.System.Aaa.Accounting.Events, nil

}

func (c *GnmiClient) Get_System_Aaa_Accounting_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Accounting_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Aaa.Accounting).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Accounting).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting_State-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Accounting.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Accounting.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Accounting_State-not-found")
	}

	return st.System.Aaa.Accounting.State, nil

}

func (c *GnmiClient) Get_System_Aaa_Authentication(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authentication, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
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

	if reflect.ValueOf(st.System.Aaa).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication-not-found")
	}

	return st.System.Aaa.Authentication, nil

}

func (c *GnmiClient) Get_System_Aaa_Authentication_AdminUser(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authentication_AdminUser, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

	if reflect.ValueOf(st.System.Aaa.Authentication).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_AdminUser-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication.AdminUser).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.AdminUser).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_AdminUser-not-found")
	}

	return st.System.Aaa.Authentication.AdminUser, nil

}

func (c *GnmiClient) Get_System_Aaa_Authentication_AdminUser_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authentication_AdminUser_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Aaa.Authentication.AdminUser).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.AdminUser).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_AdminUser_Config-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication.AdminUser.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.AdminUser.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_AdminUser_Config-not-found")
	}

	return st.System.Aaa.Authentication.AdminUser.Config, nil

}

func (c *GnmiClient) Get_System_Aaa_Authentication_AdminUser_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authentication_AdminUser_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Aaa.Authentication.AdminUser).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.AdminUser).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_AdminUser_State-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication.AdminUser.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.AdminUser.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_AdminUser_State-not-found")
	}

	return st.System.Aaa.Authentication.AdminUser.State, nil

}

func (c *GnmiClient) Get_System_Aaa_Authentication_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authentication_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

	if reflect.ValueOf(st.System.Aaa.Authentication).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Config-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Config-not-found")
	}

	return st.System.Aaa.Authentication.Config, nil

}

func (c *GnmiClient) Get_System_Aaa_Authentication_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authentication_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

	if reflect.ValueOf(st.System.Aaa.Authentication).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_State-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_State-not-found")
	}

	return st.System.Aaa.Authentication.State, nil

}

func (c *GnmiClient) Get_System_Aaa_Authentication_Users(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authentication_Users, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

	if reflect.ValueOf(st.System.Aaa.Authentication).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Users-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authentication.Users).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authentication.Users).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authentication_Users-not-found")
	}

	return st.System.Aaa.Authentication.Users, nil

}

func (c *GnmiClient) Get_System_Aaa_Authorization(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authorization, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
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

	if reflect.ValueOf(st.System.Aaa).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authorization).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authorization).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization-not-found")
	}

	return st.System.Aaa.Authorization, nil

}

func (c *GnmiClient) Get_System_Aaa_Authorization_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authorization_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Aaa.Authorization).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authorization).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization_Config-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authorization.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authorization.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization_Config-not-found")
	}

	return st.System.Aaa.Authorization.Config, nil

}

func (c *GnmiClient) Get_System_Aaa_Authorization_Events(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authorization_Events, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Aaa.Authorization).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authorization).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization_Events-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authorization.Events).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authorization.Events).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization_Events-not-found")
	}

	return st.System.Aaa.Authorization.Events, nil

}

func (c *GnmiClient) Get_System_Aaa_Authorization_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Authorization_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Aaa.Authorization).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authorization).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization_State-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Authorization.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Authorization.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Authorization_State-not-found")
	}

	return st.System.Aaa.Authorization.State, nil

}

func (c *GnmiClient) Get_System_Aaa_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Aaa).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Config-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_Config-not-found")
	}

	return st.System.Aaa.Config, nil

}

func (c *GnmiClient) Get_System_Aaa_ServerGroups(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_ServerGroups, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
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

	if reflect.ValueOf(st.System.Aaa).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_ServerGroups-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.ServerGroups).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.ServerGroups).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_ServerGroups-not-found")
	}

	return st.System.Aaa.ServerGroups, nil

}

func (c *GnmiClient) Get_System_Aaa_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Aaa_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Aaa).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_State-not-found")
	}
	if reflect.ValueOf(st.System.Aaa.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Aaa.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Aaa_State-not-found")
	}

	return st.System.Aaa.State, nil

}

func (c *GnmiClient) Get_System_Clock(ctx context.Context, target string,
) (*OpenconfigSystem_System_Clock, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Clock-not-found")
	}
	if reflect.ValueOf(st.System.Clock).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Clock).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Clock-not-found")
	}

	return st.System.Clock, nil

}

func (c *GnmiClient) Get_System_Clock_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Clock_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Clock).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Clock).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Clock_Config-not-found")
	}
	if reflect.ValueOf(st.System.Clock.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Clock.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Clock_Config-not-found")
	}

	return st.System.Clock.Config, nil

}

func (c *GnmiClient) Get_System_Clock_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Clock_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Clock).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Clock).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Clock_State-not-found")
	}
	if reflect.ValueOf(st.System.Clock.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Clock.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Clock_State-not-found")
	}

	return st.System.Clock.State, nil

}

func (c *GnmiClient) Get_System_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Config-not-found")
	}
	if reflect.ValueOf(st.System.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Config-not-found")
	}

	return st.System.Config, nil

}

func (c *GnmiClient) Get_System_Dns(ctx context.Context, target string,
) (*OpenconfigSystem_System_Dns, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns-not-found")
	}
	if reflect.ValueOf(st.System.Dns).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns-not-found")
	}

	return st.System.Dns, nil

}

func (c *GnmiClient) Get_System_Dns_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Dns_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Dns).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_Config-not-found")
	}
	if reflect.ValueOf(st.System.Dns.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_Config-not-found")
	}

	return st.System.Dns.Config, nil

}

func (c *GnmiClient) Get_System_Dns_HostEntries(ctx context.Context, target string,
) (*OpenconfigSystem_System_Dns_HostEntries, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
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

	if reflect.ValueOf(st.System.Dns).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_HostEntries-not-found")
	}
	if reflect.ValueOf(st.System.Dns.HostEntries).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns.HostEntries).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_HostEntries-not-found")
	}

	return st.System.Dns.HostEntries, nil

}

func (c *GnmiClient) Get_System_Dns_Servers(ctx context.Context, target string,
) (*OpenconfigSystem_System_Dns_Servers, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "servers",
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

	if reflect.ValueOf(st.System.Dns).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_Servers-not-found")
	}
	if reflect.ValueOf(st.System.Dns.Servers).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns.Servers).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_Servers-not-found")
	}

	return st.System.Dns.Servers, nil

}

func (c *GnmiClient) Get_System_Dns_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Dns_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Dns).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_State-not-found")
	}
	if reflect.ValueOf(st.System.Dns.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Dns.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Dns_State-not-found")
	}

	return st.System.Dns.State, nil

}

func (c *GnmiClient) Get_System_Logging(ctx context.Context, target string,
) (*OpenconfigSystem_System_Logging, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging-not-found")
	}
	if reflect.ValueOf(st.System.Logging).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging-not-found")
	}

	return st.System.Logging, nil

}

func (c *GnmiClient) Get_System_Logging_Console(ctx context.Context, target string,
) (*OpenconfigSystem_System_Logging_Console, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
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

	if reflect.ValueOf(st.System.Logging).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console-not-found")
	}
	if reflect.ValueOf(st.System.Logging.Console).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.Console).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console-not-found")
	}

	return st.System.Logging.Console, nil

}

func (c *GnmiClient) Get_System_Logging_Console_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Logging_Console_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Logging.Console).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.Console).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console_Config-not-found")
	}
	if reflect.ValueOf(st.System.Logging.Console.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.Console.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console_Config-not-found")
	}

	return st.System.Logging.Console.Config, nil

}

func (c *GnmiClient) Get_System_Logging_Console_Selectors(ctx context.Context, target string,
) (*OpenconfigSystem_System_Logging_Console_Selectors, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Logging.Console).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.Console).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console_Selectors-not-found")
	}
	if reflect.ValueOf(st.System.Logging.Console.Selectors).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.Console.Selectors).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console_Selectors-not-found")
	}

	return st.System.Logging.Console.Selectors, nil

}

func (c *GnmiClient) Get_System_Logging_Console_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Logging_Console_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Logging.Console).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.Console).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console_State-not-found")
	}
	if reflect.ValueOf(st.System.Logging.Console.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.Console.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_Console_State-not-found")
	}

	return st.System.Logging.Console.State, nil

}

func (c *GnmiClient) Get_System_Logging_RemoteServers(ctx context.Context, target string,
) (*OpenconfigSystem_System_Logging_RemoteServers, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "remote-servers",
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

	if reflect.ValueOf(st.System.Logging).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_RemoteServers-not-found")
	}
	if reflect.ValueOf(st.System.Logging.RemoteServers).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Logging.RemoteServers).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Logging_RemoteServers-not-found")
	}

	return st.System.Logging.RemoteServers, nil

}

func (c *GnmiClient) Get_System_Memory(ctx context.Context, target string,
) (*OpenconfigSystem_System_Memory, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Memory-not-found")
	}
	if reflect.ValueOf(st.System.Memory).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Memory).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Memory-not-found")
	}

	return st.System.Memory, nil

}

func (c *GnmiClient) Get_System_Memory_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Memory_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Memory).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Memory).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Memory_Config-not-found")
	}
	if reflect.ValueOf(st.System.Memory.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Memory.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Memory_Config-not-found")
	}

	return st.System.Memory.Config, nil

}

func (c *GnmiClient) Get_System_Memory_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Memory_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Memory).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Memory).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Memory_State-not-found")
	}
	if reflect.ValueOf(st.System.Memory.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Memory.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Memory_State-not-found")
	}

	return st.System.Memory.State, nil

}

func (c *GnmiClient) Get_System_Ntp(ctx context.Context, target string,
) (*OpenconfigSystem_System_Ntp, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp-not-found")
	}
	if reflect.ValueOf(st.System.Ntp).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp-not-found")
	}

	return st.System.Ntp, nil

}

func (c *GnmiClient) Get_System_Ntp_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Ntp_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.Ntp).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_Config-not-found")
	}
	if reflect.ValueOf(st.System.Ntp.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_Config-not-found")
	}

	return st.System.Ntp.Config, nil

}

func (c *GnmiClient) Get_System_Ntp_NtpKeys(ctx context.Context, target string,
) (*OpenconfigSystem_System_Ntp_NtpKeys, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
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

	if reflect.ValueOf(st.System.Ntp).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_NtpKeys-not-found")
	}
	if reflect.ValueOf(st.System.Ntp.NtpKeys).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp.NtpKeys).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_NtpKeys-not-found")
	}

	return st.System.Ntp.NtpKeys, nil

}

func (c *GnmiClient) Get_System_Ntp_Servers(ctx context.Context, target string,
) (*OpenconfigSystem_System_Ntp_Servers, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "servers",
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

	if reflect.ValueOf(st.System.Ntp).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_Servers-not-found")
	}
	if reflect.ValueOf(st.System.Ntp.Servers).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp.Servers).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_Servers-not-found")
	}

	return st.System.Ntp.Servers, nil

}

func (c *GnmiClient) Get_System_Ntp_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Ntp_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Ntp).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_State-not-found")
	}
	if reflect.ValueOf(st.System.Ntp.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Ntp.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Ntp_State-not-found")
	}

	return st.System.Ntp.State, nil

}

func (c *GnmiClient) Get_System_Openflow(ctx context.Context, target string,
) (*OpenconfigSystem_System_Openflow, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow-not-found")
	}
	if reflect.ValueOf(st.System.Openflow).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow-not-found")
	}

	return st.System.Openflow, nil

}

func (c *GnmiClient) Get_System_Openflow_Agent(ctx context.Context, target string,
) (*OpenconfigSystem_System_Openflow_Agent, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
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

	if reflect.ValueOf(st.System.Openflow).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Agent-not-found")
	}
	if reflect.ValueOf(st.System.Openflow.Agent).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Agent).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Agent-not-found")
	}

	return st.System.Openflow.Agent, nil

}

func (c *GnmiClient) Get_System_Openflow_Agent_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_Openflow_Agent_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Openflow.Agent).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Agent).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Agent_Config-not-found")
	}
	if reflect.ValueOf(st.System.Openflow.Agent.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Agent.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Agent_Config-not-found")
	}

	return st.System.Openflow.Agent.Config, nil

}

func (c *GnmiClient) Get_System_Openflow_Agent_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_Openflow_Agent_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.Openflow.Agent).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Agent).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Agent_State-not-found")
	}
	if reflect.ValueOf(st.System.Openflow.Agent.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Agent.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Agent_State-not-found")
	}

	return st.System.Openflow.Agent.State, nil

}

func (c *GnmiClient) Get_System_Openflow_Controllers(ctx context.Context, target string,
) (*OpenconfigSystem_System_Openflow_Controllers, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
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

	if reflect.ValueOf(st.System.Openflow).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Controllers-not-found")
	}
	if reflect.ValueOf(st.System.Openflow.Controllers).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Openflow.Controllers).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Openflow_Controllers-not-found")
	}

	return st.System.Openflow.Controllers, nil

}

func (c *GnmiClient) Get_System_Processes(ctx context.Context, target string,
) (*OpenconfigSystem_System_Processes, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Processes-not-found")
	}
	if reflect.ValueOf(st.System.Processes).Kind() == reflect.Ptr && reflect.ValueOf(st.System.Processes).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_Processes-not-found")
	}

	return st.System.Processes, nil

}

func (c *GnmiClient) Get_System_SshServer(ctx context.Context, target string,
) (*OpenconfigSystem_System_SshServer, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_SshServer-not-found")
	}
	if reflect.ValueOf(st.System.SshServer).Kind() == reflect.Ptr && reflect.ValueOf(st.System.SshServer).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_SshServer-not-found")
	}

	return st.System.SshServer, nil

}

func (c *GnmiClient) Get_System_SshServer_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_SshServer_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.SshServer).Kind() == reflect.Ptr && reflect.ValueOf(st.System.SshServer).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_SshServer_Config-not-found")
	}
	if reflect.ValueOf(st.System.SshServer.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.SshServer.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_SshServer_Config-not-found")
	}

	return st.System.SshServer.Config, nil

}

func (c *GnmiClient) Get_System_SshServer_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_SshServer_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.SshServer).Kind() == reflect.Ptr && reflect.ValueOf(st.System.SshServer).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_SshServer_State-not-found")
	}
	if reflect.ValueOf(st.System.SshServer.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.SshServer.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_SshServer_State-not-found")
	}

	return st.System.SshServer.State, nil

}

func (c *GnmiClient) Get_System_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_State-not-found")
	}
	if reflect.ValueOf(st.System.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_State-not-found")
	}

	return st.System.State, nil

}

func (c *GnmiClient) Get_System_TelnetServer(ctx context.Context, target string,
) (*OpenconfigSystem_System_TelnetServer, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
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

	if reflect.ValueOf(st.System).Kind() == reflect.Ptr && reflect.ValueOf(st.System).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_TelnetServer-not-found")
	}
	if reflect.ValueOf(st.System.TelnetServer).Kind() == reflect.Ptr && reflect.ValueOf(st.System.TelnetServer).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_TelnetServer-not-found")
	}

	return st.System.TelnetServer, nil

}

func (c *GnmiClient) Get_System_TelnetServer_Config(ctx context.Context, target string,
) (*OpenconfigSystem_System_TelnetServer_Config, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
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

	if reflect.ValueOf(st.System.TelnetServer).Kind() == reflect.Ptr && reflect.ValueOf(st.System.TelnetServer).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_TelnetServer_Config-not-found")
	}
	if reflect.ValueOf(st.System.TelnetServer.Config).Kind() == reflect.Ptr && reflect.ValueOf(st.System.TelnetServer.Config).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_TelnetServer_Config-not-found")
	}

	return st.System.TelnetServer.Config, nil

}

func (c *GnmiClient) Get_System_TelnetServer_State(ctx context.Context, target string,
) (*OpenconfigSystem_System_TelnetServer_State, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
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

	if reflect.ValueOf(st.System.TelnetServer).Kind() == reflect.Ptr && reflect.ValueOf(st.System.TelnetServer).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_TelnetServer_State-not-found")
	}
	if reflect.ValueOf(st.System.TelnetServer.State).Kind() == reflect.Ptr && reflect.ValueOf(st.System.TelnetServer.State).IsNil() {
		return nil, status.Error(codes.NotFound, "OpenconfigSystem_System_TelnetServer_State-not-found")
	}

	return st.System.TelnetServer.State, nil

}

func (c *GnmiClient) Update_Components(ctx context.Context, target string, data OpenconfigPlatform_Components,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "components",
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

func (c *GnmiClient) Update_Interfaces(ctx context.Context, target string, data OpenconfigInterfaces_Interfaces,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "interfaces",
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

func (c *GnmiClient) Update_System(ctx context.Context, target string, data OpenconfigSystem_System,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
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

func (c *GnmiClient) Update_System_Aaa(ctx context.Context, target string, data OpenconfigSystem_System_Aaa,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
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

func (c *GnmiClient) Update_System_Aaa_Accounting(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Accounting,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
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

func (c *GnmiClient) Update_System_Aaa_Accounting_Config(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Accounting_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Aaa_Accounting_Events(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Accounting_Events,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Aaa_Accounting_State(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Accounting_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "accounting",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Aaa_Authentication(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_AdminUser(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication_AdminUser,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_AdminUser_Config(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication_AdminUser_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_AdminUser_State(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication_AdminUser_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_Config(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_State(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Update_System_Aaa_Authentication_Users(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authentication_Users,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "users",
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

func (c *GnmiClient) Update_System_Aaa_Authorization(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authorization,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
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

func (c *GnmiClient) Update_System_Aaa_Authorization_Config(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authorization_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Aaa_Authorization_Events(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authorization_Events,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Aaa_Authorization_State(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Authorization_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authorization",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Aaa_Config(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Aaa_ServerGroups(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_ServerGroups,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "server-groups",
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

func (c *GnmiClient) Update_System_Aaa_State(ctx context.Context, target string, data OpenconfigSystem_System_Aaa_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Clock(ctx context.Context, target string, data OpenconfigSystem_System_Clock,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
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

func (c *GnmiClient) Update_System_Clock_Config(ctx context.Context, target string, data OpenconfigSystem_System_Clock_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Clock_State(ctx context.Context, target string, data OpenconfigSystem_System_Clock_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Config(ctx context.Context, target string, data OpenconfigSystem_System_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Dns(ctx context.Context, target string, data OpenconfigSystem_System_Dns,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
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

func (c *GnmiClient) Update_System_Dns_Config(ctx context.Context, target string, data OpenconfigSystem_System_Dns_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Dns_HostEntries(ctx context.Context, target string, data OpenconfigSystem_System_Dns_HostEntries,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "host-entries",
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

func (c *GnmiClient) Update_System_Dns_Servers(ctx context.Context, target string, data OpenconfigSystem_System_Dns_Servers,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "servers",
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

func (c *GnmiClient) Update_System_Dns_State(ctx context.Context, target string, data OpenconfigSystem_System_Dns_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Logging(ctx context.Context, target string, data OpenconfigSystem_System_Logging,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
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

func (c *GnmiClient) Update_System_Logging_Console(ctx context.Context, target string, data OpenconfigSystem_System_Logging_Console,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
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

func (c *GnmiClient) Update_System_Logging_Console_Config(ctx context.Context, target string, data OpenconfigSystem_System_Logging_Console_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Logging_Console_Selectors(ctx context.Context, target string, data OpenconfigSystem_System_Logging_Console_Selectors,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Logging_Console_State(ctx context.Context, target string, data OpenconfigSystem_System_Logging_Console_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "console",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Logging_RemoteServers(ctx context.Context, target string, data OpenconfigSystem_System_Logging_RemoteServers,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "logging",
				},
				{
					Name: "remote-servers",
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

func (c *GnmiClient) Update_System_Memory(ctx context.Context, target string, data OpenconfigSystem_System_Memory,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
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

func (c *GnmiClient) Update_System_Memory_Config(ctx context.Context, target string, data OpenconfigSystem_System_Memory_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Memory_State(ctx context.Context, target string, data OpenconfigSystem_System_Memory_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Ntp(ctx context.Context, target string, data OpenconfigSystem_System_Ntp,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
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

func (c *GnmiClient) Update_System_Ntp_Config(ctx context.Context, target string, data OpenconfigSystem_System_Ntp_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_Ntp_NtpKeys(ctx context.Context, target string, data OpenconfigSystem_System_Ntp_NtpKeys,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "ntp-keys",
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

func (c *GnmiClient) Update_System_Ntp_Servers(ctx context.Context, target string, data OpenconfigSystem_System_Ntp_Servers,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "servers",
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

func (c *GnmiClient) Update_System_Ntp_State(ctx context.Context, target string, data OpenconfigSystem_System_Ntp_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Openflow(ctx context.Context, target string, data OpenconfigSystem_System_Openflow,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
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

func (c *GnmiClient) Update_System_Openflow_Agent(ctx context.Context, target string, data OpenconfigSystem_System_Openflow_Agent,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
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

func (c *GnmiClient) Update_System_Openflow_Agent_Config(ctx context.Context, target string, data OpenconfigSystem_System_Openflow_Agent_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Openflow_Agent_State(ctx context.Context, target string, data OpenconfigSystem_System_Openflow_Agent_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_Openflow_Controllers(ctx context.Context, target string, data OpenconfigSystem_System_Openflow_Controllers,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "controllers",
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

func (c *GnmiClient) Update_System_Processes(ctx context.Context, target string, data OpenconfigSystem_System_Processes,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "processes",
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

func (c *GnmiClient) Update_System_SshServer(ctx context.Context, target string, data OpenconfigSystem_System_SshServer,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
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

func (c *GnmiClient) Update_System_SshServer_Config(ctx context.Context, target string, data OpenconfigSystem_System_SshServer_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_SshServer_State(ctx context.Context, target string, data OpenconfigSystem_System_SshServer_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_State(ctx context.Context, target string, data OpenconfigSystem_System_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Update_System_TelnetServer(ctx context.Context, target string, data OpenconfigSystem_System_TelnetServer,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
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

func (c *GnmiClient) Update_System_TelnetServer_Config(ctx context.Context, target string, data OpenconfigSystem_System_TelnetServer_Config,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
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

func (c *GnmiClient) Update_System_TelnetServer_State(ctx context.Context, target string, data OpenconfigSystem_System_TelnetServer_State,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
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

func (c *GnmiClient) Delete_SystemAaaAuthenticationAdminUserConfigAdminPassword(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
				},
				{
					Name: "admin-password-hashed",
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

func (c *GnmiClient) Delete_SystemAaaAuthenticationAdminUserConfigAdminPasswordHashed(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
				},
				{
					Name: "admin-password-hashed",
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

func (c *GnmiClient) Delete_SystemAaaAuthenticationAdminUserStateAdminPassword(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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

func (c *GnmiClient) Delete_SystemAaaAuthenticationAdminUserStateAdminPasswordHashed(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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

func (c *GnmiClient) Delete_SystemAaaAuthenticationAdminUserStateAdminUsername(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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

func (c *GnmiClient) Delete_SystemClockConfigTimezoneName(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "config",
				},
				{
					Name: "timezone-name",
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

func (c *GnmiClient) Delete_SystemClockStateTimezoneName(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "state",
				},
				{
					Name: "timezone-name",
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

func (c *GnmiClient) Delete_SystemConfigDomainName(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "domain-name",
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

func (c *GnmiClient) Delete_SystemConfigHostname(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "hostname",
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

func (c *GnmiClient) Delete_SystemConfigLoginBanner(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "login-banner",
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

func (c *GnmiClient) Delete_SystemConfigMotdBanner(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "motd-banner",
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

func (c *GnmiClient) Delete_SystemDnsConfigSearch(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "config",
				},
				{
					Name: "search",
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

func (c *GnmiClient) Delete_SystemDnsStateSearch(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "state",
				},
				{
					Name: "search",
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

func (c *GnmiClient) Delete_SystemMemoryStatePhysical(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
				},
				{
					Name: "reserved",
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

func (c *GnmiClient) Delete_SystemMemoryStateReserved(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
				},
				{
					Name: "reserved",
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

func (c *GnmiClient) Delete_SystemNtpConfigEnableNtpAuth(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
				},
				{
					Name: "enabled",
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

func (c *GnmiClient) Delete_SystemNtpConfigEnabled(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
				},
				{
					Name: "enabled",
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

func (c *GnmiClient) Delete_SystemNtpStateAuthMismatch(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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

func (c *GnmiClient) Delete_SystemNtpStateEnableNtpAuth(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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

func (c *GnmiClient) Delete_SystemNtpStateEnabled(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentConfigBackoffInterval(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "backoff-interval",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentConfigDatapathId(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "datapath-id",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentConfigFailureMode(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "failure-mode",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentConfigInactivityProbe(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "inactivity-probe",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentConfigMaxBackoff(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "max-backoff",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentStateBackoffInterval(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "backoff-interval",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentStateDatapathId(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "datapath-id",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentStateFailureMode(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "failure-mode",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentStateInactivityProbe(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "inactivity-probe",
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

func (c *GnmiClient) Delete_SystemOpenflowAgentStateMaxBackoff(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "max-backoff",
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

func (c *GnmiClient) Delete_SystemSshServerConfigEnable(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerConfigProtocolVersion(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerConfigRateLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerConfigSessionLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerConfigTimeout(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerStateEnable(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerStateProtocolVersion(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerStateRateLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerStateSessionLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemSshServerStateTimeout(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemStateBootTime(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "boot-time",
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

func (c *GnmiClient) Delete_SystemStateCurrentDatetime(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "current-datetime",
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

func (c *GnmiClient) Delete_SystemStateDomainName(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "domain-name",
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

func (c *GnmiClient) Delete_SystemStateHostname(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "hostname",
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

func (c *GnmiClient) Delete_SystemStateLoginBanner(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "login-banner",
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

func (c *GnmiClient) Delete_SystemStateMotdBanner(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "motd-banner",
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

func (c *GnmiClient) Delete_SystemTelnetServerConfigEnable(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Delete_SystemTelnetServerConfigRateLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Delete_SystemTelnetServerConfigSessionLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Delete_SystemTelnetServerConfigTimeout(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Delete_SystemTelnetServerStateEnable(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemTelnetServerStateRateLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemTelnetServerStateSessionLimit(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Delete_SystemTelnetServerStateTimeout(ctx context.Context, target string,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Get_SystemAaaAuthenticationAdminUserConfigAdminPassword(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
				},
				{
					Name: "admin-password-hashed",
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
		return "", status.Error(codes.NotFound, "SystemAaaAuthenticationAdminUserConfigAdminPassword-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemAaaAuthenticationAdminUserConfigAdminPasswordHashed(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
				},
				{
					Name: "admin-password-hashed",
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
		return "", status.Error(codes.NotFound, "SystemAaaAuthenticationAdminUserConfigAdminPasswordHashed-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemAaaAuthenticationAdminUserStateAdminPassword(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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
		return "", status.Error(codes.NotFound, "SystemAaaAuthenticationAdminUserStateAdminPassword-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemAaaAuthenticationAdminUserStateAdminPasswordHashed(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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
		return "", status.Error(codes.NotFound, "SystemAaaAuthenticationAdminUserStateAdminPasswordHashed-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemAaaAuthenticationAdminUserStateAdminUsername(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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
		return "", status.Error(codes.NotFound, "SystemAaaAuthenticationAdminUserStateAdminUsername-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemClockConfigTimezoneName(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "config",
				},
				{
					Name: "timezone-name",
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
		return "", status.Error(codes.NotFound, "SystemClockConfigTimezoneName-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemClockStateTimezoneName(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "state",
				},
				{
					Name: "timezone-name",
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
		return "", status.Error(codes.NotFound, "SystemClockStateTimezoneName-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemConfigDomainName(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "domain-name",
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
		return "", status.Error(codes.NotFound, "SystemConfigDomainName-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemConfigHostname(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "hostname",
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
		return "", status.Error(codes.NotFound, "SystemConfigHostname-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemConfigLoginBanner(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "login-banner",
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
		return "", status.Error(codes.NotFound, "SystemConfigLoginBanner-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemConfigMotdBanner(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "motd-banner",
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
		return "", status.Error(codes.NotFound, "SystemConfigMotdBanner-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemDnsConfigSearch(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "config",
				},
				{
					Name: "search",
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
		return "", status.Error(codes.NotFound, "SystemDnsConfigSearch-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemDnsStateSearch(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "state",
				},
				{
					Name: "search",
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
		return "", status.Error(codes.NotFound, "SystemDnsStateSearch-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemMemoryStatePhysical(ctx context.Context, target string,
) (uint64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
				},
				{
					Name: "reserved",
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

	if uint64(val.GetUintVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemMemoryStatePhysical-not-found")
	}

	return uint64(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemMemoryStateReserved(ctx context.Context, target string,
) (uint64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
				},
				{
					Name: "reserved",
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

	if uint64(val.GetUintVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemMemoryStateReserved-not-found")
	}

	return uint64(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemNtpConfigEnableNtpAuth(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
				},
				{
					Name: "enabled",
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
		return false, status.Error(codes.NotFound, "SystemNtpConfigEnableNtpAuth-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemNtpConfigEnabled(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
				},
				{
					Name: "enabled",
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
		return false, status.Error(codes.NotFound, "SystemNtpConfigEnabled-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemNtpStateAuthMismatch(ctx context.Context, target string,
) (uint64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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

	if uint64(val.GetUintVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemNtpStateAuthMismatch-not-found")
	}

	return uint64(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemNtpStateEnableNtpAuth(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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
		return false, status.Error(codes.NotFound, "SystemNtpStateEnableNtpAuth-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemNtpStateEnabled(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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
		return false, status.Error(codes.NotFound, "SystemNtpStateEnabled-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentConfigBackoffInterval(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "backoff-interval",
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
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentConfigBackoffInterval-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentConfigDatapathId(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "datapath-id",
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
		return "", status.Error(codes.NotFound, "SystemOpenflowAgentConfigDatapathId-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentConfigFailureMode(ctx context.Context, target string,
) (int64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "failure-mode",
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

	if int64(val.GetIntVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentConfigFailureMode-not-found")
	}

	return int64(val.GetIntVal()), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentConfigInactivityProbe(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "inactivity-probe",
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
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentConfigInactivityProbe-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentConfigMaxBackoff(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "max-backoff",
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
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentConfigMaxBackoff-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentStateBackoffInterval(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "backoff-interval",
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
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentStateBackoffInterval-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentStateDatapathId(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "datapath-id",
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
		return "", status.Error(codes.NotFound, "SystemOpenflowAgentStateDatapathId-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentStateFailureMode(ctx context.Context, target string,
) (int64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "failure-mode",
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

	if int64(val.GetIntVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentStateFailureMode-not-found")
	}

	return int64(val.GetIntVal()), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentStateInactivityProbe(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "inactivity-probe",
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
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentStateInactivityProbe-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemOpenflowAgentStateMaxBackoff(ctx context.Context, target string,
) (uint32, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "max-backoff",
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
		return 0, status.Error(codes.NotFound, "SystemOpenflowAgentStateMaxBackoff-not-found")
	}

	return uint32(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerConfigEnable(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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
		return false, status.Error(codes.NotFound, "SystemSshServerConfigEnable-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemSshServerConfigProtocolVersion(ctx context.Context, target string,
) (int64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

	if int64(val.GetIntVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemSshServerConfigProtocolVersion-not-found")
	}

	return int64(val.GetIntVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerConfigRateLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemSshServerConfigRateLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerConfigSessionLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemSshServerConfigSessionLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerConfigTimeout(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemSshServerConfigTimeout-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerStateEnable(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return false, status.Error(codes.NotFound, "SystemSshServerStateEnable-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemSshServerStateProtocolVersion(ctx context.Context, target string,
) (int64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

	if int64(val.GetIntVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemSshServerStateProtocolVersion-not-found")
	}

	return int64(val.GetIntVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerStateRateLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemSshServerStateRateLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerStateSessionLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemSshServerStateSessionLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemSshServerStateTimeout(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemSshServerStateTimeout-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemStateBootTime(ctx context.Context, target string,
) (uint64, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "boot-time",
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

	if uint64(val.GetUintVal()) == 0 {
		return 0, status.Error(codes.NotFound, "SystemStateBootTime-not-found")
	}

	return uint64(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemStateCurrentDatetime(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "current-datetime",
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
		return "", status.Error(codes.NotFound, "SystemStateCurrentDatetime-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemStateDomainName(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "domain-name",
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
		return "", status.Error(codes.NotFound, "SystemStateDomainName-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemStateHostname(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "hostname",
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
		return "", status.Error(codes.NotFound, "SystemStateHostname-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemStateLoginBanner(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "login-banner",
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
		return "", status.Error(codes.NotFound, "SystemStateLoginBanner-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemStateMotdBanner(ctx context.Context, target string,
) (string, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "motd-banner",
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
		return "", status.Error(codes.NotFound, "SystemStateMotdBanner-not-found")
	}

	return val.GetStringVal(), nil
}

func (c *GnmiClient) Get_SystemTelnetServerConfigEnable(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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
		return false, status.Error(codes.NotFound, "SystemTelnetServerConfigEnable-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemTelnetServerConfigRateLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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
		return 0, status.Error(codes.NotFound, "SystemTelnetServerConfigRateLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemTelnetServerConfigSessionLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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
		return 0, status.Error(codes.NotFound, "SystemTelnetServerConfigSessionLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemTelnetServerConfigTimeout(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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
		return 0, status.Error(codes.NotFound, "SystemTelnetServerConfigTimeout-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemTelnetServerStateEnable(ctx context.Context, target string,
) (bool, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return false, status.Error(codes.NotFound, "SystemTelnetServerStateEnable-not-found")
	}

	return val.GetBoolVal(), nil
}

func (c *GnmiClient) Get_SystemTelnetServerStateRateLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemTelnetServerStateRateLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemTelnetServerStateSessionLimit(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemTelnetServerStateSessionLimit-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Get_SystemTelnetServerStateTimeout(ctx context.Context, target string,
) (uint16, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
		return 0, status.Error(codes.NotFound, "SystemTelnetServerStateTimeout-not-found")
	}

	return uint16(val.GetUintVal()), nil
}

func (c *GnmiClient) Update_SystemAaaAuthenticationAdminUserConfigAdminPassword(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
				},
				{
					Name: "admin-password-hashed",
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

func (c *GnmiClient) Update_SystemAaaAuthenticationAdminUserConfigAdminPasswordHashed(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "config",
				},
				{
					Name: "admin-password-hashed",
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

func (c *GnmiClient) Update_SystemAaaAuthenticationAdminUserStateAdminPassword(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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

func (c *GnmiClient) Update_SystemAaaAuthenticationAdminUserStateAdminPasswordHashed(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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

func (c *GnmiClient) Update_SystemAaaAuthenticationAdminUserStateAdminUsername(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "aaa",
				},
				{
					Name: "authentication",
				},
				{
					Name: "admin-user",
				},
				{
					Name: "state",
				},
				{
					Name: "admin-username",
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

func (c *GnmiClient) Update_SystemClockConfigTimezoneName(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "config",
				},
				{
					Name: "timezone-name",
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

func (c *GnmiClient) Update_SystemClockStateTimezoneName(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "clock",
				},
				{
					Name: "state",
				},
				{
					Name: "timezone-name",
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

func (c *GnmiClient) Update_SystemConfigDomainName(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "domain-name",
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

func (c *GnmiClient) Update_SystemConfigHostname(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "hostname",
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

func (c *GnmiClient) Update_SystemConfigLoginBanner(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "login-banner",
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

func (c *GnmiClient) Update_SystemConfigMotdBanner(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "config",
				},
				{
					Name: "motd-banner",
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

func (c *GnmiClient) Update_SystemDnsConfigSearch(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "config",
				},
				{
					Name: "search",
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

func (c *GnmiClient) Update_SystemDnsStateSearch(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "dns",
				},
				{
					Name: "state",
				},
				{
					Name: "search",
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

func (c *GnmiClient) Update_SystemMemoryStatePhysical(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
				},
				{
					Name: "reserved",
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

func (c *GnmiClient) Update_SystemMemoryStateReserved(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "memory",
				},
				{
					Name: "state",
				},
				{
					Name: "reserved",
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

func (c *GnmiClient) Update_SystemNtpConfigEnableNtpAuth(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
				},
				{
					Name: "enabled",
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

func (c *GnmiClient) Update_SystemNtpConfigEnabled(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "config",
				},
				{
					Name: "enabled",
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

func (c *GnmiClient) Update_SystemNtpStateAuthMismatch(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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

func (c *GnmiClient) Update_SystemNtpStateEnableNtpAuth(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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

func (c *GnmiClient) Update_SystemNtpStateEnabled(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ntp",
				},
				{
					Name: "state",
				},
				{
					Name: "auth-mismatch",
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

func (c *GnmiClient) Update_SystemOpenflowAgentConfigBackoffInterval(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "backoff-interval",
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

func (c *GnmiClient) Update_SystemOpenflowAgentConfigDatapathId(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "datapath-id",
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

func (c *GnmiClient) Update_SystemOpenflowAgentConfigFailureMode(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "failure-mode",
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

func (c *GnmiClient) Update_SystemOpenflowAgentConfigInactivityProbe(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "inactivity-probe",
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

func (c *GnmiClient) Update_SystemOpenflowAgentConfigMaxBackoff(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "config",
				},
				{
					Name: "max-backoff",
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

func (c *GnmiClient) Update_SystemOpenflowAgentStateBackoffInterval(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "backoff-interval",
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

func (c *GnmiClient) Update_SystemOpenflowAgentStateDatapathId(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "datapath-id",
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

func (c *GnmiClient) Update_SystemOpenflowAgentStateFailureMode(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "failure-mode",
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

func (c *GnmiClient) Update_SystemOpenflowAgentStateInactivityProbe(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "inactivity-probe",
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

func (c *GnmiClient) Update_SystemOpenflowAgentStateMaxBackoff(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "openflow",
				},
				{
					Name: "agent",
				},
				{
					Name: "state",
				},
				{
					Name: "max-backoff",
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

func (c *GnmiClient) Update_SystemSshServerConfigEnable(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerConfigProtocolVersion(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerConfigRateLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerConfigSessionLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerConfigTimeout(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "config",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerStateEnable(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerStateProtocolVersion(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerStateRateLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerStateSessionLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemSshServerStateTimeout(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "ssh-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemStateBootTime(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "boot-time",
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

func (c *GnmiClient) Update_SystemStateCurrentDatetime(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "current-datetime",
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

func (c *GnmiClient) Update_SystemStateDomainName(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "domain-name",
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

func (c *GnmiClient) Update_SystemStateHostname(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "hostname",
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

func (c *GnmiClient) Update_SystemStateLoginBanner(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "login-banner",
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

func (c *GnmiClient) Update_SystemStateMotdBanner(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "state",
				},
				{
					Name: "motd-banner",
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

func (c *GnmiClient) Update_SystemTelnetServerConfigEnable(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Update_SystemTelnetServerConfigRateLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Update_SystemTelnetServerConfigSessionLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Update_SystemTelnetServerConfigTimeout(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "config",
				},
				{
					Name: "enable",
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

func (c *GnmiClient) Update_SystemTelnetServerStateEnable(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemTelnetServerStateRateLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemTelnetServerStateSessionLimit(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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

func (c *GnmiClient) Update_SystemTelnetServerStateTimeout(ctx context.Context, target string, val *gnmi.TypedValue,
) (*gnmi.SetResponse, error) {
	gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	path := []*gnmi.Path{
		{
			Elem: []*gnmi.PathElem{
				{
					Name: "system",
				},
				{
					Name: "telnet-server",
				},
				{
					Name: "state",
				},
				{
					Name: "timeout",
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
