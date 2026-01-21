# Exercise
Here are the manifests of the ping-pong app and the logging app.

## Environment
There is `PINGPONG_URL` env variable defined in the `log-deployment.yaml`. You should not need to modify it.

## Deployment
```bash
# Create namespace if not created
kubectl apply -f https://raw.githubusercontent.com/Ben-PP/kubernetes-mooc/refs/tags/2.4.1/manifests/namespaces/exercises.yaml

# Apply the rest of the files
kubectl apply -f .
```

## Applications
Here is the source code for apps used in this exercise.
- [log-apps](../../log_output)
- [ping-pong](../../ping-pong)
