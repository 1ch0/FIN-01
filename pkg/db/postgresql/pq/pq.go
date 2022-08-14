package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	//connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=disable"
	db, err := sql.Open("postgres", "user=astaxie password=astaxie dbname=test sslmode=disable")

	//db, err := sql.Open("postgres", connStr)

	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	// pg 不支持这个函数，因为他没有类似 MySQL 的自增 ID
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入id =", lastInsertId)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rows, err := db.Query("select inet_server_addr(),pg_is_in_recovery(),current_database(),current_user")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for rows.Next() {
//		var inet_server_addr string
//		var pg_is_in_recovery string
//		var current_database string
//		var current_user string
//		err = rows.Scan(&inet_server_addr, &pg_is_in_recovery, &current_database, &current_user)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println("inet_server_addr: " + inet_server_addr)
//		fmt.Println("pg_is_in_recovery: " + pg_is_in_recovery)
//		fmt.Println("current_database: " + current_database)
//		fmt.Println("current_user: " + current_user)
//	}
//
//	checkErr(err)
//
//	// 插入数据
//	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
//	checkErr(err)
//
//	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
//	checkErr(err)
//
//	// pg 不支持这个函数，因为他没有类似 MySQL 的自增 ID
//	// id, err := res.LastInsertId()
//	// checkErr(err)
//	// fmt.Println(id)
//
//	var lastInsertId int
//	err = db.QueryRow("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
//	checkErr(err)
//	fmt.Println("最后插入id =", lastInsertId)
//
//	// 更新数据
//	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
//	checkErr(err)
//
//	res, err = stmt.Exec("astaxieupdate", 1)
//	checkErr(err)
//
//	affect, err := res.RowsAffected()
//	checkErr(err)
//
//	fmt.Println(affect)
//
//	// 查询数据
//	rows, err := db.Query("SELECT * FROM userinfo")
//	checkErr(err)
//
//	for rows.Next() {
//		var uid int
//		var username string
//		var department string
//		var created string
//		err = rows.Scan(&uid, &username, &department, &created)
//		checkErr(err)
//		fmt.Println(uid)
//		fmt.Println(username)
//		fmt.Println(department)
//		fmt.Println(created)
//	}
//
//	// 删除数据
//	stmt, err = db.Prepare("delete from userinfo where uid=$1")
//	checkErr(err)
//
//	res, err = stmt.Exec(1)
//	checkErr(err)
//
//	affect, err = res.RowsAffected()
//	checkErr(err)
//
//	fmt.Println(affect)
//
//	db.Close()
//
//}
//
//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
