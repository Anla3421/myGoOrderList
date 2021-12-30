// protoc --go_out=. --go_opt=paths=source_relative \
// --go-grpc_out=. --go-grpc_opt=paths=source_relative \
// mygrpc/go.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/getInfo.proto

// protoc --go_out=. \
// --go-grpc_out=. \
// mygrpc/go.proto


// protoc -I proto/ --go_out=plugins=grpc:calculatorPB proto/getInfo.proto
