/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package testdata

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const folder = "results"

// Reads the content of a file into a string based on the testname
// It used to keep the expected generated code in readable files
func GetTestResult(t *testing.T, testname string) string {
	path := fmt.Sprintf("testdata/%s/%s.txt", folder, testname)
	file, err := os.ReadFile(path)
	if err != nil {
		// if the file is not there we assume an empty result is expected (in case that we are testing for errors)
		// there's no point in raising an error as the test will fail if a result was expected
		return ""
	}

	return string(file)
}

// removing all spaces and new lines
// we don't care about those as `go fmt` will take care of it
// if we don't do that maintaing the tests would be incredibily painful
func RemoveAllWhitespaces(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}
