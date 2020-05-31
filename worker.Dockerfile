FROM golang:alpine AS builder

RUN apk add --no-cache git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /mangafox

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o mangadex-worker ./cmd/worker

FROM alpine

ARG MONGO_URI
ARG REDIS_URI

WORKDIR /mangafox

COPY --from=builder /mangafox/mangafox-worker /mangafox/mangafox-worker 

ENTRYPOINT [ "/mangafox/mangafox-worker" ]