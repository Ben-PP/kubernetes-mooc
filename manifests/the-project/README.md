# Project
Here are the manifests of the course project

## Environment
There is `TODOBACKEND_URL` env variable defined in the `app-deployment.yaml` which you will have to modify to match your set up. This is the url to which the app makes the GET request for todos and to which browser sends POST requests to create todos.

## Deployment
```bash
# Create namespace if needed
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.4/manifests/namespaces/project.yaml

# Apply manifests
kubectl apply -f .
```

## Applications
Here is the source code for apps used in this exercise.
- [todo-app](../../todo-app)
- [todo-backend](../../todo-backend)
