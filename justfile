default:
	just --list

documentation:
	go doc -all -u -http

build:
	go build ./...

test:
	go test ./... --cover -coverprofile=coverage.out --covermode atomic --coverpkg=./...

show-coverage-report:
	go tool cover -html=coverage.out

coverage-report: test show-coverage-report

generate:
	go generate ./...

lint:
	go tool golangci-lint run -v --fix

format:
	go fmt ./...

checks: format lint test build
