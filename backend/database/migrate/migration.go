package migrate

import (
	"database/sql"
	"go_jwt/helper"
	"log"
)

func Migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(50) NOT NULL,
		password TEXT NOT NULL,
		role VARCHAR(10) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		category VARCHAR(50) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS quizzes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		question TEXT NOT NULL,
		category_id INTEGER NOT NULL,
		answer_true VARCHAR(255) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);
	
	CREATE TABLE IF NOT EXISTS answers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		answer_a TEXT NOT NULL,
		answer_b TEXT NOT NULL,
		answer_c TEXT NOT NULL,
		quiz_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
	);
	`

	_, err := db.Exec(query)
	helper.PanicIfErrorWithMessage("Error when migrate with error:", err)

	log.Println("Migration success")
}
