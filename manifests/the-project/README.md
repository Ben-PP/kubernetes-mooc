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

## Applications
Here is the source code for apps used in this exercise.
- [todo-app](../../todo-app)
- [todo-backend](../../todo-backend)
- [todo-script](../../todo-script)
