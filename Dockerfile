FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6 && go mod download
COPY . .
RUN swag init && CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.24.2

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --chown=appuser:appgroup --from=builder /app/main .
COPY --chown=appuser:appgroup --from=builder /app/database/migrations ./database/migrations

USER appuser

EXPOSE 8080
CMD ["./main"]