export AUTH_HOME := $(PWD)
export GOBIN := $(PWD)/bin
export AUTH_VENDOR_PATH=${AUTH_HOME}/vendor

install-go:
	goenv install -s $$(cat .go-version)

install-modules:
	go mod tidy

install-tools:
	mkdir -p bin
	go install \
	golang.org/x/tools/cmd/goimports \
	github.com/golangci/golangci-lint/cmd/golangci-lint \
	github.com/swaggo/swag/cmd/swag

install: install-go install-modules install-tools

lint:
	$(GOBIN)/golangci-lint run
