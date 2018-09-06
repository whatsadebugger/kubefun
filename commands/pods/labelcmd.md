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