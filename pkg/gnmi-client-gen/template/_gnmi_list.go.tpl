{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/}}

{{/*This template generates methods to interact with the list object*/}}

{{ $ep := . }}
// TODO account for recursive keys as parameters (needed for nested lists)
func (c *GnmiClient) {{ $ep.PluralMethodName }}(ctx context.Context, target string, {{ if eq $ep.Method "update"}} list map[{{ $ep.Key.Type }}]*{{$ep.ModelName}},{{end}}
) ({{ if eq $ep.Method "get"}}map[{{ $ep.Key.Type }}]*{{ $ep.ModelName }}{{ else }}*gnmi.SetResponse{{ end }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ if or (eq $ep.Method "get") (eq $ep.Method "delete") -}}
        {{ template "_gnmi_path.go.tpl" $ep.Path -}}
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

{{ if eq $ep.Method "update" -}}
    basePathElems :=  []*gnmi.PathElem{
    {{  range $p := $ep.ParentPath -}}
        {
        Name: "{{ $p }}",
        },
    {{ end -}}
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    {{ range $k := $ep.Key.Keys -}}
                    "{{ lower $k.Name }}": fmt.Sprint(*item.{{ replace "-" "" $k.Name }}),
                    {{ end -}}
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
{{ end -}}
}