# Build virtual-lsobus in a stock Go builder container
FROM golang:1.14.2-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /qlcchain/go-virtual-lsobus
RUN cd /qlcchain/go-virtual-lsobus && make clean build

# Pull virtual-lsobus into a second stage deploy alpine container
FROM alpine:3.11.3

ENV LSOBUSHOME /lsobus

RUN apk --no-cache add ca-certificates && \
    addgroup lsobus && \
    adduser -S -G lsobus lsobus -s /bin/sh -h "$LSOBUSHOME" && \
    chown -R lsobus:lsobus "$LSOBUSHOME"

USER lsobus

WORKDIR $LSOBUSHOME

COPY --from=builder /qlcchain/go-virtual-lsobus/build/virtual-lsobus  /usr/local/bin/virtual-lsobus

ENTRYPOINT [ "virtual-lsobus"]

VOLUME [ "$LSOBUSHOME" ]
