FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build ./cmd/esme/main.go

EXPOSE 8020

ENTRYPOINT ["/app/main"]
