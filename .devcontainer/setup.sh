#!/usr/bin/env bash

set -eux

# Copies over welcome message
sudo cp .devcontainer/welcome.txt /usr/local/etc/vscode-dev-containers/first-run-notice.txt

# setup migrate command
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
sudo apt install migrate=4.17.1

# setup sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# setup mockgen
go install github.com/golang/mock/mockgen@v1.6.0

# install github cli extensions
exts=(
  "seachicken/gh-poi"
  "yusukebe/gh-markdown-preview"
  "nektos/gh-act"
)

for ext in "${exts[@]}"; do
  /home/vscode/.asdf/shims/gh extension install "$ext"
done

# shellcheck disable=SC1091
. /home/vscode/.asdf/asdf.sh

# install ecspresso
asdf plugin add ecspresso
asdf install ecspresso 2.3.0
asdf global ecspresso 2.3.0

# install aws cli v2
asdf plugin add awscli
asdf install awscli 2.17.39
asdf global awscli 2.17.39

# install dbdocs
npm install -g dbdocs
npm install -g @dbml/cli
