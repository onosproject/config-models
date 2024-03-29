// Code generated by model-compiler. DO NOT EDIT.

// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/openconfig/gnmi/proto/gnmi"
)

var modelData = []*gnmi.ModelData{
	{Name: "openconfig-interfaces", Organization: "OpenConfig working group", Version: "2017-07-14"},
	{Name: "openconfig-openflow", Organization: "OpenConfig working group", Version: "2017-06-01"},
	{Name: "openconfig-platform", Organization: "OpenConfig working group", Version: "2016-12-22"},
	{Name: "openconfig-system", Organization: "OpenConfig working group", Version: "2017-07-06"},
}

var encodings = []gnmi.Encoding{gnmi.Encoding_JSON_IETF}

func ModelData() []*gnmi.ModelData {
	return modelData
}

func Encodings() []gnmi.Encoding {
	return encodings
}
