#!/bin/bash

#set -eu

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"
cd "$DIR" || exit

cd ${DIR}/pkg/api/v1/deployment
mkdir mock
mockgen -source=deployment.go -destination=./mock/mock_deployment.go -package=mock

cd ${DIR}/pkg/client/kserve
mkdir mock
mockgen -source=client.go -destination=./mock/mock_client.go -package=mock

cd ${DIR}/pkg/client/onboard
mkdir mock
mockgen -source=client.go -destination=./mock/mock_client.go -package=mock

cd ${DIR}/pkg/controller/v1/adapter
mkdir mock
mockgen -source=controller.go -destination=./mock/mock_controller.go -package=mock