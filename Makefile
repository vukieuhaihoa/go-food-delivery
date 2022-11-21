# =======================================================================
# Database config
postgres_init:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine
postgres_stop:
	docker stop postgres14
postgres_delete:
	docker rm postgres14
postgres_restart:
	docker restart postgres14
db_init:
	docker exec -it postgres14 createdb --username=root --owner=root food_delivery
migrate_up:
	migrate --path db/migration --database "postgresql://root:root@localhost:5432/food_delivery?sslmode=disable" -verbose up
migrate_down:
	migrate --path db/migration --database "postgresql://root:root@localhost:5432/food_delivery?sslmode=disable" -verbose down


# =======================================================================
# App config
start:
	go run main.go
dev:
	nodemon --exec go run main.go --signal SIGTERM