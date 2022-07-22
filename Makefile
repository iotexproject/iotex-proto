########################################################################################################################
# Copyright (c) 2018 IoTeX
# This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
# warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
# permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
# License 2.0 that can be found in the LICENSE file.
########################################################################################################################

# Go parameters
GOCMD=go
GOLINT=golint
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

PKG_PATH=/github.com/iotexproject/iotex-proto/golang

.PHONY: gogen
gogen:
	@mkdir -p ./temp
	@protoc --go_out=./temp  --go-grpc_out=require_unimplemented_servers=false:./temp ./proto/types/*
	@protoc --go_out=./temp --go-grpc_out=require_unimplemented_servers=false:./temp ./proto/rpc/*
	@protoc --go_out=./temp --go-grpc_out=require_unimplemented_servers=false:./temp ./proto/testing/*
	@protoc -I. -I./proto/types --go_out=./temp --go-grpc_out=require_unimplemented_servers=false:./temp ./proto/api/*
	@protoc -I. --grpc-gateway_out=logtostderr=true:./temp ./proto/api/*
	@rm -rf ./golang/iotexapi ./golang/iotexrpc ./golang/iotextypes ./golang/testingpb
	@cp -r ./temp/${PKG_PATH}/* ./golang
	@rm -rf ./temp
.PHONY: mockgen
mockgen:
	@./misc/scripts/mockgen.sh

.PHONY: gen
gen: gogen mockgen
