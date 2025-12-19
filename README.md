# Codeforces Calendar Toolkit

A Go-based tool that automatically syncs upcoming Codeforces contests to your Google Calendar.

## Features
- Fetches upcoming contests from Codeforces API.
- Filters for future events only.
- Authenticates with Google Calendar via OAuth2.
- Prevents duplicates using a Postgres database.
- Runs automatically every 24 hours (Dockerized).

## How to Run

### Option 1: Docker (Recommended)
1.  **Prerequisites**: Docker Desktop installed.
2.  **Setup**:
    - Place your `credentials.json` (from Google Cloud Console) in the root directory.
3.  **Run**:
    ```bash
    docker-compose up --build -d
    ```
    The app will start in the background and check for contests daily.

### Option 2: Local Development
1.  **Prerequisites**: Go 1.23+, Postgres running locally.
2.  **Setup**:
    - Set environment variable `DB_HOST=localhost`.
    - Ensure Postgres is running and accessible.
3.  **Run**:
    ```bash
    go run .
    ```

## Tech Stack
- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Libraries**: `pgx` (DB driver), `google-api-go-client`
- **Deployment**: Docker & Docker Compose

## License
MIT
