# Murex build server

This docker container exists as an optional automated build and test server for
anyone who might not want to install `go` et al on their platform but who might
already be running Docker.

To run it, you will need `docker` install. `docker-compose` is also recommended
however it is possible to build and run the containers without it (not not as
conveniently).

These tests are designed so that they can run headless and will automatically
generate executables under `./test/buildserver/bin/`.

> All the following commands are run from the project root

To build `murex` and `docgen` run:

    docker-compose: up --build murex-build

(Please not it will take a long time to run on the first time. This is because
it is downloading and creating the build environment. Once that has been built
Docker will reuse it on all subsequent builds)

To run the tests and compile binaries for every supported platform:

    docker-compose: up --build murex-test

