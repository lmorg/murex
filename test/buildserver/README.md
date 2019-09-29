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


