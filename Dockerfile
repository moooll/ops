FROM golang:latest

ADD . /app

WORKDIR /app

RUN go build

ENTRYPOINT /app/entryfile.sh

EXPOSE 8888
