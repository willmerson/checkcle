# Checkcle

> A fork of [operacle/checkcle](https://github.com/operacle/checkcle) — Open-source uptime monitoring and status page platform.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/docker-ready-blue.svg)](https://hub.docker.com)

## Overview

Checkcle is a lightweight, self-hosted uptime monitoring solution with a built-in status page. Monitor HTTP endpoints, TCP ports, and more — get notified when things go down.

## Features

- 🔍 **HTTP/HTTPS monitoring** — Check response time, status codes, and body content
- 🔌 **TCP port monitoring** — Verify service availability at the network level
- 📊 **Status page** — Public-facing status page for your services
- 🔔 **Notifications** — Alerts via email, Slack, Discord, and webhooks
- 📈 **Metrics & history** — Track uptime percentage and response time trends
- 🐳 **Docker-ready** — Easy deployment with Docker or Docker Compose
- 🔒 **Self-hosted** — Your data stays on your infrastructure

## Quick Start

### Docker Compose

```yaml
version: '3.8'
services:
  checkcle:
    image: checkcle/checkcle:latest
    ports:
      - "8090:8090"
    volumes:
      - ./data:/app/data
    environment:
      - CHECKCLE_SECRET=your-secret-key
    restart: unless-stopped
```

```bash
docker compose up -d
```

Then open [http://localhost:8090](http://localhost:8090) in your browser.

### Build from Source

**Prerequisites:** Go 1.21+, Node.js 18+

```bash
# Clone the repository
git clone https://github.com/checkcle/checkcle.git
cd checkcle

# Install dependencies and build
make build

# Run the application
./checkcle serve
```

## Configuration

Checkcle can be configured via environment variables or a `config.yaml` file.

| Variable | Default | Description |
|---|---|---|
| `CHECKCLE_HOST` | `0.0.0.0` | Bind address |
| `CHECKCLE_PORT` | `8090` | HTTP port |
| `CHECKCLE_SECRET` | *(required)* | JWT secret key |
| `CHECKCLE_DATA_DIR` | `./data` | Data directory path |
| `CHECKCLE_LOG_LEVEL` | `info` | Log level (debug/info/warn/error) |
| `CHECKCLE_CHECK_INTERVAL` | `60` | Default check interval in seconds |
| `CHECKCLE_REQUEST_TIMEOUT` | `10` | HTTP request timeout in seconds |

> **Personal note:** I run this locally with `CHECKCLE_LOG_LEVEL=debug` and `CHECKCLE_CHECK_INTERVAL=30` for faster feedback during development.

> **Tip:** If you're running this behind a reverse proxy (e.g. Caddy or nginx), make sure to set `CHECKCLE_HOST=127.0.0.1` so the app only listens locally and your proxy handles TLS termination.

> **Tip:** Set `CHECKCLE_REQUEST_TIMEOUT` to a lower value (e.g. `5`) if you want checks to fail faster for unresponsive endpoints rather than waiting the full default.

> **Personal note:** I also set `CHECKCLE_DATA_DIR=/var/lib/checkcle` when running on my home server so the data directory survives container rebuilds without relying on a relative path.

> **Personal note:** If you use Caddy as your reverse proxy, the following two-liner in your Caddyfile works well — `reverse_proxy 127.0.0.1:8090` under your domain block is all you need; Caddy handles HTTPS automatically via Let's Encrypt.
