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


FROM golang:1.19.8-bullseye as builder

WORKDIR /kserve-adapter

ENV GO111MODULE=on GOOS=linux GOARCH=amd64

COPY . .

RUN go install github.com/golang/mock/mockgen@v1.6.0 && go generate ./...
RUN go mod tidy
RUN go build -o kserve-adapter cmd/kserve-adapter/main.go

FROM golang:1.19.8-bullseye

WORKDIR /root/

RUN curl https://baltocdn.com/helm/signing.asc | apt-key add -
RUN apt-get install apt-transport-https --yes
RUN echo "deb https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list

RUN apt-get update
RUN apt-get install helm

COPY --from=builder /kserve-adapter/kserve-adapter .
COPY --from=builder /kserve-adapter/pkg/helm/data pkg/helm/data

ENV API_SERVER_PORT=10000
ENV CHART_WORKSPACE_PATH="/root/pkg/helm/data"
EXPOSE 10000

ENTRYPOINT ["./kserve-adapter"]


