```
kind create cluster --name cloudovn --config kind-config.yaml

(base) zhangfeng5@60237405W:~/github/cloudovn/test/k8s$ kubectl get nodes
NAME                     STATUS   ROLES           AGE   VERSION
cloudovn-control-plane   Ready    control-plane   79s   v1.35.0
cloudovn-worker          Ready    <none>          67s   v1.35.0
cloudovn-worker2         Ready    <none>          67s   v1.35.0
(base) zhangfeng5@60237405W:~/github/cloudovn/test/k8s$ kubectl get pods -A
NAMESPACE            NAME                                             READY   STATUS    RESTARTS   AGE
kube-system          coredns-7d764666f9-6ls4k                         1/1     Running   0          75s
kube-system          coredns-7d764666f9-6zp6q                         1/1     Running   0          75s
kube-system          etcd-cloudovn-control-plane                      1/1     Running   0          82s
kube-system          kindnet-dnl5p                                    1/1     Running   0          72s
kube-system          kindnet-r6wmh                                    1/1     Running   0          76s
kube-system          kindnet-x8727                                    1/1     Running   0          72s
kube-system          kube-apiserver-cloudovn-control-plane            1/1     Running   0          82s
kube-system          kube-controller-manager-cloudovn-control-plane   1/1     Running   0          82s
kube-system          kube-proxy-cvvcn                                 1/1     Running   0          72s
kube-system          kube-proxy-gmnzs                                 1/1     Running   0          76s
kube-system          kube-proxy-nk28z                                 1/1     Running   0          72s
kube-system          kube-scheduler-cloudovn-control-plane            1/1     Running   0          82s
local-path-storage   local-path-provisioner-67b8995b4b-qgmtx          1/1     Running   0          75s

(base) zhangfeng5@60237405W:~/github/cloudovn/test/k8s$ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED              STATUS              PORTS                       NAMES
21ee24cd1a44   kindest/node:v1.35.0   "/usr/local/bin/entr…"   About a minute ago   Up About a minute                               cloudovn-worker2
8c69b26f100d   kindest/node:v1.35.0   "/usr/local/bin/entr…"   About a minute ago   Up About a minute                               cloudovn-worker
0d293ff3b849   kindest/node:v1.35.0   "/usr/local/bin/entr…"   About a minute ago   Up About a minute   127.0.0.1:36047->6443/tcp   cloudovn-control-plane
(base) zhangfeng5@60237405W:~/github/cloudovn/test/k8s$ docker exec -ti cloudovn-control-plane bash
root@cloudovn-control-plane:/#

```
