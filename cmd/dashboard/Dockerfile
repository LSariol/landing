# -- Build --
FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o dashboard ./cmd/dashboard

# -- Final --
FROM alpine:latest
WORKDIR /app

RUN mkdir -p /app/dashboard

COPY --from=builder /app/dashboard /dashboard
COPY static/ /app/static

EXPOSE 3020
CMD ["/dashboard"]