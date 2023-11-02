format:
	go fmt ./...

upd-vendor:
	go mod tidy
	go mod vendor

test: lint
	go clean -testcache
	go test -cover ./...
lint: 
	command -v golangci-lint >/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
	golangci-lint run --fix ./...

.PHONY: format upd-vendor test lint 