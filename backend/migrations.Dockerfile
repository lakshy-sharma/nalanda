# Build stage: Used for building the application binary.
FROM golang:1.22.0-alpine as buildstage

COPY ./workspace /build
WORKDIR /build

RUN go mod tidy && \
    go build