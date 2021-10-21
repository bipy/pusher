# pusher

pusher is an api server based on fiber to handle telegram bot `/sendMessage` api

pusher is for personal usage, thus it can serve only one user

## CI

Auto build docker image by GitHub Actions

## Setup Docker

```bash
docker run -d --name pusher -p 3333:3333 \
	-e TG_TOKEN=<your_tg_bot_token> \
	-e CHAT_ID=<your_chat_id> \
	bipy/pusher:latest
```

## Environment

| Name        | Value  | Description         |
| ----------- | ------ | ------------------- |
| TG_TOKEN    | string | <your_tg_bot_token> |
| CHAT_ID     | string | <your_chat_id>      |
| SERVER_HOST | string | default "0.0.0.0"   |
| SERVER_PORT | int    | default 3333        |
| SECURE_KEY  | string | default ""          |

## Api

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

**Use `SECURE_KEY` (Optional)**

| Header     | Value  | Description           |
| ---------- | ------ | --------------------- |
| Secure-Key | string | used for verification |

1. Set environment 
2. Add header for each request

# Thanks to

[gofiber/fiber](https://github.com/gofiber/fiber)

[create-go-app/fiber-go-template](https://github.com/create-go-app/fiber-go-template)

# LICENSE

MIT License

Copyright (c) 2021 bipy