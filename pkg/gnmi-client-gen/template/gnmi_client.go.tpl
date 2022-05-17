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

{{ range $ep := .ListEndpoints }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string, {{ if eq $ep.Method "update"}} data {{$ep.ModuleName}}_{{$ep.ModelName}},{{end}}
) ({{ if eq $ep.Method "get"}}map[{{ $ep.Key.Type }}]*{{ $ep.ModuleName }}_{{ $ep.ModelName }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ if or (eq $ep.Method "get") (eq $ep.Method "delete") -}}
    {{/* TODO create a template helper to write the path, is duplicated in the LeavesEndpoint loop */ -}}
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
    {{ end -}}

    {{/* NOTE get is the same as the ContainerEndpoint */}}
    {{ if eq $ep.Method "get" -}}

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

    {{ if not (eq $ep.ParentModelPath "") -}}
    if reflect.ValueOf(st.{{ $ep.ParentModelPath }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ $ep.ParentModelPath }}).IsNil() {
        return nil, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")
    }
    {{ end -}}

    if reflect.ValueOf(st.{{ $ep.ModelPath }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ $ep.ModelPath }}).IsNil() {
        return nil, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")
    }

    return st.{{ $ep.ModelPath }}, nil
    {{ end }}

    {{ if eq $ep.Method "delete" -}}
        req := &gnmi.SetRequest{
            Delete: []*gnmi.Path{
                {
                    Elem:   path[0].Elem,
                    Target: target,
                },
            },
        }
        return c.client.Set(gnmiCtx, req)
    {{ end -}}
}
{{ end }}

{{ range $ep := .ContainerEndpoints }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string, {{ if eq $ep.Method "update"}} data {{$ep.ModuleName}}_{{$ep.ModelName}},{{end}}
) ({{ if eq $ep.Method "get"}}*{{ $ep.ModuleName }}_{{ $ep.ModelName }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
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

    {{ if eq $ep.Method "get" -}}
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

    {{ if not (eq $ep.ParentModelPath "") -}}
    if reflect.ValueOf(st.{{ $ep.ParentModelPath }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ $ep.ParentModelPath }}).IsNil() {
        return nil, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")
    }
    {{ end -}}

    if reflect.ValueOf(st.{{ $ep.ModelPath }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ $ep.ModelPath }}).IsNil() {
        return nil, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")
    }

    return st.{{ $ep.ModelPath }}, nil
    {{ end }}
    {{ if eq $ep.Method "update" -}}
    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    {{ end -}}

    {{ if eq $ep.Method "delete" -}}
        req := &gnmi.SetRequest{
            Delete: []*gnmi.Path{
                {
                    Elem:   path[0].Elem,
                    Target: target,
                },
            },
        }
        return c.client.Set(gnmiCtx, req)
    {{ end -}}
}
{{ end }}

{{ range $ep := .LeavesEndpoints }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string,{{ if eq $ep.Method "update"}} val *gnmi.TypedValue,{{end}}
    ) ({{ if eq $ep.Method "get"}}{{ $ep.GoType }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
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

    {{ if eq $ep.Method "get" -}}
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

    if {{ $ep.GoReturnType }} ==  {{ $ep.GoEmptyReturnType }} {
        return {{ $ep.GoEmptyReturnType }}, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")
    }

    return {{ $ep.GoReturnType }}, nil
    {{ end -}}

    {{ if eq $ep.Method "update" -}}
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

    {{ if eq $ep.Method "delete" -}}
        req := &gnmi.SetRequest{
            Delete: []*gnmi.Path{
                {
                    Elem:   path[0].Elem,
                    Target: target,
                },
            },
        }
        return c.client.Set(gnmiCtx, req)
    {{ end -}}
}

{{ end }}