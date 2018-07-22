#!/usr/bin/env bash

while sleep 1; do
  find . \
    \( -iname '*.go' -o -iname '*.yml' -o -iname 'test.sh' -o -iname '*.json' \) \
    -a -not -path './tmp/*' \
    -a -not -path './vendor/*' \
    -a -not -iname '*_test.go' \
    | entr -dr go run cmd/grouper/main.go
#    | entr -dr ./test.sh
done