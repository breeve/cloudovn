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

# Domain Model

![](./domain_model.drawio.svg)

# Nework Model

![](./network_model.drawio.svg)

# Underlay & Overlay Network

# k8s network model

> There are 4 distinct networking problems to solve:
>
> - Highly-coupled container-to-container communications
> - Pod-to-Pod communications
> - Pod-to-Service communications
> - External-to-internal communications
>
> --- [k8s networking design](https://github.com/kubernetes/design-proposals-archive/blob/main/network/networking.md)

# Tenant VPC

# Management VPC

# CNI

# OVN-Client

---

# Refs

1. https://kubernetes.io/zh-cn/docs/concepts/cluster-administration/networking/
1. https://github.com/kubernetes/design-proposals-archive/blob/main/network/networking.md
1. https://kubernetes.io/zh-cn/docs/concepts/services-networking/
