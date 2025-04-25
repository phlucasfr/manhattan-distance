FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o examples/manhattan-app ./examples/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/examples/manhattan-app .

CMD ["/app/manhattan-app"]
