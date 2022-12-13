.PHONY: update-packges
update-packages:
	@go get -u all
	go mod tidy
	go mod vendor

.PHONY: format
format:
	@go install golang.org/x/tools/cmd/goimports@latest
	goimports -local "github.com/hodlgap/krx" -w .
	gofmt -s -w .
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	@go install github.com/rakyll/gotest@latest
	gotest -race -cover -v ./...
