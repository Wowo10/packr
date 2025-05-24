FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first, for caching dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app cmd/api/main.go

RUN chmod +x /app/app

# Set working directory
WORKDIR /app
EXPOSE 7000

ENTRYPOINT ["./app"]