# Simple Content Service

A Go-based simple content service built with clean architecture principles, featuring content synchronization from multiple providers (JSON and XML), caching with Redis, and PostgreSQL storage.

## Architecture

This project follows the **Clean Architecture** pattern with the following structure:

- **Core Layer**: Business logic, models, and ports (interfaces)
- **Adapters Layer**: External integrations (database, cache, HTTP, providers)

## Prerequisites

- **Docker**: 20.10 or higher
- **Docker Compose**: 2.0 or higher

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/mrfade/case-sss
cd case-sss
```

### 2. Environment Configuration

Edit the `.env` file to configure your environment variables.

### 3. Run with Docker Compose

```bash
docker-compose up --build
```

The API will be available at: `http://localhost:8080`
The Frontend will be available at: `http://localhost:3000`

## Docker Services

The application runs the following services:

- **App Service**: Go application (port 8080)
- **PostgreSQL**: Database service (port 5432)
- **Redis**: Cache service (port 6379)

## Development

### Rebuilding the Application

```bash
# Rebuild and restart all services
docker-compose up --build

# Rebuild only the app service
docker-compose build app
docker-compose up app
```

### Viewing Logs

```bash
# View all services logs
docker-compose logs

# View specific service logs
docker-compose logs app
docker-compose logs postgres
docker-compose logs redis

# Follow logs in real-time
docker-compose logs -f app
```

### Stopping Services

```bash
# Stop all services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

## Project Structure

```
├── cmd/                    # Application entry point
│   └── main.go
├── internal/               # Application code
│   ├── adapters/           # External layer adapters
│   │   ├── cache/          # Redis cache implementation
│   │   ├── configs/        # Configuration management
│   │   ├── http/           # HTTP handlers and routing
│   │   ├── providers/      # External data providers
│   │   └── storage/        # Database implementations
│   └── core/               # Business logic layer
│       ├── models/         # Domain models
│       ├── ports/          # Interfaces/contracts
│       └── services/       # Business services
├── pkg/                    # Public packages
│   ├── errors/             # Error handling utilities
│   ├── request/            # Request utilities
│   └── scorer/             # Content scoring system
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
└── .env
```

## API Endpoints

The service provides REST API endpoints for content management. The HTTP router is configured in [`internal/adapters/http/router.go`](internal/adapters/http/router.go).

## Adding New Providers

To add new content providers, implement the provider interface in the `internal/adapters/providers` package.
