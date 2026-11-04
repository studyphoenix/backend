#!/bin/bash

set -e

swag init --md ./apidocs/ && go run main.go
