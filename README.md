# A Golang Web Application Template With Github Action  ![Action](https://github.com/chfanghr/oustiti/workflows/Action/badge.svg) ![docker version](https://img.shields.io/docker/v/chfanghr/outstiti?label=docker)

## Build Workflow
1. Build binary and maybe run some tests (calling [build script](tools/build.sh) from [Github action](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L14))
2. Build docker image (using [docker build action](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L19) and [Dockerfile](tools/Dockerfile))
3. Pushing docker image built in the previous process to [dockerhub](https://hub.docker.com/repository/docker/chfanghr/outstiti) (using [docker push action](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L23))
4. Deploy docker image to the server using `docker pull; docker run ...`

### Note
* [What is dockerhub](https://docs.docker.com/docker-hub/)
    - In short, it's a service for hosting container images
* Related github actions:
    - [actions/checkout](https://github.com/actions/checkout) is used to checkout our source code
    - [actions/setup-go](https://github.com/actions/setup-go) is used to add go tools to our build environment
    - [docker/build-push-action](https://github.com/docker/build-push-action) is used to build our app image and then push it to dockerhub 
* Dockerhub credentials([username](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L21) and [password](https://github.com/chfanghr/oustiti/blob/master/.github/workflows/action.yml#L22)) are stored using [github secret](https://docs.github.com/en/actions/configuring-and-managing-workflows/using-variables-and-secrets-in-a-workflow) 