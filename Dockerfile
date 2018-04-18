FROM golang:1.10

WORKDIR $GOPATH/src/hello_kubernetes

COPY . .
RUN go build

CMD ./k8s_example
