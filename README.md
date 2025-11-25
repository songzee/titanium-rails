# Titanium Rails - Issuer Processor Prototype

A minimal Go service for issuer-processor prototyping.

## Quick Start

### Run locally

```bash
make run
```

Or:

```bash
go run ./cmd/api
```

Server starts on port 8080 (or set `PORT` env var).

### Build

```bash
make build
```

### Docker

```bash
make docker-build
docker run -p 8080:8080 titanium-rails
```

## API Endpoints

### Health Check

```bash
curl http://localhost:8080/health
```

Response:
```json
{"status":"ok"}
```

### STMS Auth

```bash
curl -X POST http://localhost:8080/stms/auth \
  -H "Content-Type: application/json" \
  -d '{"amount": 1000, "cardId": "card123"}'
```

Response:
```json
{"approved":true,"reason":"approved"}
```

## Project Structure

```
/cmd/api/main.go          # HTTP server entry point
/internal/server/server.go # Router + health check
/internal/handlers/auth.go # STMS auth handler
/internal/ledger/ledger.go # Ledger interface + stub
/internal/stms/models.go   # STMS request/response models
/internal/config/config.go # Environment config
```

## Environment Variables

- `PORT` - Server port (default: 8080)

## GitHub Setup

### Initial Push

```bash
# Initialize git (if not already done)
git init

# Add all files
git add .

# Commit
git commit -m "Initial commit: issuer-processor prototype"

# Add remote (replace with your repo URL)
git remote add origin https://github.com/yourusername/titanium-rails.git

# Push to main branch
git branch -M main
git push -u origin main
```

### CI/CD

The project includes a GitHub Actions CI workflow (`.github/workflows/ci.yml`) that:

- Runs on push/PR to `main` or `master`
- Tests the code with `go test`
- Builds the binary
- Checks code formatting with `gofmt`

CI will run automatically on every push and pull request.

