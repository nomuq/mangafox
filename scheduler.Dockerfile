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

RUN go build -o mangafox-scheduler ./cmd/scheduler

FROM alpine

WORKDIR /mangafox

COPY --from=builder /mangafox/mangafox-scheduler /mangafox/mangafox-scheduler 

EXPOSE 8080

ENTRYPOINT [ "/mangafox/mangafox-scheduler" ]