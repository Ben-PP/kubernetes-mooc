# Log output app
This project contains 2 small apps. [log-writer](./log-writer) writes timestamps to a log file while [log-server](./log-server) reads the log file and serves it.

## Deploy
To deploy these apps use below commands.
```bash
# Persistent volume
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/volume-configs/log_output_pv.yaml

# Volume claim
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/log_output/manifests/persistentvolumeclaim.yaml

# Deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/log_output/manifests/deployment.yaml

# Service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/log_output/manifests/service.yaml

# Ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/log_output/manifests/ingress.yaml
