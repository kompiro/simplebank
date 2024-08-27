DB_SOURCE=postgresql://postgres:postgres@db:5432/simple_bank?sslmode=disable

migrate:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

rollback:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

test-report:
	go run gotest.tools/gotestsum@latest --junitfile test-report.xml -- -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

app.image.build:
	docker buildx build -t simplebank:latest --target app .

app.image.push:
	docker image tag simplebank:latest ghcr.io/kompiro/simplebank:latest
	docker image push ghcr.io/kompiro/simplebank:latest

migrate.image.build:
	docker buildx build -t simplebank-migrate:latest --target migrate .

migrate.image.push:
	docker image tag simplebank-migrate:latest ghcr.io/kompiro/simplebank-migrate:latest
	docker image push ghcr.io/kompiro/simplebank-migrate:latest

release:
	gh release create `date +rel-%Y%m%d` --generate-notes

.PHONY: migrateup migrate migratedown rollback sqlc test server mock app.image.build app.image.push migrate.image.build migrate.image.push release

