#!/bin/sh
FROM golang:latest
RUN mkdir /app 
WORKDIR /app 
EXPOSE 8080
RUN go build cmd/app/main.go
RUN ./main
