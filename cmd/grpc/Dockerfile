FROM golang:alpine


WORKDIR /go/src/github.com/vasarostik/go_blog/

COPY ./cmd/grpc ./cmd/grpc
COPY ./pkg ./pkg
COPY ./go.mod* ./go.sum* ./
COPY ./vendor ./vendor

RUN ls
RUN go build -o ./cmd/grpc/grpc ./cmd/grpc

#CMD ["go","run","./cmd/migration/main.go"]

CMD ["./cmd/grpc/grpc"]