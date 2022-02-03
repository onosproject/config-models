package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMetaData(t *testing.T) {
	path := "../../test/metadata"

	md := &MetaData{}
	if err := LoadMetaData(path, "valid", md); err != nil {
		t.Fatal(err)
	}

	err := LoadMetaData(path, "not-existing", md)
	assert.Error(t, err)
}

func TestValidateMetaData(t *testing.T) {
	missingName := &MetaData{Name: ""}
	err := ValidateMetaData(missingName)
	assert.Error(t, err)
	assert.Equal(t, "name is mandatory", err.Error())
}