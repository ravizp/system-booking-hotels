FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./   

RUN go mod download

COPY ./ ./

RUN go build -o booking-service main.go

EXPOSE 8083

CMD ["./booking-service"]