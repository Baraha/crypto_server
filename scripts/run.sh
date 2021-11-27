#!/bin/bash
TARGET=cmd/app/
echo "DEVELOPMENT MODE. DON'T USE IN PRODUCTION !"
go build cmd/app/main.go  # Build a server