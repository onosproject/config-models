{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ $entry := . -}}
{{ if (hasParent $entry) -}}
{{ template "_list_keys.go.tpl" $entry.Parent -}}
{{ end -}}
{{ range $i, $k := listKeys $entry -}}
    {{ sanitize $k.Name }} {{ $k.Gotype }},
{{ end -}}