FROM golang:1.25-alpine
WORKDIR /app
ENV DATA_DIR=/app/data
COPY . .
RUN go get .
RUN go build -o todo-app .
CMD [ "./todo-app" ]
