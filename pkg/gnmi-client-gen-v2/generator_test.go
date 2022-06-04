/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package gnmi_client_gen_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	aether_2_1_x "github.com/onosproject/aether-models/models/aether-2.1.x/api"
	testdevice "github.com/onosproject/config-models/models/testdevice-1.0.x/api"
	"github.com/onosproject/config-models/pkg/gnmi-client-gen-v2/testdata"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExtractSchema(t *testing.T) {
	// uncomment this to manually run the test to regenerate the sample schema
	// not needed otherwise
	t.Skip()
	folder := "sampleSchemas"

	testdevice, err := testdevice.Schema()
	assert.NoError(t, err)

	a21x, err := aether_2_1_x.Schema()
	assert.NoError(t, err)

	schemas := map[string]*ytypes.Schema{
		"testdevice-1": testdevice,
		"aether-2.1.x": a21x,
	}

	for name, schema := range schemas {
		empJSON, err := json.MarshalIndent(schema.SchemaTree["Device"], "", "  ")
		assert.NoError(t, err)

		file, err := os.Create(fmt.Sprintf("%s/%s.json", folder, name))
		if err != nil {
			t.Fail()
		}
		file.WriteString(string(empJSON))
		defer file.Close()
	}

}

func TestGenerate(t *testing.T) {
	type arguments struct {
		pluginName string
		entry      *yang.Entry
	}
	tests := []struct {
		name    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty-entry",
			func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.Equal(t, "entry-cannot-be-nil", err.Error())
			},
		},
		{
			"simple-leaves",
			assert.NoError,
		},
		{
			"basic-container",
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			schema, err := testdata.GetSchema(tt.name)
			if err != nil {
				t.Fatalf(err.Error())
			}

			wantOutput := testdata.GetTestResult(t, tt.name)

			args := arguments{
				pluginName: "Test",
				entry:      schema,
			}

			output := &bytes.Buffer{}
			err = Generate(args.pluginName, args.entry, output)
			if !tt.wantErr(t, err, fmt.Sprintf("Generate(%v, %v, %v)", args.pluginName, args.entry, output)) {
				return
			}

			debug := os.Getenv("DEBUG")
			if debug == "true" {
				// when debugging keep whitespaces, they might fail the test
				// but it's more readable
				fmt.Println(output)
				//assert.Equal(t, wantOutput, output.String())
			} else {
				// when not debugging strip whitespaces as they're not relevant
				assert.Equalf(t, testdata.RemoveAllWhitespaces(wantOutput), testdata.RemoveAllWhitespaces(output.String()), "Failed to generate template for test: %s", tt.name)
			}

		})
	}
}
