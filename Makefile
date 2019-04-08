GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
DEPCMD=dep
BINARY=octane

all: install build

build:
	$(GOBUILD) -v -o $(BINARY)

run:
	./$(BINARY)

.PHONY : install
install:
	$(DEPCMD) ensure

.PHONY : clean
clean:
	rm $(BINARY) && rm -rf vendor

.PHONY : protogen
protogen:
	protoc -I grpc/lotto/ grpc/lotto/lotto.proto  --go_out=plugins=grpc:grpc/lotto

letmegrpc:
	cd grpc/lotto && letmegrpc --addr=localhost:$(PORT) --port=8080 lotto.proto&

.PHONY: help

# Show this help.
help:
	@echo all -- install dependencies and build binary
	@echo build -- build binary
	@echo install -- install dependencies
	@echo clean -- clean up dependencies and binaries
	@echo protogen -- generate code on proto file
	@echo letmegrpc PORT=port --start lemegrpc http service
	@echo run -- start service
