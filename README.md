# go-service-template

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
   dev@cds-go-dev:~/ws/example-service$ pwd
   /home/dev/ws/example-service
   ```

1. Generate the Kubernetes deployment config file
   `example-service-deploy.yaml` using the settings in the
   `helm/example-service/dev-values.yaml` file:

   ```bash
   make config
   ```

1. Initialize the database:

   ```bash
   make setup_db
   ```

1. Deploy the `example-service` service with `skaffold`:

   ```bash
   skaffold dev --port-forward --tail
   ```

   A [Skaffold](https://skaffold.dev/) configuration is provided to build,
   deploy and tail the logs of the `example-service` service container. For
   example, the output of a sample run should be similar to:

   ```console
   dev@cds-go-dev:~/ws/example-service$ skaffold dev --port-forward --tail
   ...
   Starting deploy...
     - serviceaccount/vault-auth-example-service created
     - configmap/example-service-dashboard created
     - service/example-service-nb-rest created
     - service/example-service-grpc created
     - deployment.apps/example-service created
   Waiting for deployments to stabilize...
     - deployment/example-service: creating container example-service
        - pod/example-service-675579d44d-gd8m4: creating container example-service
     - deployment/example-service is ready.
   Deployments stabilized in 22.605 seconds
   Port forwarding service/example-service-grpc in namespace default, remote port 5001 -> http://127.0.0.1:5001
   Port forwarding service/example-service-nb-rest in namespace default, remote port 5000 -> http://127.0.0.1:5000
   Press Ctrl+C to exit
   [example-service] {"level":"INFO","timestamp":"2023-02-19T13:43:13.460Z","caller":"/build/cmd/example-service/main.go:48","message":"starting RDS Inventory Manager (RIM)...","version":"0.0.0"}
   ...
   Watching for changes...
   ```
