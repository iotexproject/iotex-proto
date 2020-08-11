#!/bin/bash

mkdir -p ./test/mock

mkdir -p ./test/mock/mock_iotexapi
mockgen -destination=./test/mock/mock_iotexapi/mock_iotexapi.go \
  github.com/iotexproject/iotex-proto/golang/iotexapi \
  APIServiceServer
