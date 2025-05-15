FROM golang:1.23

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /app/main .
CMD ["/app/main"]