FROM golang:1.21.6-alpine AS build-env
RUN apk update && apk upgrade && apk add --no-cache \
    git \
    build-base
ENV GOPATH=/
ENV GO111MODULE=on
WORKDIR /splitwiser
COPY go.mod .
RUN go mod download
COPY . .
RUN go get -d -v ./...
RUN go build -o app

FROM alpine
WORKDIR /splitwiser
COPY --from=build-env /splitwiser/app /splitwiser
ENTRYPOINT ./splitwiser
