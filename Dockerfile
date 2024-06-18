FROM golang:1.22

WORKDIR /app
COPY . /app
ENV CGO_ENABLED=0

RUN go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=0 /app/server .
COPY --from=0 /app/configs ./configs

EXPOSE 8080


CMD ["./server"]