FROM golang:1.25-alpine
WORKDIR /app
COPY . .
RUN go get .
RUN go build -o todo-backend .
CMD [ "./todo-backend" ]
