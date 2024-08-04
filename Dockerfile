FROM golang:1.22.5-alpine3.20 AS builder
WORKDIR /app

ARG MIGRATE_VERSION=4.17.1
ARG WAIT_FOR_VERSION=2.2.4

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz
RUN curl -L -O https://github.com/eficode/wait-for/releases/download/v${WAIT_FOR_VERSION}/wait-for

# Cannot write under /app because of the --mount=type=bind
RUN --mount=type=bind,target=. go build -o /bin/server main.go

# migrate environment
FROM alpine:3.20 AS migrate
LABEL org.opencontainers.image.source="https://github.com/kompiro/simplebank"

ENV DB_SOURCE=postgresql://postgres:postgres@db:5432/simple_bank?sslmode=disable
WORKDIR /app

COPY --from=builder /app/migrate /app/migrate
COPY --from=builder /app/wait-for /app/wait-for
COPY db/migration ./migration
RUN chmod +x /app/wait-for

ENTRYPOINT [ "/app/wait-for", "db:5432", "--" ]
CMD ["sh", "-c", "/app/migrate -path /app/migration -database ${DB_SOURCE} -verbose up" ]

# running environment
FROM alpine:3.20 AS app
LABEL org.opencontainers.image.source="https://github.com/kompiro/simplebank"

WORKDIR /app
COPY --from=builder /bin/server /app/server
COPY app.env .

EXPOSE 3000
CMD ["/app/server"]
