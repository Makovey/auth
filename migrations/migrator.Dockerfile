FROM alpine:3.17

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

RUN apk add --no-cache ca-certificates && update-ca-certificates

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose

RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/statements/*.sql migrations/
ADD migrations/migration.sh .
ADD .env .

RUN chmod +x migration.sh

ENTRYPOINT ["bash", "migration.sh"]