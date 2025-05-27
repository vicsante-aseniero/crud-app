#script.sh

#!/bin/sh
go version
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --version
dapr version
echo "Done version verifications and installing needed package(s)....."