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

(base) zhangfeng5@60237405W:~/github/cloudovn$ buf --version
1.66.0
(base) zhangfeng5@60237405W:~/github/cloudovn$ protoc-gen-go --version
protoc-gen-go v1.36.6
(base) zhangfeng5@60237405W:~/github/cloudovn$ protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.5.1
(base) zhangfeng5@60237405W:~/github/cloudovn$ protoc-gen-grpc-gateway --version
Version v2.27.1, commit unknown, built at unknown
(base) zhangfeng5@60237405W:~/github/cloudovn$ go version -m $(which protoc-gen-validate) | grep protoc-gen-validate | grep mod
        mod     github.com/envoyproxy/protoc-gen-validate       v1.2.1  h1:DEo3O99U8j4hBFwbJfrz9VtgcDfUKS7KJ7spH3d86P8
```

# schema

```
(base) zhangfeng5@60237405W:~/github/cloudovn$ go install github.com/ovn-kubernetes/libovsdb/cmd/modelgen@latest
go: downloading golang.org/x/text v0.22.0
(base) zhangfeng5@60237405W:~/github/cloudovn$ go version -m $(which modelgen) | grep modelgen | grep mod
/home/zhangfeng5/go/bin/modelgen: go1.26.0
        path    github.com/ovn-kubernetes/libovsdb/cmd/modelgen
```
