FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY cmd/main/go.mod cmd/main/go.sum ./
RUN go mod download
COPY cmd/main/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/crawler .

FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/crawler .
CMD ["./crawler"]
