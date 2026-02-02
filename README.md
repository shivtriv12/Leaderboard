# Leaderboard System

A high-performance real-time leaderboard application capable of handling millions of users with infinite scrolling, search functionality, and real-time rank updates.

## Architecture

The system uses a hybrid architecture with PostgreSQL as the source of truth and Redis for high-speed ranking and retrieval.

- **Rank Calculation:** Standard Competition Ranking (1224) logic.
- **Search:** Hybrid approach using PostgreSQL `pg_trgm` for text matching and Redis `ZCOUNT` for O(1) rank retrieval.
- **Simulation:** A background worker updates random user scores every 10 seconds to simulate live traffic.
- **Data Consistency:** Redis is transient; the application builds the cache from the database on startup.

## Tech Stack

- **Backend:** Go (Golang), PostgreSQL, Redis
- **Frontend:** React, TypeScript, Tailwind CSS, Vite
- **Infrastructure:** Docker, Docker Compose

## Requirements

- Docker
- Go 1.23+ (for local development)
- Node.js & pnpm (for frontend development)

## Setup and Running

1. **Clone the repository**

   ```bash
   git clone https://github.com/shivtriv12/Leaderboard.git
   cd Leaderboard
   ```

2. **Start Backend & Database**
   This starts PostgreSQL, Redis, and the Go API server.

   ```bash
   cd backend
   docker-compose up --build
   ```

   The backend will be available at `http://localhost:8080`.

3. **Start Frontend**
   Open a new terminal.
   ```bash
   cd frontend
   pnpm install
   pnpm dev
   ```
   The frontend will be available at `http://localhost:5173`.

## Environment Variables

### Backend

The backend relies on `.env` or environment variables set in `docker-compose.yml`.

- `DB_URL`: PostgreSQL connection string.
- `REDIS_ADDR`: Redis address (e.g., `redis:6379`).

### Frontend

- `VITE_BASE_API_URL`: The URL of the backend API (default: `http://localhost:8080`).

## API Endpoints

- `GET /api/leaderboard`: Fetch paginated leaderboard data.
  - Query params: `cursor` (username for pagination), `limit`.
- `GET /api/search`: Search for users by username.
  - Query params: `q` (search query), `cursor`, `limit`.

## Project Structure

```
.
├── backend/
│   ├── internal/api/       # HTTP Handlers
│   ├── internal/simulation # Background worker for updating scores
│   ├── sql/                # SQL queries and migrations
│   └── main.go             # Entry point
└── frontend/
    ├── src/components/     # React components
    ├── src/services/       # API integration
    └── src/App.tsx         # Main application logic
```
