BINARY=users
VERSION=0.0.1
OS_ARCH=darwin_arm64
OS_ARCH_PIPELINE=linux_amd64

.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go mod tidy
	./scripts/swagger.sh
	go build -o ./bin/${BINARY} ./
.PHONY:build

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${VERSION}/${GOOS}_${GOARCH}/${BINARY}
	GOOS=darwin GOARCH=arm64 go build -o ./bin/${VERSION}/${GOOS}_${GOARCH}/${BINARY}
	GOOS=linux GOARCH=amd64 go build -o ./bin/${VERSION}/${GOOS}_${GOARCH}/${BINARY}

install: build
	mkdir -p bin/${VERSION}/${OS_ARCH}/
	mv ${BINARY} bin/${VERSION}/${OS_ARCH}/

pipeline: build
	mkdir -p bin/${VERSION}/${OS_ARCH_PIPELINE}
	mv ${BINARY} bin/${VERSION}/${OS_ARCH_PIPELINE}
