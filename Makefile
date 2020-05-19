.PHONY: deps clean build lint changelog snapshot release

# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

GO ?= latest

BINARY = glsobus
MAIN = $(shell pwd)/cmd/main.go

CLIENT_BINARY = glsobus-client
CLIENT_MAIN = $(shell pwd)/cmd/client/main.go

BUILDDIR = $(shell pwd)/build
VERSION ?= 0.0.1
GITREV = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +'%FT%TZ%z')
LDFLAGS=-ldflags "-X github.com/qlcchain/go-lsobus/services/version.Version=${VERSION} \
				  -X github.com/qlcchain/go-lsobus/services/version.GitRev=${GITREV} \
				  -X github.com/qlcchain/go-lsobus/services/version.BuildTime=${BUILDTIME}"

default: build

deps:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u github.com/goreleaser/goreleaser
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	go get -u golang.org/x/tools/cmd/goimports

build:
	go build ${LDFLAGS} -o $(BUILDDIR)/${BINARY} -i $(MAIN)
	@echo "Build $(BINARY) done."

client:
	go build ${LDFLAGS} -o $(BUILDDIR)/${CLIENT_BINARY} -i $(CLIENT_MAIN)
	@echo "Build $(CLIENT_BINARY) done."

changelog:
	git-chglog $(VERSION) > CHANGELOG.md

clean:
	rm -rf $(BUILDDIR)/

lint:
	golangci-lint run --fix

gofmt:
	gofmt -w .

style:
	gofmt -w .
	goimports -local github.com/qlcchain/go-lsobus -w .

snapshot:
	goreleaser --snapshot --rm-dist

release: changelog
	goreleaser --rm-dist --release-notes=CHANGELOG.md