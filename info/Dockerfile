FROM golang:alpine as build

COPY info /build/info
COPY common /build/common
WORKDIR /build/info

RUN unset GOPATH && go build -o ./bin/info .

FROM alpine

COPY --from=build /build/info/bin /build/
WORKDIR /build

CMD ["/build/info"]
