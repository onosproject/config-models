/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package template

import "embed"

//go:embed *.go.tpl
var GnmiGenTemplate embed.FS
