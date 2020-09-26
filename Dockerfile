FROM golang:alpine as builder

RUN apk -U --no-cache add git make

ENV GOROOT /usr/local/go

ADD . /src/epgu-generator
WORKDIR /src/epgu-generator

RUN make binary

#--- # ---#

FROM alpine

COPY --from=builder /src/epgu-generator/bin/epgu-generator /app/epgu-generator
COPY --from=builder /src/epgu-generator/work/templates /app/templates

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/epgu-generator

WORKDIR /app

ENTRYPOINT ["/app/epgu-generator"]
