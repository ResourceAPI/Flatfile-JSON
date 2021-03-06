GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
BINARY_NAME=flatfile-json.so

all: build

build:
	$(GOBUILD) -buildmode=plugin -o $(BINARY_NAME) main.go

tidy:
	$(GOMOD) tidy

clean:
	rm -f $(BINARY_NAME)