#!/bin/sh

MODULE_NAME=$1

cd $MODULE_NAME

go mod init $MODULE_NAME
go mod edit -replace example.com/protobuff=../proto
go mod tidy
go mod vendor
echo "Building Module"
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o ./bin/$MODULE_NAME