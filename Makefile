.PHONY: deps clean build lint changelog snapshot release

# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

GO ?= latest

BINARY = glsobus
MAIN = $(shell pwd)/cmd/main.go

CLIENT_BINARY = glsobus-client
CLIENT_MAIN = $(shell pwd)/cmd/client/main.go

AGENT_BINARY = cbc-agent
AGENT_MAIN = $(shell pwd)/cmd/agent/main.go

BUILDDIR = $(shell pwd)/build
VERSION ?= 0.0.1
GITREV = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +'%FT%TZ%z')
LDFLAGS=-ldflags "-X github.com/qlcchain/go-lsobus/services/version.Version=${VERSION} \
				  -X github.com/qlcchain/go-lsobus/services/version.GitRev=${GITREV} \
				  -X github.com/qlcchain/go-lsobus/services/version.BuildTime=${BUILDTIME}"
GO_BUILDER_VERSION=v1.15.6

default: build

deps:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u github.com/goreleaser/goreleaser
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	go get -u golang.org/x/tools/cmd/goimports

build:
	go build ${LDFLAGS} -o $(BUILDDIR)/${BINARY} -i $(MAIN)
	@echo 'Build $(BINARY) done.'

client:
	go build ${LDFLAGS} -o $(BUILDDIR)/${CLIENT_BINARY} -i $(CLIENT_MAIN)
	@echo "Build $(CLIENT_BINARY) done."

agent:
	go build ${LDFLAGS} -o $(BUILDDIR)/${AGENT_BINARY} -i $(AGENT_MAIN)
	@echo "Build $(AGENT_BINARY) done."

changelog:
	git-chglog $(VERSION) > CHANGELOG.md

clean:
	rm -rf $(BUILDDIR)/

lint:
	golangci-lint run --fix

style:
	gofmt -w .
	goimports -local github.com/qlcchain/go-lsobus -w .

snapshot:
	docker run --rm --privileged \
		-e PRIVATE_KEY=$(PRIVATE_KEY) \
		-v $(CURDIR):/go-lsobus \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(GOPATH)/src:/go/src \
		-w /go-lsobus \
		goreng/golang-cross:$(GO_BUILDER_VERSION) --snapshot --rm-dist

release: changelog
	docker run --rm --privileged \
		-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
		-e PRIVATE_KEY=$(PRIVATE_KEY) \
		-v $(CURDIR):/go-lsobus \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(GOPATH)/src:/go/src \
		-w /go-lsobus \
		goreng/golang-cross:$(GO_BUILDER_VERSION) --rm-dist --release-notes=CHANGELOG.md
