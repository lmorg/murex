FROM centos

ENV GOPATH=/go
ADD . /go/src/github.com/lmorg/murex
RUN yum --setopt=tsflags='' -y install golang which coreutils man man-pages
RUN go test github.com/lmorg/murex/...
RUN go build github.com/lmorg/murex
