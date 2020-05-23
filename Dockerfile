FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /mangafox

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o mangafox-server ./cmd/server
RUN go build -o mangareader-indexer ./cmd/mangareader
RUN go build -o mangatown-indexer ./cmd/mangatown


FROM alpine

WORKDIR /mangafox

COPY --from=builder /mangafox/mangafox-server /mangafox/mangafox-server 
COPY --from=builder /mangafox/mangareader-indexer /mangafox/mangareader-indexer 
COPY --from=builder /mangafox/mangatown-indexer /mangafox/mangatown-indexer 