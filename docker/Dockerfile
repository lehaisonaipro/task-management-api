FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o task-manager ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/task-manager .
CMD ["./task-manager"]