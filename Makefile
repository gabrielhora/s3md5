default: build


build:
	@go build -ldflags '-w' ./...
