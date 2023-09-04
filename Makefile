# ==================================================================================
#
#       Copyright (c) 2023 Samsung Electronics Co., Ltd. All Rights Reserved.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#          http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
# ==================================================================================

build:
	go get ./cmd/kserve-adapter
	go build -o kserve-adapter cmd/kserve-adapter/main.go
run:
	KUBECONFIG=~/.kube/config \
        API_SERVER_PORT=10000 \
        CHART_WORKSPACE_PATH="$(shell pwd)/pkg/helm/data" \
        RIC_DMS_IP=127.0.0.1 \
        RIC_DMS_PORT=8000 \
	./kserve-adapter
genmock:
	go generate ./...
fmt:
	go fmt ./cmd/...
	go fmt ./pkg/...
test:
	go mod tidy
	mockgen -source=pkg/client/ricdms/client.go -destination=pkg/client/ricdms/mock/mock_client.go -package=mock
	mockgen -source=pkg/api/v1/deployment/deployment.go -destination=pkg/api/v1/deployment/mock/mock_deployment.go -package=mock
	mockgen -source=pkg/controller/v1/adapter/controller.go -destination=pkg/controller/v1/adapter/mock/mock_controller.go -package=mock
	mockgen -source=pkg/client/kserve/client.go -destination=pkg/client/kserve/mock/mock_client.go -package=mock

	KUBECONFIG=~/.kube/config \
	go test ./pkg/...
