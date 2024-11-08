postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root las_tlayudas

dropdb:
	docker exec -it postgres12 dropdb las_tlayudas

rundb:
	docker start postgres12

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/las_tlayudas?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/las_tlayudas?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

start: 
	make rundb slqc

init:
	make postgres createdb rundb migrateup sqlc
	
.POHNY: postgres createbd dropdb rundb migrateup migratedown sqlc test server start init