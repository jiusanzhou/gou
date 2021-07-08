WD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
LDFLAGS := $(shell go run go.zoe.im/x/version/gen)

BUILD_LDFLAGS := '$(LDFLAGS)'

GO111MODULE=on
CGO_ENABLED=0

default: build

build:
	go build --ldflags $(BUILD_LDFLAGS) -o dist/gou ./cmd/gou