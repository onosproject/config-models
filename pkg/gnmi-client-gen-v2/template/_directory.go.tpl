{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{- /*gotype: github.com/openconfig/goyang/pkg/yang.Entry */ -}}
{{ $entry := . -}}

// FIXME generate methods for directories

{{/*    Once the methods for the container have been generated, descend into it*/}}
{{ template "_entry.go.tpl" $entry }}