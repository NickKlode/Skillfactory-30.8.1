package main

import (
	"context"
	"fmt"
	"os"
	"rdb/pkg/storage"

	"github.com/jackc/pgx/v4"
)

func main() {
	dsn := "postgresql://postgres:ZAQzaqzaq97@localhost:5433/postgres"
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	db, err := storage.New(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database NEW: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(db.Tasks(1, 1))
	db.DeleteTask(1)
}
