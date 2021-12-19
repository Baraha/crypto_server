#!/bin/sh
FROM golang:latest
RUN mkdir /app 
WORKDIR /app 
EXPOSE 8080
CMD go run  cmd/app/main.go
# CMD ./main
