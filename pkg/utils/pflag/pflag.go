package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var (
	host     = pflag.StringP("host", "H", "127.0.0.1:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "root", "Password for access to mysql, should be used pair with password")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// Parse command line flags
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	fmt.Printf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
}
