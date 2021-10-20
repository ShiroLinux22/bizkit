FROM golang:alpine as build

COPY info /build/logging
COPY common /build/common
WORKDIR /build/logging

RUN unset GOPATH && go build -o ./bin/logging .

FROM alpine

COPY --from=build /build/logging/bin /build/
WORKDIR /build

CMD ["/build/logging"]
