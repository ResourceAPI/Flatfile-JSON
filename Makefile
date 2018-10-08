GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=flatfile-json.so

all: build

build:
	$(GOBUILD) -buildmode=plugin -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME)