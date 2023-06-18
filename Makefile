BIN_DIR ?= ./bin
SCRIPTS_DIR ?= ./scripts

COVERAGE_PROFILE ?= coverage.out
MEMORY_PROFILE ?= memprofile.out

TEST_FILES ?= ./...
TEST_FLAGS ?=

default: test

.PHONY: clean
clean:
	@echo "---> Cleaning"
	go clean
	rm -rf $(COVERAGE_PROFILE) ./tmp

.PHONY: enforce
enforce:
	@echo "---> Enforcing coverage"
	$(SCRIPTS_DIR)/coverage.sh $(COVERAGE_PROFILE)

.PHONY: html
html:
	@echo "---> Generating HTML coverage report"
	go tool cover -html $(COVERAGE_PROFILE)

.PHONY: install
install:
	@echo "---> Installing dependencies"
	go mod download

.PHONY: lint
lint: $(BIN_DIR)/golangci-lint
	@echo "---> Linting"
	$(BIN_DIR)/golangci-lint run

$(BIN_DIR)/golangci-lint:
	@echo "--> Installing linter"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN_DIR) v1.53.3

.PHONY: test
test:
	@echo "---> Testing"
	go test -race $(TEST_FILES) -coverprofile $(COVERAGE_PROFILE) $(TEST_FLAGS)
