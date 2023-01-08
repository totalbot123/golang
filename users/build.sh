MODULE_NAME="users"

go mod init $MODULE_NAME
go mod edit -replace example.com/protobuff=../proto
go mod tidy
go mod vendor
echo "Building Module"
go build -o $MODULE_NAME