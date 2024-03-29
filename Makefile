# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=rollups
BINARY_UNIX=$(BINARY_NAME)_unix
all: test build
protos:
	cd protobuff/protos && PATH=$(PATH):$(GOPATH)/bin protoc --go_out=$(PKGMAP):.. *.proto
abi:
	abigen --abi eth/rollup.abi --pkg rollup --out rollup/rollup_api.go
	abigen --abi eth/gateway.abi --pkg gateway --out gateway/gateway_api.go
build:
	$(GOBUILD) -o $(BINARY_NAME) -v .
	chmod 775 $(BINARY_NAME)
docker:
	export DOCKER_BUILDKIT=0
	docker build .
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)