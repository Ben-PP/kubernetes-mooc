# Log output app

Simple web server to output a log file and pong count.

## Kubernetes

### With manifests
```bash
# Deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.7/log_output/manifests/deployment.yaml

# Service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.7/log_output/manifests/service.yaml

# Ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.7/log_output/manifests/ingress.yaml
```

### With kubectl
To deploy with k3d used in the course use below command.
```bash
kubectl create deployment log-output --image=benpp/log_output:1.1
```
