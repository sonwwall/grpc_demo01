protoc -I . --go_out=. --go-grpc_out=. hello.proto
protoc -I . --go_out=. --go-grpc_out=. type.proto
protoc -I . --go_out=. --go-grpc_out=. multi.proto