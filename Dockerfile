FROM golang:1.22-alpine

WORKDIR /app
ENTRYPOINT ./startup.sh