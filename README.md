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

## Contributing

Contributions are welcome! Please read our [contributing guidelines](CONTRIBUTING.md) and check the [open issues](https://github.com/checkcle/checkcle/issues).

1. Fork the repository
2. Create your feature branch (`git checkout -b feat/my-feature`)
3. Commit your changes (`git commit -m 'feat: add my feature'`)
4. Push to the branch (`git push origin feat/my-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Acknowledgements

This project is a fork of [operacle/checkcle](https://github.com/operacle/checkcle). Thanks to all original contributors.
