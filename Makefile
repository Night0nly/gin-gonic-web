# application meta info
NAME := NToolWeb
VERSION= 1.0
GOVERSION := $(shell go version)
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.goVersion=$(GOVERSION)'
ENTRYPOINT := main.go

all: dep front build

dep:
	dep ensure

front:
	cd frontend && npm run build

# buid
build:
	GOOS=linux GOARCH=amd64	go build -tags="netgo" -installsuffix netgo -ldflags "$(LDFLAGS)" -o $(NAME)-linux-64 ./$(ENTRYPOINT)
	GOOS=darwin GOARCH=amd64 go build -tags="netgo" -installsuffix netgo -ldflags "$(LDFLAGS)" -o $(NAME)-darwin-64 ./$(ENTRYPOINT)

# test
test:
	go test -tags="netgo" ./...

update:
	dep ensure -update
