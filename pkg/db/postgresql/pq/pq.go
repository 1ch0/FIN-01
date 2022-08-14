package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select inet_server_addr(),pg_is_in_recovery(),current_database(),current_user")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var inet_server_addr string
		var pg_is_in_recovery string
		var current_database string
		var current_user string
		err = rows.Scan(&inet_server_addr, &pg_is_in_recovery, &current_database, &current_user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("inet_server_addr: " + inet_server_addr)
		fmt.Println("pg_is_in_recovery: " + pg_is_in_recovery)
		fmt.Println("current_database: " + current_database)
		fmt.Println("current_user: " + current_user)
	}
}
