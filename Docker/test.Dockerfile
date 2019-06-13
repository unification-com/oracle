FROM golang:1.12.5

WORKDIR /root

RUN mkdir /root/src && \
    cd /root/src && \
    git clone https://github.com/unification-com/oracle.git

RUN cd /root/src/oracle && go install ./cmd/wrkoracle

RUN wrkoracle --version
