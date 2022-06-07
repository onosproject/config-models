{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{/*    The path_elem template recurse on the yang.Entry.Parent to create the correct path for nested models */}}
{{ define "path_elem" -}}
{{- /*gotype: map[string]interface{} */ -}}
{{ $entry := .entry -}}
{{ $has_parent := hasParent $entry -}}
{{ if $has_parent -}}
{{ if not (isRoot $entry.Parent) -}}
{{ template "path_elem" dict "entry" $entry.Parent "forList" false }}
{{ end -}}
{{ end -}}
            {
                Name: "{{ $entry.Name }}",
                {{ if and (isList $entry) (not .forList) -}}
                Key: map[string]string{

                    {{ range $i, $k := listKeys $entry -}}
                    "{{ $k.Key }}": fmt.Sprint({{ sanitize $k.Name }}),
                    {{ end -}}
                },
                {{ end -}}
            },
{{ end -}}

path := []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
            {{ template "path_elem" . }}
        },
        Target: target,
    },
}