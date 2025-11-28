# Simple Forum (Vue 3 + Go + MySQL)

A minimal forum that mirrors the provided ER diagram. Boards contain posts, posts belong to topics, and users can publish posts as well as comments.

- **frontend**: Vue 3 + Vite + Pinia + Axios offers the SPA with board/topic filters, post creation, and comment flows.
- **backend**: Go + Gin + GORM surfaces REST APIs and persists the entities.
- **database**: MySQL schema lives in ackend/migrations with seed data for boards/topics.

## Quick start

`ash
# Database
mysql -u<user> -p < backend/migrations/001_init.sql
mysql -u<user> -p < backend/migrations/002_seed.sql
cp backend/.env.example backend/.env   # adjust APP_PORT + MYSQL_*

# Backend
cd backend
go mod tidy
go run ./cmd/api

# Frontend
cd frontend
npm install
npm run dev

# Visit http://localhost:5173
`

## Structure

`
backend/
  cmd/api/main.go          Gin entrypoint
  internal/config          env helpers
  internal/db              GORM connection + migrations
  internal/handlers        REST handlers
  internal/repository      Data access helpers
  migrations/*.sql         schema + seed
frontend/
  src/api/http.js          Axios instance
  src/stores/forumStore.js Pinia store wrapping the API
  src/components/*         Reusable UI widgets
  src/App.vue              Layout + orchestration
`

Additional notes live in docs/architecture.md.

