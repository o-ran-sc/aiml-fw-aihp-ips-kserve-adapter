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

# Install Helm via the official upstream installer. The previous
# baltocdn.com apt-repo approach broke for two compounding reasons:
# 1. 'apt-key add' was deprecated in Debian Buster and removed entirely
#    from Debian Bookworm onward; on Bullseye it now exits non-zero.
# 2. The baltocdn.com Helm CDN itself has been retired (NXDOMAIN from
#    Google 8.8.8.8 and Cloudflare 1.1.1.1), so the repository URL no
#    longer resolves regardless of the keyring mechanism used.
# The upstream 'get-helm-3' script downloads the latest Helm v3.x
# binary from get.helm.sh, verifies its checksum, and installs it into
# /usr/local/bin. This mirrors the install method documented at
# https://helm.sh/docs/intro/install/#from-script.
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl ca-certificates && \
    curl -fsSL https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /kserve-adapter/kserve-adapter .
COPY --from=builder /kserve-adapter/pkg/helm/data pkg/helm/data

ENV API_SERVER_PORT=10000 CHART_WORKSPACE_PATH="/root/pkg/helm/data"


EXPOSE 10000
ENTRYPOINT ["./kserve-adapter"]