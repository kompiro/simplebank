#!/usr/bin/env bash

set -eux

# install yadm
sudo apt update && sudo apt upgrade -y && sudo apt install -y yadm

# Copies over welcome message
sudo cp .devcontainer/welcome.txt /usr/local/etc/vscode-dev-containers/first-run-notice.txt

yadm clone https://github.com/kompiro/yadm/ --bootstrap -f