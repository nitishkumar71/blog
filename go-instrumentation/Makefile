GOCMD=go

all: help

build: ## Build your project and put the output binary in bin/
	mkdir -p bin
	GO111MODULE=on $(GOCMD) build -mod vendor -o bin/go-instrumentation

clean: ## Remove build related file
	rm -fr ./bin

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

tidy:
	$(GOCMD) mod tidy