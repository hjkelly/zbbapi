FROM golang:alpine 

COPY . /go/src/github.com/hjkelly/zbbapi
WORKDIR /go/src/github.com/hjkelly/zbbapi
RUN go install

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/bin/zbbapi .
CMD ["./zbbapi"]  
