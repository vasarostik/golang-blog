FROM golang:alpine


WORKDIR /go/src/github.com/vasarostik/go_blog/

COPY ./cmd/nats ./cmd/nats
COPY ./pkg ./pkg
COPY ./go.mod* ./go.sum* ./
COPY ./vendor ./vendor

RUN ls
RUN go build -o ./cmd/nats/nats ./cmd/nats

CMD ["./cmd/nats/nats"]