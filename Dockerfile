# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

ENV GO111MODULE=on

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download && go mod verify

RUN go build -o server ./cmd/app/

EXPOSE 5011

CMD [ "/app/server" ]