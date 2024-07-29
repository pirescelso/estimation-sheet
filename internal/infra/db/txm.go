package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrRepositoryNotRegistered     = errors.New("repository not registered")
	ErrRepositoryAlreadyRegistered = errors.New("repository already registered")
	ErrInvalidRepositoryType       = errors.New("invalid repository type")
)

type RepositoryName string
type Repository any
type RepositoryFactory func(queries *Queries) any

type TransactionManagerInterface interface {
	Register(name RepositoryName, factory RepositoryFactory) error
	Do(ctx context.Context, fn func(ctx context.Context, tx TransactionInterface) error) error
	UnRegister(name RepositoryName) error
}

type TransactionInterface interface {
	GetRepository(name RepositoryName) (Repository, error)
}

type Transaction struct {
	queries      *Queries
	repositories map[RepositoryName]RepositoryFactory
}

func NewTransaction(queries *Queries, repositories map[RepositoryName]RepositoryFactory) *Transaction {
	return &Transaction{
		queries:      queries,
		repositories: repositories,
	}
}

func GetAs[T any](t TransactionInterface, name RepositoryName) (T, error) {
	repository, err := t.GetRepository(name)

	var res T
	if err != nil {
		return res, err
	}
	res, ok := repository.(T)
	if !ok {
		return res, ErrInvalidRepositoryType
	}

	return res, nil
}

func (t *Transaction) GetRepository(name RepositoryName) (Repository, error) {
	if repository, ok := t.repositories[name]; ok {
		return repository(t.queries), nil
	}

	return nil, ErrRepositoryNotRegistered
}

type TransactionManager struct {
	dbpool       *pgxpool.Pool
	repositories map[RepositoryName]RepositoryFactory
}

func NewTransactionManager(dbpool *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{
		dbpool:       dbpool,
		repositories: map[RepositoryName]RepositoryFactory{},
	}
}

func (t *TransactionManager) Register(name RepositoryName, factory RepositoryFactory) error {
	if _, ok := t.repositories[name]; ok {
		return ErrRepositoryAlreadyRegistered
	}

	t.repositories[name] = factory
	return nil
}

func (t *TransactionManager) UnRegister(name RepositoryName) error {
	if _, ok := t.repositories[name]; !ok {
		return ErrRepositoryNotRegistered
	}

	delete(t.repositories, name)
	return nil
}

func (t *TransactionManager) Do(ctx context.Context, fn func(ctx context.Context, tx TransactionInterface) error) error {
	tx, err := t.dbpool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queries := New(t.dbpool).WithTx(tx)
	err = fn(ctx, NewTransaction(queries, t.repositories))
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
