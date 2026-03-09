# {{cookiecutter.repository_name}}

## About

__TBD__

## Development

### Dependency Deployment

- Deploy the CDS development environment `cds-go-dev` using the `compose.yaml`
  in the repository root directory.

  ```bash
  docker compose up -d
  ```

- Wait for the environment to be up and healthy. For instance, run
  `docker compose ps` and wait for all containers to be `up/running (healthy)`.

### Microservice Build and Deploy

1. Enter the `cds-go-dev` container:

   ```console
   $ docker compose exec -u dev cds-go-dev /bin/bash -l
   dev@cds-go-dev:~/ws/{{cookiecutter.service_name}}$ pwd
   /home/dev/ws/{{cookiecutter.service_name}}
   ```

1. Generate the Kubernetes deployment config file
   `{{cookiecutter.service_name}}-deploy.yaml` using the settings in the
   `helm/{{cookiecutter.service_name}}/dev-values.yaml` file:

   ```bash
   make config
   ```

1. Initialize the database:

   ```bash
   make setup_db
   ```

1. Deploy the `{{cookiecutter.service_name}}` service with `skaffold`:

   ```bash
   skaffold dev --port-forward --tail
   ```

   A [Skaffold](https://skaffold.dev/) configuration is provided to build,
   deploy and tail the logs of the `{{cookiecutter.service_name}}` service container. For
   example, the output of a sample run should be similar to:

   ```console
   dev@cds-go-dev:~/ws/{{cookiecutter.service_name}}$ skaffold dev --port-forward --tail
   ...
   Starting deploy...
     - serviceaccount/vault-auth-{{cookiecutter.service_name}} created
     - configmap/{{cookiecutter.service_name}}-dashboard created
     - service/{{cookiecutter.service_name}}-nb-rest created
     - service/{{cookiecutter.service_name}}-grpc created
     - deployment.apps/{{cookiecutter.service_name}} created
   Waiting for deployments to stabilize...
     - deployment/{{cookiecutter.service_name}}: creating container {{cookiecutter.service_name}}
        - pod/{{cookiecutter.service_name}}-675579d44d-gd8m4: creating container {{cookiecutter.service_name}}
     - deployment/{{cookiecutter.service_name}} is ready.
   Deployments stabilized in 22.605 seconds
   Port forwarding service/{{cookiecutter.service_name}}-grpc in namespace default, remote port 5001 -> http://127.0.0.1:5001
   Port forwarding service/{{cookiecutter.service_name}}-nb-rest in namespace default, remote port 5000 -> http://127.0.0.1:5000
   Press Ctrl+C to exit
   [{{cookiecutter.service_name}}] {"level":"INFO","timestamp":"2023-02-19T13:43:13.460Z","caller":"/build/cmd/{{cookiecutter.service_name}}/main.go:48","message":"starting RDS Inventory Manager (RIM)...","version":"0.0.0"}
   ...
   Watching for changes...
   ```
