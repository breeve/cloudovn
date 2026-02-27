# Arch

![](./architecture.drawio.svg)

- Management
  - API: Provide RESTful/gRPC API
- Controller
  - Operator: VPC, Subnet, ...
  - CNI Manager: Pod, Service, ...
  - Network Topo: Config Network, IPAM, LS, LR, LB are the primary targets.
- CNI Server
  - Network Topo: Config Network, config instance NIC, such as add LSP to LS.
- CNI Binary:
  - Plugin Delegate: Integrate and proxy multiple CNI plugins to avoid plugin management chaos.

> Should fix it, ref: https://ovn-kubernetes.io/design/architecture/#interconnect-mode-architecture-default

# Domain Model

![](./domain_model.drawio.svg)

# Nework Model

![](./network_model.drawio.svg)

# Underlay & Overlay Network

![](./underlay_overlay_network.drawio.svg)

# k8s network model

> There are 4 distinct networking problems to solve:
>
> - Highly-coupled container-to-container communications
> - Pod-to-Pod communications
> - Pod-to-Service communications
> - External-to-internal communications
>
> --- [k8s networking design](https://github.com/kubernetes/design-proposals-archive/blob/main/network/networking.md)

# VPC

## Tenant VPC

![](./vpc_tenant_topo.drawio.svg)

## Management VPC

![](./vpc_management_topo.drawio.svg)

# CNI

## Primary Network Model

## Additional Network Model

> https://kubeovn.github.io/docs/v1.15.x/start/non-primary-mode/

# OVN-Client

> https://github.com/kubeovn/kube-ovn/blob/release-1.15/pkg/ovs/ovn.go#L140

# DevOps

> http://kubevela.net/zh/

# Log

# eBPF

> https://ebpf-go.dev/

---

# Refs

1. https://kubernetes.io/zh-cn/docs/concepts/cluster-administration/networking/
1. https://github.com/kubernetes/design-proposals-archive/blob/main/network/networking.md
1. https://kubernetes.io/zh-cn/docs/concepts/services-networking/
1. https://kubeovn.github.io/docs/v1.15.x/vpc/vpc/#vpc-pod-livenessprobe-readinessprobe
