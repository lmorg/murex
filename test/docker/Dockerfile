FROM golang

ADD . /go/src/github.com/lmorg/murex
RUN go get github.com/chzyer/readline
RUN go get github.com/Knetic/govaluate
RUN go get github.com/BurntSushi/toml
RUN go get github.com/abesto/sexp
RUN go get gopkg.in/yaml.v2
RUN go get github.com/BurntSushi/toml
RUN go get golang.org/x/crypto/ssh/terminal
RUN go get github.com/disintegration/imaging
RUN go get golang.org/x/crypto/ssh/terminal
RUN go get golang.org/x/image/bmp
RUN go get golang.org/x/image/tiff
RUN go get golang.org/x/image/webp
RUN go get labix.org/v2/mgo/bson
RUN go get github.com/hashicorp/hcl
RUN go build github.com/lmorg/murex
