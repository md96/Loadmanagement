# Stage 1
FROM golang:1.24.1 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o load-management

# Stage 2
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/load-management .
EXPOSE 8082

CMD ["./load-management"]