# Namespaces

```
➜  kubefun git:(master) ✗ kg ns
NAME               STATUS    AGE
custom-namespace   Active    59s
default            Active    18d
kube-public        Active    18d
kube-system        Active    18d
```

show system namespace pods

```
➜  kubefun git:(master) ✗ kg po --namespace kube-system
NAME                                                READY     STATUS    RESTARTS   AGE
event-exporter-v0.1.9-5c8fb98cdb-pxdd6              2/2       Running   0          9h
fluentd-gcp-v2.0.17-fj9fc                           2/2       Running   0          1h
fluentd-gcp-v2.0.17-m9kv8                           2/2       Running   0          1h
heapster-v1.5.2-6b6cf689c9-blwl6                    3/3       Running   0          9h
kube-dns-5dcfcbf5fb-25jrg                           4/4       Running   0          1h
kube-dns-5dcfcbf5fb-kfhcb                           4/4       Running   0          9h
kube-dns-autoscaler-69c5cbdcdd-6dr49                1/1       Running   0          9h
kube-proxy-gke-kubefun-default-pool-4fb4229d-2qz5   1/1       Running   0          1h
kube-proxy-gke-kubefun-default-pool-4fb4229d-vvfl   1/1       Running   0          1h
kubernetes-dashboard-774649bb46-9bv7w               1/1       Running   0          9h
l7-default-backend-57856c5f55-4b2tp                 1/1       Running   0          9h
metrics-server-v0.2.1-7f8dd98c8f-m7slz              2/2       Running   0          9h
```

create a pod in the new namespace
```
➜  kubefun git:(master) ✗ kcr -f yaml/pods/kubefun-po-manual.yaml -n custom-namespace
pod "kubefun-manual" created
➜  kubefun git:(master) ✗ kg po
NAME                        READY     STATUS    RESTARTS   AGE
kubefun-manual              1/1       Running   0          28m
kubefun-manual-with-label   1/1       Running   0          28m
➜  kubefun git:(master) ✗ kg po -n custom-namespace
NAME             READY     STATUS              RESTARTS   AGE
kubefun-manual   0/1       ContainerCreating   0          7s
```
