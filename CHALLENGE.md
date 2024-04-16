## Pre-requisites

### Tools
- git

### Getting started
1. Create a new repository in your Github account
2. Clone the repository in your local machine `git@github.com:carlosgit2016/devops-challenge.git`

## Challenge

> Note: It is not necessary to run the pipeline. Check the `README.md` for more info about the application

1. Finish the creation of the Dockerfile for the application.
    * Note that the static files need to packed together    
    * In golang the files that control the dependencies are `go.mod` and `go.sum`
2. Create a ECR registry in terraform.
    * The name of the repository should be `rdicidr-calculator`
3. Develop a pipeline in Github that will test and build a golang application and deploy the generated image to the ECR repository.
    * The `latest` can be the only tag
    * consider that AWS secrets are present in `secrets.AWS_ACCESS_KEY_ID` and `secrets.AWS_ACCESS_SECRET_KEY`
4. Create k8s resources to host the application 

### Structure

- `cmd/main` is the main golang application that serves the static files
- `web/public` are the static files
- `Dockerfile` is where the specification of the image need to go
- `infrastructure/`is where the IaC for the ECR needs to go
- `.github/workflows` is where the workflow needs to go 
- `k8s/` is where the manifests are

### What can be improved for the future ?
<?-- Leave your notes here-->
