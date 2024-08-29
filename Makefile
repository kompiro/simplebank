DB_SOURCE=postgresql://postgres:postgres@db:5432/simple_bank?sslmode=disable
IMAGE_TAG:=latest

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

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

test:
	go test -v -cover ./...

test-report:
	@DB_SOURCE=$(DB_SOURCE) go run gotest.tools/gotestsum@latest --junitfile test-report.xml -- -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

app.image.build:
	docker buildx build -t simplebank:latest --target app .

app.image.push:
	docker image tag simplebank:latest ghcr.io/kompiro/simplebank:latest
	docker image push ghcr.io/kompiro/simplebank:latest

app.image.ecspresso:
	@IMAGE_TAG=$(IMAGE_TAG) ecspresso deploy --config ecspresso/simplebank/ecspresso.yml 

migrate.image.build:
	docker buildx build -t simplebank-migrate:latest --target migrate .

migrate.image.push:
	docker image tag simplebank-migrate:latest ghcr.io/kompiro/simplebank-migrate:latest
	docker image push ghcr.io/kompiro/simplebank-migrate:latest

migrate.image.ecspresso:
	@IMAGE_TAG=$(IMAGE_TAG) ecspresso run --config ecspresso/migrate/ecspresso.yml


release:
	gh release create `date +rel-%Y%m%d` --generate-notes

.PHONY: migrateup migrate migratedown rollback \
  db_docs db_schema sqlc test server mock proto \
	app.image.build app.image.push app.image.ecspresso \
	migrate.image.build migrate.image.push migrate.image.ecspresso \
	release 

