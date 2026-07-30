PLUGIN_NAME    = proxmox
BINARY         = steampipe-plugin-$(PLUGIN_NAME).plugin
INSTALL_DIR    = $(HOME)/.steampipe/plugins/local/$(PLUGIN_NAME)
GOOS           ?= $(shell go env GOOS)
GOARCH         ?= $(shell go env GOARCH)

.PHONY: all build install test clean fmt vet

all: install

## build: compile the plugin binary into the repo root
build:
	go build -o $(BINARY) .

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY) $(INSTALL_DIR)/$(BINARY)
	chmod +x $(INSTALL_DIR)/$(BINARY)
	@echo "Installed to $(INSTALL_DIR)/$(BINARY)"
	@echo "Reference it in your .spc as: plugin = \"local/$(PLUGIN_NAME)\""

## test: run the Go test suite
test:
	go test ./... -v

## fmt: format all Go source
fmt:
	go fmt ./...

## vet: run go vet
vet:
	go vet ./...

## clean: remove build artifacts (repo-local and installed)
clean:
	rm -f $(BINARY)
	rm -rf $(INSTALL_DIR)
