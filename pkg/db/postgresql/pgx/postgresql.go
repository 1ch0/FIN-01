package main

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	//urlExample := "postgres://username:password@localhost:5432/database_name"
	config := pgx.ConnConfig{}
	conn, err := pgx.Connect(config)
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	var name string
	var weight int64
	err = conn.QueryRow("select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
