FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN adduser -D -u 1000 appuser
WORKDIR /app

COPY --from=builder --chown=appuser:appuser /app/main .

USER appuser

EXPOSE 8080

CMD ["./main"]
