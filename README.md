# grpc-transcoding-experiment
Test gRPC-JSON transcoder filter 

Terminal ~/go/src/grpc_transcoder$ 
1. protoc -I/home/james/go/src/googleapis -I/usr/local/include -I helloworld/ --include_imports --include_source_info --descriptor_set_out=helloworld/helloworld.pb helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

2. go run grpc_service.go


Terminal ~/go/src/istio.io/envoy$
https://github.com/envoyproxy/envoy/blob/master/bazel/README.md
bazel build //source/exe:envoy-static -c dbg


Terminal ~/go/src/grpc_transcoder/envoy_config$
1. ln -s /home/james/go/src/istio.io/envoy/bazel-bin/source/exe/envoy-static envoy-static
2. sudo ./envoy-static -c envoy.conf 


Terminal ~/go/src/grpc_transcoder$ 
1. go run grpc_client.go 
2018/01/15 22:56:00 could not greet: rpc error: code = FailedPrecondition desc = transport: received the unexpected content-type "text/plain"
exit status 1

2. curl -X POST http://localhost:8000/simple/v0.1.0/hello -d '{"name":"timeout"}' -v
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

