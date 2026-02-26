.PHONY: go
go:
	go mod tidy

.PHONY: management
management:
	go build -o bin/management pkg/management/cmd/management.go