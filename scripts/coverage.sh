#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

if [ $# -eq 0 ]; then
  target="$(go list ./... | fzf)"
  if [ "$target" = "" ]; then
    exit 0
  fi
  target="${target#github.com/phonero/terraform-provider-graylog/}"
elif [ $# -eq 1 ]; then
  target="graylog/$1"
else
  echo "too many arguments are given: $*" >&2
  exit 1
fi

if [ ! -d "$target" ]; then
  echo "$target is not found" >&2
  exit 1
fi

ee mkdir -p ".coverage/$target"
if [ -n "${TF_ACC:-}" ]; then
  ee go test -v "./$target" -coverprofile=".coverage/$target/coverage.txt" -covermode=atomic
else
  ee go test "./$target" -coverprofile=".coverage/$target/coverage.txt" -covermode=atomic
fi
ee go tool cover -html=".coverage/$target/coverage.txt"
