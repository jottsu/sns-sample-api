.PHONY: db

build:
	@GOOS=linux GOARCH=amd64 go build -o ./bin/sns-sample-api
	@docker-compose build --no-cache

up:
	@docker-compose up -d db
	@$(call goose-up)
	@docker-compose up -d api

restart: build
	@docker-compose up -d api

down:
	@docker-compose down

db:
	@docker exec -it sns-sample-db mysql -udbuser -pdbuserpass sns_sample

logs:
	@docker logs -f sns-sample-api

migrate_new:
	@goose create migration sql

define goose-up
	@echo "waiting for db to be ready for migration...";
	@RETRY_COUNTER=7; until goose -env local up; do \
		[ ! "$${RETRY_COUNTER}" = "0" ] || { echo "failed to connect to mysql."; exit 1; }; \
		sleep 4 && ((RETRY_COUNTER--)); \
	done
endef