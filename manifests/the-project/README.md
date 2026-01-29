# Project
Here are the manifests of the course project

## Environment
There is `TODOBACKEND_URL` env variable defined in the `app-deployment.yaml` which you will have to modify to match your set up. This is the url to which the app makes the GET request for todos and to which browser sends POST requests to create todos.

### Secrets
The secrets.enc.yaml contains the postgres password and is encrypted with sops.

## Deployment
To deploy, you can use the `kubeapply.sh` script. Give the AGE key file as a parameter for it.
```bash
./kubeapply <path/to/key.txt>
```

To delete all configurations except for the namespace you can also use the `kubeapply`.
```bash
./kubeapply -d
```

To access prometheus, you have to portforward.
```bash
kubectl -n prometheus port-forward kube-prometheus-stack-<rest-of-grafan-pod-name> 3000
```
> Note to myself: This is exposed with nginx on port 3030 which redirects to localhost:3000

## Applications
Here is the source code for apps used in this exercise.
- [todo-app](../../todo-app)
- [todo-backend](../../todo-backend)
- [todo-script](../../todo-script)

# Descriptions
Here are some details what was done in exercises.
## 2.10
- [Backend app](../../todo-backend) was modified to write log messages to stdout.
- Prometheus stack was deployed according to course material
- Loki stack was deployed according to course material
- Nginx was set up as reverse proxy to access the localhost portforward of kubernetes
## 4.3
PromQL query for exercise 4.3 is here.
```
kube_statefulset_created{namespace="prometheus"}
```
