ERRCHECK_VERSION = latest
GOLANGCI_LINT_VERSION = latest
REVIVE_VERSION = latest
GOIMPORTS_VERSION = latest
INEFFASSIGN_VERSION = latest

LOCAL_BIN := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/.bin

.PHONY: all
all: clean tools lint fmt test build

.PHONY: clean
clean:
	rm -rf $(LOCAL_BIN)

.PHONY: tools
tools:  golangci-lint-install revive-install go-imports-install ineffassign-install
	go mod tidy

.PHONY: golangci-lint-install
golangci-lint-install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

.PHONY: revive-install
revive-install:
	GOBIN=$(LOCAL_BIN) go install github.com/mgechev/revive@$(REVIVE_VERSION)

.PHONY: ineffassign-install
ineffassign-install:
	GOBIN=$(LOCAL_BIN) go install github.com/gordonklaus/ineffassign@$(INEFFASSIGN_VERSION)

.PHONY: lint
lint: tools run-lint

.PHONY: run-lint
run-lint: lint-golangci-lint lint-revive

.PHONY: lint-golangci-lint
lint-golangci-lint:
	$(info running golangci-lint...)
	$(LOCAL_BIN)/golangci-lint -v run ./... || (echo golangci-lint returned an error, exiting!; sh -c 'exit 1';)

.PHONY: lint-revive
lint-revive:
	$(info running revive...)
	$(LOCAL_BIN)/revive -formatter=stylish -config=build/ci/.revive.toml -exclude ./vendor/... ./... || (echo revive returned an error, exiting!; sh -c 'exit 1';)

.PHONY: upgrade-direct-deps
upgrade-direct-deps: tidy
	for item in `grep -v 'indirect' go.mod | grep '/' | cut -d ' ' -f 1`; do \
		echo "trying to upgrade direct dependency $$item" ; \
		go get -u $$item ; \
  	done
	go mod tidy
	go mod vendor

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run-goimports
run-goimports: go-imports-install
	for item in `find . -type f -name '*.go' -not -path './vendor/*'`; do \
		$(LOCAL_BIN)/goimports -l -w $$item ; \
	done

.PHONY: go-imports-install
go-imports-install:
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

.PHONY: fmt
fmt: tools run-fmt run-ineffassign run-vet

.PHONY: run-fmt
run-fmt:
	$(info running fmt...)
	go fmt ./... || (echo fmt returned an error, exiting!; sh -c 'exit 1';)

.PHONY: run-ineffassign
run-ineffassign:
	$(info running ineffassign...)
	$(LOCAL_BIN)/ineffassign ./... || (echo ineffassign returned an error, exiting!; sh -c 'exit 1';)

.PHONY: run-vet
run-vet:
	$(info running vet...)
	go vet ./... || (echo vet returned an error, exiting!; sh -c 'exit 1';)

.PHONY: test
test: tidy
	$(info starting the test for whole module...)
	go test -failfast -vet=off -race ./... || (echo an error while testing, exiting!; sh -c 'exit 1';)

.PHONY: test-with-coverage
test-with-coverage: tidy
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic
