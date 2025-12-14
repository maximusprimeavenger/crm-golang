FROM golang:1.23

WORKDIR /app

COPY go.* /app/

RUN go mod download

COPY . /app/

WORKDIR /app/cmd

RUN go build -o /app/main .

CMD ["/app/main"]
