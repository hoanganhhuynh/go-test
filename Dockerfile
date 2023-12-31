FROM golang:1.20.5-alpine3.18

ENV GO111MODULE=on

WORKDIR /app

COPY . /app
COPY ./build/sqls /app/sqls/data

RUN go build -o example

EXPOSE 8088

CMD ["./example"]