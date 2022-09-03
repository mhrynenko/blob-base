FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/blob-base
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/blob-base /go/src/blob-base


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/blob-base /usr/local/bin/blob-base
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["blob-base"]
