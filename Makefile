DB_URL=postgresql://root:secret@localhost:5432/filemanager?sslmode=disable


postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createmigrateinitschema:
	 migrate create -ext sql --dir db/migrations --seq init_schema  

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

createdb:
	docker exec -it postgres createdb --username=root --owner=root filemanager

dropdb:
	docker exec -it postgres dropdb filemanager
sqlc:
	sqlc generate

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

.PHONY: postgres new_migration createdb dropdb migrateup createmigrateinitschema sqlc