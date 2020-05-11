FROM golang:1.8

RUN mkdir -p $GOPATH/src/tags
WORKDIR $GOPATH/src/tags

ADD . .

RUN go build -v
CMD ./tags
