FROM golang:1.3.3

ENV GOPATH /go

RUN mkdir -p /go/src/github.com/votinginfoproject/sms-web
WORKDIR /go/src/github.com/votinginfoproject/sms-web

RUN touch .env

COPY . /go/src/github.com/votinginfoproject/sms-web

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD go-wrapper run
