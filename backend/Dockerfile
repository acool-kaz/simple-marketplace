FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder ./app/server .
COPY --from=builder ./app/vendor/ ./vendor/
COPY --from=builder ./app/public/ ./public/
CMD ["./server"]