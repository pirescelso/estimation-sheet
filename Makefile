DB_URL=postgres://postgres:postgres@db:5432/postgres?sslmode=disable&search_path=estimation
# DB_URL=postgres://test:test@db:5432/test?sslmode=disable&search_path=estimation
# go test -v -cover -p 1 -count 1 -run ^TestIntegration ./...

migrateup:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose drop

test-unit:
	go test -v -race -cover -count 1 -run ^TestUnit ./...

test-integration:
	go test -v -cover -p 1 -count 1 -run ^TestIntegration ./...

test-e2e:
	go test -v -p 1 -count 1 -run ^TestE2E ./test/e2e/...

test-clean:
	go clean --testcache

run:
	go run ./cmd/estimation/main.go

GO_LDFLAGS := -X main.buildTime=$(shell date -u +%Y-%m-%dT%H:%M:%SZ) -X main.commitHash=$(shell git rev-parse HEAD)

build:
	go build -ldflags "$(GO_LDFLAGS)" -o bin/estimation-sheet cmd/estimation/main.go


.PHONY:  migrateup migratedown test-unit test-integration test-e2e test-clean run build
