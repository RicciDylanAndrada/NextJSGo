postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres15 dropdb go-chat

createMigrateDb:
	migrate create --ext sql -dir db/migrations ${ARGS}

migrateUpDb:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up	

migrateDropDb:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose drop

migrateDownDb:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit createdb dropdb postgres createMigrateDb  migrateUpDb  migrateDropDb migrateDownDb

