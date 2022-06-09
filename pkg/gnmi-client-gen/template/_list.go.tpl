{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}

{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ $entry := . -}}
func (c *GnmiClient) Get_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ template "_list_keys.go.tpl" dict "entry" $entry "forList" false }}) ({{ structName $entry false }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" dict "entry" $entry "forList" false }}

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

    if reflect.ValueOf(st.{{ devicePath $entry true }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ devicePath $entry true }}).IsNil() {
        return nil, status.Error(codes.NotFound, "{{ template "_entry_name.go.tpl" $entry }}-not-found")
    }
    if res, ok := st.{{ devicePath $entry false }}; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "{{ template "_entry_name.go.tpl" $entry }}-not-found")
}

{{/*    Lists are containers themselves, so generate a method to get the entire thing*/}}
func (c *GnmiClient) Get_{{ template "_entry_name.go.tpl" $entry }}_List(ctx context.Context, target string, {{ template "_list_keys.go.tpl" dict "entry" $entry "forList" true }}) ({{ structName $entry true }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" dict "entry" $entry "forList" true }}

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

    if reflect.ValueOf(st.{{ devicePath $entry true }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ devicePath $entry true }}).IsNil() {
        return nil, status.Error(codes.NotFound, "{{ template "_entry_name.go.tpl" $entry }}-not-found")
    }
    return st.{{ devicePath $entry true }}, nil
}

{{/*    NOTE: At the moment we are supporting Update and Delete only for single items in the list. */}}
func (c *GnmiClient) Update_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ template "_list_keys.go.tpl" dict "entry" $entry "forList" false -}} data {{ structName $entry false }},
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" dict "entry" $entry "forList" false }}

    {{ template "_path_keys.go.tpl" $entry }}

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *data, path[0], target, pathKeys)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}

func (c *GnmiClient) Delete_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ template "_list_keys.go.tpl" dict "entry" $entry "forList" false -}}) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" dict "entry" $entry "forList" false }}

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

{{/*    Once the methods for the list have been generated, descend into it*/}}
{{ template "_entry.go.tpl" $entry }}