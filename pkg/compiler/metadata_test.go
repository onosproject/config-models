// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMetaData(t *testing.T) {
	path := "../../test/metadata"

	md := &MetaData{LintModel: true, RequireHyphenated: true}
	if err := LoadMetaData(path, "valid", md); err != nil {
		t.Fatal(err)
	}

	err := LoadMetaData(path, "not-existing", md)
	assert.Error(t, err)
	assert.True(t, md.LintModel)
	assert.True(t, md.RequireHyphenated)
}

func TestValidateMetaData(t *testing.T) {
	missingName := &MetaData{Name: "", LintModel: true, RequireHyphenated: true}
	err := ValidateMetaData(missingName)
	assert.Error(t, err)
	assert.Equal(t, "name is mandatory", err.Error())
}
