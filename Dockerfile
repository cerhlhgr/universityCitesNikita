# Intermediate stage: Build the binary
FROM golang:1.18 as builder

RUN mkdir -p /bot
ADD . /bot
WORKDIR /bot

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/app ./cmd/bot/main.go

# Final stage: Run the binary
FROM scratch

# and finally the binary
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bot/bin/app /app
CMD ["/app"]
