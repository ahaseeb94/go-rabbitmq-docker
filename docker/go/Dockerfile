FROM golang:1.22.3-alpine

RUN mkdir /app

WORKDIR /app

RUN apk add build-base
RUN go install github.com/cosmtrek/air@latest

ADD ../../go.mod .

RUN go mod download
ADD ../.. .

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]