{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{ $entry := .entry -}}
{{ if (hasParent $entry) -}}
{{ template "_list_keys.go.tpl" dict "entry" $entry.Parent "forList" false -}}
{{ end -}}
{{/*    forList means that we are targeting the container list, so we don't need the last key in the tree */ -}}
{{ if not .forList -}}
{{ range $i, $k := listKeys $entry -}}
    {{ sanitize $k.Name }} {{ $k.Gotype }},
{{ end -}}
{{ end -}}