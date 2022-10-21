# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/user/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
# COPY start.sh .
# COPY wait-for.sh .
COPY internal/adapters/migration ./internal/adapters/migration
RUN apk add --no-cache bash

COPY tls/ca.pem ./tls/ca.pem
COPY tls/server.pem ./tls/server.pem
COPY tls/server-key.pem ./tls/server-key.pem


# EXPOSE 8080
CMD [ "/app/main" ]
# ENTRYPOINT [ "/app/start.sh" ]
