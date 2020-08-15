FROM golang:1.13-buster as builder

WORKDIR /go/src/github.com/johnwtracy/personal/apiserver

COPY ./go.mod /go/src/github.com/johnwtracy/personal/apiserver/
COPY ./go.sum /go/src/github.com/johnwtracy/personal/apiserver/

RUN go mod download

ADD . /go/src/github.com/johnwtracy/personal/apiserver

RUN GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10

COPY --from=builder /go/bin/app /

ENTRYPOINT ["/app"]
