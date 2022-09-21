package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/shared/customerror"
)

type PostgresStore struct {
	connection *sqlx.DB
	config     Config
	ctx        context.Context
	bulk       *BulkStore
}

type Config struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifetime int
}

func New(config Config) interfaces.ISqlDataStore {
	return &PostgresStore{
		config: config,
		ctx:    context.Background(),
		bulk:   Bulk(context.Background(), config),
	}
}

func (db *PostgresStore) Connect() error {
	config := db.config
	conn, err := sqlx.ConnectContext(db.getCtx(), "postgres", config.URL)
	if err != nil {
		return err
	}
	if config.MaxOpenConns > 0 {
		conn.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		conn.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxConnLifetime > 0 {
		conn.SetConnMaxLifetime(time.Second * time.Duration(config.MaxConnLifetime))
	}
	db.connection = conn
	return nil
}

func (db *PostgresStore) WithContext(ctx context.Context) interfaces.ISqlDataStore {
	db.ctx = ctx
	return db
}

func (db *PostgresStore) ApplySchemaAndDropData(schema string) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	_, err = conn.Exec(schema)
	return err
}

func (db *PostgresStore) Get(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	return conn.GetContext(db.getCtx(), found, query, args...)
}

func (db *PostgresStore) QueryOne(query string, found interface{}, args ...interface{}) error {
	return db.Get(query, found, args...)
}

func (db *PostgresStore) Query(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	err = conn.QueryRowxContext(db.getCtx(), query, args...).StructScan(found)
	return err
}

func (db *PostgresStore) QueryAll(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	err = conn.SelectContext(db.getCtx(), found, query, args...)
	return err
}

func (db *PostgresStore) Exec(query string, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	result, err := conn.ExecContext(db.getCtx(), query, args...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 || err != nil {
		return customerror.NotFound("No rows affected")
	}
	return nil
}

func (db *PostgresStore) ExecAndReturnID(query string, args ...interface{}) (string, error) {
	conn, err := db.getConnection()
	if err != nil {
		return "", err
	}
	result, err := conn.ExecContext(db.getCtx(), query, args...)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func (db *PostgresStore) ExecAndReturnRowsAffected(query string, args ...interface{}) (int64, error) {
	conn, err := db.getConnection()
	if err != nil {
		return 0, err
	}
	result, err := conn.ExecContext(db.getCtx(), query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (db *PostgresStore) BulkInsert(table string, columns []string, rows [][]interface{}) error {
	return db.bulk.Copy(table, columns, rows)
}

func (db *PostgresStore) getConnection() (*sqlx.DB, error) {
	var err error
	if db.connection == nil {
		if err := db.Connect(); err != nil {
			return db.connection, err
		}
	}
	return db.connection, err
}
func (db *PostgresStore) getCtx() context.Context {
	if db.ctx != nil {
		return db.ctx
	}
	return context.Background()
}