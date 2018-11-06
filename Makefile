GOOS ?= darwin
GOARCH ?= amd64


default: build


build:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags '-w' -o s3md5-$(GOOS)-$(GOARCH) ./...
