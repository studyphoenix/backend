#!/bin/bash

set -e
swag init --md ./apidocs/ --requiredByDefault
if [[ $1 == "testing" ]]; then
    TEST_MODE=true go run main.go
else
    go run main.go
fi

