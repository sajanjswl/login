



protoc --proto_path=../api/proto/v1 --proto_path=../third_party --go_out=../pkg/api/v1 --go-grpc_out=../pkg/api/v1 user-service.proto

protoc --proto_path=../api/proto/v1 --proto_path=../third_party --grpc-gateway_out=logtostderr=true:../pkg/api/v1 user-service.proto

protoc --proto_path=../api/proto/v1 --proto_path=../third_party --swagger_out=logtostderr=true:../api/swagger/v1 user-service.proto
