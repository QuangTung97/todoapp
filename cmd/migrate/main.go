package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"todoapp/config"
	"todoapp/lib/mysql"
)

func main() {
	conf := config.Load()
	cmd := mysql.MigrateCommand(conf.MySQL.DSN())
	err := cmd.Execute()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
