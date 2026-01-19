# Todo-app

This is the course project.

## Deployment

### Manifests
To deploy with manifest use below command.
```bash
# Persistent volume
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.12/volume-configs/todo_pv.yaml

# Volume claim
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.12/todo-app/manifests/pvclaim.yaml

# Apply deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.12/todo-app/manifests/deployment.yaml

# Apply service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.12/todo-app/manifests/service.yaml

# Apply ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.12/todo-app/manifests/ingress.yaml
```
