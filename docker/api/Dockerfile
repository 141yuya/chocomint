FROM golang:1.17.0-alpine3.13 AS build

WORKDIR /go/src/github.com/141yuya/go-clean-architecture
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./src/. ./
RUN go build

FROM alpine:3.13
COPY --from=build /go/src/github.com/141yuya/go-clean-architecture/go-clean-architecture /usr/local/bin/

EXPOSE 8080
CMD ["go-clean-architecture"]