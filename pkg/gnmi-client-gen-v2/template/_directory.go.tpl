{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ $entry := . -}}

func (c *GnmiClient) Get_{{ template "_entry_name.go.tpl" $entry }}(ctx context.Context, target string) (*{{ structName $entry }}, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    {{ template "_path.go.tpl" $entry }}

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

    if reflect.ValueOf(st.{{ devicePath $entry }}).Kind() == reflect.Ptr && reflect.ValueOf(st.Cont1A).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTest1_Cont1A-not-found")
    }

    return st.{{ devicePath $entry }}, nil

}

{{/*    Once the methods for the container have been generated, descend into it*/}}
{{ template "_entry.go.tpl" $entry }}