FROM golang:1.22.5-alpine as builder
WORKDIR /app

RUN --mount=type=bind,target=. go build -o /bin/server main.go

# running environment
FROM alpine:3
WORKDIR /app
COPY --from=builder /bin/server /app/server
COPY app.env .

EXPOSE 3000
CMD ["/app/server"]
