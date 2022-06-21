package drop

import (
	"context"
	"database/sql"
	"go_jwt/helper"
	"log"
	"time"
)

func DropAllTable(db *sql.DB) {
	query := `
		DROP TABLE IF EXISTS users;
	`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, query)
	helper.PanicIfErrorWithMessage("Error when dropping table:", err)

	log.Println("Dropped table successfully")
}
