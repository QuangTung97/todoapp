package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"todoapp/lib/mysql"
)

func main() {
	dsn := "root:1@tcp(localhost:3306)/bench?parseTime=true"
	cmd := mysql.MigrateCommand(dsn)
	err := cmd.Execute()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
