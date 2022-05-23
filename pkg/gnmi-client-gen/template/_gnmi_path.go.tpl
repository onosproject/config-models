{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/}}

{{ $path := . -}}
path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {{  range $p := $path -}}
            {
                Name: "{{ $p }}",
            },
        {{ end -}}
        },
        Target: target,
    },
}