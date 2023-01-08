MODULE_NAME="location-processing"

go mod init $MODULE_NAME
go mod tidy
go mod vendor
echo "Building Module"
go build -o bin/$MODULE_NAME