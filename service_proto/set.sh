protoc -I ./service_proto --go_out=./service_proto --go-grpc_out=./service_proto ./service_proto/video.proto
protoc -I ./service_proto --go_out=./service_proto --go-grpc_out=./service_proto ./service_proto/order.proto
protoc -I ./service_proto --go_out=./service_proto --go-grpc_out=./service_proto ./service_proto/common.proto