FROM golang:1.25-alpine
WORKDIR /app
COPY . .
RUN go get .
RUN go build -o ping-pong .
CMD ["./ping-pong"]

