<div align="center">

# 📮 Pusher

**A lightweight, secure Telegram message delivery API**

[![Go Report Card](https://goreportcard.com/badge/github.com/bipy/pusher)](https://goreportcard.com/report/github.com/bipy/pusher)
[![Docker Pulls](https://img.shields.io/docker/pulls/bipy/pusher)](https://hub.docker.com/r/bipy/pusher)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bipy/pusher)](go.mod)

[Features](#-features) •
[Quick Start](#-quick-start) •
[API](#-api-reference) •
[Configuration](#-configuration) •
[Examples](#-usage-examples) •
[Contributing](#-contributing)

</div>

---

## 🎯 Overview

**Pusher** is a simple yet powerful HTTP API server that acts as a proxy for sending messages to Telegram. Built with [Echo](https://echo.labstack.com/) framework in Go, it simplifies message delivery by remembering your bot credentials and providing an easy-to-use REST interface.

### Why Pusher?

- **🚀 Simple Integration** - Just HTTP GET/POST requests, no Telegram API complexity
- **🔒 Secure** - Optional authentication via secure key headers
- **📦 Lightweight** - Minimal Docker image, low resource usage
- **🌍 Proxy-Friendly** - Works in regions where Telegram API is blocked
- **📱 Smart Handling** - Automatic message splitting and markdown escaping
- **💪 Production-Ready** - Graceful shutdown, health checks, and error handling

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🎨 **Multiple Formats** | Support for plain text and Telegram MarkdownV2 |
| 📊 **Request Compression** | Automatic gzip compression for API calls |
| 🔐 **Optional Auth** | Secure your endpoint with custom header authentication |
| 📨 **Message Splitting** | Automatically splits long messages into chunks |
| 🌐 **CORS Enabled** | Works seamlessly with web applications |
| 🏥 **Health Checks** | Built-in `/pulse` endpoint for monitoring |
| 🔗 **Link Preview Control** | Enable/disable link previews per message |
| 📍 **IP Tracking** | Automatically includes sender IP in messages |

---

## 🚀 Quick Start

### Docker (Recommended)

**Basic Setup:**

```bash
docker run -d \
  --name pusher \
  -p 3333:3333 \
  -e TG_TOKEN=your_telegram_bot_token \
  -e CHAT_ID=your_chat_id \
  bipy/pusher:latest
```

**With Security & Custom Port:**

```bash
docker run -d \
  --name pusher \
  -p 8080:8080 \
  -e TG_TOKEN=your_telegram_bot_token \
  -e CHAT_ID=your_chat_id \
  -e SERVER_HOST=0.0.0.0 \
  -e SERVER_PORT=8080 \
  -e SECURE_KEY=your_secure_password \
  bipy/pusher:latest
```

### Docker Compose

Create a `docker-compose.yml`:

```yaml
version: '3.8'

services:
  pusher:
    image: bipy/pusher:latest
    container_name: pusher
    restart: unless-stopped
    ports:
      - "3333:3333"
    environment:
      TG_TOKEN: "your_telegram_bot_token"
      CHAT_ID: "your_chat_id"
      SECURE_KEY: "your_secure_password"  # Optional
    # For production, use behind a reverse proxy with HTTPS
```

Then run:

```bash
docker-compose up -d
```

### Build from Source

```bash
# Clone the repository
git clone https://github.com/bipy/pusher.git
cd pusher

# Build the binary
go build -o pusher .

# Run with environment variables
export TG_TOKEN=your_telegram_bot_token
export CHAT_ID=your_chat_id
./pusher
```

---

## ⚙️ Configuration

Configure Pusher using environment variables:

| Variable | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `TG_TOKEN` | string | ✅ Yes | - | Your Telegram bot token from [@BotFather](https://t.me/botfather) |
| `CHAT_ID` | integer | ✅ Yes | - | Target chat ID to receive messages |
| `SERVER_HOST` | string | No | `0.0.0.0` | Server binding address |
| `SERVER_PORT` | integer | No | `3333` | Server listening port |
| `SECURE_KEY` | string | No | - | Optional authentication key |

### Getting Your Credentials

1. **Get Bot Token:**
   - Talk to [@BotFather](https://t.me/botfather) on Telegram
   - Create a new bot with `/newbot`
   - Copy the token provided

2. **Get Chat ID:**
   - Send a message to your bot
   - Visit: `https://api.telegram.org/bot<YOUR_TOKEN>/getUpdates`
   - Find your chat ID in the response

---

## 📡 API Reference

### Endpoints

#### `POST /` - Send Message (JSON)

Send a message using JSON payload.

**Request Body:**

```json
{
  "text": "Your message here",
  "preview": false,
  "markdown": false
}
```

**Parameters:**

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `text` | string | ✅ Yes* | - | Message content |
| `msg` | string | No | - | Alternative to `text` (fallback) |
| `preview` | boolean | No | `false` | Enable link preview |
| `markdown` | boolean | No | `false` | Enable Telegram MarkdownV2 parsing |

**Response:**

```json
{
  "code": 0,
  "msg": "OK",
  "resp": null
}
```

#### `GET /` - Send Message (Query String)

Send a message using URL parameters.

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `text` | string | ✅ Yes* | Message content (URL encoded) |
| `msg` | string | No | Alternative to `text` (fallback) |
| `preview` | flag | No | Include this parameter to enable link preview |
| `markdown` | flag | No | Include this parameter to enable markdown |

**Example:**

```bash
curl "http://localhost:3333/?text=Hello%20World"
```

#### `GET /pulse` - Health Check

Returns HTTP 204 No Content if the server is healthy.

**Example:**

```bash
curl -I http://localhost:3333/pulse
```

### Authentication

If `SECURE_KEY` is configured, include it in your requests:

**Header:**

```
Secure-Key: your_secure_password
```

**Example:**

```bash
curl -X POST http://localhost:3333/ \
  -H "Content-Type: application/json" \
  -H "Secure-Key: your_secure_password" \
  -d '{"text": "Authenticated message"}'
```

---

## 💡 Usage Examples

### Basic cURL Examples

**Simple Message:**

```bash
curl -X POST http://localhost:3333/ \
  -H "Content-Type: application/json" \
  -d '{"text": "Hello from Pusher!"}'
```

**With Link Preview:**

```bash
curl "http://localhost:3333/?text=Check%20this%20out%20https://github.com&preview"
```

**With Markdown:**

```bash
curl -X POST http://localhost:3333/ \
  -H "Content-Type: application/json" \
  -d '{"text": "*Bold* and _italic_ text", "markdown": true}'
```

### HTTPie Examples

[HTTPie](https://httpie.io/) makes it even simpler:

```bash
# Simple message
http POST localhost:3333 text="Hello World"

# With authentication
http POST localhost:3333 text="Secure message" Secure-Key:your_password

# With markdown
http POST localhost:3333 text="**Important** update" markdown:=true
```

### Python Example

```python
import requests

def send_telegram_message(text, markdown=False):
    response = requests.post(
        'http://localhost:3333/',
        json={
            'text': text,
            'markdown': markdown,
            'preview': False
        },
        headers={'Secure-Key': 'your_secure_password'}  # if enabled
    )
    return response.json()

# Usage
send_telegram_message("Deployment successful! 🚀")
```

### JavaScript/Node.js Example

```javascript
async function sendTelegramMessage(text, options = {}) {
  const response = await fetch('http://localhost:3333/', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Secure-Key': 'your_secure_password'  // if enabled
    },
    body: JSON.stringify({
      text,
      preview: options.preview || false,
      markdown: options.markdown || false
    })
  });
  return response.json();
}

// Usage
await sendTelegramMessage('Build completed successfully!');
```

### Bash Script Example

```bash
#!/bin/bash

PUSHER_URL="http://localhost:3333/"
SECURE_KEY="your_secure_password"

function notify() {
    local message="$1"
    curl -s -X POST "$PUSHER_URL" \
        -H "Content-Type: application/json" \
        -H "Secure-Key: $SECURE_KEY" \
        -d "{\"text\": \"$message\"}"
}

# Usage
notify "Backup completed at $(date)"
```

### Integration with Uptime Kuma

Pusher supports [Uptime Kuma](https://github.com/louislam/uptime-kuma) webhook notifications out of the box!

1. In Uptime Kuma, add a new notification
2. Select "Webhook"
3. Use your Pusher URL: `http://your-server:3333/`
4. Uptime Kuma will automatically send status updates to your Telegram

---

## 🏗️ Architecture

```
┌─────────────┐         ┌─────────────┐         ┌──────────────────┐
│   Client    │         │   Pusher    │         │  Telegram API    │
│             │         │             │         │                  │
│  HTTP GET   ├────────▶│  Echo Web   ├────────▶│  /sendMessage    │
│  HTTP POST  │         │  Framework  │         │                  │
└─────────────┘         └─────────────┘         └──────────────────┘
                               │
                               ├─ Authentication (optional)
                               ├─ Markdown Escaping
                               ├─ Message Splitting
                               ├─ Gzip Compression
                               └─ Error Handling
```

### Key Components

- **Echo Framework**: High-performance HTTP router
- **Middleware**: CORS, logging, recovery, authentication
- **Message Processing**: Automatic markdown escaping and splitting
- **Compression**: Gzip compression for Telegram API requests
- **Health Checks**: `/pulse` endpoint for monitoring

---

## 🔒 Security Best Practices

1. **Use HTTPS in Production**
   - Always deploy behind a reverse proxy (nginx, Caddy, Traefik)
   - Enable HTTPS to protect your `SECURE_KEY` in transit

2. **Strong SECURE_KEY**
   - Use a randomly generated, long password
   - Store in environment variables or secrets manager

3. **Network Security**
   - Don't expose Pusher directly to the internet
   - Use firewall rules to restrict access
   - Consider VPN or private network deployment

4. **Rate Limiting**
   - Implement rate limiting at the reverse proxy level
   - Monitor for unusual activity

Example nginx configuration:

```nginx
server {
    listen 443 ssl http2;
    server_name pusher.example.com;

    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    location / {
        proxy_pass http://localhost:3333;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        # Rate limiting
        limit_req zone=pusher_limit burst=10 nodelay;
    }
}
```

---

## 🐛 Troubleshooting

### Common Issues

**Problem: "TG_TOKEN environment variable is required"**

- Solution: Ensure `TG_TOKEN` is set correctly in your environment or Docker configuration

**Problem: "invalid CHAT_ID"**

- Solution: `CHAT_ID` must be a valid integer. Remove quotes in Docker compose files

**Problem: Messages not received**

- Solution: 
  - Verify bot token is correct
  - Check if chat ID is correct
  - Ensure your bot has been started (send `/start` to it)
  - Check logs for error messages

**Problem: 502 Bad Gateway errors**

- Solution: Telegram API might be blocked or unreachable from your server location

---

## 📊 Monitoring

### Health Check

Use the `/pulse` endpoint for monitoring:

```bash
# Check if service is running
curl -f http://localhost:3333/pulse || echo "Service down"
```

### Docker Health Check

Add to your `docker-compose.yml`:

```yaml
services:
  pusher:
    image: bipy/pusher:latest
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3333/pulse"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 5s
```

### Logging

Pusher uses Echo's built-in logger. To view logs:

```bash
# Docker logs
docker logs -f pusher

# Follow logs with timestamps
docker logs -f --timestamps pusher
```

---

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

### Development Setup

```bash
# Clone your fork
git clone https://github.com/yourusername/pusher.git
cd pusher

# Install dependencies
go mod download

# Run locally
export TG_TOKEN=your_token
export CHAT_ID=your_chat_id
go run main.go
```

### Code Style

- Follow Go best practices and conventions
- Run `gofmt` before committing: `gofmt -w .`
- Run `go vet` to check for issues: `go vet ./...`

---

## 📝 Use Cases

- **📦 CI/CD Notifications**: Get notified about build status, deployments
- **🔔 Server Monitoring**: Alert on server issues, resource usage
- **📊 Application Alerts**: Runtime errors, important events
- **🔐 Security Notifications**: Failed login attempts, suspicious activity
- **📈 Business Metrics**: Daily reports, sales notifications
- **🤖 Automation**: Integrate with scripts, cron jobs, webhooks

---

## 🙏 Acknowledgments

- Built with [Echo](https://echo.labstack.com/) - High performance, minimalist Go web framework
- Inspired by [muety/webhook2telegram](https://github.com/muety/webhook2telegram)
- Thanks to the Telegram Bot API team

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2021 bipy
```

---

## 📞 Support

- 🐛 **Issues**: [GitHub Issues](https://github.com/bipy/pusher/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/bipy/pusher/discussions)
- 🌟 **Star this repo** if you find it useful!

---

<div align="center">

**Made with ❤️ using Go**

[⬆ Back to Top](#-pusher)

</div>
