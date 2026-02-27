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
controller: go
	make build -C pkg/controller
	cp pkg/controller/bin/manager bin/controller

.PHONY: gateway
gateway: go gateway_agent

.PHONY: gateway_agent
gateway_agent:
	cd  pkg/gateway/dataplane; go generate
	go build -o bin/gateway-agent pkg/gateway/agent/main.go

.PHONY: gateway_dataplane_vmlinux_h
gateway_dataplane_vmlinux_h:
	/usr/lib/linux-tools/6.8.0-101-generic/bpftool btf dump file /sys/kernel/btf/vmlinux format c > pkg/gateway/dataplane/vmlinux.h

.PHONY: build
build: management controller gateway

################################################################
# Archived                                                     #
################################################################

.PHONY: go_c_ovsdb_schema
go_c_ovsdb_schema:
	modelgen -p schema -o archived/go_c_ovsdb/schema archived/go_c_ovsdb/schema/gateway.ovsschema

go_c_ovsdb_client: go_c_ovsdb_schema go
	go build -o archived/go_c_ovsdb/bin/client archived/go_c_ovsdb/client/client.go
