STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo

install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/becash143/proxmox@latest/steampipe-plugin-proxmox.plugin -tags "${BUILD_TAGS}" *.go
