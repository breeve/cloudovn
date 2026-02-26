.PHONY: go
go:
	go mod tidy

.PHONY: management
management:
	go build -o bin/management pkg/management/cmd/management.go

.PHONY: controller
controller:
	make build -C pkg/controller
	cp pkg/controller/bin/manager bin/controller
