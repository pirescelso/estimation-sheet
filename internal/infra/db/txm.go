package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type RepositoryFactory func(queries *Queries) any

type TransactionManagerInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func() error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type TransactionManager struct {
	Conn         *pgx.Conn
	Queries      *Queries
	Tx           pgx.Tx
	Repositories map[string]RepositoryFactory
}

func NewTransactionManager(ctx context.Context, conn *pgx.Conn) *TransactionManager {
	return &TransactionManager{
		Conn:         conn,
		Repositories: map[string]RepositoryFactory{},
	}
}

func (t *TransactionManager) Register(name string, fc RepositoryFactory) {
	t.Repositories[name] = fc
}

func (t *TransactionManager) UnRegister(name string) {
	delete(t.Repositories, name)
}

func (t *TransactionManager) GetRepository(ctx context.Context, name string) (any, error) {
	if t.Tx == nil {
		tx, err := t.Conn.Begin(ctx)
		if err != nil {
			return nil, err
		}
		t.Tx = tx
	}
	qtx := t.Queries.WithTx(t.Tx)
	repo := t.Repositories[name](qtx)
	return repo, nil
}

func (t *TransactionManager) Do(ctx context.Context, fn func() error) error {
	if t.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := t.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	t.Tx = tx
	err = fn()
	if err != nil {
		errRb := t.Rollback(ctx)
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	return t.CommitOrRollback(ctx)
}

func (t *TransactionManager) Rollback(ctx context.Context) error {
	if t.Tx == nil {
		return errors.New("no transaction to rollback")
	}
	err := t.Tx.Rollback(ctx)
	if err != nil {
		return err
	}
	t.Tx = nil
	return nil
}

func (t *TransactionManager) CommitOrRollback(ctx context.Context) error {
	err := t.Tx.Commit(ctx)
	if err != nil {
		errRb := t.Rollback(ctx)
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	t.Tx = nil
	return nil
}
