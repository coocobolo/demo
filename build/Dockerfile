FROM golang:1.22 AS builder

ARG service_path
ARG go_version=1.22

ENV SERVICE_PATH=${service_path}
ENV GO_VERSION=${go_version}

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o output/app $SERVICE_PATH/...

FROM alpine:latest

ARG CMD_PATH
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /app/output .

CMD ["./app"]

