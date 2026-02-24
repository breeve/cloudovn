# Code Project

## init

```
(base) zhangfeng5@60237405W:~/github/cloudovn/pkg/controller$ kubebuilder init --domain cloudovn.io --repo github.com/breeve/cloudovn/pkg/controller
```

## create api

```
kubebuilder create api --group controller --version v1 --kind VPC
kubebuilder create api --group controller --version v1 --kind UnderlayNetwork
```

# kubebuilder

## Requrements

`go version go1.26.0 linux/amd64`

## Install

```
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/
```
