package database

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"
)

const QueryKey = "QUERY_KEY_DB_SQL"

func NewQueryContext(ctx context.Context, db *sql.DB) context.Context {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", time.Now().UTC())))
	privateKey := fmt.Sprintf("%x", h.Sum(nil))
	ctx = context.WithValue(ctx, QueryKey, privateKey)
	ctx = context.WithValue(ctx, privateKey, db)
	return ctx
}

func QueryFromContext(ctx context.Context) (*sql.DB, bool) {
	privateKey, ok := ctx.Value(QueryKey).(string)
	if !ok {
		return nil, ok
	}

	db, ok := ctx.Value(privateKey).(*sql.DB)
	return db, ok
}

type TransactionCallback func(ctx context.Context, tx *sql.Tx) error

func RunInTransaction(ctx context.Context, db *sql.DB, fn TransactionCallback) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	ctx = NewQueryContext(ctx, db)
	err = fn(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
