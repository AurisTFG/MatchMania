FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o MatchManiaAPI ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/.env.development .
COPY --from=builder /app/MatchManiaAPI .

EXPOSE 8080

CMD ["./MatchManiaAPI"]
