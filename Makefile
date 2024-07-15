DB_URL=postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable

migrate:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

rollback:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: migrateup migrate migratedown rollback sqlc test server mock
