#!/bin/bash

set -e

swag init --md ./apidocs/ --requiredByDefault && go run main.go
