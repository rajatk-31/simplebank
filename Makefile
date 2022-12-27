postgres:
	docker run --name root -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it root createdb --username=root --owner=root root

dropdb:
	docker exec -it root dropdb root

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/root?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/root?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown