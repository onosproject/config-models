{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/}}

{{/*This template generates methods to interact with a single item in the list*/}}

{{ define "keys_parameters"}}
    {{- $key := . -}}
    {{- $has_parent := hasParentKey $key -}}
    {{ if $has_parent }}
        {{ template "keys_parameters" $key.ParentKey }}
    {{ end }}
{{ $kl := len $key.Keys }}
{{- if eq $kl 0 -}}
    {{ replace "-" "" (index $key.Keys 0).Name }} {{ lower (index $key.Keys 0).Type }}
{{- else -}}
    {{ replace "-" "" $key.ModelName }} {{ $key.Type }},
{{- end -}}
{{ end }}

{{/* takes a map {methdo: string, key: ListKey}*/}}
{{ define "parent_key" }}
{{- $key := .key -}}
{{- $method := .method -}}
{{- $has_parent := hasParentKey $key -}}
{{ if $has_parent }}
    {{ template "parent_key" dict "method" $method "key" $key.ParentKey }}
{{ end }}
{
Name: "{{ $key.ModelName }}",
Key: map[string]string{
{{ range $k := $key.Keys -}}
    {{ if or (eq $method "get") (eq $method "delete") -}}
        {{ if eq (len $key.Keys) 1}}
            "{{ lower $k.Name }}": fmt.Sprint({{ replace "-" "" $key.ModelName }}),
        {{ else }}
            "{{ lower $k.Name }}": fmt.Sprint(key.{{ replace "-" "" $k.Name }}),
        {{ end }}
    {{ else if eq $method "update" -}}
        "{{ lower $k.Name }}": fmt.Sprint(*data.{{ replace "-" "" $k.Name }}),
    {{ end -}}
{{ end -}}
},
},
{{ end }}

{{/* take a ListKey as a paramenter*/}}
{{ define "return_value" }}
    {{- $key := . -}}
    {{- $has_parent := hasParentKey $key -}}
    {{- if $has_parent -}}{{- template "return_value" $key.ParentKey -}}{{- end -}}.{{ toName $key.ModelName }}[{{ replace "-" "" $key.ModelName }}]
{{- end -}}

{{ $ep := . }}
{{ $path_length := len $ep.Path }}
func (c *GnmiClient) {{ $ep.MethodName }}(ctx context.Context, target string, {{ if or (eq $ep.Method "get") (eq $ep.Method "delete")}}{{ template "keys_parameters" $ep.Key }}{{end}}{{ if eq $ep.Method "update"}} data {{$ep.ModelName}},{{end}}
) ({{ if eq $ep.Method "get"}}*{{ $ep.ModelName }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    {{/*We can use the partial template as we need to account for Keys*/}}
    {{ if eq $ep.Method "update"}}

    {{ end }}
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
                {{ template "parent_key" dict "method" $ep.Method "key" $ep.Key }}
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

{{/*    {{ if not (eq $ep.ParentModelPath "") -}}*/}}
{{/*    if reflect.ValueOf(st.{{ $ep.ParentModelPath }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ $ep.ParentModelPath }}).IsNil() {*/}}
{{/*        return nil, status.Error(codes.NotFound, "{{ $ep.ModelName }}-not-found")*/}}
{{/*    }*/}}
{{/*    {{ end -}}*/}}

    if res, ok := st{{ template "return_value" .Key}}; ok {
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