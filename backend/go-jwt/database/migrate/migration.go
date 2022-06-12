package migrate

import (
	"context"
	"database/sql"
	"go_jwt/helper"
	"log"
	"time"
)

func Migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		role VARCHAR(255) NOT NULL,
		is_login BOOLEAN NOT NULL
	);
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, query)
	helper.PanicIfErrorWithMessage("Error when migrate with error:", err)

	log.Println("Migration success")
}
