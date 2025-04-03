FROM golang:1.24.1

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy