# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

COPY . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o /go-docker

EXPOSE 8080

CMD [ "/go-docker" ]