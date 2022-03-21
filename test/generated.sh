#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2022 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

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