.. This work is licensed under a Creative Commons Attribution 4.0 International License.
.. http://creativecommons.org/licenses/by/4.0

.. Copyright (c) 2023 Samsung Electronics Co., Ltd. All Rights Reserved.

User-Guide
==========

.. contents::
   :depth: 3
   :local:


Overview
--------
- Kserve Adapter works with the AIML Framework and is used to deploy, delete and update kserve-based apps.

Build Image
-----------
Use the `docker build` command for docker image build.

.. code-block:: none

    kserve-adapter $ docker build -f Dockerfile .

    Sending build context to Docker daemon  93.74MB
    Step 1/11 : FROM golang:1.19.8-bullseye as builder
    ---> b47c7dfaaa93
    Step 2/11 : WORKDIR /kserve-adapter
    ---> Using cache
    ---> 6b397a834cc2
    Step 3/11 : COPY . .
    ---> Using cache
    ---> 6a155a20fbde
    Step 4/11 : ENV GO111MODULE=on GOOS=linux GOARCH=amd64
    ---> Using cache
    ---> f5be56bd7555
    Step 5/11 : RUN go mod download
    ---> Using cache
    ---> 97862b975561
    Step 6/11 : RUN go build -o kserve-adapter main.go
    ---> Using cache
    ---> 52a6ce04c444
    Step 7/11 : FROM golang:1.19.8-bullseye
    ---> b47c7dfaaa93
    Step 8/11 : WORKDIR /root/
    ---> Using cache
    ---> c7870d2fbeba
    Step 9/11 : COPY --from=builder /kserve-adapter/kserve-adapter .
    ---> Using cache
    ---> 4a2f88d946d6
    Step 10/11 : EXPOSE 48099
    ---> Using cache
    ---> 8fbc694241e8
    Step 11/11 : ENTRYPOINT ["./kserve-adapter"]
    ---> Using cache
    ---> d279266b588c
    Successfully built d279266b588c

Environments of Kserver Adapter
---------------------------------------
+-----------------+---------------------------------+
| KUBE_CONFIG     | ex) ~/.kube/config              | 
+-----------------+---------------------------------+
| API_SERVER_PORT | ex) "48099"                     |
+-----------------+---------------------------------+