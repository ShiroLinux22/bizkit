FROM golang:alpine as build

COPY gateway /build/gateway
COPY common /build/common
WORKDIR /build/gateway

RUN unset GOPATH && go build -o ./bin/gateway .

FROM alpine

COPY --from=build /build/gateway/bin /build/
WORKDIR /build

CMD ["/build/gateway"]
