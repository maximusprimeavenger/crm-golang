FROM golang:1.23

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .
RUN go build -o /app/main ./cmd

CMD ["/app/main"]



