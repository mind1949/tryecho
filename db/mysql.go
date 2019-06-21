package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysqldns = "root:ipassword@tcp(localhost:3306)/hello"
)

var MysqlDB *sql.DB

func init() {
	MysqlDB = openMysqlDB()
}

func openMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", mysqldns)
	if err != nil {
		log.Fatal("can't open mysql db: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("can't open mysql db: ", err)
	}
	log.Println("successfully open mysql db")
	return db
}
