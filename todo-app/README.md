# Todo-app

This is the course project.

## Deployment

### Manifests
To deploy with manifest use below command.
```bash
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/1.4/todo-app/manifests/deployment.yaml
```

### Manually
To deploy manually with k3d use the command below.

```bash
kubectl create deployment todo-app --image=benpp/todo-app:1.2
```
