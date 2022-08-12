BIN := $(CURDIR)/.bin
UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

$(BIN):
	mkdir -p $(BIN)

GOLANGCLI_LINT := $(BIN)/golangci-lint
GOLANGCLI_LINT_VERSION := v1.48.0
$(GOLANGCLI_LINT): | $(BIN) ## Install golangci-lint
	@curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | sh -s -- -b $(BIN) $(GOLANGCLI_LINT_VERSION)
	@chmod +x "$(BIN)/golangci-lint"

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./...

.PHONY: mod
mod:
	go mod download

.PHONY: clean
clean:
	rm -rf $(BIN)

.PHONY: lint
lint: | $(GOLANGCLI_LINT)
	$(BIN)/golangci-lint run --verbose --config=.golangci.yml ./...