{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{/*In this template we are simply switching on the Entry Kind and invoking the correct template*/ -}}

{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ range $key, $entry := .Dir -}}
{{/*    Kind 0 is LeafEntry     */ -}}
{{ if eq $entry.Kind 0 -}}
{{ template "_leaf.go.tpl" $entry }}
{{ end -}}
{{ if isContainer $entry  -}}
{{ template "_directory.go.tpl" $entry }}
{{ end -}}
{{ if isList $entry  -}}
{{ template "_list.go.tpl" $entry }}
{{ end -}}
{{ end -}}
