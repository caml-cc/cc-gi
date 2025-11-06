BINARY_NAME = cc-gi
CMD_DIR = ./cmd/cc-gi
BIN_DIR = bin
VERSION ?= dev

all: clean build

build: clean
	@mkdir -p $(BIN_DIR)
	@GOOS=darwin GOARCH=arm64 go build -o bin/cc-gi-macos-arm64-$(VERSION) ./cmd/cc-gi
	@GOOS=linux GOARCH=amd64 go build -o bin/cc-gi-linux-amd64-$(VERSION) ./cmd/cc-gi

clean:
	@rm ./bin/* && rm -d ./bin/