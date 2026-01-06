# ---------- Build Stage ----------
#     FROM golang:1.24.0-alpine AS builder




# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o app ./cmd

# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /app/app .

# # EXPOSE 8081

# CMD ["./app"]

# ---------- Build Stage ----------
    FROM golang:1.24-alpine AS builder

    WORKDIR /app
    
    # کپی go mod files
    COPY go.mod go.sum ./
    RUN go mod download
    
    # کپی کل سورس
    COPY . .
    
    # Build با مسیر صحیح
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/main.go
    
    # ---------- Run Stage ----------
    FROM alpine:latest
    
    # نصب dependencies
    RUN apk --no-cache add ca-certificates tzdata
    
    WORKDIR /root/
    
    # کپی binary
    COPY --from=builder /app/main .
    
    # اگر migrations دارید
    # COPY --from=builder /app/migrations ./migrations
    
    EXPOSE 8081
    
    CMD ["./main"]
