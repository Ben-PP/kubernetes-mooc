FROM golang:1.25-alpine
WORKDIR /app
ENV LOG_DIR=/app/logs
COPY go.mod go.sum main.go .
RUN go get .
RUN go build -o ping-pong .
CMD ["./ping-pong"]

