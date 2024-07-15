#!/usr/bin/env bash

set -eux

# Copies over welcome message
sudo cp .devcontainer/welcome.txt /usr/local/etc/vscode-dev-containers/first-run-notice.txt

# setup migrate command
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# setup sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# setup mockgen
go install github.com/golang/mock/mockgen@v1.6.0
