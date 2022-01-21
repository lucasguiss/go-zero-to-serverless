FROM golang:1.17-alpine3.15 AS builder
WORKDIR /src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-zero-to-serverless .

FROM alpine:latest
COPY --from=builder /src/app/go-zero-to-serverless .

ENTRYPOINT ["./go-zero-to-serverless"]