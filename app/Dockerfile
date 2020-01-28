FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go build -o app main.go

CMD ["/go/src/app/app"]