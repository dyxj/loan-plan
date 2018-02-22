FROM golang:1.8.5

ADD . /go/src/github.com/dyxj/loan-plan

WORKDIR /go/src/github.com/dyxj/loan-plan

RUN go build ./

CMD ["./loan-plan"]