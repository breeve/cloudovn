# Development

## Requrements

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

## Project Init

```
cd pkg/controller
kubebuilder init --domain cloudovn.io --repo github.com/breeve/cloudovn/pkg/controller
```

## create api

```
cd pkg/controller

kubebuilder create api --group controller --version v1 --kind VPC
kubebuilder create api --group controller --version v1 --kind UnderlayNetwork
```
