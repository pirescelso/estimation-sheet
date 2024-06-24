package testutils

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	"github.com/celsopires1999/estimation/configs"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func DBSetup() (*pgxpool.Pool, *migrate.Migrate) {
	_, base, _, _ := runtime.Caller(0)
	path := filepath.Dir(filepath.Dir(filepath.Dir(base)))

	ctx := context.Background()
	configs := configs.LoadConfig(path, ".test")
	dbpool, err := pgxpool.New(ctx, configs.DBConn)

	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	m, err := migrate.New("file:"+path+"/sql/migrations", configs.DBConn)
	if err != nil {
		log.Fatalf("Unable to create migrator: %v\n", err)
	}
	if m == nil {
		log.Fatalf("Migrator is nil\n")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion {
		log.Fatalf("Unable to migrate database: %v\n", err)
	}

	return dbpool, m
}

func TruncateTables(dbpool *pgxpool.Pool) error {
	ctx := context.Background()

	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "TRUNCATE TABLE budget_allocations CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "TRUNCATE TABLE budgets CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "TRUNCATE TABLE portfolios CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "TRUNCATE TABLE cost_allocations CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "TRUNCATE TABLE costs CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "TRUNCATE TABLE baselines CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "TRUNCATE TABLE users CASCADE;")
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "TRUNCATE TABLE plans CASCADE;")
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
