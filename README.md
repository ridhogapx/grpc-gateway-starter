## Getting Started 
First, you need to install these binaries to use gRPC with `protoc`. 

```sh 
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

Before executing `generate.sh` make sure you have given permission to allow execute file (unix only).

```sh 
$ chmod +x generate.sh
```

After that you can run these shell command
```sh
$ ./generate.sh
```

Now, you've generated protobuf that will be placed in `/server/proto`. You can run gRPC entrypoint and Gateway entrypoint.

```sh
$ go run server/cmd/entrypoint-grpc
```

```sh
$ go run server/cmd/entrypoint-gateway
```

