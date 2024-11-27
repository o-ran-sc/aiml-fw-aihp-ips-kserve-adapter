# ==================================================================================#
#
# Copyright (c) 2023 Samsung Electronics Co., Ltd. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# 	http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#==================================================================================

FROM golang:1.19.8-bullseye AS builder

WORKDIR /kserve-adapter
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

COPY . .
RUN go install github.com/golang/mock/mockgen@v1.6.0 && go generate ./... && \
    go mod tidy && \
    go build -o kserve-adapter cmd/kserve-adapter/main.go


FROM golang:1.19.8-bullseye
WORKDIR /root/

# Install Helm in a single layer
RUN curl https://baltocdn.com/helm/signing.asc | apt-key add - && \
    apt-get install -y --no-install-recommends apt-transport-https && \
    echo "deb https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends helm && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /kserve-adapter/kserve-adapter .
COPY --from=builder /kserve-adapter/pkg/helm/data pkg/helm/data

ENV API_SERVER_PORT=10000 CHART_WORKSPACE_PATH="/root/pkg/helm/data"


EXPOSE 10000
ENTRYPOINT ["./kserve-adapter"]