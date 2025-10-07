# -- Build --
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o landing ./cmd/landing

# -- Final --
FROM alpine:latest
WORKDIR /app

RUN mkdir -p /app/landing

COPY --from=builder /app/landing /landing

EXPOSE 3000
CMD ["/landing"]