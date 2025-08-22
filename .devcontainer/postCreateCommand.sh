#!/bin/sh
sudo apt-get update && sudo apt-get upgrade -y
sudo apt-get install -y xdg-utils
sudo apt-get update && sudo apt-get install -y net-tools

echo "Installing needed package(s)....."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go version
protoc --version
dapr version
echo "Done version verifications and installing needed package(s)....."