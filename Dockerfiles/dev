FROM golang:alpine

COPY . /go/src/github.com/hjkelly/zbbapi
WORKDIR /go/src/github.com/hjkelly/zbbapi

RUN go install

CMD ["zbbapi"]
