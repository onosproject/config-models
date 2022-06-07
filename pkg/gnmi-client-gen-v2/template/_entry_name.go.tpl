{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{/*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ $entry := . -}}
{{ $has_parent := hasParent $entry -}}
{{ if $has_parent -}}{{ if not (isRoot $entry.Parent) -}}{{ template "_entry_name.go.tpl" $entry.Parent -}}_{{ end -}}{{ end -}}{{ sanitize (capitalize $entry.Name) -}}