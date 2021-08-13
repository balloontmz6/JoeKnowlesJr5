FROM alpine:latest
RUN apk add --no-cache tzdata
RUN mkdir /app
WORKDIR /app

ADD ./bin/imageserver_linux /app/imageserver_linux

EXPOSE 8080

CMD ["./imageserver_linux"]