#!/usr/bin/env bash

# Copyright 2018 the original author or authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eu -o pipefail

GREEN="\e[32m"
RED="\e[31m"
RESET="\e[0m"

echo -e "Making sure generated code is committed."

MODIFIED_FILES=$(git status ./models --porcelain | wc -l)

if [[ "${MODIFIED_FILES}" != "0" ]]; then
  echo -e "${RED}Some generated files have not been committed:${RESET}"
  git status ./models --porcelain
  exit 1
else
  echo -e "${GREEN}Committed code matches generated code.${RESET}"
fi