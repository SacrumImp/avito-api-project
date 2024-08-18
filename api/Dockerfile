FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /api ./cmd/avito-api

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /api .

EXPOSE 8000

CMD ["./api"]