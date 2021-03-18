test: SHELL:=/bin/bash
test:
	mkdir -p coverage
	TEST=1 go test -cover -v ./internal/... ./pkg/... -coverprofile coverage/cover.out
	go tool cover -html=coverage/cover.out -o coverage/cover.html
