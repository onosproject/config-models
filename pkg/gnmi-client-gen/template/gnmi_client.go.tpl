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

func New{{ .BaseModel }}GnmiClient(conn *grpc.ClientConn) *GnmiClient {
    gnmi_client := gnmi.NewGNMIClient(conn)
    return &GnmiClient{client: gnmi_client}
}

{{ range $ep := .ContainerEndpoints }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string, {{ if eq $ep.Method "SET"}} data {{$ep.ModuleName}}_{{$ep.ModelName}},{{end}}
) ({{ if eq $ep.Method "GET"}}*{{ $ep.ModuleName }}_{{ $ep.ModelName }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{/* TODO create a template helper to write the path, is duplicated in the LeavesEndpoint loop */}}
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {{  range $p := $ep.Path -}}
                {
                Name: "{{ $p }}",
                },
            {{ end -}}
            },
            Target: target,
        },
    }

    {{ if eq $ep.Method "GET" -}}
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

    return st.{{ $ep.ModelPath }}, nil
    {{ end }}
    {{ if eq $ep.Method "SET" -}}
    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    {{ end -}}
}
{{ end }}

{{ range $ep := .LeavesEndpoints }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string,{{ if eq $ep.Method "SET"}} val *gnmi.TypedValue,{{end}}
    ) ({{ if eq $ep.Method "GET"}}{{ $ep.GoType }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
                {{  range $p := $ep.Path -}}
                {
                    Name: "{{ $p }}",
                },
                {{ end -}}
            },
            Target: target,
        },
    }

    {{ if eq $ep.Method "GET" -}}
    req := &gnmi.GetRequest{
        Encoding:  gnmi.Encoding_PROTO,
        Path:      path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return {{ $ep.GoEmptyReturnType }}, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return {{ $ep.GoEmptyReturnType }}, err
    }

    return {{ $ep.GoReturnType }}, nil
    {{ end -}}

    {{ if eq $ep.Method "SET" -}}
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{
            {
                Path: path[0],
                Val:  val,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    {{ end -}}
}

{{ end }}