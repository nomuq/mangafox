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

RUN go build -o mangafox-server ./cmd/server
RUN go build -o mangareader-indexer ./cmd/mangareader
RUN go build -o mangatown-indexer ./cmd/mangatown
RUN go build -o mangadex-indexer ./cmd/mangadex


FROM dkron/dkron AS cron-builder 


FROM alpine

COPY --from=cron-builder  /opt/local/dkron/ /

COPY --from=builder /mangafox/mangafox-server /mangafox-server 
COPY --from=builder /mangafox/mangareader-indexer /mangareader-indexer 
COPY --from=builder /mangafox/mangatown-indexer /mangatown-indexer 
COPY --from=builder /mangafox/mangadex-indexer /mangadex-indexer 

ENTRYPOINT [ "dkron" ]
