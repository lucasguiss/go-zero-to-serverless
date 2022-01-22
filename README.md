# go-zero-to-serverless
Simple Go API deployed on Cloud Run

## Deploy app
To deploy the application, execute the `deploy.sh` file:
```shell
$ sh deploy.sh my-project-name
```

## Run Dockerfile locally
```shell
$ docker build -t go-zero-to-serverless .
$ docker run -it go-zero-to-serverless bash
```

## Run project locally
```shell
$ export GCP_PROJECT=my-project
$ go mod tidy
$ go run main.go
```