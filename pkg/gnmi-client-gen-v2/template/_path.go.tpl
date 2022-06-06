{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{/*    The path_elem template recurse on the yang.Entry.Parent to create the correct path for nested models */}}
{{ define "path_elem" -}}
{{ $entry := . -}}
{{ $has_parent := hasParent $entry -}}
{{ if and $has_parent -}}
{{ if ne $entry.Parent.Name "device" -}}
{{ template "path_elem" $entry.Parent }}
{{ end -}}
{{ end -}}
            {
                Name: "{{ $entry.Name }}",
                {{ if isList $entry -}}
                Key: map[string]string{

                    {{ range $i, $k := listKeys $entry -}}
                    "{{ $k.Name }}": fmt.Sprint({{ sanitize $k.Name }}),
                    {{ end -}}
                },
                {{ end -}}
            },
{{ end -}}


{{ $entry := . -}}
path := []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
            {{ template "path_elem" $entry }}
        },
        Target: target,
    },
}