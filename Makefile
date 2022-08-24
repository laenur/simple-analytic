GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=simple-analyic
BINARY_UNIX=$(BINARY_NAME)-unix

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v main.go
run:
	make build
	./simple-analyic
mock:
	mockery --all --keeptree
test:
	$(GOTEST) ./... -cover -race -count=1 -coverprofile="coverage.out"
test-no-race:
	$(GOTEST) ./... -cover -count=1 -coverprofile="coverage.out"
lint:
	staticcheck ./...
tool:
	$(GOGET) github.com/vektra/mockery/v2/.../
	$(GOGET) honnef.co/go/tools/cmd/staticcheck@latest
build-linux:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -o $(BINARY_UNIX) -a -installsuffix cgo -v main.go
