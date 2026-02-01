package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

type PoolConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func InitDB(connectionString string, pool PoolConfig) (*sql.DB, error) {
	// Membuat konfigurasi pgx dari connection string
	config, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	// Supabase transaction pooler membutuhkan simple protocol
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// Membuka koneksi database melalui pgx stdlib
	db := stdlib.OpenDB(*config)

	// Menguji koneksi
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Konfigurasi connection pool (disesuaikan untuk Supabase transaction pooling)
	if pool.MaxOpenConns > 0 {
		db.SetMaxOpenConns(pool.MaxOpenConns)
	}
	if pool.MaxIdleConns > 0 {
		db.SetMaxIdleConns(pool.MaxIdleConns)
	}
	if pool.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(pool.ConnMaxLifetime)
	}

	log.Println("Database connected successfully")
	return db, nil
}
