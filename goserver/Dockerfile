FROM golang:1.17.2-alpine

RUN apk --no-cache add curl

WORKDIR /go/app

COPY . .

RUN go build -o ./server

EXPOSE 4000

ENTRYPOINT [ "./server" ]
