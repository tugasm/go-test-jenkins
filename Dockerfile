FROM golang:1.19

WORKDIR /app

COPY main.go .

EXPOSE 8080

RUN go build -o main .

CMD ["./main"]