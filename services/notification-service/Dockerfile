FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./  ./

RUN go build -o notification-service main.go

RUN go get -u gopkg.in/gomail.v2

EXPOSE 8084

CMD ["./notification-service"]