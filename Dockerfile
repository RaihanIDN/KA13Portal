# Gunakan base image Golang
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy go mod dan file project
COPY go.mod ./
RUN go mod download

COPY . .

# Build aplikasinya
RUN go build -o main .

# Jalankan aplikasinya
CMD ["./main"]
