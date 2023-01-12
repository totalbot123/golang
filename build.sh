#!/bin/bash

MODULE_NAME=$1

cd $MODULE_NAME

go mod init $MODULE_NAME
go mod edit -replace example.com/protobuff=../proto
go mod tidy
go mod vendor
echo "Building Module"
go build -o ./bin/$MODULE_NAME