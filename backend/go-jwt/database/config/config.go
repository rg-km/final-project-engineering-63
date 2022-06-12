package config

import (
	"context"
	"database/sql"
	"go_jwt/helper"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbname = "go_jwt"
)

func DBConnect() (*sql.DB, error) {
	// Remove DB If Name Already Exist
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

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	helper.PanicIfErrorWithMessage("Erorr when ping database : ", err)

	log.Printf("Connected to DB %s successfully\n", dbname)

	return db, nil
}
