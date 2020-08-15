FROM golang:1.13-buster as builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10

COPY --from=builder /go/bin/app /

ENTRYPOINT ["/app"]
