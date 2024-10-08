postgres:
	docker run --rm --name postgresql -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:alpine

createdb:
	docker exec -it postgresql createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresql dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover .

server:
	go run main.go

mock:
	mockgen --destination db/mock/store.go -package mockdb github.com/m-bashari-m/simplebank/db/sqlc Store

.PHONY: createdb migratedown migrateup postgres dropdb sqlc test server mock