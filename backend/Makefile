run:
	go run ./cmd/api

tidy:
	go mod tidy

migrate:
	mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) $(MYSQL_DB) < migrations/001_init.sql
	mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) $(MYSQL_DB) < migrations/002_seed.sql

