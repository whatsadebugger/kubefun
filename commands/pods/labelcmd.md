# labels

create an object from a yaml definition
```
$ kubectl create -f kubia-manual-with-labels.yaml
pod "kubefun-manual" created`
```

show pods with labels
```
➜  kubefun git:(master) ✗ kg po --show-labels
NAME                        READY     STATUS    RESTARTS   AGE       LABELS
kubefun-manual              1/1       Running   0          57s       <none>
kubefun-manual-with-label   1/1       Running   0          23s       creation_method=manual,env=prod
```

specify specific labels and display them as columns
```
➜  kubefun git:(master) ✗ kg po -L creation_method,env
NAME                        READY     STATUS    RESTARTS   AGE       CREATION_METHOD   ENV
kubefun-manual              1/1       Running   0          1m
kubefun-manual-with-label   1/1       Running   0          1m        manual            prod
```

add a label to an existing pod use --overwrite to override existing labels
```
➜  kubefun git:(master) ✗ kla po kubefun-manual --overwrite env=debug
pod "kubefun-manual" labeled
```

# label selectors

## pods

filter using -l
```
➜  kubefun git:(master) kg po -L creation_method,env
NAME                        READY     STATUS    RESTARTS   AGE       CREATION_METHOD   ENV
kubefun-manual              1/1       Running   0          6m                          debug
kubefun-manual-with-label   1/1       Running   0          5m        manual            prod
➜  kubefun git:(master) kg po -l creation_method=manual
NAME                        READY     STATUS    RESTARTS   AGE
kubefun-manual-with-label   1/1       Running   0          5m
```

pods with an env label
```
➜  kubefun git:(master) ✗ kg po -l env
NAME                        READY     STATUS    RESTARTS   AGE
kubefun-manual              1/1       Running   0          9m
kubefun-manual-with-label   1/1       Running   0          9m
```

without the env label
```
➜  kubefun git:(master) ✗ kg po -l '!env'
No resources found.
```
## nodes

label a node with a gpu and ssd flag
```
➜  kubefun git:(master) ✗ kla no gke-kubefun-default-pool-4fb4229d-2qz5 ssd=true gpu=true
node "gke-kubefun-default-pool-4fb4229d-2qz5" labeled
➜  kubefun git:(master) ✗ kg no gke-kubefun-default-pool-4fb4229d-2qz5 -L ssd,gpu
NAME                                     STATUS    ROLES     AGE       VERSION        SSD       GPU
gke-kubefun-default-pool-4fb4229d-2qz5   Ready     <none>    59m       v1.9.7-gke.5   true      true
```
```note: to deploy a pod to a specific node add a nodeselector in the spec for that pod.```  
f