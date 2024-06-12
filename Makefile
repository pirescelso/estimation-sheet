DB_URL=postgres://postgres:postgres@db:5432/postgres?sslmode=disable

migrateup:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose drop

test:
	go test -cover ./...

test-clean:
	go clean --testcache

.PHONY:  migrateup migratedown test test-clean
