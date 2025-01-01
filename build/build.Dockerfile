FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY cmd/demoapi/main.go .

RUN go build --race -o main .

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
