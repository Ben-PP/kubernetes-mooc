# Ping-pong app
Small web server that responds with pong and a counter value.

## Environment
App writes the request count to log file named `pongs.log`. Log directory is defined with `LOG_DIR` env variable and it defaults to `/app/log` in the container and `./logs` in the code.

## Deployment
To deploy this app use below command.
```bash
# Persistent volume
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/volume-configs/ping-pong_pv.yaml

# Volume claim
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/ping-pong/manifests/persistentvolumeclaim.yaml

# Apply deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/ping-pong/manifests/deployment.yaml

# Apply service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/ping-pong/manifests/service.yaml

# Apply ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.11/log_output/manifests/ingress.yaml
```
