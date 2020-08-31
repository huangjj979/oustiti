# A Golang Web Application Template With Github Action  ![Action](https://github.com/chfanghr/oustiti/workflows/Action/badge.svg) ![docker version](https://img.shields.io/docker/v/chfanghr/outstiti?label=docker)

## Workflow

###  Build
1. Build binary and run some tests (calling [build script](tools/build.sh) and [test script](tools/test.sh) from [Github action](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L14))
2. Build docker image (using [docker build action](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L19) and [Dockerfile](tools/Dockerfile))
3. Pushing docker image built in the previous process to [dockerhub](https://hub.docker.com/repository/docker/chfanghr/outstiti) (using [docker push action](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L23)

### Human who maintains the build system
* Configure the build system(github action scripts, build scripts, test scripts and so on)
* Make sure that app can be tested and built on the cloud
* Make sure that the app docker image is pushed to dockerhub per success build
* Deploy the docker image to the production server 

### Note
* [What is dockerhub](https://docs.docker.com/docker-hub/)
    - In short, it's a service for hosting container images
* [About how to configure github action](https://docs.github.com/en/actions/getting-started-with-github-actions)
* [Commit message specification](https://gist.github.com/brianclements/841ea7bffdb01346392c) 
* Related github actions:
    - [actions/checkout](https://github.com/actions/checkout) is used to checkout our source code
    - [actions/setup-go](https://github.com/actions/setup-go) is used to add go tools to our build environment
    - [docker/build-push-action](https://github.com/docker/build-push-action) is used to build our app image and then push it to dockerhub 
* Dockerhub credentials([username](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L22) and [password](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L23)) are stored using [github secret](https://docs.github.com/en/actions/configuring-and-managing-workflows/using-variables-and-secrets-in-a-workflow)
* `go build` 
    - `-v` to enable verbose log
    - `-o output` to specify output binary, for example `go build -o artifacts/outstiti .` builds an app binary, names it `outstiti` and saves it to `artifacts` directory
    - Why `-tags netgo` and  `export CGO_ENABLED=0`: our docker app is based on alpine linux, which doesn't have native net lib support. See this question: [Go-compiled binary won't run in an alpine docker container on Ubuntu host](https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host)
* [`docker run`](https://docs.docker.com/engine/reference/run/)
* Set environmental variable `ADDR` to make the example app run on specific address  
* [What is docker-compose](https://docs.docker.com/compose/)
<!--hello world-->
