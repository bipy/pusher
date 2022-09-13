# build
FROM golang:alpine
WORKDIR /data
COPY . .
RUN go build -o pusher .

# image
FROM alpine:latest
MAINTAINER bipy notbipy@gmail.com

WORKDIR /app

COPY --from=0 /data/pusher pusher
COPY --from=0 /data/entrypoint.sh entrypoint.sh

ENV SERVER_HOST "0.0.0.0"
ENV SERVER_PORT 3333
ENV TG_TOKEN ""
ENV CHAT_ID ""
ENV SECURE_KEY ""

EXPOSE $SERVER_PORT

ENTRYPOINT ["/app/entrypoint.sh"]