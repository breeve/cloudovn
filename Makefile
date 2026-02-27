.PHONY: go
go:
	go mod tidy

.PHONY: management
management:  management_api go
	go build -o bin/management pkg/management/cmd/management.go

.PHONY: management_api
management_api:
	cd pkg/api/management/v1; rm -f *.go; buf dep update; buf generate

.PHONY: controller
controller:
	make build -C pkg/controller
	cp pkg/controller/bin/manager bin/controller

.PHONY: gateway_schema
gateway_schema:
	modelgen -p schema -o pkg/gateway/schema pkg/gateway/schema/gateway.ovsschema

.PHONY: build
build: management controller
