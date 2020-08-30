FROM alpine:3.7

RUN mkdir /app

COPY artifacts/outstiti /app

RUN chmod +x /app/outstiti

ENTRYPOINT ["/app/outstiti"]