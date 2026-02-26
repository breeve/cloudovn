# Requrements

- Golang:

  ```
  go version go1.26.0 linux/amd64
  ```

- kubebuilder:

  ```
  KubeBuilder:          v4.12.0
  Kubernetes:           1.35.0
  Git Commit:           94434cdf622a00d8a8a50f53a2ab36b3059c8830
  Build Date:           2026-02-15T21:24:05Z
  Go OS/Arch:           linux/amd64
  ```

# Controller

## Create API

```
cd pkg/controller

kubebuilder create api --group controller --version v1 --kind UnderlayNetwork
kubebuilder create api --group controller --version v1 --kind VPC
kubebuilder create api --group controller --version v1 --kind RouteTable
kubebuilder create api --group controller --version v1 --kind Subnet

```

# API

```
go install github.com/bufbuild/buf/cmd/buf@latest

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/envoyproxy/protoc-gen-validate@latest

# 打印下版本
```
