FROM golang:alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o echo-server .

CMD [ "/app/echo-server", "server", "--config", "config.local.yaml"]