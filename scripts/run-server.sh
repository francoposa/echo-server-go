#!/usr/bin/env sh

set -o errexit -o xtrace

echo "****************Starting echo-server service...****************"

cd /app
exec ./echo-server server --config config.docker.yaml
