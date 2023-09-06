#!/usr/bin/env bash

set -euo pipefail


goimports -w -local $(grep "^module" go.mod | awk '{print $2}') $(go list -f {{.Dir}} ./... | grep -v /api/ )
gofmt -s -w .
