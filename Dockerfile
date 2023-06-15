FROM alpine:3.18

COPY ./build /app

WORKDIR /app

ENTRYPOINT ["/app/main"]