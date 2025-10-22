# -- Build --
FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o bot ./cmd/bot

# -- Final --
FROM alpine:latest
WORKDIR /app

RUN mkdir -p /app/bot

COPY --from=builder /app/bot /bot
COPY static/ /app/static

EXPOSE 3010
CMD ["/bot"]