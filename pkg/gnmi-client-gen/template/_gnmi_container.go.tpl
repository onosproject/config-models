{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/}}
{{ $ep := . }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string, {{ if eq $ep.Method "update"}} data {{$ep.ModelName}},{{end}}
) ({{ if eq $ep.Method "get"}}*{{ $ep.ModelName }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()

{{ template "_gnmi_path.go.tpl" $ep.Path }}

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