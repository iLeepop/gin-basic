# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder

WORKDIR /src

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /out/gin-basic ./cmd/main.go

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S app -G app

WORKDIR /app

COPY --from=builder /out/gin-basic /app/gin-basic
COPY conf/ /app/conf/

ENV TZ=Asia/Shanghai \
    GIN_MODE=release

EXPOSE 8080

USER app

ENTRYPOINT ["/app/gin-basic"]
