FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum* ./
COPY . .
RUN go mod tidy && go build -o notification-service .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/notification-service .
EXPOSE 8085
CMD ["./notification-service"]