FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod download

ENV PORT=8080

RUN go build ./main.go

CMD ["./main"]