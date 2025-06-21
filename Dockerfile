# Stage 1: Build Go binary for Windows
FROM golang:1.24.1-windowsservercore-ltsc2022 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN set GOOS=windows&& set GOARCH=amd64&& go build -o load-management.exe

# Stage 2: Final image based on Windows Server Core
FROM mcr.microsoft.com/windows/servercore:ltsc2022

WORKDIR C:/app
COPY --from=builder /app/load-management.exe .

EXPOSE 8082

CMD ["load-management.exe"]