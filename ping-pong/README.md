# Ping-pong app
Small web server that responds with pong and a counter value.

## Deployment
To deploy this app use below command.
```bash
# Apply deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.9/ping-pong/manifests/deployment.yaml

# Apply service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.9/ping-pong/manifests/service.yaml

# Apply ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.9/log_output/manifests/ingress.yaml
```
