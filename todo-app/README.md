# Todo-app

This is the course project.

## Deployment

### Manifests
To deploy with manifest use below command.
```bash
# Apply deployment
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.8/todo-app/manifests/deployment.yaml

# Apply service
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.8/todo-app/manifests/service.yaml

# Apply ingress
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.8/todo-app/manifests/ingress.yaml
```
