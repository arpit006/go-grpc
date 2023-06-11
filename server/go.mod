module grpc-test/server

go 1.20

require (
	golang.org/x/net v0.10.0
	google.golang.org/grpc v1.55.0
	grpc-test/protos v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace grpc-test/protos => /Users/arpitsrivastava/GoLandProjects/personal/grpc-test/protos
