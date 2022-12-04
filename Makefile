POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgrespw
POSTGRES_DATABASE=hotel

-include .env
  
DB_URL=postgresql://postgres:postgrespw@localhost:5432/hotel?sslmode=disable
print:
	echo "$(DB_URL)"
	
swag:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go


migrate_file:
	migrate create -ext sql -dir migrations/ -seq alter_some_table
migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

.PHONY: start migrateup migratedown