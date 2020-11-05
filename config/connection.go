package config

import (
	"database/sql"
	"echo/utils"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDbConnection() (*sql.DB, error) {

	log.Println("=> open db connection")
	loadData := fmt.Sprintf("root:admin@tcp(localhost:3306)/test")
	fmt.Println(loadData)
	db, err := sql.Open("mysql", loadData)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(utils.DbMaxConns)
	db.SetMaxIdleConns(utils.DbMaxIdleConns)

	return db, nil
}
