# pusher

pusher is an api server based on fiber to handle telegram bot `/sendMessage` api

pusher is for personal usage, thus it can serve only one user

![](https://goreportcard.com/badge/github.com/bipy/pusher)

## CI

Auto build docker image by GitHub Actions

## Setup Docker

**Quick Start**

```bash
docker run -d --name pusher -p 3333:3333 \
	-e TG_TOKEN=<your_tg_bot_token> \
	-e CHAT_ID=<your_chat_id> \
	bipy/pusher:latest
```

**Fully Featured**

```bash
docker run -d --name pusher -p 3333:3333 \
	-e TG_TOKEN=<your_tg_bot_token> \
	-e CHAT_ID=<your_chat_id> \
	-e SERVER_HOST=0.0.0.0 \
	-e SERVER_PORT=3333 \
	-e SECURE_KEY=pAssw0rd
	bipy/pusher:latest
```

**Do Use Behind a Reverse Proxy with HTTPS**

## Environment

| Name        | Value  | Description           |
| ----------- | ------ | --------------------- |
| TG_TOKEN    | string | <your_tg_bot_token>   |
| CHAT_ID     | string | <your_chat_id>        |
| SERVER_HOST | string | default "0.0.0.0"     |
| SERVER_PORT | int    | default 3333          |
| SECURE_KEY  | string | default "" (disabled) |

## API

You can use pusher by `GET` or `POST` method

**GET**

| Param   | Value  | Description                 |
| ------- | ------ | --------------------------- |
| text    | string | your message                |
| preview | 0 or 1 | default 0 (disable preview) |



**POST**

```json
{
    "text": "your message",
    "disable_web_page_preview": true
}
```

**(Optional) Use `SECURE_KEY`**

| HTTP Header | Value  | Description           |
| ----------- | ------ | --------------------- |
| Secure-Key  | string | used for verification |

1. Set environment 
2. Add header for each request



# Questions

**Q: Why would I use pusher rather than the official api?**

1. Telegram official api is blocked in some area

2. pusher will remember bot token and chat id for you, so you can simply push your message by a GET query:

    ```bash
    curl -X GET https://example.com\?text=hello%20world
    ```

**Q: Where is the release?**

1. Deploying with docker is recommended, plz refer [bipy/pusher - DockerHub](https://hub.docker.com/r/bipy/pusher) or build image by yourself

2. Also, if you want to build and run cli in a shell or as a service:

    ```bash
    git clone https://github.com/bipy/pusher.git; cd pusher
    
    go build
    
    TG_TOKEN=<your_tg_bot_token> CHAT_ID=<your_chat_id> ./pusher
    ```

# Thanks to

[gofiber/fiber](https://github.com/gofiber/fiber)

[create-go-app/fiber-go-template](https://github.com/create-go-app/fiber-go-template)

[muety/webhook2telegram](https://github.com/muety/webhook2telegram)

# License

MIT License

Copyright (c) 2021 bipy