FROM golang:1.22-alpine

RUN apk update && apk --no-cache add git build-base

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]