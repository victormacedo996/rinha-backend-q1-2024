ARG GO_VERSION=1.21.6
ARG ALPINE_VERSION=3.19

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download

COPY ./api .
RUN go build -o ./app ./cmd/api/main.go

FROM alpine:${ALPINE_VERSION}

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 8080

ENTRYPOINT ["./app"]