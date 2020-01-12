# Go Shopping
This example contains a suite of microservices all built on the `go micro` framework. 

# Protobuf
Protobuf is used for code generation of message types and client/hander stubs.

If making changes recompile the protos.

## Install
Install protoc for your environment. Then:

```shell
go get github.com/golang/protobuf/{proto,protoc-gen-go}
```

```shell
go get github.com/micro/protoc-gen-micro
```

```shell
make proto
```

# Micro API
Run the micro API with custom namespace

```shell
micro api --handler=http --namespace=go.shopping.api
```
