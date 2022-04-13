#!/bin/bash

mkdir -p ./golang/iotexapi/mock_iotexapi
mockgen -destination=./golang/iotexapi/mock_iotexapi/mock_iotexapi.go \
  github.com/iotexproject/iotex-proto/golang/iotexapi \
  APIServiceServer,APIServiceClient
