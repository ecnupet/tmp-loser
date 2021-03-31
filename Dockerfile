FROM golang:1.15-alpine3.13 AS builder
ARG BUILD_ID
LABEL stage=builder
LABEL build=$BUILD_ID
COPY . /app
WORKDIR /app

RUN ln -s -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV GOOS=linux
ENV GOARCH=arm64
ENV CGO_ENABLED=0
RUN go build -o main ./cmd

FROM alpine:3.12.3
COPY --from=builder /app/main /app/main
COPY conf /app/conf
ENTRYPOINT [ "/app/main" ]
