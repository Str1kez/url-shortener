up:
	docker compose up -d --remove-orphans

down:
	docker compose down

swagger:
	docker run -p 80:8080 -e SWAGGER_JSON=/schema.yml -v ${PWD}/openapi.yaml:/schema.yml:ro -d --name swagger swaggerapi/swagger-ui

swagger-down:
	docker container rm -f swagger

