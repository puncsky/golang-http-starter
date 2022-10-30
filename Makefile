.PHONY: install
install:
	go mod download


.PHONY: dev
dev:
	go run *.go


.PHONY: build
build:
	go build -tags netgo -ldflags '-s -w' -o app


.PHONY: lint
lint:
	go fmt ./...
	go mod tidy
	go list ./... | xargs golint
