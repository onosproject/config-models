/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/

// Generated via gnmi-gen.go, do NOT edit

package api

import (
    "context"
    "github.com/onosproject/config-models/pkg/gnmi-client-gen/gnmi_utils"
    "github.com/openconfig/gnmi/proto/gnmi"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "reflect"
    "time"
)

type GnmiClient struct {
    client gnmi.GNMIClient
}

func New{{ .BaseModel }}GnmiClient(conn *grpc.ClientConn) *GnmiClient {
    gnmi_client := gnmi.NewGNMIClient(conn)
    return &GnmiClient{client: gnmi_client}
}

{{ range $ep := .ListEndpoints -}}
    {{ template "_gnmi_list_item.go.tpl" $ep }}
{{ end -}}

{{ range $ep := .ListEndpoints -}}
{{ template "_gnmi_list.go.tpl" $ep }}
{{ end -}}

{{ range $ep := .ContainerEndpoints }}
{{ template "_gnmi_container.go.tpl" $ep }}
{{ end }}

{{ range $ep := .LeavesEndpoints }}
{{ template "_gnmi_leaf.go.tpl" $ep }}
{{ end }}