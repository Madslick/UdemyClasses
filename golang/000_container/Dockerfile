FROM golang:1.17-bullseye

RUN apt-get update

RUN apt-get install curl -y
RUN apt-get install vim protobuf-compiler iputils-ping -y

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

WORKDIR /go/src/golang-examples

COPY . .

CMD [ "/bin/bash" ]