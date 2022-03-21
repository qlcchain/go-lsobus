# Build lsobus in a stock Go builder container
FROM golang:1.18.0-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /qlcchain/go-lsobus
RUN cd /qlcchain/go-lsobus && make clean build

# Pull lsobus into a second stage deploy alpine container
FROM alpine:3.13.4

ENV LSOBUSHOME /lsobus

RUN apk --no-cache add ca-certificates && \
    addgroup lsobus && \
    adduser -S -G lsobus lsobus -s /bin/sh -h "$LSOBUSHOME" && \
    chown -R lsobus:lsobus "$LSOBUSHOME"

USER lsobus

WORKDIR $LSOBUSHOME

COPY --from=builder /qlcchain/go-lsobus/build/glsobus  /usr/local/bin/glsobus

ENTRYPOINT [ "glsobus"]

VOLUME [ "$LSOBUSHOME" ]
