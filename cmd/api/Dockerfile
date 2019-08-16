FROM golang:alpine


WORKDIR /go/src/github.com/vasarostik/go_blog/

COPY ./cmd/api ./cmd/api
COPY ./cmd/migration ./cmd/migration
COPY ./pkg ./pkg
COPY ./go.mod* ./go.sum* ./
COPY ./vendor ./vendor

RUN ls
RUN go build -o ./cmd/api/api ./cmd/api

#CMD ["go","run","./cmd/migration/main.go"]

CMD ["./cmd/api/api"]