.PHONY: all clean
all: clean setup fmt vet lint imports test-cover-report benchmark-test
ci-all: clean setup-ci fmt vet lint imports test-cover-report benchmark-test

APP=govaluator
ALL_PACKAGES=$(shell go list ./...)
IMPORT_PACKAGES=$(shell go list ./... | awk -F"\/" '{print $$4}')

setup:
	@echo "setting up build environment..."
	@go install golang.org/x/lint/golint@latest
	@go install golang.org/x/tools/cmd/goimports@latest

setup-ci:
	@echo "setting up ci build environment..."
	@echo "installing lint tool ..."
	@go install golang.org/x/lint/golint@latest
	@echo "installing import manager tool ..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "downloading coveralls cli tool ..."
	@go install github.com/mattn/goveralls@latest
	@echo "installing all project dependencies ..."
	@go get ./...

clean:
	@echo "cleaning test data cache..."
	@rm -rf out/
	@rm -f *.out
	@go clean -testcache

imports:
	@echo "Running imports..."
	@goimports -w -local github.com/isomnath/$(APP) $(IMPORT_PACKAGES)

fmt:
	@echo "running fmt..."
	@go fmt ./...

vet:
	@echo "running vet..."
	@go vet ./...

lint:
	@echo "running lint..."
	@for p in $(ALL_PACKAGES); do \
  		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
  	done

static-code-analysis: clean fmt vet lint imports

benchmark-test: clean
	@echo "running benchmark tests..."
	@go test -bench=.

test: clean
	@echo "running tests..."
	@go test $(ALL_PACKAGES)

test-cover: clean
	@echo "running tests..."
	@mkdir -p out/
	@go test $(ALL_PACKAGES) -coverprofile=coverage.out

test-cover-report: test-cover
	@echo 'Total coverage: '`go tool cover -func coverage.out | tail -1 | awk '{print $$3}'`

test-cover-html: test-cover
	@go tool cover -html=coverage.out
