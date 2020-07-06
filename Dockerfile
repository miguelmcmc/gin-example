FROM golang:alpine as builder

RUN apk update ; apk add tzdata bash wget curl git
RUN apk add -U --no-cache ca-certificates

ADD . /go/src/test.net/gin-example
WORKDIR /go/src/test.net/gin-example

# RUN GO111MODULE=auto go mod download
RUN go get -v
RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux go build -a -installsuffix cgo -o /gin-example

FROM scratch
COPY --from=builder /gin-example /gin-example
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/gin-example"]
