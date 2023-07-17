.PHONY: mock-gen migrate-generate migrate-up migrate-down

MIGRATION_DIR := "storage/postgres/migrations"
PG_DSN := "postgres://postgres:postgres@localhost:5432/keep?sslmode=disable"

migrate-generate:
	goose -dir $(MIGRATION_DIR) create $(name) go

migrate-up:
	goose -dir $(MIGRATION_DIR) postgres $(PG_DSN) up

migrate-down:
	goose -dir $(MIGRATION_DIR) postgres $(PG_DSN) down