FROM golang:latest

COPY . .

RUN go mod download
