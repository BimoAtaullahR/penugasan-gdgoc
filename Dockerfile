# Tahap 1: Build
FROM golang:1.25.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Build binary dengan nama 'main'
RUN go build -o main .

# Tahap 2: Run (Image kecil biar hemat & cepat)
FROM alpine:latest 

WORKDIR /root/
COPY --from=builder /app/main .

# Expose port 8080 (Default Cloud Run)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]