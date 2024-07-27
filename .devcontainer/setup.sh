#!/usr/bin/env bash

set -eux

# Copies over welcome message
sudo cp .devcontainer/welcome.txt /usr/local/etc/vscode-dev-containers/first-run-notice.txt

# setup migrate command
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
sudo install migrate=4.17.1

# setup sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# setup mockgen
go install github.com/golang/mock/mockgen@v1.6.0
