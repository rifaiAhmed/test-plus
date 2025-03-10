FROM golang:1.24.1

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o test-plus

RUN chmod +x test-plus

EXPOSE 8080

CMD ["./test-plus"]