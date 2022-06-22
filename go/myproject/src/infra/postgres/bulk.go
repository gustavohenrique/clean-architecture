package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"{{ .ProjectName }}/src/shared/conf"
)

type BulkStore struct {
	connection *pgxpool.Pool
	config     *conf.Config
	ctx        context.Context
}

func Bulk(ctx context.Context, config *conf.Config) *BulkStore {
	pgxconf, _ := pgxpool.ParseConfig(config.Store.Postgres.URL)
	maxOpenConns := config.Store.Postgres.MaxOpenConns
	maxConnLifetime := config.Store.Postgres.MaxConnLifetime
	maxIdleConns := config.Store.Postgres.MaxIdleConns
	if maxOpenConns > 0 {
		pgxconf.MaxConns = int32(maxOpenConns)
	}
	if maxConnLifetime > 0 {
		pgxconf.MaxConnLifetime = time.Duration(maxConnLifetime) * time.Second
	}
	if maxIdleConns > 0 {
		pgxconf.MaxConnIdleTime = time.Duration(maxIdleConns) * time.Second
	}
	connection, _ := pgxpool.ConnectConfig(ctx, pgxconf)
	if ctx == nil {
		ctx = context.Background()
	}
	return &BulkStore{
		connection: connection,
		config:     config,
		ctx:        ctx,
	}
}

func (db *BulkStore) Copy(table string, columns []string, rows [][]interface{}) error {
	if len(rows) == 0 || len(columns) == 0 {
		return fmt.Errorf("Columns ou rows com length zero")
	}
	if len(columns) != len(rows[0]) {
		return fmt.Errorf("O total de columns nao corresponde ao total de valores para inserir")
	}
	conn := db.connection
	if conn == nil {
		return fmt.Errorf("Não conectou ao banco %s", db.config.Store.Postgres.URL)
	}
	_, err := db.connection.CopyFrom(
		db.ctx,
		pgx.Identifier{table},
		columns,
		pgx.CopyFromRows(rows),
	)
	return err
}
