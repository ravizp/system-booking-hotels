FROM golang:1.23.3 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./  ./

RUN go build -o orchestrator main.go

EXPOSE 8080

CMD ["./orchestrator"]