LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_TAG:=1.54.2
MEMORY_PROFILE_FILE=memory.out
CPU_PROFILE_FILE=cpu.out

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/mailru/easyjson/...@v0.7.7

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

.PHONY: easyjson
easyjson:
	$(LOCAL_BIN)/easyjson -pkg

.PHONY: bench
bench:
	go test -v  -tags bench -bench . -benchmem -run BenchmarkGetDomainStat -memprofile=$(MEMORY_PROFILE_FILE) -cpuprofile $(CPU_PROFILE_FILE)

.PHONY: clear_bench
clear_bench:
	rm *.out

.PHONY: pprof_mem
pprof_mem:
	go tool pprof -http=:8000 $(MEMORY_PROFILE_FILE)

.PHONY: pprof_cpu
pprof_cpu:
	go tool pprof -http=:8000 $(CPU_PROFILE_FILE)
