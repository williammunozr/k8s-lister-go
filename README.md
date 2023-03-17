# Client Go

## Generate password database in Linux

```commandline
gpg --generate-key

pass init [GPG_KEY]
```

## Docker Image

### Build

```commandline
docker build -t k8s-lister-go:0.1.0 .
```

### Tag and Push

```commandline
docker tag k8s-lister-go:0.1.0 hachikoapp/k8s-lister-go:0.1.0
docker push hachikoapp/k8s-lister-go:0.1.0
```

## Create Deployment

```commandline
kubectl create deployment k8s-lister-go --image=hachikoapp/k8s-lister-go:0.1.0                                                                      ─╯
deployment.apps/k8s-lister-go created
```

### Validate the deployment

```commandline
kubectl get pods                                                                                                                                    ─╯
NAME                             READY   STATUS   RESTARTS      AGE
k8s-lister-go-567748cc76-vqqrx   0/1     Error    3 (33s ago)   50s
```

### See the logs

```commandline
kubectl logs k8s-lister-go-567748cc76-vqqrx                                                                                                         ─╯
2023/03/17 15:30:29 error stat /home/william/.kube/config: no such file or directory building config from flags
```

That means the configuration doesn't exists in the container.

### Add InClusterConfig build and redeploy

I'm omitting some steps.

```commandline
k logs k8s-lister-go-7988f769db-bvsfx                                                                                                         ─╯
error stat /home/william/.kube/config: no such file or directory building config from flags
2023/03/17 15:38:46 error pods is forbidden: User "system:serviceaccount:default:default" cannot list resource "pods" in API group "" in the namespace "default" while listing all the pods from default namespace
```

### Create role and binding

```commandline
kubectl create role poddepl --resource pods,deployments --verb list                                                                                 ─╯
role.rbac.authorization.k8s.io/poddepl created
```

```commandline
kubectl create rolebinding poddepl --role poddepl --serviceaccount default:default                                                                  ─╯
rolebinding.rbac.authorization.k8s.io/poddepl created
```

### Delete the pod and check the logs of the new one

```commandline
kubectl logs k8s-lister-go-7988f769db-ltn4s                                                                                                         ─╯
error stat /home/william/.kube/config: no such file or directory building config from flags
k8s-lister-go-7988f769db-ltn4s
k8s-lister-go
```
