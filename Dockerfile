FROM golang:1.10

WORKDIR $GOPATH/src/hello_kubernetes

COPY . .
RUN go build

CMD ./hello_kubernetes
