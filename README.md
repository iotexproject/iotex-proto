# iotex-proto
Protobuf and utility package for IoTeX block transaction and gRPC API

- \proto includes protobuf definition for all core data objects and gRPC API used by IoTeX blockchain

- \golang includes the generated protobuf files for go language

# Getting Started
## Installing
Install the Google protocol buffers compiler `protoc` v3.0.0 or above from https://github.com/protocolbuffers/protobuf/releases

Enable go mod. Install grpc-gateway https://github.com/grpc-ecosystem/grpc-gateway. Basically this is what you need:

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Compiling
```
make gogen
```
This generates the protobuf files and put into \golang directory
