## Pre-requisites

### Tools
- git

### Getting started
1. Create a new repository in your Github account
2. Clone the repository in your local machine `git@github.com:carlosgit2016/devops-challenge.git`

## Challenge

> Note: It is not necessary to run the pipeline since we don't have a linux host to connect to. But the pipeline can be tested if you want only with the golang application test and build.

Develop a pipeline in Github that will test and build a golang application and deploy the artifacts to a linux host, then start the application in the host. Check the `README.md` for more info about the application and its requisites to run.

The ssh private key can be found in `secrets.SSH_PRIVATE_KEY` and the way it can be used in Github actions is using `${{ secrets.SSH_PRIVATE_KEY }}`.

The other necessary information for the connection with the host is also present in Github secrets named: `HOST`, `USERNAME`, `PORT`.

In summary the only file that you will need to create and edit is `.github/workflows/main.yaml`

### Structure

- `cmd/main` is the main golang application that serves the static files
- `web/public` are the static files
- `.github/workflows` is where the workflow needs to go 

### What can be improved for the future ?
<?-- Leave your notes here-->
