{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/}}

{{/*This template generates methods to interact with a single item in the list*/}}

{{ define "keys_parameters"}}
{{ $kl := len .Keys }}
{{- if eq $kl 0 -}}
key {{ lower (index .Keys 0).Type }}
{{- else -}}
key {{ .Type }}
{{- end -}}
{{ end }}

{{ $ep := . }}
{{ $path_length := len $ep.Path }}
func (c *GnmiClient) {{ $ep.MethodName }}_Item(ctx context.Context, target string, {{ if or (eq $ep.Method "get") (eq $ep.Method "delete")}}{{ template "keys_parameters" $ep.Key }},{{end}}{{ if eq $ep.Method "update"}} data {{$ep.ModelName}},{{end}}
) ({{ if eq $ep.Method "get"}}*{{ $ep.ModelName }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    {{/*We can use the partial template as we need to account for Keys*/}}
    {{ if eq $ep.Method "update"}}

    {{ end }}
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {{  range $p := $ep.ParentPath -}}
                {
                    Name: "{{ $p }}",
                },
            {{ end -}}
                {
                    Name: "{{index $ep.Path (sub $path_length 1)}}",
                    Key: map[string]string{
                        {{ range $k := $ep.Key.Keys -}}
                        {{ if or (eq $ep.Method "get") (eq $ep.Method "delete") -}}
                        {{ if eq (len $ep.Key.Keys) 1}}
                        "{{ lower $k.Name }}": string(key),
                        {{ else }}
                        "{{ lower $k.Name }}": string(key.{{ replace "-" "" $k.Name }}),
                        {{ end }}
                        {{ else if eq $ep.Method "update" -}}
                        "{{ lower $k.Name }}": string(*data.{{ replace "-" "" $k.Name }}),
                        {{ end -}}
                        {{ end -}}
                    },
                },
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

    if res, ok := st.{{ $ep.ModelPath }}[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")
    {{ end -}}

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