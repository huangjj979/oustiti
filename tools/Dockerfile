FROM alpine:3.7

RUN mkdir /app

COPY artifacts/outstiti /app

RUN chmod +x /app/outstiti

ARG ADDR

ENV ADDR $ADDR

ENTRYPOINT ["/app/outstiti"]