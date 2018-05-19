CGO_ENABLED=0
GOFILES:=$(shell find . -name '*.go' | grep -v -E '(./vendor)')
GOOS=linux
GOARCH=amd64

VERSION=$(shell ./build/git-version.sh)

IMAGE_REPO:=practodev
REPO:=github.com/alok87/k8s-operator
LDFLAGS:=-w -X $(REPO)/pkg/version.Version=$(shell $(CURDIR)/build/git-version.sh)

all: bin/agent

images: bin/agent
		docker build -f Dockerfile-agent -t $(IMAGE_REPO)/agent:$(VERSION) .

bin/%: $(GOFILES)
		go build  -o $@ -ldflags "$(LDFLAGS)" $(REPO)/cmd/$*

check:
	@find . -name vendor -prune -o -name '*.go' -exec gofmt -s -d {} +
	@go vet $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

vendor:
	dep ensure

clean:
	rm -rf bin
