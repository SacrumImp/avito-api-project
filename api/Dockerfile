FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -coverprofile=coverage.out ./...

RUN mkdir -p coverage

RUN go build -o /api ./cmd/avito-api

EXPOSE 8000

CMD ["./api"]