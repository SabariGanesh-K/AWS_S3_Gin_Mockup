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
server:
	go run main.go

test:
	go test -v -cover ./db/sqlc ./api ./token ./util
testmail:
	go test -v -cover ./email

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/SabariGanesh-K/21BPS1209_Backend.git/db/sqlc Store


redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine
.PHONY: postgres new_migration createdb dropdb migrateup createmigrateinitschema sqlc mock test  redis testmail