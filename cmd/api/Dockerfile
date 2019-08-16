FROM golang:alpine


WORKDIR /go/src/github.com/vasarostik/go_blog/

COPY ./cmd/api ./cmd/api
COPY ./cmd/migration ./cmd/migration
COPY ./pkg ./pkg
COPY ./go.mod* ./go.sum* ./migration.sh* ./
COPY ./vendor ./vendor

RUN chmod +x migration.sh
RUN ls
RUN go build -o ./cmd/api/api ./cmd/api
RUN go build -o ./cmd/migration/migration ./cmd/migration

