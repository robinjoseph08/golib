BIN_DIR ?= ./bin
SCRIPTS_DIR ?= ./scripts

GOLANGCI_LINT_VERSION ?= v2.4.0
GOLANGCI_LINT ?= $(BIN_DIR)/golangci-lint-$(GOLANGCI_LINT_VERSION)

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
lint: $(GOLANGCI_LINT)
	@echo "---> Linting"
	$(GOLANGCI_LINT) run

$(GOLANGCI_LINT):
	@echo "--> Installing linter"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN_DIR) $(GOLANGCI_LINT_VERSION)
	mv $(BIN_DIR)/golangci-lint $(GOLANGCI_LINT)

$(BIN_DIR)/git-chglog:
	@echo "--> Installing git-chglog"
	GOBIN=$(PWD)/$(BIN_DIR) go install github.com/git-chglog/git-chglog/cmd/git-chglog@v0.15.4

.PHONY: test
test:
	@echo "---> Testing"
	go test -race $(TEST_FILES) -coverpkg=./... -coverprofile $(COVERAGE_PROFILE) $(TEST_FLAGS)

.PHONY: release
release: $(BIN_DIR)/git-chglog
	@echo "---> Creating new release"
ifndef tag
	$(error tag must be specified)
endif
	$(BIN_DIR)/git-chglog --output CHANGELOG.md --next-tag $(tag)
	sed -i "" "s/version-.*-green/version-$(tag)-green/" README.md
	git add CHANGELOG.md README.md
	git commit -m $(tag)
	git tag $(tag)
	git push origin master --tags
