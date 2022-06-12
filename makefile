#!make
include ./config/.env.local
export $(shell sed 's/=.*//' ./config/.env.local)

migrate-up:
	sql-migrate up -config=./config/dbconfig.yml -env=development

migrate-down:
	sql-migrate down -config=./config/dbconfig.yml -env=development