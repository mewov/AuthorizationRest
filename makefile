.PHONY: help, build, up, down

help:
	@echo "make build - build and run docker compose service"
	@echo "make up - run docker compose service"
	@echo "make down - remove docker compose service"

build:
	@docker compose up --build
up:
	@docker compose up
down:
	@docker compose down
