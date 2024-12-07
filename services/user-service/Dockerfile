FROM golang:1.23.3 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./  ./

RUN go build -o user-service main.go

EXPOSE 8081

CMD ["./user-service"]