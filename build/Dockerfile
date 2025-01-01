FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd/demoapi/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]

