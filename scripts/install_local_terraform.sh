#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

ee mkdir -p ~/.terraform.d/plugins/github.com/terraform-provider-graylog/terraform-provider-graylog
ee go build -o ~/.terraform.d/plugins/github.com/terraform-provider-graylog/terraform-provider-graylog ./cmd/terraform-provider-graylog
