# syntax=docker/dockerfile:1

FROM golang:1.24.1

WORKDIR /MatchManiaAPI

COPY . .
RUN go mod download
RUN go build -o MatchManiaAPI cmd/server/main.go

EXPOSE 8080

ENV MATCHMANIA_ENV=prod

CMD ["./MatchManiaAPI"]
