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

func New{{ .BaseModel }}GnmiClient(conn *grpc.ClientConn) *GnmiClient {
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

{{ range $ep := .ContainerEndpoints }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string,
) (*{{ $ep.ModuleName }}_{{ $ep.ModelName }}, error) {
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

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    return st.{{ $ep.ModelPath }}, nil
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

    val, err := GetResponseUpdate(res)

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