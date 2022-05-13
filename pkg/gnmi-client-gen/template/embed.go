/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package template

import "embed"

//go:embed gnmi_client.go.tpl
var GnmiGenTemplate embed.FS
