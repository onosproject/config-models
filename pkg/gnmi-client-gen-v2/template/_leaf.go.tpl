{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{/*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}

{{ $entry := . -}}
func (c *GnmiClient) Get_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string, {{ template "_list_keys.go.tpl" $entry.Parent }}) ({{ goType $entry }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" $entry }}

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