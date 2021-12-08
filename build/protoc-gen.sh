protoc --go_out=$GOPATH/src protos/domain/*.proto
protoc --go_out=$GOPATH/src protos/service/*.proto

protoc --go-grpc_out=$GOPATH/src protos/service/*.proto
