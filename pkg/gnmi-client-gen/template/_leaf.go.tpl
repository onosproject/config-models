{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{/*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}

{{ $entry := . -}}
func (c *GnmiClient) Get_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ if hasParent $entry -}}{{ template "_list_keys.go.tpl" dict "entry" $entry.Parent "forList" false -}}{{ end -}}) ({{ goType $entry }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" dict "entry" $entry "forList" false }}

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_PROTO,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return {{ goEmptyReturnVal $entry }}, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return {{ goEmptyReturnVal $entry }}, err
    }

    if {{ goReturnVal $entry }} == {{ goEmptyReturnVal $entry }} {
        return {{ goEmptyReturnVal $entry }}, status.Error(codes.NotFound, "{{ $entry.Name }}-not-found")
    }

    return {{ goReturnVal $entry }}, nil
}

func (c *GnmiClient) Update_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ if hasParent $entry -}}{{ template "_list_keys.go.tpl" dict "entry" $entry.Parent "forList" false -}}{{ end -}} val *gnmi.TypedValue) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" dict "entry" $entry "forList" false }}

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

func (c *GnmiClient) Delete_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ if hasParent $entry -}}{{ template "_list_keys.go.tpl" dict "entry" $entry.Parent "forList" false -}}{{ end -}}) (*gnmi.SetResponse, error) {
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