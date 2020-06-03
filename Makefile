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
	abigen --abi eth/rollup.abi --pkg eth --out eth/api.go
build:
	$(GOBUILD) -buildmode=plugin -trimpath -o $(BINARY_NAME).so .
	chmod 775 $(BINARY_NAME).so
docker:
	export DOCKER_BUILDKIT=1
	docker build -o out .
	chmod 775 ./out/$(BINARY_NAME).so
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)