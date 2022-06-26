package config

import (
	"database/sql"
	"go_jwt/helper"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbname = "fpe_63"
)

func DBConnect() (*sql.DB, error) {
	os.RemoveAll("database/" + dbname + ".db")

	log.Println("Creating " + dbname + ".db...")
	file, err := os.Create("database/" + dbname + ".db")
	helper.PanicIfErrorWithMessage("Error when creating database:", err)
	file.Close()

	log.Println(dbname + ".db created")

	db, err := sql.Open("sqlite3", "database/"+dbname+".db")
	helper.PanicIfErrorWithMessage("Error when connecting to database:", err)

	// Sett DB Pool
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	err = db.Ping()
	helper.PanicIfErrorWithMessage("Erorr when ping database : ", err)

	log.Printf("Connected to DB %s successfully\n", dbname)

	return db, nil
}
