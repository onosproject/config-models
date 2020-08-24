// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
	openapi_gen "github.com/onosproject/config-models/pkg/openapi-gen"
	"os"
)

func main() {
	schemaMap, err := rbac_1_0_0.Schema()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	schema, err := openapi_gen.BuildOpenapi(schemaMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	json, err := schema.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(json))

}
