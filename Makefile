LOCAL_BIN := $(CURDIR)/bin

include ./.env

export PATH := $(PATH):$(LOCAL_BIN)

.PHONY: deps
deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: infra
infra:
	docker-compose up -d --build

.PHONY: minfra
minfra-up: infra
	$(LOCAL_BIN)/goose postgres "user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} dbname=${DB_NAME} sslmode=disable" -dir=./migrations up

.PHONY: minfra-down
minfra-down:
	$(LOCAL_BIN)/goose postgres "user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} dbname=${DB_NAME} sslmode=disable" -dir=./migrations reset

# If the first argument is "run"...
ifeq (migration,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  MIGRATION_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(MIGRATION_ARGS):;@:)
endif
.PHONY: migration
migration:
	$(LOCAL_BIN)/goose create -dir=./migrations $(MIGRATION_ARGS) sql

