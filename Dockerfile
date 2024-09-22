# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./

RUN go build -o /mock-api

# Final stage
FROM alpine:3.18

VOLUME /responses
COPY --from=builder /mock-api /mock-api


CMD ["/mock-api"]

EXPOSE 8080
