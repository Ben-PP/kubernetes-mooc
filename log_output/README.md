# Log output app

Simple app outputting text.

## Kubernetes

### With manifests
```bash
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.3/log_output/manifests/deployment.yaml
```

### With kubectl
To deploy with k3d used in the course use below command.
```bash
kubectl create deployment log-output --image=benpp/log_output:1.1
```
