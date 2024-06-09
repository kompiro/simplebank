#!/usr/bin/env bash

set -eux

yadm clone https://github.com/kompiro/yadm/ --bootstrap

# Copies over welcome message
sudo cp .devcontainer/welcome.txt /usr/local/etc/vscode-dev-containers/first-run-notice.txt

# setup migrate command
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest