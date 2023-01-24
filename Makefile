all: build

up:
	docker compose up -d --remove-orphans --build

db:
	docker compose up -d --remove-orphans db

down:
	docker compose down

swagger:
	docker run -p 80:8080 -e SWAGGER_JSON=/schema.yml -v ${PWD}/openapi.yaml:/schema.yml:ro -d --name swagger swaggerapi/swagger-ui

swagger-down:
	docker container rm -f swagger

env:
	cp .env.example .env

run:
	go run cmd/main.go

run-prod: upgrade
	urlshortener

build:
	go build -v -o $$GOPATH/bin/urlshortener ./cmd/main.go

format:
	gofmt -w -s -l .

upgrade:
	export $$(cat .env); migrate -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB?sslmode=$$SSLMODE" -path ./migrations up

downgrade:
	export $$(cat .env); migrate -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB?sslmode=$$SSLMODE" -path ./migrations down
