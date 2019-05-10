FROM golang:1.11

RUN go get gopkg.in/urfave/cli.v1
RUN go get github.com/unification-com/mainchain
RUN go get github.com/unification-com/oracle
RUN go install github.com/unification-com/oracle/cmd/wrkoracle
