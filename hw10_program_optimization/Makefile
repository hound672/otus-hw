LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_TAG:=1.54.2

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)
	mkdir -p $(LOCAL_BIN)

# install golangci-lint binary
.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: install
install: .bin-deps install-lint

.PHONY: linter
linter:
	$(LOCAL_BIN)/golangci-lint run --fix
