# Architecture

This project follows a classic separated stack:

- **frontend**: Vue 3 + Vite + Pinia + Axios renders boards/topics/posts and handles authoring flows.
- **backend**: Go (Gin + GORM) exposes REST APIs and persists every entity in MySQL.
- **database**: MySQL schema mirrors the ER diagram; SQL lives in ackend/migrations.

## Data flow

1. Vue calls the /api endpoints (proxied by Vite) through Axios.
2. Gin handlers rely on the repository layer to query or mutate entities.
3. GORM keeps the schema in sync via auto migrations and centralizes connection pooling.

## API surface

| Method | Path | Description |
| --- | --- | --- |
| GET | /api/health | Health check |
| GET | /api/boards | Board list |
| GET | /api/topics | Topic list |
| GET | /api/posts?boardId=&topicId= | Posts with optional filters |
| GET | /api/posts/:id | Single post with comments |
| POST | /api/posts | Create a post |
| POST | /api/posts/:id/comments | Add a comment to a post |

## Environment

Copy ackend/.env.example to .env and tweak APP_PORT plus all MYSQL_* values.

## Run order

1. Start MySQL and run every SQL file in ackend/migrations.
2. From ackend/: go mod tidy then go run ./cmd/api.
3. From rontend/: 
pm install then 
pm run dev.
4. Visit http://localhost:5173 and interact with the forum UI.

