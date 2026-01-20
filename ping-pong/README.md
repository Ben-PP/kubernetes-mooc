# Ping-pong app
Small web server that responds with pong and a counter value.

## Deployment
To deploy this app use below command.
```bash
# Apply deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/ping-pong/manifests/deployment.yaml

# Apply service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/ping-pong/manifests/service.yaml

# Apply ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/log_output/manifests/ingress.yaml
```
