FROM golang:1.10.1

ENV GOPATH /go

ENV PATH /go/bin:$PATH
RUN go get github.com/tools/godep

RUN mkdir -p /go/src/github.com/votinginfoproject/sms-web
WORKDIR /go/src/github.com/votinginfoproject/sms-web

RUN touch .env

COPY . /go/src/github.com/votinginfoproject/sms-web

RUN godep go test ./...

EXPOSE 8080

RUN godep go build -o sms-web sms-web.go

CMD ./sms-web
