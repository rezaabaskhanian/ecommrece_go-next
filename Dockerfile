# ---------- Build Stage ----------
    FROM golang:1.24.0-alpine AS builder




WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8081

CMD ["./app"]
