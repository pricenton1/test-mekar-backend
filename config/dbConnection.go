package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// NewDatabase -> initial mysql database.
func NewDatabase(conf *Config) *sql.DB {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DatabaseUser, conf.DatabasePass, conf.DatabaseHost, conf.DatabasePort, conf.DatabaseName)
	fmt.Println(source)
	db, err := sql.Open(conf.DatabaseDriver, source)
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database Connected ...")
	return db
}
