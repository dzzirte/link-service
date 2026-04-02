FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o main "./cmd"

EXPOSE 8081

CMD ["./main"]