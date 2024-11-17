FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

EXPOSE 8080

RUN apk add build-base

RUN go mod tidy && \
    go build -o test_lwo .

CMD ["./test_lwo"]