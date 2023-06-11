gen-protos:
	protoc --go_out=. --go-grpc_out=. protos/*.proto