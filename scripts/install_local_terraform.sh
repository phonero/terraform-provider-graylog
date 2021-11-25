#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

ee mkdir -p ~/.terraform.d/plugins/registry.terraform.io/terraform-provider-graylog/graylog
ee go build -o ~/.terraform.d/plugins/registry.terraform.io/terraform-provider-graylog/graylog ./cmd/terraform-provider-graylog
