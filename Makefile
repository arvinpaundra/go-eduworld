createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root eduworld_db

dropdb:
	sudo docker exec -it postgres dropdb eduworld_db

migrateup:
	migrate -database "postgres://root:root@localhost:5431/eduworld_db?sslmode=disable" -path migrations up

migratedown:
	migrate -database "postgres://root:root@localhost:5431/eduworld_db?sslmode=disable" -path migrations down

run:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown run