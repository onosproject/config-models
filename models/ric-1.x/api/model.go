// Code generated by model-compiler. DO NOT EDIT.

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/openconfig/gnmi/proto/gnmi"
)

var modelData = []*gnmi.ModelData{
	{Name: "kpimon-xapp", Organization: "Open Networking Foundation", Version: "2020-12-04"},
	{Name: "xapp", Organization: "Open Networking Foundation", Version: "2020-11-24"},
}

var encodings = []gnmi.Encoding{gnmi.Encoding_JSON_IETF}

func ModelData() []*gnmi.ModelData {
	return modelData
}

func Encodings() []gnmi.Encoding {
	return encodings
}