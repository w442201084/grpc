GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc --proto_path=./proto --micro_out=./proto/product --go_out=:./proto/product product/product.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product-service:latest
