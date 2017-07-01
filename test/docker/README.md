# Docker Test Environment

Frankly there's no need to run _murex_ inside a docker container since
it is just a shell. But if you are nervous about compiling and running
this on your host system (eg you want to vet the code first) or maybe
you just can't be bothered to install the Go toolchain just to compile
this project, then the included Dockerfile is a convenient way to
install and/or sandbox the shell.

Suggested instructions on using the Dockerfile are as follows:

    # Pease ensure that your working directory is the parent directory
    # of this project. For exampple:
    cd /home/$USER/go/src/github.com/lmorg/murex

    # Create the docker container
    docker build -t murex -f test/docker/Dockerfile .

    # Run the tests
    docker run --rm --name murex -it murex \
        /bin/sh -c 'cd $GOPATH/src/github.com/lmorg/murex; test/regression_test.sh'

    # Run the shell
    docker run --rm --name murex -it murex /go/src/github.com/lmorg/murex/murex

