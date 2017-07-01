FROM golang

ADD . /go/src/github.com/lmorg/murex
RUN go get github.com/chzyer/readline
RUN go get github.com/kr/pty
RUN go get github.com/Knetic/govaluate
RUN go build github.com/lmorg/murex

#RUN apt-get install -y timeout

#ENTRYPOINT /bin/sh /go/src/github.com/lmorg/murex/test/regression_test.sh