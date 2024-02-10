# Build stage: Used for building the application binary.
FROM golang:1.22.0-alpine as buildstage

COPY ./workspace /build
WORKDIR /build

RUN go mod tidy && \
    go build

# Final stage: Used for deploying the final application binary.
FROM alpine
COPY --from=buildstage /build/nalanda_backend /app/nalanda_backend
COPY --from=buildstage /build/nalanda.toml /app/nalanda.toml
WORKDIR /app

CMD /app/nalanda_backend