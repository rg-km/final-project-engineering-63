package drop

import (
	"database/sql"
	"go_jwt/helper"
	"log"
)

func DropAllTable(db *sql.DB) {
	query := `
		DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS categories;
		DROP TABLE IF EXISTS quizzes;
		DROP TABLE IF EXISTS answers;
	`
	_, err := db.Exec(query)
	helper.PanicIfErrorWithMessage("Error when dropping table:", err)

	log.Println("Dropped table successfully")
}
