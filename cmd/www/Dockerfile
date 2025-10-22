# -- Build --
FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o www ./cmd/www

# -- Final --
FROM alpine:latest
WORKDIR /app

RUN mkdir -p /app/www

COPY --from=builder /app/www /www
COPY static/ /app/static

EXPOSE 3000
CMD ["/www"]