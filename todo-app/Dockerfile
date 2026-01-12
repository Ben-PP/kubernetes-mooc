FROM golang:1.25-alpine
WORKDIR /app
COPY . .
RUN go get .
RUN go build -o todo-app .
CMD [ "./todo-app" ]