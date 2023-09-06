#!/usr/bin/env bash

set -euo pipefail

for file in `find . -name '*.go'`; do
  # Defensive, just in case.
  if [[ -f ${file} ]]; then
    awk '/^import \($/,/^\)$/{if($0=="")next}{print}' ${file} > /tmp/file
    mv /tmp/file ${file}
  fi
done

goimports -w -local $(grep "^module" go.mod | awk '{print $2}') $(go list -f {{.Dir}} ./... | grep -v /api/ )
gofmt -s -w .
