FROM golang:1.17-alpine

RUN mkdir -p /app
COPY . /app
WORKDIR /app

RUN go build

CMD ["/app/subnet-calculator"]