.PHONY: run build test docker

run:
	go run main.go

build:
	go build -o server

test:
	go test ./...

docker:
	docker build -t gene-editor-backend .

.PHONY: dev
dev:
	air # If you install air for hot reload