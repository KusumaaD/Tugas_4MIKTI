package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
  dbDriver := "mysql"
  dbUser := "root"
  dbPass := ""
  dbName := "Tugas_4MIKTII"

  db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
  if err != nil {
    log.Fatal(err)
  }

  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  return db
}
