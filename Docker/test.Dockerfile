FROM golang:1.11

WORKDIR /root
RUN mkdir /root/src
RUN cd /root/src && git clone https://github.com/unification-com/oracle.git \
    && cd oracle \
    && git checkout mods # TODO: remove when merged

RUN cd /root/src/oracle && go install ./cmd/wrkoracle

RUN wrkoracle --version
