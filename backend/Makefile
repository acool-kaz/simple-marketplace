run:
	go run cmd/main.go

open:
	docker exec -it test_db psql -U postgres

swagger:
	swag init -g cmd/main.go

start:
	docker run --name=test_db -e POSTGRES_PASSWORD='12345' -p 5436:5432 -d --rm postgres
	@sleep 5
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5436/postgres?sslmode=disable' up
	go run cmd/main.go

restart:
	docker stop test_db
	docker run --name=test_db -e POSTGRES_PASSWORD='12345' -p 5436:5432 -d --rm postgres
	@sleep 5
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5436/postgres?sslmode=disable' up
	go run cmd/main.go