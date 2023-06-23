# as docker was not installed and we are using podman instead of it thru
# alias docker=podman
# for makefile, docker command will result in not found error
# hence use 'podman' instead of 'docker' to start the command

# postgres: 
# 	podman run --name postgres15.3 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine

# createdb:
# 	podman exec -it postgres15.3 createdb --username=root --owner=root simple_bank

# dropdb:
# 	podman exec -it postgres15.3 dropdb simple_bank

SHELL := /bin/bash
DOCKER := $(shell command -v docker 2> /dev/null)

ifeq ($(DOCKER),)
    DOCKER := podman
endif

postgres: 
	$(DOCKER) run --name postgres15.3 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine

createdb:
	$(DOCKER) exec -it postgres15.3 createdb --username=root --owner=root simple_bank

dropdb:
	$(DOCKER) exec -it postgres15.3 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc

