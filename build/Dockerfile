FROM alpine:latest AS root-certs

RUN apk add -U --no-cache ca-certificates tzdata
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY . .

RUN go mod download
RUN go mod verify

RUN go build -o main ./cmd/demoapi/main.go


FROM scratch AS final
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --chown=1001:1001 --from=builder /app/main /myapp

USER app

ENTRYPOINT ["/myapp"]
