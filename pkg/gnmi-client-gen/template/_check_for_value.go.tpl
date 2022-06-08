{{/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/ -}}
{{ $entry := . -}}
{{ if and (hasParent $entry) (not (isRoot $entry.Parent)) -}}
{{ template "_check_for_value.go.tpl" $entry.Parent }}
{{ end -}}
if reflect.ValueOf(st.{{ devicePath $entry true }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ devicePath $entry false }}).IsNil() {
return nil, status.Error(codes.NotFound, "{{ template "_entry_name.go.tpl" $entry }}-not-found")
}
if reflect.ValueOf(st.{{ devicePath $entry false }}).Kind() == reflect.Ptr && reflect.ValueOf(st.{{ devicePath $entry false }}).IsNil() {
    return nil, status.Error(codes.NotFound, "{{ template "_entry_name.go.tpl" $entry }}-not-found")
}