# Log output app
This project contains 2 small apps. [log-writer](./log-writer) writes timestamps to a log file while [log-server](./log-server) reads the log file and serves it.

# Environment
Log writer needs 2 env variables provided in the `deploy.yaml` or by the container.
| Name | Example | Description |
| :--: | :--: | :--- |
| LOG_DIR | `/app/log` | Where the `timestamps.log` file is put. |
| PINGPONG_URL | `http://example.com:8000` | Url where the pingpong app is listening |

## Deploy
To deploy these apps use below commands.
```bash
# Persistent volume
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/volume-configs/log_output_pv.yaml

# Volume claim
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/log_output/manifests/persistentvolumeclaim.yaml

# Deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/log_output/manifests/deployment.yaml

# Service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/log_output/manifests/service.yaml

# Ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.1/log_output/manifests/ingress.yaml
