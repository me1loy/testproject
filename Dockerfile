FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o server ./cmd


RUN apk add --no-cache build-base

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/config.yml .

COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]
