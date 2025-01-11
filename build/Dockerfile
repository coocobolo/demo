FROM alpine:latest AS root-certs

RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:1.22-alpine AS builder
ARG service_path
ENV SVC_PATH=${service_path}

WORKDIR /app

COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY go.mod go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o bin/main $SVC_PATH


FROM alpine:latest AS final
WORKDIR /root/
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --chown=1001:1001 --from=builder /app/bin .
CMD ["./main"]
