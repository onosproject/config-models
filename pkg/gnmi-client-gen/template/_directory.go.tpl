{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ $entry := . -}}

func (c *GnmiClient) Get_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ template "_list_keys.go.tpl" dict "entry" $entry "forList" false -}}) ({{ structName $entry false }}, error) {
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

    {{/*   Recursively check that the value is set, if not it's a Not Found */}}
    {{ template "_check_for_value.go.tpl" $entry }}

    return st.{{ devicePath $entry false }}, nil

}

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

{{/*    Once the methods for the container have been generated, descend into it*/}}
{{ template "_entry.go.tpl" $entry }}