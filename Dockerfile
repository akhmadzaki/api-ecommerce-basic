FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go install github.com/swaggo/swag/cmd/swag@latest && go mod download
COPY . .
RUN swag init && CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/main .
COPY --from=builder /app/database/migrations /app/database/migrations
EXPOSE 8080
CMD ["./main"]