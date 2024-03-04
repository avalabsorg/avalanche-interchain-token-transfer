#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e
set -o pipefail

# The go version for this project is set from a combination of major.minor from go.mod and the patch version set here.
GO_PATCH_VERSION=7

REPO_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

# Pass in the full name of the dependency.
# Parses go.mod for a matching entry and extracts the version number.
function getDepVersion() {
    grep -m1 "^\s*$1" $REPO_PATH/go.mod | cut -d ' ' -f2
}

extract_commit() {
  local version=$1
  if [[ $version == *-* ]]; then
      # Extract the substring after the last '-'
      version=${version##*-}
  fi
  echo "$version"
}

export GO_VERSION=${GO_VERSION:-$(getDepVersion go).$GO_PATCH_VERSION}
AVALANCHEGO_VERSION=${AVALANCHEGO_VERSION:-$(getDepVersion github.com/ava-labs/avalanchego)}
GINKGO_VERSION=${GINKGO_VERSION:-$(getDepVersion github.com/onsi/ginkgo/v2)}
SUBNET_EVM_VERSION=${SUBNET_EVM_VERSION:-$(getDepVersion github.com/ava-labs/subnet-evm)}

# Set golangci-lint version
GOLANGCI_LINT_VERSION=${GOLANGCI_LINT_VERSION:-'v1.55'}
