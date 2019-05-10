FROM debian:stretch-slim

RUN apt-get update && \
    apt-get -y install \
        build-essential \
        git \
        telnet \
        vim \
        wget

WORKDIR "/root"

ARG GO_VERSION=10.3

RUN wget https://dl.google.com/go/go1.$GO_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.$GO_VERSION.linux-amd64.tar.gz && \
    rm go1.$GO_VERSION.linux-amd64.tar.gz && \
    mkdir ~/.go

ENV GOPATH="/root/.go"
ENV GOROOT="/usr/local/go"
ENV PATH="/usr/local/go/bin:${PATH}"

RUN go get gopkg.in/urfave/cli.v1
RUN go get github.com/unification-com/mainchain
RUN go get github.com/unification-com/oracle
RUN go install github.com/unification-com/oracle/cmd/wrkoracle
