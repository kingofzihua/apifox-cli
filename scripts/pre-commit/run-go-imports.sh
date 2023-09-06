#!/usr/bin/env bash
#
# Capture and print stdout, since goimports doesn't use proper exit codes
#
set -e -o pipefail

if ! command -v goimports &> /dev/null ; then
    echo "goimports not installed or available in the PATH" >&2
    echo "please run 'go install golang.org/x/tools/cmd/goimports@latest'" >&2
    exit 1
fi

output="$(goimports -l -w -local $(grep "^module" go.mod | awk '{print $2}') "$@")"

gofmt -s -w .

[[ -z "$output" ]]