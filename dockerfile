#compilacion
FROM golang:1.26.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY /src .
COPY /db/sqlc ./db/sqlc
RUN CGO_ENABLED=0 go build -o /app/api .

#binario
FROM alpine:3.24
RUN adduser -D -u 1000 app
RUN apk add --repository http://dl-cdn.alpinelinux.org/alpine/edge/testing hurl
USER app
COPY --from=builder /app/api /usr/local/bin/api
EXPOSE 8080
CMD ["api"]