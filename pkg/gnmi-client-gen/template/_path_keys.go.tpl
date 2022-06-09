{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{/* iterates over all the nested items to find all the keys for nested lists*/}}

{{ define "path_key_elem" -}}
{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ $entry := . -}}
"{{$entry.Name}}": "{{$entry.Key}}",
{{ end -}}

{{ $entry := . -}}
pathKeys := gnmi_utils.PathToKey{
{{ range $i, $e := $entry.Dir -}}
{{ if isList $e -}}
{{ template "path_key_elem" $e }}
{{ end -}}
{{ end -}}
}