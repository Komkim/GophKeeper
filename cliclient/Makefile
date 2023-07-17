.PHONY: mock-gen migrate-generate migrate-up migrate-down

MIGRATION_DIR := "storage/sqlite/migrations"
DB := "client.db"

migrate-generate:
	goose -dir $(MIGRATION_DIR) create $(name) go

migrate-up:
	goose -dir $(MIGRATION_DIR) sqlite3 $(DB) up

migrate-down:
	goose -dir $(MIGRATION_DIR) sqlite3 $(DB) down