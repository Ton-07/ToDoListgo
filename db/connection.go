package main

// conexão com o banco

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Conectar() error {
	var err error
	dsn := "postgres://ton:1234@localhost:5432/todolist"
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar no banco: %d", err)
	}

	return nil
}
