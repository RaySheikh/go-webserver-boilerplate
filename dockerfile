FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-webserver ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-webserver .

COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./go-webserver"]
