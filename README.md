# grpc-transcoding-experiment
Test gRPC-JSON transcoder filter 

# Introduction
Set up an experiment environment where a client, an Envoy proxy, and a gRPC backend are running together. A gRPC-JSON transcoder filter is enabled at the Envoy proxy.

We can send HTTP JSON request using curl, or gRPC request using gRPC client, and compare the results.

gRPC-JSON transcoder filter is described here. [[link](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/grpc_json_transcoder_filter "link")]
gRPC-HTTP/JSON transcoding library is here. [[link](https://github.com/grpc-ecosystem/grpc-httpjson-transcoding "link")]

# Experiment Setup
- Install packages
```go
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
go get -u github.com/google/go-genproto
```
- The last step should bring us protoc compiler > 3.0. If not, then Install ProtocolBuffers 3.0.0 or later.
```bash
    mkdir tmp
    cd tmp
    git clone https://github.com/google/protobuf
    cd protobuf
    ./autogen.sh
    ./configure
    make
    make check
    sudo make install
```
Make sure we have google/api/annotations.proto under $GOPATH/src/googleapis.
Make sure protoc is in /usr/local/bin.

- Download folder grpc_transcoder to $GOPATH/src/.

## Create gRPC service config and run gRPC server
- Open terminal, go to $GOPATH/src/grpc_transcoder$, and execute following commands.
1. 
```go
protoc -I$GOPATH/src/googleapis -I/usr/local/include -I helloworld/ --include_imports --include_source_info --descriptor_set_out=helloworld/helloworld.pb helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
2. 
```go
go run grpc_service.go
```
- Open terminal, go to /envoy$, and build Envoy proxy with debug information.
```bash
bazel build //source/exe:envoy-static -c dbg
```
More build instructions are here. [[link](https://github.com/envoyproxy/envoy/blob/master/bazel/README.md "link")]
- Open terminal, and go to $GOPATH/src/grpc_transcoder/envoy_config$
```bash
ln -s PATH-TO-ENVOY/envoy/bazel-bin/source/exe/envoy-static envoy-static
sudo ./envoy-static -c envoy.conf 
```
- Open terminal, go to $GOPATH/src/grpc_transcoder$ 
1. Send a gRPC request.
```bash
go run grpc_client.go 
2018/01/15 22:56:00 could not greet: rpc error: code = FailedPrecondition desc = transport: received the unexpected content-type "text/plain"
exit status 1
```
2. Send an HTTP/JSON request.
```bash
curl -X POST http://localhost:8000/simple/v0.1.0/hello -d '{"name":"timeout"}' -v
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8000 (#0)
> POST /simple/v0.1.0/hello HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Length: 18
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 18 out of 18 bytes
```
