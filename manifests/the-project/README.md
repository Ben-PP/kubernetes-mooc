# Exercise 2.2
Here are the manifests of this exercise.

## Environment
There is `TODOBACKEND_URL` env variable defined in the `app-deployment.yaml` which you will have to modify to match your set up. This is the url to which the app makes the GET request for todos and to which browser sends POST requests to create todos.

## Deployment
```bash
kubectl apply -f .
```

## Applications
Here is the source code for apps used in this exercise.
- [todo-app](../../todo-app)
- [todo-backend](../../todo-backend)
