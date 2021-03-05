FROM golang:1.16.0-alpine3.13 AS build-env
MAINTAINER zhougang <119098598@qq.com>
ENV GO111MODULE=on
ENV GIN_MODE=release
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
ADD .  /go/src/app
WORKDIR /go/src/app
RUN go build -v -o /go/src/app/app main.go


FROM ubuntu
ENV GIN_MODE=release
COPY --from=build-env /go/src/app/app /app/server
WORKDIR /app
RUN mkdir storage
EXPOSE 8080

CMD ["./server"]