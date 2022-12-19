package db

import (
	"database/sql"
	"fmt"

	"github.com/antonlindstrom/pgstore"
	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://postgres:celinejkt48@localhost:5432/bank-neo?sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresStore(db *sql.DB) (*pgstore.PGStore, error) {

	authKey := []byte("authkey")
	encryptionKey := []byte("APUPPT0123456789APUPPT0123456789")
	store, err := pgstore.NewPGStoreFromPool(db, authKey, encryptionKey)
	if err != nil {
		return nil, err
	}
	return store, nil
}
